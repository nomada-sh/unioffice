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
)

type CT_Shd struct {
	// Shading Pattern
	ValAttr ST_Shd
	// Shading Pattern Color
	ColorAttr *ST_HexColor
	// Shading Pattern Theme Color
	ThemeColorAttr ST_ThemeColor
	// Shading Pattern Theme Color Tint
	ThemeTintAttr *string
	// Shading Pattern Theme Color Shade
	ThemeShadeAttr *string
	// Shading Background Color
	FillAttr *ST_HexColor
	// Shading Background Theme Color
	ThemeFillAttr ST_ThemeColor
	// Shading Background Theme Color Tint
	ThemeFillTintAttr *string
	// Shading Background Theme Color Shade
	ThemeFillShadeAttr *string
}

func NewCT_Shd() *CT_Shd {
	ret := &CT_Shd{}
	ret.ValAttr = ST_Shd(1)
	return ret
}
func (m *CT_Shd) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "w:val"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	if m.ColorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:color"},
			Value: fmt.Sprintf("%v", *m.ColorAttr)})
	}
	if m.ThemeColorAttr != ST_ThemeColorUnset {
		attr, err := m.ThemeColorAttr.MarshalXMLAttr(xml.Name{Local: "w:themeColor"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ThemeTintAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeTint"},
			Value: fmt.Sprintf("%v", *m.ThemeTintAttr)})
	}
	if m.ThemeShadeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeShade"},
			Value: fmt.Sprintf("%v", *m.ThemeShadeAttr)})
	}
	if m.FillAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:fill"},
			Value: fmt.Sprintf("%v", *m.FillAttr)})
	}
	if m.ThemeFillAttr != ST_ThemeColorUnset {
		attr, err := m.ThemeFillAttr.MarshalXMLAttr(xml.Name{Local: "w:themeFill"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ThemeFillTintAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeFillTint"},
			Value: fmt.Sprintf("%v", *m.ThemeFillTintAttr)})
	}
	if m.ThemeFillShadeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeFillShade"},
			Value: fmt.Sprintf("%v", *m.ThemeFillShadeAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Shd) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = ST_Shd(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			m.ValAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "color" {
			parsed, err := ParseUnionST_HexColor(attr.Value)
			if err != nil {
				return err
			}
			m.ColorAttr = &parsed
		}
		if attr.Name.Local == "themeColor" {
			m.ThemeColorAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "themeTint" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ThemeTintAttr = &parsed
		}
		if attr.Name.Local == "themeShade" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ThemeShadeAttr = &parsed
		}
		if attr.Name.Local == "fill" {
			parsed, err := ParseUnionST_HexColor(attr.Value)
			if err != nil {
				return err
			}
			m.FillAttr = &parsed
		}
		if attr.Name.Local == "themeFill" {
			m.ThemeFillAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "themeFillTint" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ThemeFillTintAttr = &parsed
		}
		if attr.Name.Local == "themeFillShade" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ThemeFillShadeAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Shd: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_Shd) Validate() error {
	return m.ValidateWithPath("CT_Shd")
}
func (m *CT_Shd) ValidateWithPath(path string) error {
	if m.ValAttr == ST_ShdUnset {
		return fmt.Errorf("%s/ValAttr is a mandatory field", path)
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	if m.ColorAttr != nil {
		if err := m.ColorAttr.ValidateWithPath(path + "/ColorAttr"); err != nil {
			return err
		}
	}
	if err := m.ThemeColorAttr.ValidateWithPath(path + "/ThemeColorAttr"); err != nil {
		return err
	}
	if m.FillAttr != nil {
		if err := m.FillAttr.ValidateWithPath(path + "/FillAttr"); err != nil {
			return err
		}
	}
	if err := m.ThemeFillAttr.ValidateWithPath(path + "/ThemeFillAttr"); err != nil {
		return err
	}
	return nil
}
