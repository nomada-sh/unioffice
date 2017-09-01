// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"log"
)

type CT_TcBorders struct {
	// Table Cell Top Border
	Top *CT_Border
	// Table Cell Leading Edge Border
	Start *CT_Border
	// Table Cell Leading Edge Border
	Left *CT_Border
	// Table Cell Bottom Border
	Bottom *CT_Border
	// Table Cell Trailing Edge Border
	End *CT_Border
	// Table Cell Trailing Edge Border
	Right *CT_Border
	// Table Cell Inside Horizontal Edges Border
	InsideH *CT_Border
	// Table Cell Inside Vertical Edges Border
	InsideV *CT_Border
	// Table Cell Top Left to Bottom Right Diagonal Border
	Tl2br *CT_Border
	// Table Cell Top Right to Bottom Left Diagonal Border
	Tr2bl *CT_Border
}

func NewCT_TcBorders() *CT_TcBorders {
	ret := &CT_TcBorders{}
	return ret
}
func (m *CT_TcBorders) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.Top != nil {
		setop := xml.StartElement{Name: xml.Name{Local: "w:top"}}
		e.EncodeElement(m.Top, setop)
	}
	if m.Start != nil {
		sestart := xml.StartElement{Name: xml.Name{Local: "w:start"}}
		e.EncodeElement(m.Start, sestart)
	}
	if m.Left != nil {
		seleft := xml.StartElement{Name: xml.Name{Local: "w:left"}}
		e.EncodeElement(m.Left, seleft)
	}
	if m.Bottom != nil {
		sebottom := xml.StartElement{Name: xml.Name{Local: "w:bottom"}}
		e.EncodeElement(m.Bottom, sebottom)
	}
	if m.End != nil {
		seend := xml.StartElement{Name: xml.Name{Local: "w:end"}}
		e.EncodeElement(m.End, seend)
	}
	if m.Right != nil {
		seright := xml.StartElement{Name: xml.Name{Local: "w:right"}}
		e.EncodeElement(m.Right, seright)
	}
	if m.InsideH != nil {
		seinsideH := xml.StartElement{Name: xml.Name{Local: "w:insideH"}}
		e.EncodeElement(m.InsideH, seinsideH)
	}
	if m.InsideV != nil {
		seinsideV := xml.StartElement{Name: xml.Name{Local: "w:insideV"}}
		e.EncodeElement(m.InsideV, seinsideV)
	}
	if m.Tl2br != nil {
		setl2br := xml.StartElement{Name: xml.Name{Local: "w:tl2br"}}
		e.EncodeElement(m.Tl2br, setl2br)
	}
	if m.Tr2bl != nil {
		setr2bl := xml.StartElement{Name: xml.Name{Local: "w:tr2bl"}}
		e.EncodeElement(m.Tr2bl, setr2bl)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TcBorders) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TcBorders:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "top":
				m.Top = NewCT_Border()
				if err := d.DecodeElement(m.Top, &el); err != nil {
					return err
				}
			case "start":
				m.Start = NewCT_Border()
				if err := d.DecodeElement(m.Start, &el); err != nil {
					return err
				}
			case "left":
				m.Left = NewCT_Border()
				if err := d.DecodeElement(m.Left, &el); err != nil {
					return err
				}
			case "bottom":
				m.Bottom = NewCT_Border()
				if err := d.DecodeElement(m.Bottom, &el); err != nil {
					return err
				}
			case "end":
				m.End = NewCT_Border()
				if err := d.DecodeElement(m.End, &el); err != nil {
					return err
				}
			case "right":
				m.Right = NewCT_Border()
				if err := d.DecodeElement(m.Right, &el); err != nil {
					return err
				}
			case "insideH":
				m.InsideH = NewCT_Border()
				if err := d.DecodeElement(m.InsideH, &el); err != nil {
					return err
				}
			case "insideV":
				m.InsideV = NewCT_Border()
				if err := d.DecodeElement(m.InsideV, &el); err != nil {
					return err
				}
			case "tl2br":
				m.Tl2br = NewCT_Border()
				if err := d.DecodeElement(m.Tl2br, &el); err != nil {
					return err
				}
			case "tr2bl":
				m.Tr2bl = NewCT_Border()
				if err := d.DecodeElement(m.Tr2bl, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TcBorders
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_TcBorders) Validate() error {
	return m.ValidateWithPath("CT_TcBorders")
}
func (m *CT_TcBorders) ValidateWithPath(path string) error {
	if m.Top != nil {
		if err := m.Top.ValidateWithPath(path + "/Top"); err != nil {
			return err
		}
	}
	if m.Start != nil {
		if err := m.Start.ValidateWithPath(path + "/Start"); err != nil {
			return err
		}
	}
	if m.Left != nil {
		if err := m.Left.ValidateWithPath(path + "/Left"); err != nil {
			return err
		}
	}
	if m.Bottom != nil {
		if err := m.Bottom.ValidateWithPath(path + "/Bottom"); err != nil {
			return err
		}
	}
	if m.End != nil {
		if err := m.End.ValidateWithPath(path + "/End"); err != nil {
			return err
		}
	}
	if m.Right != nil {
		if err := m.Right.ValidateWithPath(path + "/Right"); err != nil {
			return err
		}
	}
	if m.InsideH != nil {
		if err := m.InsideH.ValidateWithPath(path + "/InsideH"); err != nil {
			return err
		}
	}
	if m.InsideV != nil {
		if err := m.InsideV.ValidateWithPath(path + "/InsideV"); err != nil {
			return err
		}
	}
	if m.Tl2br != nil {
		if err := m.Tl2br.ValidateWithPath(path + "/Tl2br"); err != nil {
			return err
		}
	}
	if m.Tr2bl != nil {
		if err := m.Tr2bl.ValidateWithPath(path + "/Tr2bl"); err != nil {
			return err
		}
	}
	return nil
}
