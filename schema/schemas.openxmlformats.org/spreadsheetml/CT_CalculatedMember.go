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

type CT_CalculatedMember struct {
	// Calculated Member Name
	NameAttr string
	// Calculated Member MDX Formula
	MdxAttr string
	// OLAP Calculated Member Name
	MemberNameAttr *string
	// Hierarchy Name
	HierarchyAttr *string
	// Parent Name
	ParentAttr *string
	// Calculated Members Solve Order
	SolveOrderAttr *int32
	// Set
	SetAttr *bool
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_CalculatedMember() *CT_CalculatedMember {
	ret := &CT_CalculatedMember{}
	return ret
}
func (m *CT_CalculatedMember) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "mdx"},
		Value: fmt.Sprintf("%v", m.MdxAttr)})
	if m.MemberNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "memberName"},
			Value: fmt.Sprintf("%v", *m.MemberNameAttr)})
	}
	if m.HierarchyAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hierarchy"},
			Value: fmt.Sprintf("%v", *m.HierarchyAttr)})
	}
	if m.ParentAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "parent"},
			Value: fmt.Sprintf("%v", *m.ParentAttr)})
	}
	if m.SolveOrderAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "solveOrder"},
			Value: fmt.Sprintf("%v", *m.SolveOrderAttr)})
	}
	if m.SetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "set"},
			Value: fmt.Sprintf("%v", *m.SetAttr)})
	}
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_CalculatedMember) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
		}
		if attr.Name.Local == "mdx" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.MdxAttr = parsed
		}
		if attr.Name.Local == "memberName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.MemberNameAttr = &parsed
		}
		if attr.Name.Local == "hierarchy" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.HierarchyAttr = &parsed
		}
		if attr.Name.Local == "parent" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ParentAttr = &parsed
		}
		if attr.Name.Local == "solveOrder" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.SolveOrderAttr = &pt
		}
		if attr.Name.Local == "set" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SetAttr = &parsed
		}
	}
lCT_CalculatedMember:
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
			break lCT_CalculatedMember
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_CalculatedMember) Validate() error {
	return m.ValidateWithPath("CT_CalculatedMember")
}
func (m *CT_CalculatedMember) ValidateWithPath(path string) error {
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
