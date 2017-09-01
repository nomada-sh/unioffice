// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"
	"log"
)

type CT_OMathArgPr struct {
	ArgSz *CT_Integer2
}

func NewCT_OMathArgPr() *CT_OMathArgPr {
	ret := &CT_OMathArgPr{}
	return ret
}
func (m *CT_OMathArgPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.ArgSz != nil {
		seargSz := xml.StartElement{Name: xml.Name{Local: "m:argSz"}}
		e.EncodeElement(m.ArgSz, seargSz)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_OMathArgPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_OMathArgPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "argSz":
				m.ArgSz = NewCT_Integer2()
				if err := d.DecodeElement(m.ArgSz, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_OMathArgPr
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_OMathArgPr) Validate() error {
	return m.ValidateWithPath("CT_OMathArgPr")
}
func (m *CT_OMathArgPr) ValidateWithPath(path string) error {
	if m.ArgSz != nil {
		if err := m.ArgSz.ValidateWithPath(path + "/ArgSz"); err != nil {
			return err
		}
	}
	return nil
}
