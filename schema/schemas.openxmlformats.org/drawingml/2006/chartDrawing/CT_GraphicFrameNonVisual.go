// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chartDrawing

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_GraphicFrameNonVisual struct {
	CNvPr             *drawingml.CT_NonVisualDrawingProps
	CNvGraphicFramePr *drawingml.CT_NonVisualGraphicFrameProperties
}

func NewCT_GraphicFrameNonVisual() *CT_GraphicFrameNonVisual {
	ret := &CT_GraphicFrameNonVisual{}
	ret.CNvPr = drawingml.NewCT_NonVisualDrawingProps()
	ret.CNvGraphicFramePr = drawingml.NewCT_NonVisualGraphicFrameProperties()
	return ret
}
func (m *CT_GraphicFrameNonVisual) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	secNvPr := xml.StartElement{Name: xml.Name{Local: "cNvPr"}}
	e.EncodeElement(m.CNvPr, secNvPr)
	secNvGraphicFramePr := xml.StartElement{Name: xml.Name{Local: "cNvGraphicFramePr"}}
	e.EncodeElement(m.CNvGraphicFramePr, secNvGraphicFramePr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_GraphicFrameNonVisual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CNvPr = drawingml.NewCT_NonVisualDrawingProps()
	m.CNvGraphicFramePr = drawingml.NewCT_NonVisualGraphicFrameProperties()
lCT_GraphicFrameNonVisual:
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
			case "cNvGraphicFramePr":
				if err := d.DecodeElement(m.CNvGraphicFramePr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GraphicFrameNonVisual
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_GraphicFrameNonVisual) Validate() error {
	return m.ValidateWithPath("CT_GraphicFrameNonVisual")
}
func (m *CT_GraphicFrameNonVisual) ValidateWithPath(path string) error {
	if err := m.CNvPr.ValidateWithPath(path + "/CNvPr"); err != nil {
		return err
	}
	if err := m.CNvGraphicFramePr.ValidateWithPath(path + "/CNvGraphicFramePr"); err != nil {
		return err
	}
	return nil
}
