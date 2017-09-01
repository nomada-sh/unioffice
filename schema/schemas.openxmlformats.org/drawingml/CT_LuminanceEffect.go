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

type CT_LuminanceEffect struct {
	BrightAttr   *ST_FixedPercentage
	ContrastAttr *ST_FixedPercentage
}

func NewCT_LuminanceEffect() *CT_LuminanceEffect {
	ret := &CT_LuminanceEffect{}
	return ret
}
func (m *CT_LuminanceEffect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.BrightAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "bright"},
			Value: fmt.Sprintf("%v", *m.BrightAttr)})
	}
	if m.ContrastAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "contrast"},
			Value: fmt.Sprintf("%v", *m.ContrastAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_LuminanceEffect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "bright" {
			parsed, err := ParseUnionST_FixedPercentage(attr.Value)
			if err != nil {
				return err
			}
			m.BrightAttr = &parsed
		}
		if attr.Name.Local == "contrast" {
			parsed, err := ParseUnionST_FixedPercentage(attr.Value)
			if err != nil {
				return err
			}
			m.ContrastAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_LuminanceEffect: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_LuminanceEffect) Validate() error {
	return m.ValidateWithPath("CT_LuminanceEffect")
}
func (m *CT_LuminanceEffect) ValidateWithPath(path string) error {
	if m.BrightAttr != nil {
		if err := m.BrightAttr.ValidateWithPath(path + "/BrightAttr"); err != nil {
			return err
		}
	}
	if m.ContrastAttr != nil {
		if err := m.ContrastAttr.ValidateWithPath(path + "/ContrastAttr"); err != nil {
			return err
		}
	}
	return nil
}
