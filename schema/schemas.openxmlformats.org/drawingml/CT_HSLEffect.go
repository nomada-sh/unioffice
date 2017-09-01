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
	"strconv"
)

type CT_HSLEffect struct {
	HueAttr *int32
	SatAttr *ST_FixedPercentage
	LumAttr *ST_FixedPercentage
}

func NewCT_HSLEffect() *CT_HSLEffect {
	ret := &CT_HSLEffect{}
	return ret
}
func (m *CT_HSLEffect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.HueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hue"},
			Value: fmt.Sprintf("%v", *m.HueAttr)})
	}
	if m.SatAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sat"},
			Value: fmt.Sprintf("%v", *m.SatAttr)})
	}
	if m.LumAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lum"},
			Value: fmt.Sprintf("%v", *m.LumAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_HSLEffect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "hue" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.HueAttr = &pt
		}
		if attr.Name.Local == "sat" {
			parsed, err := ParseUnionST_FixedPercentage(attr.Value)
			if err != nil {
				return err
			}
			m.SatAttr = &parsed
		}
		if attr.Name.Local == "lum" {
			parsed, err := ParseUnionST_FixedPercentage(attr.Value)
			if err != nil {
				return err
			}
			m.LumAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_HSLEffect: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_HSLEffect) Validate() error {
	return m.ValidateWithPath("CT_HSLEffect")
}
func (m *CT_HSLEffect) ValidateWithPath(path string) error {
	if m.HueAttr != nil {
		if *m.HueAttr < 0 {
			return fmt.Errorf("%s/m.HueAttr must be >= 0 (have %v)", path, *m.HueAttr)
		}
		if *m.HueAttr >= 21600000 {
			return fmt.Errorf("%s/m.HueAttr must be < 21600000 (have %v)", path, *m.HueAttr)
		}
	}
	if m.SatAttr != nil {
		if err := m.SatAttr.ValidateWithPath(path + "/SatAttr"); err != nil {
			return err
		}
	}
	if m.LumAttr != nil {
		if err := m.LumAttr.ValidateWithPath(path + "/LumAttr"); err != nil {
			return err
		}
	}
	return nil
}
