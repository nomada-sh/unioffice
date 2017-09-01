// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package extended_properties

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml"
)

type CT_DigSigBlob struct {
	Blob string
}

func NewCT_DigSigBlob() *CT_DigSigBlob {
	ret := &CT_DigSigBlob{}
	return ret
}
func (m *CT_DigSigBlob) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	seblob := xml.StartElement{Name: xml.Name{Local: "vt:blob"}}
	gooxml.AddPreserveSpaceAttr(&seblob, m.Blob)
	e.EncodeElement(m.Blob, seblob)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_DigSigBlob) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_DigSigBlob:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "blob":
				if err := d.DecodeElement(m.Blob, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DigSigBlob
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_DigSigBlob) Validate() error {
	return m.ValidateWithPath("CT_DigSigBlob")
}
func (m *CT_DigSigBlob) ValidateWithPath(path string) error {
	return nil
}
