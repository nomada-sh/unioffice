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

type CT_RecipientData struct {
	// Record Is Included in Mail Merge
	Active *CT_OnOff
	// Index of Column Containing Unique Values for Record
	Column *CT_DecimalNumber
	// Unique Value for Record
	UniqueTag *CT_Base64Binary
}

func NewCT_RecipientData() *CT_RecipientData {
	ret := &CT_RecipientData{}
	ret.Column = NewCT_DecimalNumber()
	ret.UniqueTag = NewCT_Base64Binary()
	return ret
}
func (m *CT_RecipientData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.Active != nil {
		seactive := xml.StartElement{Name: xml.Name{Local: "w:active"}}
		e.EncodeElement(m.Active, seactive)
	}
	secolumn := xml.StartElement{Name: xml.Name{Local: "w:column"}}
	e.EncodeElement(m.Column, secolumn)
	seuniqueTag := xml.StartElement{Name: xml.Name{Local: "w:uniqueTag"}}
	e.EncodeElement(m.UniqueTag, seuniqueTag)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_RecipientData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Column = NewCT_DecimalNumber()
	m.UniqueTag = NewCT_Base64Binary()
lCT_RecipientData:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "active":
				m.Active = NewCT_OnOff()
				if err := d.DecodeElement(m.Active, &el); err != nil {
					return err
				}
			case "column":
				if err := d.DecodeElement(m.Column, &el); err != nil {
					return err
				}
			case "uniqueTag":
				if err := d.DecodeElement(m.UniqueTag, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_RecipientData
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_RecipientData) Validate() error {
	return m.ValidateWithPath("CT_RecipientData")
}
func (m *CT_RecipientData) ValidateWithPath(path string) error {
	if m.Active != nil {
		if err := m.Active.ValidateWithPath(path + "/Active"); err != nil {
			return err
		}
	}
	if err := m.Column.ValidateWithPath(path + "/Column"); err != nil {
		return err
	}
	if err := m.UniqueTag.ValidateWithPath(path + "/UniqueTag"); err != nil {
		return err
	}
	return nil
}
