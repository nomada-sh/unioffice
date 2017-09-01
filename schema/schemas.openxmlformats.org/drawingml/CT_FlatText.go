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
)

type CT_FlatText struct {
	ZAttr *ST_Coordinate
}

func NewCT_FlatText() *CT_FlatText {
	ret := &CT_FlatText{}
	return ret
}
func (m *CT_FlatText) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.ZAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "z"},
			Value: fmt.Sprintf("%v", *m.ZAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_FlatText) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "z" {
			parsed, err := ParseUnionST_Coordinate(attr.Value)
			if err != nil {
				return err
			}
			m.ZAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_FlatText: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_FlatText) Validate() error {
	return m.ValidateWithPath("CT_FlatText")
}
func (m *CT_FlatText) ValidateWithPath(path string) error {
	if m.ZAttr != nil {
		if err := m.ZAttr.ValidateWithPath(path + "/ZAttr"); err != nil {
			return err
		}
	}
	return nil
}
