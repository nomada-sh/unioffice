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

type CT_DataBinding struct {
	// XML Namespace Prefix Mappings
	PrefixMappingsAttr *string
	// XPath
	XpathAttr string
	// Custom XML Data Storage ID
	StoreItemIDAttr string
}

func NewCT_DataBinding() *CT_DataBinding {
	ret := &CT_DataBinding{}
	return ret
}
func (m *CT_DataBinding) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.PrefixMappingsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:prefixMappings"},
			Value: fmt.Sprintf("%v", *m.PrefixMappingsAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:xpath"},
		Value: fmt.Sprintf("%v", m.XpathAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:storeItemID"},
		Value: fmt.Sprintf("%v", m.StoreItemIDAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_DataBinding) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "prefixMappings" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PrefixMappingsAttr = &parsed
		}
		if attr.Name.Local == "xpath" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.XpathAttr = parsed
		}
		if attr.Name.Local == "storeItemID" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StoreItemIDAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_DataBinding: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_DataBinding) Validate() error {
	return m.ValidateWithPath("CT_DataBinding")
}
func (m *CT_DataBinding) ValidateWithPath(path string) error {
	return nil
}
