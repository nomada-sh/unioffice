// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_MemberProperties struct {
	// OLAP Member Properties Count
	CountAttr *uint32
	// OLAP Member Property
	Mp []*CT_MemberProperty
}

func NewCT_MemberProperties() *CT_MemberProperties {
	ret := &CT_MemberProperties{}
	return ret
}
func (m *CT_MemberProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	semp := xml.StartElement{Name: xml.Name{Local: "x:mp"}}
	e.EncodeElement(m.Mp, semp)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_MemberProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
		}
	}
lCT_MemberProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "mp":
				tmp := NewCT_MemberProperty()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Mp = append(m.Mp, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_MemberProperties
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_MemberProperties) Validate() error {
	return m.ValidateWithPath("CT_MemberProperties")
}
func (m *CT_MemberProperties) ValidateWithPath(path string) error {
	for i, v := range m.Mp {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Mp[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
