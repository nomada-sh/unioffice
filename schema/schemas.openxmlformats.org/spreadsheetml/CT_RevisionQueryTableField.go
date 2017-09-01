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

type CT_RevisionQueryTableField struct {
	// Sheet Id
	SheetIdAttr uint32
	// QueryTable Reference
	RefAttr string
	// Field Id
	FieldIdAttr uint32
}

func NewCT_RevisionQueryTableField() *CT_RevisionQueryTableField {
	ret := &CT_RevisionQueryTableField{}
	return ret
}
func (m *CT_RevisionQueryTableField) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sheetId"},
		Value: fmt.Sprintf("%v", m.SheetIdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ref"},
		Value: fmt.Sprintf("%v", m.RefAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fieldId"},
		Value: fmt.Sprintf("%v", m.FieldIdAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_RevisionQueryTableField) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "sheetId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.SheetIdAttr = uint32(parsed)
		}
		if attr.Name.Local == "ref" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RefAttr = parsed
		}
		if attr.Name.Local == "fieldId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.FieldIdAttr = uint32(parsed)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_RevisionQueryTableField: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_RevisionQueryTableField) Validate() error {
	return m.ValidateWithPath("CT_RevisionQueryTableField")
}
func (m *CT_RevisionQueryTableField) ValidateWithPath(path string) error {
	return nil
}
