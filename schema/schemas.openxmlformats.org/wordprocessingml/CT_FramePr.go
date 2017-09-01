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
	"strconv"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type CT_FramePr struct {
	// Drop Cap Frame
	DropCapAttr ST_DropCap
	// Drop Cap Vertical Height in Lines
	LinesAttr *int64
	// Frame Width
	WAttr *sharedTypes.ST_TwipsMeasure
	// Frame Height
	HAttr *sharedTypes.ST_TwipsMeasure
	// Vertical Frame Padding
	VSpaceAttr *sharedTypes.ST_TwipsMeasure
	// Horizontal Frame Padding
	HSpaceAttr *sharedTypes.ST_TwipsMeasure
	// Text Wrapping Around Frame
	WrapAttr ST_Wrap
	// Frame Horizontal Positioning Base
	HAnchorAttr ST_HAnchor
	// Frame Vertical Positioning Base
	VAnchorAttr ST_VAnchor
	// Absolute Horizontal Position
	XAttr *ST_SignedTwipsMeasure
	// Relative Horizontal Position
	XAlignAttr sharedTypes.ST_XAlign
	// Absolute Vertical Position
	YAttr *ST_SignedTwipsMeasure
	// Relative Vertical Position
	YAlignAttr sharedTypes.ST_YAlign
	// Frame Height Type
	HRuleAttr ST_HeightRule
	// Lock Frame Anchor to Paragraph
	AnchorLockAttr *sharedTypes.ST_OnOff
}

func NewCT_FramePr() *CT_FramePr {
	ret := &CT_FramePr{}
	return ret
}
func (m *CT_FramePr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.DropCapAttr != ST_DropCapUnset {
		attr, err := m.DropCapAttr.MarshalXMLAttr(xml.Name{Local: "w:dropCap"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.LinesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:lines"},
			Value: fmt.Sprintf("%v", *m.LinesAttr)})
	}
	if m.WAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:w"},
			Value: fmt.Sprintf("%v", *m.WAttr)})
	}
	if m.HAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:h"},
			Value: fmt.Sprintf("%v", *m.HAttr)})
	}
	if m.VSpaceAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:vSpace"},
			Value: fmt.Sprintf("%v", *m.VSpaceAttr)})
	}
	if m.HSpaceAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:hSpace"},
			Value: fmt.Sprintf("%v", *m.HSpaceAttr)})
	}
	if m.WrapAttr != ST_WrapUnset {
		attr, err := m.WrapAttr.MarshalXMLAttr(xml.Name{Local: "w:wrap"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.HAnchorAttr != ST_HAnchorUnset {
		attr, err := m.HAnchorAttr.MarshalXMLAttr(xml.Name{Local: "w:hAnchor"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.VAnchorAttr != ST_VAnchorUnset {
		attr, err := m.VAnchorAttr.MarshalXMLAttr(xml.Name{Local: "w:vAnchor"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.XAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:x"},
			Value: fmt.Sprintf("%v", *m.XAttr)})
	}
	if m.XAlignAttr != sharedTypes.ST_XAlignUnset {
		attr, err := m.XAlignAttr.MarshalXMLAttr(xml.Name{Local: "w:xAlign"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.YAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:y"},
			Value: fmt.Sprintf("%v", *m.YAttr)})
	}
	if m.YAlignAttr != sharedTypes.ST_YAlignUnset {
		attr, err := m.YAlignAttr.MarshalXMLAttr(xml.Name{Local: "w:yAlign"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.HRuleAttr != ST_HeightRuleUnset {
		attr, err := m.HRuleAttr.MarshalXMLAttr(xml.Name{Local: "w:hRule"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.AnchorLockAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:anchorLock"},
			Value: fmt.Sprintf("%v", *m.AnchorLockAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_FramePr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "dropCap" {
			m.DropCapAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "lines" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.LinesAttr = &parsed
		}
		if attr.Name.Local == "w" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.WAttr = &parsed
		}
		if attr.Name.Local == "h" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.HAttr = &parsed
		}
		if attr.Name.Local == "vSpace" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.VSpaceAttr = &parsed
		}
		if attr.Name.Local == "hSpace" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.HSpaceAttr = &parsed
		}
		if attr.Name.Local == "wrap" {
			m.WrapAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "hAnchor" {
			m.HAnchorAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "vAnchor" {
			m.VAnchorAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "x" {
			parsed, err := ParseUnionST_SignedTwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.XAttr = &parsed
		}
		if attr.Name.Local == "xAlign" {
			m.XAlignAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "y" {
			parsed, err := ParseUnionST_SignedTwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.YAttr = &parsed
		}
		if attr.Name.Local == "yAlign" {
			m.YAlignAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "hRule" {
			m.HRuleAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "anchorLock" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.AnchorLockAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_FramePr: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_FramePr) Validate() error {
	return m.ValidateWithPath("CT_FramePr")
}
func (m *CT_FramePr) ValidateWithPath(path string) error {
	if err := m.DropCapAttr.ValidateWithPath(path + "/DropCapAttr"); err != nil {
		return err
	}
	if m.WAttr != nil {
		if err := m.WAttr.ValidateWithPath(path + "/WAttr"); err != nil {
			return err
		}
	}
	if m.HAttr != nil {
		if err := m.HAttr.ValidateWithPath(path + "/HAttr"); err != nil {
			return err
		}
	}
	if m.VSpaceAttr != nil {
		if err := m.VSpaceAttr.ValidateWithPath(path + "/VSpaceAttr"); err != nil {
			return err
		}
	}
	if m.HSpaceAttr != nil {
		if err := m.HSpaceAttr.ValidateWithPath(path + "/HSpaceAttr"); err != nil {
			return err
		}
	}
	if err := m.WrapAttr.ValidateWithPath(path + "/WrapAttr"); err != nil {
		return err
	}
	if err := m.HAnchorAttr.ValidateWithPath(path + "/HAnchorAttr"); err != nil {
		return err
	}
	if err := m.VAnchorAttr.ValidateWithPath(path + "/VAnchorAttr"); err != nil {
		return err
	}
	if m.XAttr != nil {
		if err := m.XAttr.ValidateWithPath(path + "/XAttr"); err != nil {
			return err
		}
	}
	if err := m.XAlignAttr.ValidateWithPath(path + "/XAlignAttr"); err != nil {
		return err
	}
	if m.YAttr != nil {
		if err := m.YAttr.ValidateWithPath(path + "/YAttr"); err != nil {
			return err
		}
	}
	if err := m.YAlignAttr.ValidateWithPath(path + "/YAlignAttr"); err != nil {
		return err
	}
	if err := m.HRuleAttr.ValidateWithPath(path + "/HRuleAttr"); err != nil {
		return err
	}
	if m.AnchorLockAttr != nil {
		if err := m.AnchorLockAttr.ValidateWithPath(path + "/AnchorLockAttr"); err != nil {
			return err
		}
	}
	return nil
}
