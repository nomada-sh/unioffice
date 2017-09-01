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

type CT_BaseStyles struct {
	ClrScheme  *CT_ColorScheme
	FontScheme *CT_FontScheme
	FmtScheme  *CT_StyleMatrix
	ExtLst     *CT_OfficeArtExtensionList
}

func NewCT_BaseStyles() *CT_BaseStyles {
	ret := &CT_BaseStyles{}
	ret.ClrScheme = NewCT_ColorScheme()
	ret.FontScheme = NewCT_FontScheme()
	ret.FmtScheme = NewCT_StyleMatrix()
	return ret
}
func (m *CT_BaseStyles) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	seclrScheme := xml.StartElement{Name: xml.Name{Local: "a:clrScheme"}}
	e.EncodeElement(m.ClrScheme, seclrScheme)
	sefontScheme := xml.StartElement{Name: xml.Name{Local: "a:fontScheme"}}
	e.EncodeElement(m.FontScheme, sefontScheme)
	sefmtScheme := xml.StartElement{Name: xml.Name{Local: "a:fmtScheme"}}
	e.EncodeElement(m.FmtScheme, sefmtScheme)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_BaseStyles) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ClrScheme = NewCT_ColorScheme()
	m.FontScheme = NewCT_FontScheme()
	m.FmtScheme = NewCT_StyleMatrix()
lCT_BaseStyles:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "clrScheme":
				if err := d.DecodeElement(m.ClrScheme, &el); err != nil {
					return err
				}
			case "fontScheme":
				if err := d.DecodeElement(m.FontScheme, &el); err != nil {
					return err
				}
			case "fmtScheme":
				if err := d.DecodeElement(m.FmtScheme, &el); err != nil {
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
			break lCT_BaseStyles
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_BaseStyles) Validate() error {
	return m.ValidateWithPath("CT_BaseStyles")
}
func (m *CT_BaseStyles) ValidateWithPath(path string) error {
	if err := m.ClrScheme.ValidateWithPath(path + "/ClrScheme"); err != nil {
		return err
	}
	if err := m.FontScheme.ValidateWithPath(path + "/FontScheme"); err != nil {
		return err
	}
	if err := m.FmtScheme.ValidateWithPath(path + "/FmtScheme"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
