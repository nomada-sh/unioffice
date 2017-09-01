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

type CT_AlphaModulateEffect struct {
	Cont *CT_EffectContainer
}

func NewCT_AlphaModulateEffect() *CT_AlphaModulateEffect {
	ret := &CT_AlphaModulateEffect{}
	ret.Cont = NewCT_EffectContainer()
	return ret
}
func (m *CT_AlphaModulateEffect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	secont := xml.StartElement{Name: xml.Name{Local: "a:cont"}}
	e.EncodeElement(m.Cont, secont)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_AlphaModulateEffect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Cont = NewCT_EffectContainer()
lCT_AlphaModulateEffect:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cont":
				if err := d.DecodeElement(m.Cont, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_AlphaModulateEffect
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_AlphaModulateEffect) Validate() error {
	return m.ValidateWithPath("CT_AlphaModulateEffect")
}
func (m *CT_AlphaModulateEffect) ValidateWithPath(path string) error {
	if err := m.Cont.ValidateWithPath(path + "/Cont"); err != nil {
		return err
	}
	return nil
}
