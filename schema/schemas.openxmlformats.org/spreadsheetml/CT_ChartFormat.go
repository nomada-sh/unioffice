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

type CT_ChartFormat struct {
	// Chart Index
	ChartAttr uint32
	// Pivot Format Id
	FormatAttr uint32
	// Series Format
	SeriesAttr *bool
	// Pivot Table Location Rule
	PivotArea *CT_PivotArea
}

func NewCT_ChartFormat() *CT_ChartFormat {
	ret := &CT_ChartFormat{}
	ret.PivotArea = NewCT_PivotArea()
	return ret
}
func (m *CT_ChartFormat) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "chart"},
		Value: fmt.Sprintf("%v", m.ChartAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "format"},
		Value: fmt.Sprintf("%v", m.FormatAttr)})
	if m.SeriesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "series"},
			Value: fmt.Sprintf("%v", *m.SeriesAttr)})
	}
	e.EncodeToken(start)
	sepivotArea := xml.StartElement{Name: xml.Name{Local: "x:pivotArea"}}
	e.EncodeElement(m.PivotArea, sepivotArea)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_ChartFormat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.PivotArea = NewCT_PivotArea()
	for _, attr := range start.Attr {
		if attr.Name.Local == "chart" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.ChartAttr = uint32(parsed)
		}
		if attr.Name.Local == "format" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.FormatAttr = uint32(parsed)
		}
		if attr.Name.Local == "series" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SeriesAttr = &parsed
		}
	}
lCT_ChartFormat:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "pivotArea":
				if err := d.DecodeElement(m.PivotArea, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ChartFormat
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_ChartFormat) Validate() error {
	return m.ValidateWithPath("CT_ChartFormat")
}
func (m *CT_ChartFormat) ValidateWithPath(path string) error {
	if err := m.PivotArea.ValidateWithPath(path + "/PivotArea"); err != nil {
		return err
	}
	return nil
}
