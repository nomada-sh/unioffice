// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_PivotFmts struct {
	PivotFmt []*CT_PivotFmt
}

func NewCT_PivotFmts() *CT_PivotFmts {
	ret := &CT_PivotFmts{}
	return ret
}
func (m *CT_PivotFmts) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.PivotFmt != nil {
		sepivotFmt := xml.StartElement{Name: xml.Name{Local: "pivotFmt"}}
		e.EncodeElement(m.PivotFmt, sepivotFmt)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_PivotFmts) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_PivotFmts:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "pivotFmt":
				tmp := NewCT_PivotFmt()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.PivotFmt = append(m.PivotFmt, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PivotFmts
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_PivotFmts) Validate() error {
	return m.ValidateWithPath("CT_PivotFmts")
}
func (m *CT_PivotFmts) ValidateWithPath(path string) error {
	for i, v := range m.PivotFmt {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/PivotFmt[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
