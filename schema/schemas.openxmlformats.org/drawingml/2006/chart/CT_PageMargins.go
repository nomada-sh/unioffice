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
	"strconv"
)

type CT_PageMargins struct {
	LAttr      float64
	RAttr      float64
	TAttr      float64
	BAttr      float64
	HeaderAttr float64
	FooterAttr float64
}

func NewCT_PageMargins() *CT_PageMargins {
	ret := &CT_PageMargins{}
	return ret
}
func (m *CT_PageMargins) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "l"},
		Value: fmt.Sprintf("%v", m.LAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r"},
		Value: fmt.Sprintf("%v", m.RAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "t"},
		Value: fmt.Sprintf("%v", m.TAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "b"},
		Value: fmt.Sprintf("%v", m.BAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "header"},
		Value: fmt.Sprintf("%v", m.HeaderAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "footer"},
		Value: fmt.Sprintf("%v", m.FooterAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_PageMargins) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "l" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.LAttr = parsed
		}
		if attr.Name.Local == "r" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.RAttr = parsed
		}
		if attr.Name.Local == "t" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.TAttr = parsed
		}
		if attr.Name.Local == "b" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.BAttr = parsed
		}
		if attr.Name.Local == "header" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.HeaderAttr = parsed
		}
		if attr.Name.Local == "footer" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.FooterAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_PageMargins: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_PageMargins) Validate() error {
	return m.ValidateWithPath("CT_PageMargins")
}
func (m *CT_PageMargins) ValidateWithPath(path string) error {
	return nil
}
