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

type CT_NonVisualGroupDrawingShapeProps struct {
	GrpSpLocks *CT_GroupLocking
	ExtLst     *CT_OfficeArtExtensionList
}

func NewCT_NonVisualGroupDrawingShapeProps() *CT_NonVisualGroupDrawingShapeProps {
	ret := &CT_NonVisualGroupDrawingShapeProps{}
	return ret
}
func (m *CT_NonVisualGroupDrawingShapeProps) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.GrpSpLocks != nil {
		segrpSpLocks := xml.StartElement{Name: xml.Name{Local: "a:grpSpLocks"}}
		e.EncodeElement(m.GrpSpLocks, segrpSpLocks)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_NonVisualGroupDrawingShapeProps) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_NonVisualGroupDrawingShapeProps:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "grpSpLocks":
				m.GrpSpLocks = NewCT_GroupLocking()
				if err := d.DecodeElement(m.GrpSpLocks, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_OfficeArtExtensionList()
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
			break lCT_NonVisualGroupDrawingShapeProps
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_NonVisualGroupDrawingShapeProps) Validate() error {
	return m.ValidateWithPath("CT_NonVisualGroupDrawingShapeProps")
}
func (m *CT_NonVisualGroupDrawingShapeProps) ValidateWithPath(path string) error {
	if m.GrpSpLocks != nil {
		if err := m.GrpSpLocks.ValidateWithPath(path + "/GrpSpLocks"); err != nil {
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
