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

type CT_PivotDimension struct {
	// Measure
	MeasureAttr *bool
	// Dimension Name
	NameAttr string
	// Dimension Unique Name
	UniqueNameAttr string
	// Dimension Display Name
	CaptionAttr string
}

func NewCT_PivotDimension() *CT_PivotDimension {
	ret := &CT_PivotDimension{}
	return ret
}
func (m *CT_PivotDimension) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.MeasureAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "measure"},
			Value: fmt.Sprintf("%v", *m.MeasureAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uniqueName"},
		Value: fmt.Sprintf("%v", m.UniqueNameAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "caption"},
		Value: fmt.Sprintf("%v", m.CaptionAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_PivotDimension) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "measure" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.MeasureAttr = &parsed
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
		}
		if attr.Name.Local == "uniqueName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UniqueNameAttr = parsed
		}
		if attr.Name.Local == "caption" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CaptionAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_PivotDimension: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_PivotDimension) Validate() error {
	return m.ValidateWithPath("CT_PivotDimension")
}
func (m *CT_PivotDimension) ValidateWithPath(path string) error {
	return nil
}
