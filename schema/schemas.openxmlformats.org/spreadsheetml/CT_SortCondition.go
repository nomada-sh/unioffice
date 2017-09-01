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

type CT_SortCondition struct {
	// Descending
	DescendingAttr *bool
	// Sort By
	SortByAttr ST_SortBy
	// Reference
	RefAttr string
	// Custom List
	CustomListAttr *string
	// Format Id
	DxfIdAttr *uint32
	// Icon Set
	IconSetAttr ST_IconSetType
	// Icon Id
	IconIdAttr *uint32
}

func NewCT_SortCondition() *CT_SortCondition {
	ret := &CT_SortCondition{}
	return ret
}
func (m *CT_SortCondition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.DescendingAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "descending"},
			Value: fmt.Sprintf("%v", *m.DescendingAttr)})
	}
	if m.SortByAttr != ST_SortByUnset {
		attr, err := m.SortByAttr.MarshalXMLAttr(xml.Name{Local: "sortBy"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ref"},
		Value: fmt.Sprintf("%v", m.RefAttr)})
	if m.CustomListAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "customList"},
			Value: fmt.Sprintf("%v", *m.CustomListAttr)})
	}
	if m.DxfIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dxfId"},
			Value: fmt.Sprintf("%v", *m.DxfIdAttr)})
	}
	if m.IconSetAttr != ST_IconSetTypeUnset {
		attr, err := m.IconSetAttr.MarshalXMLAttr(xml.Name{Local: "iconSet"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.IconIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "iconId"},
			Value: fmt.Sprintf("%v", *m.IconIdAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_SortCondition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "descending" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DescendingAttr = &parsed
		}
		if attr.Name.Local == "sortBy" {
			m.SortByAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "ref" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RefAttr = parsed
		}
		if attr.Name.Local == "customList" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CustomListAttr = &parsed
		}
		if attr.Name.Local == "dxfId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DxfIdAttr = &pt
		}
		if attr.Name.Local == "iconSet" {
			m.IconSetAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "iconId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.IconIdAttr = &pt
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_SortCondition: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_SortCondition) Validate() error {
	return m.ValidateWithPath("CT_SortCondition")
}
func (m *CT_SortCondition) ValidateWithPath(path string) error {
	if err := m.SortByAttr.ValidateWithPath(path + "/SortByAttr"); err != nil {
		return err
	}
	if err := m.IconSetAttr.ValidateWithPath(path + "/IconSetAttr"); err != nil {
		return err
	}
	return nil
}
