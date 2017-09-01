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

type CT_LinearShadeProperties struct {
	AngAttr    *int32
	ScaledAttr *bool
}

func NewCT_LinearShadeProperties() *CT_LinearShadeProperties {
	ret := &CT_LinearShadeProperties{}
	return ret
}
func (m *CT_LinearShadeProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.AngAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ang"},
			Value: fmt.Sprintf("%v", *m.AngAttr)})
	}
	if m.ScaledAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "scaled"},
			Value: fmt.Sprintf("%v", *m.ScaledAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_LinearShadeProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "ang" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.AngAttr = &pt
		}
		if attr.Name.Local == "scaled" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ScaledAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_LinearShadeProperties: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_LinearShadeProperties) Validate() error {
	return m.ValidateWithPath("CT_LinearShadeProperties")
}
func (m *CT_LinearShadeProperties) ValidateWithPath(path string) error {
	if m.AngAttr != nil {
		if *m.AngAttr < 0 {
			return fmt.Errorf("%s/m.AngAttr must be >= 0 (have %v)", path, *m.AngAttr)
		}
		if *m.AngAttr >= 21600000 {
			return fmt.Errorf("%s/m.AngAttr must be < 21600000 (have %v)", path, *m.AngAttr)
		}
	}
	return nil
}
