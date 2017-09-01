// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"log"
)

type CT_SerTx struct {
	Choice *CT_SerTxChoice
}

func NewCT_SerTx() *CT_SerTx {
	ret := &CT_SerTx{}
	ret.Choice = NewCT_SerTxChoice()
	return ret
}
func (m *CT_SerTx) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	m.Choice.MarshalXML(e, start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_SerTx) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Choice = NewCT_SerTxChoice()
lCT_SerTx:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "strRef":
				m.Choice = NewCT_SerTxChoice()
				if err := d.DecodeElement(&m.Choice.StrRef, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "v":
				m.Choice = NewCT_SerTxChoice()
				if err := d.DecodeElement(&m.Choice.V, &el); err != nil {
					return err
				}
				_ = m.Choice
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SerTx
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_SerTx) Validate() error {
	return m.ValidateWithPath("CT_SerTx")
}
func (m *CT_SerTx) ValidateWithPath(path string) error {
	if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
		return err
	}
	return nil
}
