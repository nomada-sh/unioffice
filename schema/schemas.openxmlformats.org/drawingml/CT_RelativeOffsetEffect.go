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

type CT_RelativeOffsetEffect struct {
	TxAttr *ST_Percentage
	TyAttr *ST_Percentage
}

func NewCT_RelativeOffsetEffect() *CT_RelativeOffsetEffect {
	ret := &CT_RelativeOffsetEffect{}
	return ret
}
func (m *CT_RelativeOffsetEffect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.TxAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "tx"},
			Value: fmt.Sprintf("%v", *m.TxAttr)})
	}
	if m.TyAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ty"},
			Value: fmt.Sprintf("%v", *m.TyAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_RelativeOffsetEffect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "tx" {
			parsed, err := ParseUnionST_Percentage(attr.Value)
			if err != nil {
				return err
			}
			m.TxAttr = &parsed
		}
		if attr.Name.Local == "ty" {
			parsed, err := ParseUnionST_Percentage(attr.Value)
			if err != nil {
				return err
			}
			m.TyAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_RelativeOffsetEffect: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_RelativeOffsetEffect) Validate() error {
	return m.ValidateWithPath("CT_RelativeOffsetEffect")
}
func (m *CT_RelativeOffsetEffect) ValidateWithPath(path string) error {
	if m.TxAttr != nil {
		if err := m.TxAttr.ValidateWithPath(path + "/TxAttr"); err != nil {
			return err
		}
	}
	if m.TyAttr != nil {
		if err := m.TyAttr.ValidateWithPath(path + "/TyAttr"); err != nil {
			return err
		}
	}
	return nil
}
