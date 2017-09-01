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

type CT_TLAnimVariantBooleanVal struct {
	// Value
	ValAttr bool
}

func NewCT_TLAnimVariantBooleanVal() *CT_TLAnimVariantBooleanVal {
	ret := &CT_TLAnimVariantBooleanVal{}
	return ret
}
func (m *CT_TLAnimVariantBooleanVal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TLAnimVariantBooleanVal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ValAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TLAnimVariantBooleanVal: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_TLAnimVariantBooleanVal) Validate() error {
	return m.ValidateWithPath("CT_TLAnimVariantBooleanVal")
}
func (m *CT_TLAnimVariantBooleanVal) ValidateWithPath(path string) error {
	return nil
}
