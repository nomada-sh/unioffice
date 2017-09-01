// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_Frameset struct {
	// Nested Frameset Size
	Sz *CT_String
	// Frameset Splitter Properties
	FramesetSplitbar *CT_FramesetSplitbar
	// Frameset Layout
	FrameLayout *CT_FrameLayout
	Title       *CT_String
	Choice      []*CT_FramesetChoice
}

func NewCT_Frameset() *CT_Frameset {
	ret := &CT_Frameset{}
	return ret
}
func (m *CT_Frameset) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.Sz != nil {
		sesz := xml.StartElement{Name: xml.Name{Local: "w:sz"}}
		e.EncodeElement(m.Sz, sesz)
	}
	if m.FramesetSplitbar != nil {
		seframesetSplitbar := xml.StartElement{Name: xml.Name{Local: "w:framesetSplitbar"}}
		e.EncodeElement(m.FramesetSplitbar, seframesetSplitbar)
	}
	if m.FrameLayout != nil {
		seframeLayout := xml.StartElement{Name: xml.Name{Local: "w:frameLayout"}}
		e.EncodeElement(m.FrameLayout, seframeLayout)
	}
	if m.Title != nil {
		setitle := xml.StartElement{Name: xml.Name{Local: "w:title"}}
		e.EncodeElement(m.Title, setitle)
	}
	if m.Choice != nil {
		for _, c := range m.Choice {
			c.MarshalXML(e, start)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Frameset) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Frameset:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "sz":
				m.Sz = NewCT_String()
				if err := d.DecodeElement(m.Sz, &el); err != nil {
					return err
				}
			case "framesetSplitbar":
				m.FramesetSplitbar = NewCT_FramesetSplitbar()
				if err := d.DecodeElement(m.FramesetSplitbar, &el); err != nil {
					return err
				}
			case "frameLayout":
				m.FrameLayout = NewCT_FrameLayout()
				if err := d.DecodeElement(m.FrameLayout, &el); err != nil {
					return err
				}
			case "title":
				m.Title = NewCT_String()
				if err := d.DecodeElement(m.Title, &el); err != nil {
					return err
				}
			case "frameset":
				tmp := NewCT_FramesetChoice()
				if err := d.DecodeElement(&tmp.Frameset, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "frame":
				tmp := NewCT_FramesetChoice()
				if err := d.DecodeElement(&tmp.Frame, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Frameset
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Frameset) Validate() error {
	return m.ValidateWithPath("CT_Frameset")
}
func (m *CT_Frameset) ValidateWithPath(path string) error {
	if m.Sz != nil {
		if err := m.Sz.ValidateWithPath(path + "/Sz"); err != nil {
			return err
		}
	}
	if m.FramesetSplitbar != nil {
		if err := m.FramesetSplitbar.ValidateWithPath(path + "/FramesetSplitbar"); err != nil {
			return err
		}
	}
	if m.FrameLayout != nil {
		if err := m.FrameLayout.ValidateWithPath(path + "/FrameLayout"); err != nil {
			return err
		}
	}
	if m.Title != nil {
		if err := m.Title.ValidateWithPath(path + "/Title"); err != nil {
			return err
		}
	}
	for i, v := range m.Choice {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Choice[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
