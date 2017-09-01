// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"log"
)

type CT_SdtPr struct {
	// Run Properties For Structured Document Tag Contents
	RPr *CT_RPr
	// Friendly Name
	Alias *CT_String
	// Programmatic Tag
	Tag *CT_String
	// Unique ID
	Id *CT_DecimalNumber
	// Locking Setting
	Lock *CT_Lock
	// Structured Document Tag Placeholder Text
	Placeholder *CT_Placeholder
	// Remove Structured Document Tag When Contents Are Edited
	Temporary *CT_OnOff
	// Current Contents Are Placeholder Text
	ShowingPlcHdr *CT_OnOff
	// XML Mapping
	DataBinding *CT_DataBinding
	// Structured Document Tag Label
	Label *CT_DecimalNumber
	// Structured Document Tag Navigation Order Index
	TabIndex *CT_UnsignedDecimalNumber
	Choice   *CT_SdtPrChoice
}

func NewCT_SdtPr() *CT_SdtPr {
	ret := &CT_SdtPr{}
	return ret
}
func (m *CT_SdtPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.RPr != nil {
		serPr := xml.StartElement{Name: xml.Name{Local: "w:rPr"}}
		e.EncodeElement(m.RPr, serPr)
	}
	if m.Alias != nil {
		sealias := xml.StartElement{Name: xml.Name{Local: "w:alias"}}
		e.EncodeElement(m.Alias, sealias)
	}
	if m.Tag != nil {
		setag := xml.StartElement{Name: xml.Name{Local: "w:tag"}}
		e.EncodeElement(m.Tag, setag)
	}
	if m.Id != nil {
		seid := xml.StartElement{Name: xml.Name{Local: "w:id"}}
		e.EncodeElement(m.Id, seid)
	}
	if m.Lock != nil {
		selock := xml.StartElement{Name: xml.Name{Local: "w:lock"}}
		e.EncodeElement(m.Lock, selock)
	}
	if m.Placeholder != nil {
		seplaceholder := xml.StartElement{Name: xml.Name{Local: "w:placeholder"}}
		e.EncodeElement(m.Placeholder, seplaceholder)
	}
	if m.Temporary != nil {
		setemporary := xml.StartElement{Name: xml.Name{Local: "w:temporary"}}
		e.EncodeElement(m.Temporary, setemporary)
	}
	if m.ShowingPlcHdr != nil {
		seshowingPlcHdr := xml.StartElement{Name: xml.Name{Local: "w:showingPlcHdr"}}
		e.EncodeElement(m.ShowingPlcHdr, seshowingPlcHdr)
	}
	if m.DataBinding != nil {
		sedataBinding := xml.StartElement{Name: xml.Name{Local: "w:dataBinding"}}
		e.EncodeElement(m.DataBinding, sedataBinding)
	}
	if m.Label != nil {
		selabel := xml.StartElement{Name: xml.Name{Local: "w:label"}}
		e.EncodeElement(m.Label, selabel)
	}
	if m.TabIndex != nil {
		setabIndex := xml.StartElement{Name: xml.Name{Local: "w:tabIndex"}}
		e.EncodeElement(m.TabIndex, setabIndex)
	}
	if m.Choice != nil {
		m.Choice.MarshalXML(e, start)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_SdtPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SdtPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "rPr":
				m.RPr = NewCT_RPr()
				if err := d.DecodeElement(m.RPr, &el); err != nil {
					return err
				}
			case "alias":
				m.Alias = NewCT_String()
				if err := d.DecodeElement(m.Alias, &el); err != nil {
					return err
				}
			case "tag":
				m.Tag = NewCT_String()
				if err := d.DecodeElement(m.Tag, &el); err != nil {
					return err
				}
			case "id":
				m.Id = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.Id, &el); err != nil {
					return err
				}
			case "lock":
				m.Lock = NewCT_Lock()
				if err := d.DecodeElement(m.Lock, &el); err != nil {
					return err
				}
			case "placeholder":
				m.Placeholder = NewCT_Placeholder()
				if err := d.DecodeElement(m.Placeholder, &el); err != nil {
					return err
				}
			case "temporary":
				m.Temporary = NewCT_OnOff()
				if err := d.DecodeElement(m.Temporary, &el); err != nil {
					return err
				}
			case "showingPlcHdr":
				m.ShowingPlcHdr = NewCT_OnOff()
				if err := d.DecodeElement(m.ShowingPlcHdr, &el); err != nil {
					return err
				}
			case "dataBinding":
				m.DataBinding = NewCT_DataBinding()
				if err := d.DecodeElement(m.DataBinding, &el); err != nil {
					return err
				}
			case "label":
				m.Label = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.Label, &el); err != nil {
					return err
				}
			case "tabIndex":
				m.TabIndex = NewCT_UnsignedDecimalNumber()
				if err := d.DecodeElement(m.TabIndex, &el); err != nil {
					return err
				}
			case "equation":
				m.Choice = NewCT_SdtPrChoice()
				if err := d.DecodeElement(&m.Choice.Equation, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "comboBox":
				m.Choice = NewCT_SdtPrChoice()
				if err := d.DecodeElement(&m.Choice.ComboBox, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "date":
				m.Choice = NewCT_SdtPrChoice()
				if err := d.DecodeElement(&m.Choice.Date, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "docPartObj":
				m.Choice = NewCT_SdtPrChoice()
				if err := d.DecodeElement(&m.Choice.DocPartObj, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "docPartList":
				m.Choice = NewCT_SdtPrChoice()
				if err := d.DecodeElement(&m.Choice.DocPartList, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "dropDownList":
				m.Choice = NewCT_SdtPrChoice()
				if err := d.DecodeElement(&m.Choice.DropDownList, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "picture":
				m.Choice = NewCT_SdtPrChoice()
				if err := d.DecodeElement(&m.Choice.Picture, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "richText":
				m.Choice = NewCT_SdtPrChoice()
				if err := d.DecodeElement(&m.Choice.RichText, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "text":
				m.Choice = NewCT_SdtPrChoice()
				if err := d.DecodeElement(&m.Choice.Text, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "citation":
				m.Choice = NewCT_SdtPrChoice()
				if err := d.DecodeElement(&m.Choice.Citation, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "group":
				m.Choice = NewCT_SdtPrChoice()
				if err := d.DecodeElement(&m.Choice.Group, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "bibliography":
				m.Choice = NewCT_SdtPrChoice()
				if err := d.DecodeElement(&m.Choice.Bibliography, &el); err != nil {
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
			break lCT_SdtPr
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_SdtPr) Validate() error {
	return m.ValidateWithPath("CT_SdtPr")
}
func (m *CT_SdtPr) ValidateWithPath(path string) error {
	if m.RPr != nil {
		if err := m.RPr.ValidateWithPath(path + "/RPr"); err != nil {
			return err
		}
	}
	if m.Alias != nil {
		if err := m.Alias.ValidateWithPath(path + "/Alias"); err != nil {
			return err
		}
	}
	if m.Tag != nil {
		if err := m.Tag.ValidateWithPath(path + "/Tag"); err != nil {
			return err
		}
	}
	if m.Id != nil {
		if err := m.Id.ValidateWithPath(path + "/Id"); err != nil {
			return err
		}
	}
	if m.Lock != nil {
		if err := m.Lock.ValidateWithPath(path + "/Lock"); err != nil {
			return err
		}
	}
	if m.Placeholder != nil {
		if err := m.Placeholder.ValidateWithPath(path + "/Placeholder"); err != nil {
			return err
		}
	}
	if m.Temporary != nil {
		if err := m.Temporary.ValidateWithPath(path + "/Temporary"); err != nil {
			return err
		}
	}
	if m.ShowingPlcHdr != nil {
		if err := m.ShowingPlcHdr.ValidateWithPath(path + "/ShowingPlcHdr"); err != nil {
			return err
		}
	}
	if m.DataBinding != nil {
		if err := m.DataBinding.ValidateWithPath(path + "/DataBinding"); err != nil {
			return err
		}
	}
	if m.Label != nil {
		if err := m.Label.ValidateWithPath(path + "/Label"); err != nil {
			return err
		}
	}
	if m.TabIndex != nil {
		if err := m.TabIndex.ValidateWithPath(path + "/TabIndex"); err != nil {
			return err
		}
	}
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	return nil
}
