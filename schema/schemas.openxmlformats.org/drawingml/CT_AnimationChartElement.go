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

type CT_AnimationChartElement struct {
	SeriesIdxAttr   *int32
	CategoryIdxAttr *int32
	BldStepAttr     ST_ChartBuildStep
}

func NewCT_AnimationChartElement() *CT_AnimationChartElement {
	ret := &CT_AnimationChartElement{}
	ret.BldStepAttr = ST_ChartBuildStep(1)
	return ret
}
func (m *CT_AnimationChartElement) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.SeriesIdxAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "seriesIdx"},
			Value: fmt.Sprintf("%v", *m.SeriesIdxAttr)})
	}
	if m.CategoryIdxAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "categoryIdx"},
			Value: fmt.Sprintf("%v", *m.CategoryIdxAttr)})
	}
	attr, err := m.BldStepAttr.MarshalXMLAttr(xml.Name{Local: "bldStep"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_AnimationChartElement) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.BldStepAttr = ST_ChartBuildStep(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "seriesIdx" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.SeriesIdxAttr = &pt
		}
		if attr.Name.Local == "categoryIdx" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.CategoryIdxAttr = &pt
		}
		if attr.Name.Local == "bldStep" {
			m.BldStepAttr.UnmarshalXMLAttr(attr)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_AnimationChartElement: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_AnimationChartElement) Validate() error {
	return m.ValidateWithPath("CT_AnimationChartElement")
}
func (m *CT_AnimationChartElement) ValidateWithPath(path string) error {
	if m.BldStepAttr == ST_ChartBuildStepUnset {
		return fmt.Errorf("%s/BldStepAttr is a mandatory field", path)
	}
	if err := m.BldStepAttr.ValidateWithPath(path + "/BldStepAttr"); err != nil {
		return err
	}
	return nil
}
