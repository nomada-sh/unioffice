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

type CT_Path2DArcTo struct {
	WRAttr    ST_AdjCoordinate
	HRAttr    ST_AdjCoordinate
	StAngAttr ST_AdjAngle
	SwAngAttr ST_AdjAngle
}

func NewCT_Path2DArcTo() *CT_Path2DArcTo {
	ret := &CT_Path2DArcTo{}
	return ret
}
func (m *CT_Path2DArcTo) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "wR"},
		Value: fmt.Sprintf("%v", m.WRAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hR"},
		Value: fmt.Sprintf("%v", m.HRAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "stAng"},
		Value: fmt.Sprintf("%v", m.StAngAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "swAng"},
		Value: fmt.Sprintf("%v", m.SwAngAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Path2DArcTo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "wR" {
			parsed, err := ParseUnionST_AdjCoordinate(attr.Value)
			if err != nil {
				return err
			}
			m.WRAttr = parsed
		}
		if attr.Name.Local == "hR" {
			parsed, err := ParseUnionST_AdjCoordinate(attr.Value)
			if err != nil {
				return err
			}
			m.HRAttr = parsed
		}
		if attr.Name.Local == "stAng" {
			parsed, err := ParseUnionST_AdjAngle(attr.Value)
			if err != nil {
				return err
			}
			m.StAngAttr = parsed
		}
		if attr.Name.Local == "swAng" {
			parsed, err := ParseUnionST_AdjAngle(attr.Value)
			if err != nil {
				return err
			}
			m.SwAngAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Path2DArcTo: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_Path2DArcTo) Validate() error {
	return m.ValidateWithPath("CT_Path2DArcTo")
}
func (m *CT_Path2DArcTo) ValidateWithPath(path string) error {
	if err := m.WRAttr.ValidateWithPath(path + "/WRAttr"); err != nil {
		return err
	}
	if err := m.HRAttr.ValidateWithPath(path + "/HRAttr"); err != nil {
		return err
	}
	if err := m.StAngAttr.ValidateWithPath(path + "/StAngAttr"); err != nil {
		return err
	}
	if err := m.SwAngAttr.ValidateWithPath(path + "/SwAngAttr"); err != nil {
		return err
	}
	return nil
}
