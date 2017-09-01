// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"log"
)

type CT_PathShadeProperties struct {
	PathAttr   ST_PathShadeType
	FillToRect *CT_RelativeRect
}

func NewCT_PathShadeProperties() *CT_PathShadeProperties {
	ret := &CT_PathShadeProperties{}
	return ret
}
func (m *CT_PathShadeProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.PathAttr != ST_PathShadeTypeUnset {
		attr, err := m.PathAttr.MarshalXMLAttr(xml.Name{Local: "path"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.FillToRect != nil {
		sefillToRect := xml.StartElement{Name: xml.Name{Local: "a:fillToRect"}}
		e.EncodeElement(m.FillToRect, sefillToRect)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_PathShadeProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "path" {
			m.PathAttr.UnmarshalXMLAttr(attr)
		}
	}
lCT_PathShadeProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "fillToRect":
				m.FillToRect = NewCT_RelativeRect()
				if err := d.DecodeElement(m.FillToRect, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PathShadeProperties
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_PathShadeProperties) Validate() error {
	return m.ValidateWithPath("CT_PathShadeProperties")
}
func (m *CT_PathShadeProperties) ValidateWithPath(path string) error {
	if err := m.PathAttr.ValidateWithPath(path + "/PathAttr"); err != nil {
		return err
	}
	if m.FillToRect != nil {
		if err := m.FillToRect.ValidateWithPath(path + "/FillToRect"); err != nil {
			return err
		}
	}
	return nil
}
