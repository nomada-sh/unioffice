// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_NumericRule struct {
	ValAttr     *float64
	FactAttr    *float64
	MaxAttr     *float64
	ExtLst      *drawingml.CT_OfficeArtExtensionList
	TypeAttr    ST_ConstraintType
	ForAttr     ST_ConstraintRelationship
	ForNameAttr *string
	PtTypeAttr  ST_ElementType
}

func NewCT_NumericRule() *CT_NumericRule {
	ret := &CT_NumericRule{}
	return ret
}
func (m *CT_NumericRule) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.ValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
			Value: fmt.Sprintf("%v", *m.ValAttr)})
	}
	if m.FactAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fact"},
			Value: fmt.Sprintf("%v", *m.FactAttr)})
	}
	if m.MaxAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "max"},
			Value: fmt.Sprintf("%v", *m.MaxAttr)})
	}
	if m.TypeAttr != ST_ConstraintTypeUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ForAttr != ST_ConstraintRelationshipUnset {
		attr, err := m.ForAttr.MarshalXMLAttr(xml.Name{Local: "for"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ForNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "forName"},
			Value: fmt.Sprintf("%v", *m.ForNameAttr)})
	}
	if m.PtTypeAttr != ST_ElementTypeUnset {
		attr, err := m.PtTypeAttr.MarshalXMLAttr(xml.Name{Local: "ptType"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_NumericRule) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.ValAttr = &parsed
		}
		if attr.Name.Local == "fact" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.FactAttr = &parsed
		}
		if attr.Name.Local == "max" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.MaxAttr = &parsed
		}
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "for" {
			m.ForAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "forName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ForNameAttr = &parsed
		}
		if attr.Name.Local == "ptType" {
			m.PtTypeAttr.UnmarshalXMLAttr(attr)
		}
	}
lCT_NumericRule:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "extLst":
				m.ExtLst = drawingml.NewCT_OfficeArtExtensionList()
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
			break lCT_NumericRule
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_NumericRule) Validate() error {
	return m.ValidateWithPath("CT_NumericRule")
}
func (m *CT_NumericRule) ValidateWithPath(path string) error {
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if err := m.ForAttr.ValidateWithPath(path + "/ForAttr"); err != nil {
		return err
	}
	if err := m.PtTypeAttr.ValidateWithPath(path + "/PtTypeAttr"); err != nil {
		return err
	}
	return nil
}
