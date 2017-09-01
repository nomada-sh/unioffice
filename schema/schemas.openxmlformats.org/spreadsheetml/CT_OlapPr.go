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

type CT_OlapPr struct {
	// Local Cube
	LocalAttr *bool
	// Local Cube Connection
	LocalConnectionAttr *string
	// Local Refresh
	LocalRefreshAttr *bool
	// Send Locale to OLAP
	SendLocaleAttr *bool
	// Drill Through Count
	RowDrillCountAttr *uint32
	// OLAP Fill Formatting
	ServerFillAttr *bool
	// OLAP Number Format
	ServerNumberFormatAttr *bool
	// OLAP Server Font
	ServerFontAttr *bool
	// OLAP Font Formatting
	ServerFontColorAttr *bool
}

func NewCT_OlapPr() *CT_OlapPr {
	ret := &CT_OlapPr{}
	return ret
}
func (m *CT_OlapPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.LocalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "local"},
			Value: fmt.Sprintf("%v", *m.LocalAttr)})
	}
	if m.LocalConnectionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "localConnection"},
			Value: fmt.Sprintf("%v", *m.LocalConnectionAttr)})
	}
	if m.LocalRefreshAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "localRefresh"},
			Value: fmt.Sprintf("%v", *m.LocalRefreshAttr)})
	}
	if m.SendLocaleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sendLocale"},
			Value: fmt.Sprintf("%v", *m.SendLocaleAttr)})
	}
	if m.RowDrillCountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rowDrillCount"},
			Value: fmt.Sprintf("%v", *m.RowDrillCountAttr)})
	}
	if m.ServerFillAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "serverFill"},
			Value: fmt.Sprintf("%v", *m.ServerFillAttr)})
	}
	if m.ServerNumberFormatAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "serverNumberFormat"},
			Value: fmt.Sprintf("%v", *m.ServerNumberFormatAttr)})
	}
	if m.ServerFontAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "serverFont"},
			Value: fmt.Sprintf("%v", *m.ServerFontAttr)})
	}
	if m.ServerFontColorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "serverFontColor"},
			Value: fmt.Sprintf("%v", *m.ServerFontColorAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_OlapPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "local" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LocalAttr = &parsed
		}
		if attr.Name.Local == "localConnection" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LocalConnectionAttr = &parsed
		}
		if attr.Name.Local == "localRefresh" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LocalRefreshAttr = &parsed
		}
		if attr.Name.Local == "sendLocale" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SendLocaleAttr = &parsed
		}
		if attr.Name.Local == "rowDrillCount" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.RowDrillCountAttr = &pt
		}
		if attr.Name.Local == "serverFill" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ServerFillAttr = &parsed
		}
		if attr.Name.Local == "serverNumberFormat" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ServerNumberFormatAttr = &parsed
		}
		if attr.Name.Local == "serverFont" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ServerFontAttr = &parsed
		}
		if attr.Name.Local == "serverFontColor" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ServerFontColorAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_OlapPr: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_OlapPr) Validate() error {
	return m.ValidateWithPath("CT_OlapPr")
}
func (m *CT_OlapPr) ValidateWithPath(path string) error {
	return nil
}
