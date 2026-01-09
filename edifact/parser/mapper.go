package parser

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/lorenzoliver/edi-tools/edifact/ast"
)

type mappingTag struct {
	Min  int
	Max  int
	Type string
	Len  int
}

func parseTags(t string) mappingTag {
	mTag := mappingTag{}
	for part := range strings.SplitSeq(t, ",") {
		kv := strings.Split(part, "=")
		if len(kv) != 2 {
			continue
		}
		switch kv[0] {
		case "min":
			mTag.Min, _ = strconv.Atoi(kv[1])
		case "max":
			mTag.Max, _ = strconv.Atoi(kv[1])
		case "type":
			mTag.Type = kv[1]
		case "len":
			mTag.Len, _ = strconv.Atoi(kv[1])
		}
	}
	return mTag
}

type cursor struct {
	unh      *ast.Segment
	segments []*ast.Segment
	unt      *ast.Segment
	pos      int
}

func newCursor(unh *ast.Segment, segments []*ast.Segment, unt *ast.Segment) *cursor {
	return &cursor{
		unh:      unh,
		segments: segments,
		unt:      unt,
		pos:      -1,
	}
}

func (c *cursor) current() *ast.Segment {
	if c.pos == -1 {
		return c.unh
	}
	if c.pos < len(c.segments) {
		return c.segments[c.pos]
	}
	if c.pos == len(c.segments) {
		return c.unt
	}
	return nil
}

func (c *cursor) advance() {
	c.pos++
}

func MapEdifact(msg *ast.Message) (any, error) {
	cursor := newCursor(msg.UNH, msg.Segments, msg.UNT)
	spec, err := GetSpec(msg.UNH.Components[1].String())
	if err != nil {
		return nil, err
	}
	specStruct := reflect.ValueOf(spec).Elem()
	err = mapEdifact(cursor, specStruct)
	if err != nil {
		return nil, err
	}
	return spec, nil
}

func validateElement(element, fieldName string, tags mappingTag) error {
	if tags.Min == 1 && element == "" {
		return fmt.Errorf("validation failed for field %s: value is required (min=1)", fieldName)
	}
	if element != "" {
		if tags.Len != 0 && utf8.RuneCountInString(element) > tags.Len {
			return fmt.Errorf("validation failed for field %s: value '%s' is longer than max_len %d", fieldName, element, tags.Len)
		}
	}
	return nil
}

func mapSegment(s *ast.Segment, v reflect.Value) error {
	if v.Kind() != reflect.Pointer {
		return fmt.Errorf("expected pointer to struct, got %T", v.Interface())
	}
	structVal := v.Elem()
	structType := structVal.Type()
	componentIdx := 0
	for i := 0; i < structType.NumField(); i++ {
		if componentIdx >= len(s.Components) {
			break
		}
		field := structType.Field(i)
		fieldTags := field.Tag.Get("edi")
		tags := parseTags(fieldTags)
		fieldVal := structVal.Field(i)
		switch fieldVal.Kind() {
		case reflect.String:
			if len(s.Components[componentIdx].Elements) > 0 {
				element := s.Components[componentIdx].Elements[0]
				if err := validateElement(element, field.Name, tags); err != nil {
					return err
				}
				fieldVal.SetString(element)

			} else {
				if err := validateElement("", field.Name, tags); err != nil {
					return err
				}
			}
		case reflect.Struct:
			newComposite := reflect.New(fieldVal.Type())
			if err := mapComposite(s.Components[componentIdx], newComposite); err != nil && tags.Min > 0 {
				return err
			}
			fieldVal.Set(newComposite.Elem())
		default:
			// TODO: repeating composites and repeating elements
			panic("unreachable")
		}
		componentIdx++
	}
	return nil
}

func mapComposite(c *ast.Component, v reflect.Value) error {
	if v.Kind() != reflect.Pointer || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("expected a pointer to a struct, got %T", v.Interface())
	}
	structVal := v.Elem()
	structType := structVal.Type()

	elementIdx := 0
	for i := 0; i < structType.NumField(); i++ {
		if elementIdx >= len(c.Elements) {
			break
		}

		field := structType.Field(i)
		tags := parseTags(field.Tag.Get("edi"))

		fieldVal := structVal.Field(i)

		switch fieldVal.Kind() {
		case reflect.String:
			element := c.Elements[elementIdx]
			if err := validateElement(element, field.Name, tags); err != nil {
				return err
			}
			fieldVal.SetString(element)
			elementIdx++
		case reflect.Slice:
			if fieldVal.Type().Elem().Kind() == reflect.String {
				fieldsLeft := structType.NumField() - i - 1
				elementsAvailable := len(c.Elements) - elementIdx - fieldsLeft
				numToConsume := tags.Max
				numToConsume = min(elementsAvailable, numToConsume)

				if numToConsume > 0 {
					sliceElements := c.Elements[elementIdx : elementIdx+numToConsume]
					for _, el := range sliceElements {
						if err := validateElement(el, field.Name, tags); err != nil {
							return err
						}
					}
					fieldVal.Set(reflect.ValueOf(sliceElements))
					elementIdx += numToConsume
				}
			}
		default:
			panic("unreachable")
		}
	}
	return nil
}

func mapEdifact(cursor *cursor, obj reflect.Value) error {
	t := obj.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tagStr := field.Tag.Get("edi")
		// Not an Edifact specified field
		if tagStr == "" {
			continue
		}
		tag := parseTags(tagStr)
		fieldVal := obj.Field(i)
		if tag.Type == "g" {
			sliceType := fieldVal.Type().Elem()
			for count := 0; count < tag.Max; count++ {
				s := cursor.current()
				if s == nil {
					break
				}
				groupStartTag := ""
				if sliceType.NumField() > 0 {
					f := sliceType.Field(0)
					groupStartTag = f.Name
				}

				if s.Tag != groupStartTag {
					break
				}

				newValue := reflect.New(sliceType)
				if err := mapEdifact(cursor, newValue.Elem()); err != nil {
					return err
				}
				fieldVal.Set(reflect.Append(fieldVal, newValue.Elem()))
			}
		} else if tag.Type == "s" {
			if fieldVal.Kind() == reflect.Slice {
				for count := 0; count < tag.Max; count++ {
					elemType := fieldVal.Type().Elem()
					s := cursor.current()
					if s == nil || s.Tag != elemType.Name() {
						break
					}
					newInstance := reflect.New(elemType)
					if err := mapSegment(s, newInstance); err != nil {
						return err
					}
					fieldVal.Set(reflect.Append(fieldVal, newInstance.Elem()))
					cursor.advance()
				}
			} else if fieldVal.Kind() == reflect.Struct {
				elemType := fieldVal.Type()
				s := cursor.current()
				if s == nil || s.Tag != elemType.Name() {
					break
				}
				if err := mapSegment(s, fieldVal.Addr()); err != nil {
					return err
				}
				cursor.advance()
			}
		} else {
			panic("unreachable")
		}
	}
	return nil
}
