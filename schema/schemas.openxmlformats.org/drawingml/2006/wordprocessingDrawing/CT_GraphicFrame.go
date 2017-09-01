// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingDrawing

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_GraphicFrame struct {
	CNvPr   *drawingml.CT_NonVisualDrawingProps
	CNvFrPr *drawingml.CT_NonVisualGraphicFrameProperties
	Xfrm    *drawingml.CT_Transform2D
	Graphic *drawingml.Graphic
	ExtLst  *drawingml.CT_OfficeArtExtensionList
}

func NewCT_GraphicFrame() *CT_GraphicFrame {
	ret := &CT_GraphicFrame{}
	ret.CNvPr = drawingml.NewCT_NonVisualDrawingProps()
	ret.CNvFrPr = drawingml.NewCT_NonVisualGraphicFrameProperties()
	ret.Xfrm = drawingml.NewCT_Transform2D()
	ret.Graphic = drawingml.NewGraphic()
	return ret
}
func (m *CT_GraphicFrame) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	secNvPr := xml.StartElement{Name: xml.Name{Local: "wp:cNvPr"}}
	e.EncodeElement(m.CNvPr, secNvPr)
	secNvFrPr := xml.StartElement{Name: xml.Name{Local: "wp:cNvFrPr"}}
	e.EncodeElement(m.CNvFrPr, secNvFrPr)
	sexfrm := xml.StartElement{Name: xml.Name{Local: "wp:xfrm"}}
	e.EncodeElement(m.Xfrm, sexfrm)
	segraphic := xml.StartElement{Name: xml.Name{Local: "a:graphic"}}
	e.EncodeElement(m.Graphic, segraphic)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "wp:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_GraphicFrame) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CNvPr = drawingml.NewCT_NonVisualDrawingProps()
	m.CNvFrPr = drawingml.NewCT_NonVisualGraphicFrameProperties()
	m.Xfrm = drawingml.NewCT_Transform2D()
	m.Graphic = drawingml.NewGraphic()
lCT_GraphicFrame:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cNvPr":
				if err := d.DecodeElement(m.CNvPr, &el); err != nil {
					return err
				}
			case "cNvFrPr":
				if err := d.DecodeElement(m.CNvFrPr, &el); err != nil {
					return err
				}
			case "xfrm":
				if err := d.DecodeElement(m.Xfrm, &el); err != nil {
					return err
				}
			case "graphic":
				if err := d.DecodeElement(m.Graphic, &el); err != nil {
					return err
				}
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
			break lCT_GraphicFrame
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_GraphicFrame) Validate() error {
	return m.ValidateWithPath("CT_GraphicFrame")
}
func (m *CT_GraphicFrame) ValidateWithPath(path string) error {
	if err := m.CNvPr.ValidateWithPath(path + "/CNvPr"); err != nil {
		return err
	}
	if err := m.CNvFrPr.ValidateWithPath(path + "/CNvFrPr"); err != nil {
		return err
	}
	if err := m.Xfrm.ValidateWithPath(path + "/Xfrm"); err != nil {
		return err
	}
	if err := m.Graphic.ValidateWithPath(path + "/Graphic"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
