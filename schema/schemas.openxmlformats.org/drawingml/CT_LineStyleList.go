// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_LineStyleList struct {
	Ln []*CT_LineProperties
}

func NewCT_LineStyleList() *CT_LineStyleList {
	ret := &CT_LineStyleList{}
	return ret
}
func (m *CT_LineStyleList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	seln := xml.StartElement{Name: xml.Name{Local: "a:ln"}}
	e.EncodeElement(m.Ln, seln)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_LineStyleList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_LineStyleList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "ln":
				tmp := NewCT_LineProperties()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Ln = append(m.Ln, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_LineStyleList
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_LineStyleList) Validate() error {
	return m.ValidateWithPath("CT_LineStyleList")
}
func (m *CT_LineStyleList) ValidateWithPath(path string) error {
	for i, v := range m.Ln {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Ln[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
