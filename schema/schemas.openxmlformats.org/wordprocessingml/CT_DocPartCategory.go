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

type CT_DocPartCategory struct {
	// Category Associated With Entry
	Name *CT_String
	// Gallery Associated With Entry
	Gallery *CT_DocPartGallery
}

func NewCT_DocPartCategory() *CT_DocPartCategory {
	ret := &CT_DocPartCategory{}
	ret.Name = NewCT_String()
	ret.Gallery = NewCT_DocPartGallery()
	return ret
}
func (m *CT_DocPartCategory) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	sename := xml.StartElement{Name: xml.Name{Local: "w:name"}}
	e.EncodeElement(m.Name, sename)
	segallery := xml.StartElement{Name: xml.Name{Local: "w:gallery"}}
	e.EncodeElement(m.Gallery, segallery)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_DocPartCategory) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Name = NewCT_String()
	m.Gallery = NewCT_DocPartGallery()
lCT_DocPartCategory:
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
			case "gallery":
				if err := d.DecodeElement(m.Gallery, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DocPartCategory
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_DocPartCategory) Validate() error {
	return m.ValidateWithPath("CT_DocPartCategory")
}
func (m *CT_DocPartCategory) ValidateWithPath(path string) error {
	if err := m.Name.ValidateWithPath(path + "/Name"); err != nil {
		return err
	}
	if err := m.Gallery.ValidateWithPath(path + "/Gallery"); err != nil {
		return err
	}
	return nil
}
