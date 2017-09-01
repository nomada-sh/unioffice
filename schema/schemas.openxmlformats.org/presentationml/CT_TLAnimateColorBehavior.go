// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_TLAnimateColorBehavior struct {
	// Color Space
	ClrSpcAttr ST_TLAnimateColorSpace
	// Direction
	DirAttr ST_TLAnimateColorDirection
	CBhvr   *CT_TLCommonBehaviorData
	// By
	By *CT_TLByAnimateColorTransform
	// From
	From *drawingml.CT_Color
	// To
	To *drawingml.CT_Color
}

func NewCT_TLAnimateColorBehavior() *CT_TLAnimateColorBehavior {
	ret := &CT_TLAnimateColorBehavior{}
	ret.CBhvr = NewCT_TLCommonBehaviorData()
	return ret
}
func (m *CT_TLAnimateColorBehavior) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.ClrSpcAttr != ST_TLAnimateColorSpaceUnset {
		attr, err := m.ClrSpcAttr.MarshalXMLAttr(xml.Name{Local: "clrSpc"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.DirAttr != ST_TLAnimateColorDirectionUnset {
		attr, err := m.DirAttr.MarshalXMLAttr(xml.Name{Local: "dir"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	secBhvr := xml.StartElement{Name: xml.Name{Local: "p:cBhvr"}}
	e.EncodeElement(m.CBhvr, secBhvr)
	if m.By != nil {
		seby := xml.StartElement{Name: xml.Name{Local: "p:by"}}
		e.EncodeElement(m.By, seby)
	}
	if m.From != nil {
		sefrom := xml.StartElement{Name: xml.Name{Local: "p:from"}}
		e.EncodeElement(m.From, sefrom)
	}
	if m.To != nil {
		seto := xml.StartElement{Name: xml.Name{Local: "p:to"}}
		e.EncodeElement(m.To, seto)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TLAnimateColorBehavior) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CBhvr = NewCT_TLCommonBehaviorData()
	for _, attr := range start.Attr {
		if attr.Name.Local == "clrSpc" {
			m.ClrSpcAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "dir" {
			m.DirAttr.UnmarshalXMLAttr(attr)
		}
	}
lCT_TLAnimateColorBehavior:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cBhvr":
				if err := d.DecodeElement(m.CBhvr, &el); err != nil {
					return err
				}
			case "by":
				m.By = NewCT_TLByAnimateColorTransform()
				if err := d.DecodeElement(m.By, &el); err != nil {
					return err
				}
			case "from":
				m.From = drawingml.NewCT_Color()
				if err := d.DecodeElement(m.From, &el); err != nil {
					return err
				}
			case "to":
				m.To = drawingml.NewCT_Color()
				if err := d.DecodeElement(m.To, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLAnimateColorBehavior
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_TLAnimateColorBehavior) Validate() error {
	return m.ValidateWithPath("CT_TLAnimateColorBehavior")
}
func (m *CT_TLAnimateColorBehavior) ValidateWithPath(path string) error {
	if err := m.ClrSpcAttr.ValidateWithPath(path + "/ClrSpcAttr"); err != nil {
		return err
	}
	if err := m.DirAttr.ValidateWithPath(path + "/DirAttr"); err != nil {
		return err
	}
	if err := m.CBhvr.ValidateWithPath(path + "/CBhvr"); err != nil {
		return err
	}
	if m.By != nil {
		if err := m.By.ValidateWithPath(path + "/By"); err != nil {
			return err
		}
	}
	if m.From != nil {
		if err := m.From.ValidateWithPath(path + "/From"); err != nil {
			return err
		}
	}
	if m.To != nil {
		if err := m.To.ValidateWithPath(path + "/To"); err != nil {
			return err
		}
	}
	return nil
}
