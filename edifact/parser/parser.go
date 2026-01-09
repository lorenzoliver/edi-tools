package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"io"

	"github.com/lorenzoliver/edi-tools/edifact/ast"
)

type Parser struct {
	r *bufio.Reader
	d Delimiters
}

func NewParser(r io.Reader) *Parser {
	return &Parser{
		r: bufio.NewReader(r),
		d: DefaultDelimiters(),
	}
}

type Delimiters struct {
	ElementSeperator   byte
	ComponentSeperator byte
	DecimalChar        byte
	ReleaseChar        byte
	Reserved           byte
	SegmentTerminator  byte
}

func DefaultDelimiters() Delimiters {
	return Delimiters{
		ElementSeperator:   ':',
		ComponentSeperator: '+',
		DecimalChar:        '.',
		ReleaseChar:        '?',
		Reserved:           ' ',
		SegmentTerminator:  '\'',
	}
}

func (p Parser) Parse() (*ast.Interchange, error) {
	interchange := &ast.Interchange{}

	peek, err := p.r.Peek(3)
	if err == nil && string(peek) == "UNA" {
		una, err := p.r.Peek(9)

		if err != nil {
			return nil, fmt.Errorf("failed to extract UNA segment %w", err)
		}
		p.d.ElementSeperator = una[3]
		p.d.ComponentSeperator = una[4]
		p.d.DecimalChar = una[5]
		p.d.ReleaseChar = una[6]
		p.d.Reserved = una[7]
		p.d.SegmentTerminator = una[8]
		interchange.UNA = ast.UNA(una)

		// this skips the UNA Segment
		_, err = p.parseSegment()
		if err != nil {
			return nil, err
		}
	}
	interchange.UNB, err = p.parseSegment()
	if err != nil || interchange.UNB.Tag != "UNB" {
		return nil,
			fmt.Errorf("failed to parse UNB segment: %w", err)
	}
	for {
		peek, err := p.r.Peek(3)
		if err != nil {
			if err == io.EOF {
				return nil, fmt.Errorf("unexpected EOF befor UNZ, %w", err)
			}
			return nil, err
		}
		if string(peek) == "UNZ" {
			break
		}
		msg, err := p.parseMessage()
		if err != nil {
			return nil, fmt.Errorf("failed to parse message: %w", err)
		}
		interchange.Messages = append(interchange.Messages, msg)
	}
	interchange.UNZ, err = p.parseSegment()
	if err != nil || interchange.UNZ.Tag != "UNZ" {
		return nil, fmt.Errorf("error parsing UNZ segment: %w", err)
	}

	return interchange, nil
}

func (p Parser) parseMessage() (*ast.Message, error) {
	msg := &ast.Message{}

	// Expect to parse UNH first
	unh, err := p.parseSegment()
	if err != nil || unh.Tag != "UNH" {
		return nil, fmt.Errorf("error parsing UNH segment, %w", err)
	}
	msg.UNH = unh
	for {
		peek, err := p.r.Peek(3)
		if err != nil {
			if err == io.EOF {
				return nil, fmt.Errorf("unexpected EOF befor UNT, %w", err)
			}
			return nil, err
		}
		if string(peek) == "UNT" {
			break
		}
		seg, err := p.parseSegment()
		if err != nil {
			return nil, err
		}
		msg.Segments = append(msg.Segments, seg)
	}
	unt, err := p.parseSegment()
	if err != nil || unt.Tag != "UNT" {
		return nil, fmt.Errorf("error parsing UNT segment, %w", err)
	}
	msg.UNT = unt
	return msg, nil
}

func (p Parser) parseSegment() (*ast.Segment, error) {
	raw, err := p.readUntil(p.d.SegmentTerminator)
	if err != nil {
		return nil, err
	}
	//Sanatize
	raw = bytes.TrimSpace(raw)

	//Split Components
	parts := p.split(raw, p.d.ComponentSeperator)
	if len(parts) == 0 {
		return nil, fmt.Errorf("empty segment found")
	}
	seg := &ast.Segment{
		Tag: string(parts[0]),
	}

	for _, part := range parts[1:] {
		c := &ast.Component{}
		el := p.split(part, p.d.ElementSeperator)
		for _, e := range el {
			c.Elements = append(c.Elements, string(p.removeReleaseChars(e)))
		}
		seg.Components = append(seg.Components, c)
	}
	p.skipNewlines()
	return seg, nil
}

func (p Parser) readUntil(terminator byte) ([]byte, error) {
	var buf bytes.Buffer
	for {
		b, err := p.r.ReadByte()
		if err != nil {
			return nil, fmt.Errorf("error reading Segment: %w", err)
		}
		if b == terminator {
			return buf.Bytes(), nil
		}
		peek, err := p.r.Peek(1)
		if err != nil {
			return nil, fmt.Errorf("segment terminated prematurely: %w", err)
		}
		if b == p.d.ReleaseChar && peek[0] == terminator {
			//this skips the release char
			b, err = p.r.ReadByte()
			if err != nil {
				return nil, fmt.Errorf("error reading Segment: %w", err)
			}
		}
		buf.WriteByte(b)
	}
}

func (p Parser) split(data []byte, seperator byte) [][]byte {
	var res [][]byte
	var cur []byte
	i := 0
	for i < len(data) {
		if data[i] == p.d.ReleaseChar && i+1 < len(data) && ( data[i+1] == seperator) {
			//skip release char
			i++
			cur = append(cur, data[i])
		} else if data[i] == p.d.ReleaseChar && i+1 < len(data) && (data[i+1] == p.d.ReleaseChar) {
			//add both release chars, will be properly cleaned in removeReleaseChars()
			cur = append(cur, data[i:i+2]...)
			i++
		} else if data[i] == seperator {
			res = append(res, cur)
			cur = nil
		} else {
			cur = append(cur, data[i])
		}
		i++
	}
	res = append(res, cur)
	return res
}

func (p Parser) removeReleaseChars(data []byte) []byte {
	i := 0
	for i < len(data){
		if data[i] == p.d.ReleaseChar && i+1 < len(data) && ( data[i+1] == p.d.ReleaseChar) {
			//remove ReleaseChar
			data = append(data[:i], data[i+1:]...)
		}
		i++
	}
	return data
}

func (p Parser) skipNewlines() {
	for {
		peek, err := p.r.Peek(1)
		if err != nil {
			return
		}
		if peek[0] != '\n' && peek[0] != '\r' {
			return
		}
		if _, err := p.r.Discard(1); err != nil {
			return
		}
	}
}
