// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_Transform2D struct {
	RotAttr   *int32
	FlipHAttr *bool
	FlipVAttr *bool
	Off       *CT_Point2D
	Ext       *CT_PositiveSize2D
}

func NewCT_Transform2D() *CT_Transform2D {
	ret := &CT_Transform2D{}
	return ret
}
func (m *CT_Transform2D) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.RotAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rot"},
			Value: fmt.Sprintf("%v", *m.RotAttr)})
	}
	if m.FlipHAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "flipH"},
			Value: fmt.Sprintf("%v", *m.FlipHAttr)})
	}
	if m.FlipVAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "flipV"},
			Value: fmt.Sprintf("%v", *m.FlipVAttr)})
	}
	e.EncodeToken(start)
	if m.Off != nil {
		seoff := xml.StartElement{Name: xml.Name{Local: "a:off"}}
		e.EncodeElement(m.Off, seoff)
	}
	if m.Ext != nil {
		seext := xml.StartElement{Name: xml.Name{Local: "a:ext"}}
		e.EncodeElement(m.Ext, seext)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Transform2D) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "rot" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.RotAttr = &pt
		}
		if attr.Name.Local == "flipH" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FlipHAttr = &parsed
		}
		if attr.Name.Local == "flipV" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FlipVAttr = &parsed
		}
	}
lCT_Transform2D:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "off":
				m.Off = NewCT_Point2D()
				if err := d.DecodeElement(m.Off, &el); err != nil {
					return err
				}
			case "ext":
				m.Ext = NewCT_PositiveSize2D()
				if err := d.DecodeElement(m.Ext, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Transform2D
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Transform2D) Validate() error {
	return m.ValidateWithPath("CT_Transform2D")
}
func (m *CT_Transform2D) ValidateWithPath(path string) error {
	if m.Off != nil {
		if err := m.Off.ValidateWithPath(path + "/Off"); err != nil {
			return err
		}
	}
	if m.Ext != nil {
		if err := m.Ext.ValidateWithPath(path + "/Ext"); err != nil {
			return err
		}
	}
	return nil
}
