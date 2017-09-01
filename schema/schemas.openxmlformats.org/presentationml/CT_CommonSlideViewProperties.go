// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_CommonSlideViewProperties struct {
	// Snap Objects to Grid
	SnapToGridAttr *bool
	// Snap Objects to Objects
	SnapToObjectsAttr *bool
	// Show Guides in View
	ShowGuidesAttr *bool
	// Base properties for Slide View
	CViewPr *CT_CommonViewProperties
	// List of Guides
	GuideLst *CT_GuideList
}

func NewCT_CommonSlideViewProperties() *CT_CommonSlideViewProperties {
	ret := &CT_CommonSlideViewProperties{}
	ret.CViewPr = NewCT_CommonViewProperties()
	return ret
}
func (m *CT_CommonSlideViewProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.SnapToGridAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "snapToGrid"},
			Value: fmt.Sprintf("%v", *m.SnapToGridAttr)})
	}
	if m.SnapToObjectsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "snapToObjects"},
			Value: fmt.Sprintf("%v", *m.SnapToObjectsAttr)})
	}
	if m.ShowGuidesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showGuides"},
			Value: fmt.Sprintf("%v", *m.ShowGuidesAttr)})
	}
	e.EncodeToken(start)
	secViewPr := xml.StartElement{Name: xml.Name{Local: "p:cViewPr"}}
	e.EncodeElement(m.CViewPr, secViewPr)
	if m.GuideLst != nil {
		seguideLst := xml.StartElement{Name: xml.Name{Local: "p:guideLst"}}
		e.EncodeElement(m.GuideLst, seguideLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_CommonSlideViewProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CViewPr = NewCT_CommonViewProperties()
	for _, attr := range start.Attr {
		if attr.Name.Local == "snapToGrid" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SnapToGridAttr = &parsed
		}
		if attr.Name.Local == "snapToObjects" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SnapToObjectsAttr = &parsed
		}
		if attr.Name.Local == "showGuides" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowGuidesAttr = &parsed
		}
	}
lCT_CommonSlideViewProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cViewPr":
				if err := d.DecodeElement(m.CViewPr, &el); err != nil {
					return err
				}
			case "guideLst":
				m.GuideLst = NewCT_GuideList()
				if err := d.DecodeElement(m.GuideLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CommonSlideViewProperties
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_CommonSlideViewProperties) Validate() error {
	return m.ValidateWithPath("CT_CommonSlideViewProperties")
}
func (m *CT_CommonSlideViewProperties) ValidateWithPath(path string) error {
	if err := m.CViewPr.ValidateWithPath(path + "/CViewPr"); err != nil {
		return err
	}
	if m.GuideLst != nil {
		if err := m.GuideLst.ValidateWithPath(path + "/GuideLst"); err != nil {
			return err
		}
	}
	return nil
}
