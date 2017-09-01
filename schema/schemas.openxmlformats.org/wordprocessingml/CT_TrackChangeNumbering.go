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
	"time"
)

type CT_TrackChangeNumbering struct {
	OriginalAttr *string
	AuthorAttr   string
	DateAttr     *time.Time
	// Annotation Identifier
	IdAttr int64
}

func NewCT_TrackChangeNumbering() *CT_TrackChangeNumbering {
	ret := &CT_TrackChangeNumbering{}
	return ret
}
func (m *CT_TrackChangeNumbering) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.OriginalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:original"},
			Value: fmt.Sprintf("%v", *m.OriginalAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:author"},
		Value: fmt.Sprintf("%v", m.AuthorAttr)})
	if m.DateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:date"},
			Value: fmt.Sprintf("%v", *m.DateAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TrackChangeNumbering) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "original" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OriginalAttr = &parsed
		}
		if attr.Name.Local == "author" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AuthorAttr = parsed
		}
		if attr.Name.Local == "date" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.DateAttr = &parsed
		}
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TrackChangeNumbering: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_TrackChangeNumbering) Validate() error {
	return m.ValidateWithPath("CT_TrackChangeNumbering")
}
func (m *CT_TrackChangeNumbering) ValidateWithPath(path string) error {
	return nil
}
