// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_TLTimeAnimateValue struct {
	// Time
	TmAttr *ST_TLTimeAnimateValueTime
	// Formula
	FmlaAttr *string
	// Value
	Val *CT_TLAnimVariant
}

func NewCT_TLTimeAnimateValue() *CT_TLTimeAnimateValue {
	ret := &CT_TLTimeAnimateValue{}
	return ret
}
func (m *CT_TLTimeAnimateValue) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.TmAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "tm"},
			Value: fmt.Sprintf("%v", *m.TmAttr)})
	}
	if m.FmlaAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fmla"},
			Value: fmt.Sprintf("%v", *m.FmlaAttr)})
	}
	e.EncodeToken(start)
	if m.Val != nil {
		seval := xml.StartElement{Name: xml.Name{Local: "p:val"}}
		e.EncodeElement(m.Val, seval)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TLTimeAnimateValue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "tm" {
			parsed, err := ParseUnionST_TLTimeAnimateValueTime(attr.Value)
			if err != nil {
				return err
			}
			m.TmAttr = &parsed
		}
		if attr.Name.Local == "fmla" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FmlaAttr = &parsed
		}
	}
lCT_TLTimeAnimateValue:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "val":
				m.Val = NewCT_TLAnimVariant()
				if err := d.DecodeElement(m.Val, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLTimeAnimateValue
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_TLTimeAnimateValue) Validate() error {
	return m.ValidateWithPath("CT_TLTimeAnimateValue")
}
func (m *CT_TLTimeAnimateValue) ValidateWithPath(path string) error {
	if m.TmAttr != nil {
		if err := m.TmAttr.ValidateWithPath(path + "/TmAttr"); err != nil {
			return err
		}
	}
	if m.Val != nil {
		if err := m.Val.ValidateWithPath(path + "/Val"); err != nil {
			return err
		}
	}
	return nil
}
