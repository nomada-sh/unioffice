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

type CT_GroupMember struct {
	// Group Member Unique Name
	UniqueNameAttr string
	// Group
	GroupAttr *bool
}

func NewCT_GroupMember() *CT_GroupMember {
	ret := &CT_GroupMember{}
	return ret
}
func (m *CT_GroupMember) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uniqueName"},
		Value: fmt.Sprintf("%v", m.UniqueNameAttr)})
	if m.GroupAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "group"},
			Value: fmt.Sprintf("%v", *m.GroupAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_GroupMember) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "uniqueName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UniqueNameAttr = parsed
		}
		if attr.Name.Local == "group" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.GroupAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_GroupMember: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_GroupMember) Validate() error {
	return m.ValidateWithPath("CT_GroupMember")
}
func (m *CT_GroupMember) ValidateWithPath(path string) error {
	return nil
}
