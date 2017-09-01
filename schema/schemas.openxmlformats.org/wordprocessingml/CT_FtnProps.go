// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"log"
)

type CT_FtnProps struct {
	// Footnote Placement
	Pos *CT_FtnPos
	// Footnote Numbering Format
	NumFmt *CT_NumFmt
	// Footnote and Endnote Numbering Starting Value
	NumStart *CT_DecimalNumber
	// Footnote and Endnote Numbering Restart Location
	NumRestart *CT_NumRestart
}

func NewCT_FtnProps() *CT_FtnProps {
	ret := &CT_FtnProps{}
	return ret
}
func (m *CT_FtnProps) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.Pos != nil {
		sepos := xml.StartElement{Name: xml.Name{Local: "w:pos"}}
		e.EncodeElement(m.Pos, sepos)
	}
	if m.NumFmt != nil {
		senumFmt := xml.StartElement{Name: xml.Name{Local: "w:numFmt"}}
		e.EncodeElement(m.NumFmt, senumFmt)
	}
	if m.NumStart != nil {
		senumStart := xml.StartElement{Name: xml.Name{Local: "w:numStart"}}
		e.EncodeElement(m.NumStart, senumStart)
	}
	if m.NumRestart != nil {
		senumRestart := xml.StartElement{Name: xml.Name{Local: "w:numRestart"}}
		e.EncodeElement(m.NumRestart, senumRestart)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_FtnProps) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_FtnProps:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "pos":
				m.Pos = NewCT_FtnPos()
				if err := d.DecodeElement(m.Pos, &el); err != nil {
					return err
				}
			case "numFmt":
				m.NumFmt = NewCT_NumFmt()
				if err := d.DecodeElement(m.NumFmt, &el); err != nil {
					return err
				}
			case "numStart":
				m.NumStart = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.NumStart, &el); err != nil {
					return err
				}
			case "numRestart":
				m.NumRestart = NewCT_NumRestart()
				if err := d.DecodeElement(m.NumRestart, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_FtnProps
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_FtnProps) Validate() error {
	return m.ValidateWithPath("CT_FtnProps")
}
func (m *CT_FtnProps) ValidateWithPath(path string) error {
	if m.Pos != nil {
		if err := m.Pos.ValidateWithPath(path + "/Pos"); err != nil {
			return err
		}
	}
	if m.NumFmt != nil {
		if err := m.NumFmt.ValidateWithPath(path + "/NumFmt"); err != nil {
			return err
		}
	}
	if m.NumStart != nil {
		if err := m.NumStart.ValidateWithPath(path + "/NumStart"); err != nil {
			return err
		}
	}
	if m.NumRestart != nil {
		if err := m.NumRestart.ValidateWithPath(path + "/NumRestart"); err != nil {
			return err
		}
	}
	return nil
}
