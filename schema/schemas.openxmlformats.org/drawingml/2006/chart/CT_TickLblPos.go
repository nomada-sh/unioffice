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
)

type CT_TickLblPos struct {
	ValAttr ST_TickLblPos
}

func NewCT_TickLblPos() *CT_TickLblPos {
	ret := &CT_TickLblPos{}
	return ret
}
func (m *CT_TickLblPos) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.ValAttr != ST_TickLblPosUnset {
		attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "val"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TickLblPos) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			m.ValAttr.UnmarshalXMLAttr(attr)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TickLblPos: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_TickLblPos) Validate() error {
	return m.ValidateWithPath("CT_TickLblPos")
}
func (m *CT_TickLblPos) ValidateWithPath(path string) error {
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
