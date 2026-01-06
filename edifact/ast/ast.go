package ast

import (
	"strings"
)

type Interchange struct {
	UNA      UNA
	UNB      *Segment
	Messages []*Message
	UNZ      *Segment
}

type Message struct {
	UNH      *Segment
	Segments []*Segment
	UNT      *Segment
}

type Segment struct {
	Tag        string
	Components []*Component
}

type Component struct {
	Elements []string
}

type UNA string

func (c Component) String() string {
	var el []string
	el = append(el, c.Elements...)
	return strings.Join(el, ":")
}

func (s Segment) String() string {
	var parts []string
	for _, c := range s.Components {
		parts = append(parts, c.String())
	}
	return s.Tag + "+" + strings.Join(parts, "+")

}
