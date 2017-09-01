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

type CT_QueryTableField struct {
	// Field Id
	IdAttr uint32
	// Name
	NameAttr *string
	// Data Bound Column
	DataBoundAttr *bool
	// Row Numbers
	RowNumbersAttr *bool
	// Fill This Formula On Refresh
	FillFormulasAttr *bool
	// Clipped Column
	ClippedAttr *bool
	// Table Column Id
	TableColumnIdAttr *uint32
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_QueryTableField() *CT_QueryTableField {
	ret := &CT_QueryTableField{}
	return ret
}
func (m *CT_QueryTableField) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	if m.DataBoundAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dataBound"},
			Value: fmt.Sprintf("%v", *m.DataBoundAttr)})
	}
	if m.RowNumbersAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rowNumbers"},
			Value: fmt.Sprintf("%v", *m.RowNumbersAttr)})
	}
	if m.FillFormulasAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fillFormulas"},
			Value: fmt.Sprintf("%v", *m.FillFormulasAttr)})
	}
	if m.ClippedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "clipped"},
			Value: fmt.Sprintf("%v", *m.ClippedAttr)})
	}
	if m.TableColumnIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "tableColumnId"},
			Value: fmt.Sprintf("%v", *m.TableColumnIdAttr)})
	}
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_QueryTableField) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.IdAttr = uint32(parsed)
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
		}
		if attr.Name.Local == "dataBound" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DataBoundAttr = &parsed
		}
		if attr.Name.Local == "rowNumbers" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RowNumbersAttr = &parsed
		}
		if attr.Name.Local == "fillFormulas" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FillFormulasAttr = &parsed
		}
		if attr.Name.Local == "clipped" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ClippedAttr = &parsed
		}
		if attr.Name.Local == "tableColumnId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.TableColumnIdAttr = &pt
		}
	}
lCT_QueryTableField:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_QueryTableField
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_QueryTableField) Validate() error {
	return m.ValidateWithPath("CT_QueryTableField")
}
func (m *CT_QueryTableField) ValidateWithPath(path string) error {
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
