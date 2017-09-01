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
	"strconv"
)

type CT_SlideSize struct {
	// Extent Length
	CxAttr int32
	// Extent Width
	CyAttr int32
	// Type of Size
	TypeAttr ST_SlideSizeType
}

func NewCT_SlideSize() *CT_SlideSize {
	ret := &CT_SlideSize{}
	ret.CxAttr = 914400
	ret.CyAttr = 914400
	return ret
}
func (m *CT_SlideSize) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cx"},
		Value: fmt.Sprintf("%v", m.CxAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cy"},
		Value: fmt.Sprintf("%v", m.CyAttr)})
	if m.TypeAttr != ST_SlideSizeTypeUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_SlideSize) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CxAttr = 914400
	m.CyAttr = 914400
	for _, attr := range start.Attr {
		if attr.Name.Local == "cx" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.CxAttr = int32(parsed)
		}
		if attr.Name.Local == "cy" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.CyAttr = int32(parsed)
		}
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_SlideSize: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_SlideSize) Validate() error {
	return m.ValidateWithPath("CT_SlideSize")
}
func (m *CT_SlideSize) ValidateWithPath(path string) error {
	if m.CxAttr < 914400 {
		return fmt.Errorf("%s/m.CxAttr must be >= 914400 (have %v)", path, m.CxAttr)
	}
	if m.CxAttr > 51206400 {
		return fmt.Errorf("%s/m.CxAttr must be <= 51206400 (have %v)", path, m.CxAttr)
	}
	if m.CxAttr < 0 {
		return fmt.Errorf("%s/m.CxAttr must be >= 0 (have %v)", path, m.CxAttr)
	}
	if m.CyAttr < 914400 {
		return fmt.Errorf("%s/m.CyAttr must be >= 914400 (have %v)", path, m.CyAttr)
	}
	if m.CyAttr > 51206400 {
		return fmt.Errorf("%s/m.CyAttr must be <= 51206400 (have %v)", path, m.CyAttr)
	}
	if m.CyAttr < 0 {
		return fmt.Errorf("%s/m.CyAttr must be >= 0 (have %v)", path, m.CyAttr)
	}
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	return nil
}
