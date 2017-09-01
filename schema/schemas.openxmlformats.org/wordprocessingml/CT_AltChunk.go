// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_AltChunk struct {
	IdAttr *string
	// External Content Import Properties
	AltChunkPr *CT_AltChunkPr
}

func NewCT_AltChunk() *CT_AltChunk {
	ret := &CT_AltChunk{}
	return ret
}
func (m *CT_AltChunk) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	e.EncodeToken(start)
	if m.AltChunkPr != nil {
		sealtChunkPr := xml.StartElement{Name: xml.Name{Local: "w:altChunkPr"}}
		e.EncodeElement(m.AltChunkPr, sealtChunkPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_AltChunk) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
		}
	}
lCT_AltChunk:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "altChunkPr":
				m.AltChunkPr = NewCT_AltChunkPr()
				if err := d.DecodeElement(m.AltChunkPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_AltChunk
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_AltChunk) Validate() error {
	return m.ValidateWithPath("CT_AltChunk")
}
func (m *CT_AltChunk) ValidateWithPath(path string) error {
	if m.AltChunkPr != nil {
		if err := m.AltChunkPr.ValidateWithPath(path + "/AltChunkPr"); err != nil {
			return err
		}
	}
	return nil
}
