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

type CT_MemberProperty struct {
	// OLAP Member Property Unique Name
	NameAttr *string
	// Show Cell
	ShowCellAttr *bool
	// Show Tooltip
	ShowTipAttr *bool
	// Show As Caption
	ShowAsCaptionAttr *bool
	// Name Length
	NameLenAttr *uint32
	// Property Name Character Index
	PPosAttr *uint32
	// Property Name Length
	PLenAttr *uint32
	// Level Index
	LevelAttr *uint32
	// Field Index
	FieldAttr uint32
}

func NewCT_MemberProperty() *CT_MemberProperty {
	ret := &CT_MemberProperty{}
	return ret
}
func (m *CT_MemberProperty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	if m.ShowCellAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showCell"},
			Value: fmt.Sprintf("%v", *m.ShowCellAttr)})
	}
	if m.ShowTipAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showTip"},
			Value: fmt.Sprintf("%v", *m.ShowTipAttr)})
	}
	if m.ShowAsCaptionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showAsCaption"},
			Value: fmt.Sprintf("%v", *m.ShowAsCaptionAttr)})
	}
	if m.NameLenAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "nameLen"},
			Value: fmt.Sprintf("%v", *m.NameLenAttr)})
	}
	if m.PPosAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "pPos"},
			Value: fmt.Sprintf("%v", *m.PPosAttr)})
	}
	if m.PLenAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "pLen"},
			Value: fmt.Sprintf("%v", *m.PLenAttr)})
	}
	if m.LevelAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "level"},
			Value: fmt.Sprintf("%v", *m.LevelAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "field"},
		Value: fmt.Sprintf("%v", m.FieldAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_MemberProperty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
		}
		if attr.Name.Local == "showCell" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowCellAttr = &parsed
		}
		if attr.Name.Local == "showTip" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowTipAttr = &parsed
		}
		if attr.Name.Local == "showAsCaption" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowAsCaptionAttr = &parsed
		}
		if attr.Name.Local == "nameLen" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.NameLenAttr = &pt
		}
		if attr.Name.Local == "pPos" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.PPosAttr = &pt
		}
		if attr.Name.Local == "pLen" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.PLenAttr = &pt
		}
		if attr.Name.Local == "level" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.LevelAttr = &pt
		}
		if attr.Name.Local == "field" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.FieldAttr = uint32(parsed)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_MemberProperty: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_MemberProperty) Validate() error {
	return m.ValidateWithPath("CT_MemberProperty")
}
func (m *CT_MemberProperty) ValidateWithPath(path string) error {
	return nil
}
