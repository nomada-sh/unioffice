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
	"strconv"
)

type CT_SmartTagPr struct {
	// Embed SmartTags
	EmbedAttr *bool
	// Show Smart Tags
	ShowAttr ST_SmartTagShow
}

func NewCT_SmartTagPr() *CT_SmartTagPr {
	ret := &CT_SmartTagPr{}
	return ret
}
func (m *CT_SmartTagPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.EmbedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "embed"},
			Value: fmt.Sprintf("%v", *m.EmbedAttr)})
	}
	if m.ShowAttr != ST_SmartTagShowUnset {
		attr, err := m.ShowAttr.MarshalXMLAttr(xml.Name{Local: "show"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_SmartTagPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "embed" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.EmbedAttr = &parsed
		}
		if attr.Name.Local == "show" {
			m.ShowAttr.UnmarshalXMLAttr(attr)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_SmartTagPr: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_SmartTagPr) Validate() error {
	return m.ValidateWithPath("CT_SmartTagPr")
}
func (m *CT_SmartTagPr) ValidateWithPath(path string) error {
	if err := m.ShowAttr.ValidateWithPath(path + "/ShowAttr"); err != nil {
		return err
	}
	return nil
}
