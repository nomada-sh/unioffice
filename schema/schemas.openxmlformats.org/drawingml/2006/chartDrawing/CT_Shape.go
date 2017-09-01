// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chartDrawing

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_Shape struct {
	MacroAttr      *string
	TextlinkAttr   *string
	FLocksTextAttr *bool
	FPublishedAttr *bool
	NvSpPr         *CT_ShapeNonVisual
	SpPr           *drawingml.CT_ShapeProperties
	Style          *drawingml.CT_ShapeStyle
	TxBody         *drawingml.CT_TextBody
}

func NewCT_Shape() *CT_Shape {
	ret := &CT_Shape{}
	ret.NvSpPr = NewCT_ShapeNonVisual()
	ret.SpPr = drawingml.NewCT_ShapeProperties()
	return ret
}
func (m *CT_Shape) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.MacroAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "macro"},
			Value: fmt.Sprintf("%v", *m.MacroAttr)})
	}
	if m.TextlinkAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "textlink"},
			Value: fmt.Sprintf("%v", *m.TextlinkAttr)})
	}
	if m.FLocksTextAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fLocksText"},
			Value: fmt.Sprintf("%v", *m.FLocksTextAttr)})
	}
	if m.FPublishedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fPublished"},
			Value: fmt.Sprintf("%v", *m.FPublishedAttr)})
	}
	e.EncodeToken(start)
	senvSpPr := xml.StartElement{Name: xml.Name{Local: "nvSpPr"}}
	e.EncodeElement(m.NvSpPr, senvSpPr)
	sespPr := xml.StartElement{Name: xml.Name{Local: "spPr"}}
	e.EncodeElement(m.SpPr, sespPr)
	if m.Style != nil {
		sestyle := xml.StartElement{Name: xml.Name{Local: "style"}}
		e.EncodeElement(m.Style, sestyle)
	}
	if m.TxBody != nil {
		setxBody := xml.StartElement{Name: xml.Name{Local: "txBody"}}
		e.EncodeElement(m.TxBody, setxBody)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Shape) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.NvSpPr = NewCT_ShapeNonVisual()
	m.SpPr = drawingml.NewCT_ShapeProperties()
	for _, attr := range start.Attr {
		if attr.Name.Local == "macro" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.MacroAttr = &parsed
		}
		if attr.Name.Local == "textlink" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TextlinkAttr = &parsed
		}
		if attr.Name.Local == "fLocksText" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FLocksTextAttr = &parsed
		}
		if attr.Name.Local == "fPublished" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FPublishedAttr = &parsed
		}
	}
lCT_Shape:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "nvSpPr":
				if err := d.DecodeElement(m.NvSpPr, &el); err != nil {
					return err
				}
			case "spPr":
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case "style":
				m.Style = drawingml.NewCT_ShapeStyle()
				if err := d.DecodeElement(m.Style, &el); err != nil {
					return err
				}
			case "txBody":
				m.TxBody = drawingml.NewCT_TextBody()
				if err := d.DecodeElement(m.TxBody, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Shape
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Shape) Validate() error {
	return m.ValidateWithPath("CT_Shape")
}
func (m *CT_Shape) ValidateWithPath(path string) error {
	if err := m.NvSpPr.ValidateWithPath(path + "/NvSpPr"); err != nil {
		return err
	}
	if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
		return err
	}
	if m.Style != nil {
		if err := m.Style.ValidateWithPath(path + "/Style"); err != nil {
			return err
		}
	}
	if m.TxBody != nil {
		if err := m.TxBody.ValidateWithPath(path + "/TxBody"); err != nil {
			return err
		}
	}
	return nil
}
