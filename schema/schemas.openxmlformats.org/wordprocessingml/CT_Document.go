// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type CT_Document struct {
	ConformanceAttr sharedTypes.ST_ConformanceClass
	// Document Background
	Background *CT_Background
	Body       *CT_Body
}

func NewCT_Document() *CT_Document {
	ret := &CT_Document{}
	ret.ConformanceAttr = sharedTypes.ST_ConformanceClass(1)
	return ret
}
func (m *CT_Document) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	attr, err := m.ConformanceAttr.MarshalXMLAttr(xml.Name{Local: "w:conformance"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	if m.Background != nil {
		sebackground := xml.StartElement{Name: xml.Name{Local: "w:background"}}
		e.EncodeElement(m.Background, sebackground)
	}
	if m.Body != nil {
		sebody := xml.StartElement{Name: xml.Name{Local: "w:body"}}
		e.EncodeElement(m.Body, sebody)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Document) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ConformanceAttr = sharedTypes.ST_ConformanceClass(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "conformance" {
			m.ConformanceAttr.UnmarshalXMLAttr(attr)
		}
	}
lCT_Document:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "background":
				m.Background = NewCT_Background()
				if err := d.DecodeElement(m.Background, &el); err != nil {
					return err
				}
			case "body":
				m.Body = NewCT_Body()
				if err := d.DecodeElement(m.Body, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Document
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Document) Validate() error {
	return m.ValidateWithPath("CT_Document")
}
func (m *CT_Document) ValidateWithPath(path string) error {
	if m.ConformanceAttr == sharedTypes.ST_ConformanceClassUnset {
		return fmt.Errorf("%s/ConformanceAttr is a mandatory field", path)
	}
	if err := m.ConformanceAttr.ValidateWithPath(path + "/ConformanceAttr"); err != nil {
		return err
	}
	if m.Background != nil {
		if err := m.Background.ValidateWithPath(path + "/Background"); err != nil {
			return err
		}
	}
	if m.Body != nil {
		if err := m.Body.ValidateWithPath(path + "/Body"); err != nil {
			return err
		}
	}
	return nil
}
