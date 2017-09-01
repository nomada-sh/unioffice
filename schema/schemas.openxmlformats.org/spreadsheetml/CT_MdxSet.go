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

type CT_MdxSet struct {
	// Set Definition Index
	NsAttr uint32
	// Sort By Member Index Count
	CAttr *uint32
	// Set Sort Order
	OAttr ST_MdxSetOrder
	// Member Unique Name Index
	N []*CT_MetadataStringIndex
}

func NewCT_MdxSet() *CT_MdxSet {
	ret := &CT_MdxSet{}
	return ret
}
func (m *CT_MdxSet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ns"},
		Value: fmt.Sprintf("%v", m.NsAttr)})
	if m.CAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "c"},
			Value: fmt.Sprintf("%v", *m.CAttr)})
	}
	if m.OAttr != ST_MdxSetOrderUnset {
		attr, err := m.OAttr.MarshalXMLAttr(xml.Name{Local: "o"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.N != nil {
		sen := xml.StartElement{Name: xml.Name{Local: "x:n"}}
		e.EncodeElement(m.N, sen)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_MdxSet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "ns" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.NsAttr = uint32(parsed)
		}
		if attr.Name.Local == "c" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CAttr = &pt
		}
		if attr.Name.Local == "o" {
			m.OAttr.UnmarshalXMLAttr(attr)
		}
	}
lCT_MdxSet:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "n":
				tmp := NewCT_MetadataStringIndex()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.N = append(m.N, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_MdxSet
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_MdxSet) Validate() error {
	return m.ValidateWithPath("CT_MdxSet")
}
func (m *CT_MdxSet) ValidateWithPath(path string) error {
	if err := m.OAttr.ValidateWithPath(path + "/OAttr"); err != nil {
		return err
	}
	for i, v := range m.N {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/N[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
