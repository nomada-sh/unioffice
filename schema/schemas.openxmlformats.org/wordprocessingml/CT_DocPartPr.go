// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"log"
)

type CT_DocPartPr struct {
	// Entry Name
	Name *CT_DocPartName
	// Associated Paragraph Style Name
	Style *CT_String
	// Entry Categorization
	Category *CT_DocPartCategory
	// Entry Types
	Types *CT_DocPartTypes
	// Entry Insertion Behaviors
	Behaviors *CT_DocPartBehaviors
	// Description for Entry
	Description *CT_String
	// Entry ID
	Guid *CT_Guid
}

func NewCT_DocPartPr() *CT_DocPartPr {
	ret := &CT_DocPartPr{}
	ret.Name = NewCT_DocPartName()
	return ret
}
func (m *CT_DocPartPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	sename := xml.StartElement{Name: xml.Name{Local: "w:name"}}
	e.EncodeElement(m.Name, sename)
	if m.Style != nil {
		sestyle := xml.StartElement{Name: xml.Name{Local: "w:style"}}
		e.EncodeElement(m.Style, sestyle)
	}
	if m.Category != nil {
		secategory := xml.StartElement{Name: xml.Name{Local: "w:category"}}
		e.EncodeElement(m.Category, secategory)
	}
	if m.Types != nil {
		setypes := xml.StartElement{Name: xml.Name{Local: "w:types"}}
		e.EncodeElement(m.Types, setypes)
	}
	if m.Behaviors != nil {
		sebehaviors := xml.StartElement{Name: xml.Name{Local: "w:behaviors"}}
		e.EncodeElement(m.Behaviors, sebehaviors)
	}
	if m.Description != nil {
		sedescription := xml.StartElement{Name: xml.Name{Local: "w:description"}}
		e.EncodeElement(m.Description, sedescription)
	}
	if m.Guid != nil {
		seguid := xml.StartElement{Name: xml.Name{Local: "w:guid"}}
		e.EncodeElement(m.Guid, seguid)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_DocPartPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Name = NewCT_DocPartName()
lCT_DocPartPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "name":
				if err := d.DecodeElement(m.Name, &el); err != nil {
					return err
				}
			case "style":
				m.Style = NewCT_String()
				if err := d.DecodeElement(m.Style, &el); err != nil {
					return err
				}
			case "category":
				m.Category = NewCT_DocPartCategory()
				if err := d.DecodeElement(m.Category, &el); err != nil {
					return err
				}
			case "types":
				m.Types = NewCT_DocPartTypes()
				if err := d.DecodeElement(m.Types, &el); err != nil {
					return err
				}
			case "behaviors":
				m.Behaviors = NewCT_DocPartBehaviors()
				if err := d.DecodeElement(m.Behaviors, &el); err != nil {
					return err
				}
			case "description":
				m.Description = NewCT_String()
				if err := d.DecodeElement(m.Description, &el); err != nil {
					return err
				}
			case "guid":
				m.Guid = NewCT_Guid()
				if err := d.DecodeElement(m.Guid, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DocPartPr
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_DocPartPr) Validate() error {
	return m.ValidateWithPath("CT_DocPartPr")
}
func (m *CT_DocPartPr) ValidateWithPath(path string) error {
	if err := m.Name.ValidateWithPath(path + "/Name"); err != nil {
		return err
	}
	if m.Style != nil {
		if err := m.Style.ValidateWithPath(path + "/Style"); err != nil {
			return err
		}
	}
	if m.Category != nil {
		if err := m.Category.ValidateWithPath(path + "/Category"); err != nil {
			return err
		}
	}
	if m.Types != nil {
		if err := m.Types.ValidateWithPath(path + "/Types"); err != nil {
			return err
		}
	}
	if m.Behaviors != nil {
		if err := m.Behaviors.ValidateWithPath(path + "/Behaviors"); err != nil {
			return err
		}
	}
	if m.Description != nil {
		if err := m.Description.ValidateWithPath(path + "/Description"); err != nil {
			return err
		}
	}
	if m.Guid != nil {
		if err := m.Guid.ValidateWithPath(path + "/Guid"); err != nil {
			return err
		}
	}
	return nil
}
