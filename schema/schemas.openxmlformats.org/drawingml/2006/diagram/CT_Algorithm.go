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

type CT_Algorithm struct {
	TypeAttr ST_AlgorithmType
	RevAttr  *uint32
	Param    []*CT_Parameter
	ExtLst   *drawingml.CT_OfficeArtExtensionList
}

func NewCT_Algorithm() *CT_Algorithm {
	ret := &CT_Algorithm{}
	ret.TypeAttr = ST_AlgorithmType(1)
	return ret
}
func (m *CT_Algorithm) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	if m.RevAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rev"},
			Value: fmt.Sprintf("%v", *m.RevAttr)})
	}
	e.EncodeToken(start)
	if m.Param != nil {
		separam := xml.StartElement{Name: xml.Name{Local: "param"}}
		e.EncodeElement(m.Param, separam)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Algorithm) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TypeAttr = ST_AlgorithmType(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "rev" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.RevAttr = &pt
		}
	}
lCT_Algorithm:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "param":
				tmp := NewCT_Parameter()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Param = append(m.Param, tmp)
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
			break lCT_Algorithm
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Algorithm) Validate() error {
	return m.ValidateWithPath("CT_Algorithm")
}
func (m *CT_Algorithm) ValidateWithPath(path string) error {
	if m.TypeAttr == ST_AlgorithmTypeUnset {
		return fmt.Errorf("%s/TypeAttr is a mandatory field", path)
	}
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	for i, v := range m.Param {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Param[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
