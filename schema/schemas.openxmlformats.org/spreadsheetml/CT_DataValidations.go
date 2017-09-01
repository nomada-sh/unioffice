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
	"log"
	"strconv"
)

type CT_DataValidations struct {
	// Disable Prompts
	DisablePromptsAttr *bool
	// Top Left Corner (X Coodrinate)
	XWindowAttr *uint32
	// Top Left Corner (Y Coordinate)
	YWindowAttr *uint32
	// Data Validation Item Count
	CountAttr *uint32
	// Data Validation
	DataValidation []*CT_DataValidation
}

func NewCT_DataValidations() *CT_DataValidations {
	ret := &CT_DataValidations{}
	return ret
}
func (m *CT_DataValidations) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.DisablePromptsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "disablePrompts"},
			Value: fmt.Sprintf("%v", *m.DisablePromptsAttr)})
	}
	if m.XWindowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xWindow"},
			Value: fmt.Sprintf("%v", *m.XWindowAttr)})
	}
	if m.YWindowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "yWindow"},
			Value: fmt.Sprintf("%v", *m.YWindowAttr)})
	}
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	sedataValidation := xml.StartElement{Name: xml.Name{Local: "x:dataValidation"}}
	e.EncodeElement(m.DataValidation, sedataValidation)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_DataValidations) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "disablePrompts" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DisablePromptsAttr = &parsed
		}
		if attr.Name.Local == "xWindow" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.XWindowAttr = &pt
		}
		if attr.Name.Local == "yWindow" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.YWindowAttr = &pt
		}
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
		}
	}
lCT_DataValidations:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "dataValidation":
				tmp := NewCT_DataValidation()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.DataValidation = append(m.DataValidation, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DataValidations
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_DataValidations) Validate() error {
	return m.ValidateWithPath("CT_DataValidations")
}
func (m *CT_DataValidations) ValidateWithPath(path string) error {
	for i, v := range m.DataValidation {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/DataValidation[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
