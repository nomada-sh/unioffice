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
)

type CT_LineJoinBevel struct {
}

func NewCT_LineJoinBevel() *CT_LineJoinBevel {
	ret := &CT_LineJoinBevel{}
	return ret
}
func (m *CT_LineJoinBevel) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_LineJoinBevel) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_LineJoinBevel: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_LineJoinBevel) Validate() error {
	return m.ValidateWithPath("CT_LineJoinBevel")
}
func (m *CT_LineJoinBevel) ValidateWithPath(path string) error {
	return nil
}
