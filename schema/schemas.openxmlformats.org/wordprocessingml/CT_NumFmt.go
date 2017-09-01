// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
)

type CT_NumFmt struct {
	// Numbering Format Type
	ValAttr ST_NumberFormat
	// Custom Defined Number Format
	FormatAttr *string
}

func NewCT_NumFmt() *CT_NumFmt {
	ret := &CT_NumFmt{}
	ret.ValAttr = ST_NumberFormat(1)
	return ret
}
func (m *CT_NumFmt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "w:val"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	if m.FormatAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:format"},
			Value: fmt.Sprintf("%v", *m.FormatAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_NumFmt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = ST_NumberFormat(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			m.ValAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "format" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FormatAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_NumFmt: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_NumFmt) Validate() error {
	return m.ValidateWithPath("CT_NumFmt")
}
func (m *CT_NumFmt) ValidateWithPath(path string) error {
	if m.ValAttr == ST_NumberFormatUnset {
		return fmt.Errorf("%s/ValAttr is a mandatory field", path)
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
