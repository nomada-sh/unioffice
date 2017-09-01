// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_AnimationChartBuildProperties struct {
	BldAttr    *ST_AnimationChartBuildType
	AnimBgAttr *bool
}

func NewCT_AnimationChartBuildProperties() *CT_AnimationChartBuildProperties {
	ret := &CT_AnimationChartBuildProperties{}
	return ret
}
func (m *CT_AnimationChartBuildProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.BldAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "bld"},
			Value: fmt.Sprintf("%v", *m.BldAttr)})
	}
	if m.AnimBgAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "animBg"},
			Value: fmt.Sprintf("%v", *m.AnimBgAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_AnimationChartBuildProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "bld" {
			parsed, err := ParseUnionST_AnimationChartBuildType(attr.Value)
			if err != nil {
				return err
			}
			m.BldAttr = &parsed
		}
		if attr.Name.Local == "animBg" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AnimBgAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_AnimationChartBuildProperties: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_AnimationChartBuildProperties) Validate() error {
	return m.ValidateWithPath("CT_AnimationChartBuildProperties")
}
func (m *CT_AnimationChartBuildProperties) ValidateWithPath(path string) error {
	if m.BldAttr != nil {
		if err := m.BldAttr.ValidateWithPath(path + "/BldAttr"); err != nil {
			return err
		}
	}
	return nil
}
