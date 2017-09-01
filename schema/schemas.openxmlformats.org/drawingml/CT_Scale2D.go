// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"log"
)

type CT_Scale2D struct {
	Sx *CT_Ratio
	Sy *CT_Ratio
}

func NewCT_Scale2D() *CT_Scale2D {
	ret := &CT_Scale2D{}
	ret.Sx = NewCT_Ratio()
	ret.Sy = NewCT_Ratio()
	return ret
}
func (m *CT_Scale2D) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Name.Local = "a:CT_Scale2D"
	e.EncodeToken(start)
	sesx := xml.StartElement{Name: xml.Name{Local: "a:sx"}}
	e.EncodeElement(m.Sx, sesx)
	sesy := xml.StartElement{Name: xml.Name{Local: "a:sy"}}
	e.EncodeElement(m.Sy, sesy)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Scale2D) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Sx = NewCT_Ratio()
	m.Sy = NewCT_Ratio()
lCT_Scale2D:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "sx":
				if err := d.DecodeElement(m.Sx, &el); err != nil {
					return err
				}
			case "sy":
				if err := d.DecodeElement(m.Sy, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Scale2D
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Scale2D) Validate() error {
	return m.ValidateWithPath("CT_Scale2D")
}
func (m *CT_Scale2D) ValidateWithPath(path string) error {
	if err := m.Sx.ValidateWithPath(path + "/Sx"); err != nil {
		return err
	}
	if err := m.Sy.ValidateWithPath(path + "/Sy"); err != nil {
		return err
	}
	return nil
}
