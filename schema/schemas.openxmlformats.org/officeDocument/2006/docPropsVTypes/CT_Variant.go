// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package docPropsVTypes

import (
	"encoding/xml"
	"fmt"
	"log"
	"time"

	"baliance.com/gooxml"
	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type CT_Variant struct {
	Variant  *Variant
	Vector   *Vector
	Array    *Array
	Blob     *string
	Oblob    *string
	Empty    *Empty
	Null     *Null
	I1       *int8
	I2       *int16
	I4       *int32
	I8       *int64
	Int      *int32
	Ui1      *uint8
	Ui2      *uint16
	Ui4      *uint32
	Ui8      *uint64
	Uint     *uint32
	R4       *float32
	R8       *float64
	Decimal  *float64
	Lpstr    *string
	Lpwstr   *string
	Bstr     *string
	Date     *time.Time
	Filetime *time.Time
	Bool     *bool
	Cy       *string
	Error    *string
	Stream   *string
	Ostream  *string
	Storage  *string
	Ostorage *string
	Vstream  *Vstream
	Clsid    *string
}

func NewCT_Variant() *CT_Variant {
	ret := &CT_Variant{}
	return ret
}
func (m *CT_Variant) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.Variant != nil {
		sevariant := xml.StartElement{Name: xml.Name{Local: "vt:variant"}}
		e.EncodeElement(m.Variant, sevariant)
	}
	if m.Vector != nil {
		sevector := xml.StartElement{Name: xml.Name{Local: "vt:vector"}}
		e.EncodeElement(m.Vector, sevector)
	}
	if m.Array != nil {
		searray := xml.StartElement{Name: xml.Name{Local: "vt:array"}}
		e.EncodeElement(m.Array, searray)
	}
	if m.Blob != nil {
		seblob := xml.StartElement{Name: xml.Name{Local: "vt:blob"}}
		gooxml.AddPreserveSpaceAttr(&seblob, *m.Blob)
		e.EncodeElement(m.Blob, seblob)
	}
	if m.Oblob != nil {
		seoblob := xml.StartElement{Name: xml.Name{Local: "vt:oblob"}}
		gooxml.AddPreserveSpaceAttr(&seoblob, *m.Oblob)
		e.EncodeElement(m.Oblob, seoblob)
	}
	if m.Empty != nil {
		seempty := xml.StartElement{Name: xml.Name{Local: "vt:empty"}}
		e.EncodeElement(m.Empty, seempty)
	}
	if m.Null != nil {
		senull := xml.StartElement{Name: xml.Name{Local: "vt:null"}}
		e.EncodeElement(m.Null, senull)
	}
	if m.I1 != nil {
		sei1 := xml.StartElement{Name: xml.Name{Local: "vt:i1"}}
		e.EncodeElement(m.I1, sei1)
	}
	if m.I2 != nil {
		sei2 := xml.StartElement{Name: xml.Name{Local: "vt:i2"}}
		e.EncodeElement(m.I2, sei2)
	}
	if m.I4 != nil {
		sei4 := xml.StartElement{Name: xml.Name{Local: "vt:i4"}}
		e.EncodeElement(m.I4, sei4)
	}
	if m.I8 != nil {
		sei8 := xml.StartElement{Name: xml.Name{Local: "vt:i8"}}
		e.EncodeElement(m.I8, sei8)
	}
	if m.Int != nil {
		seint := xml.StartElement{Name: xml.Name{Local: "vt:int"}}
		e.EncodeElement(m.Int, seint)
	}
	if m.Ui1 != nil {
		seui1 := xml.StartElement{Name: xml.Name{Local: "vt:ui1"}}
		e.EncodeElement(m.Ui1, seui1)
	}
	if m.Ui2 != nil {
		seui2 := xml.StartElement{Name: xml.Name{Local: "vt:ui2"}}
		e.EncodeElement(m.Ui2, seui2)
	}
	if m.Ui4 != nil {
		seui4 := xml.StartElement{Name: xml.Name{Local: "vt:ui4"}}
		e.EncodeElement(m.Ui4, seui4)
	}
	if m.Ui8 != nil {
		seui8 := xml.StartElement{Name: xml.Name{Local: "vt:ui8"}}
		e.EncodeElement(m.Ui8, seui8)
	}
	if m.Uint != nil {
		seuint := xml.StartElement{Name: xml.Name{Local: "vt:uint"}}
		e.EncodeElement(m.Uint, seuint)
	}
	if m.R4 != nil {
		ser4 := xml.StartElement{Name: xml.Name{Local: "vt:r4"}}
		e.EncodeElement(m.R4, ser4)
	}
	if m.R8 != nil {
		ser8 := xml.StartElement{Name: xml.Name{Local: "vt:r8"}}
		e.EncodeElement(m.R8, ser8)
	}
	if m.Decimal != nil {
		sedecimal := xml.StartElement{Name: xml.Name{Local: "vt:decimal"}}
		e.EncodeElement(m.Decimal, sedecimal)
	}
	if m.Lpstr != nil {
		selpstr := xml.StartElement{Name: xml.Name{Local: "vt:lpstr"}}
		gooxml.AddPreserveSpaceAttr(&selpstr, *m.Lpstr)
		e.EncodeElement(m.Lpstr, selpstr)
	}
	if m.Lpwstr != nil {
		selpwstr := xml.StartElement{Name: xml.Name{Local: "vt:lpwstr"}}
		gooxml.AddPreserveSpaceAttr(&selpwstr, *m.Lpwstr)
		e.EncodeElement(m.Lpwstr, selpwstr)
	}
	if m.Bstr != nil {
		sebstr := xml.StartElement{Name: xml.Name{Local: "vt:bstr"}}
		gooxml.AddPreserveSpaceAttr(&sebstr, *m.Bstr)
		e.EncodeElement(m.Bstr, sebstr)
	}
	if m.Date != nil {
		sedate := xml.StartElement{Name: xml.Name{Local: "vt:date"}}
		e.EncodeElement(m.Date, sedate)
	}
	if m.Filetime != nil {
		sefiletime := xml.StartElement{Name: xml.Name{Local: "vt:filetime"}}
		e.EncodeElement(m.Filetime, sefiletime)
	}
	if m.Bool != nil {
		sebool := xml.StartElement{Name: xml.Name{Local: "vt:bool"}}
		e.EncodeElement(m.Bool, sebool)
	}
	if m.Cy != nil {
		secy := xml.StartElement{Name: xml.Name{Local: "vt:cy"}}
		gooxml.AddPreserveSpaceAttr(&secy, *m.Cy)
		e.EncodeElement(m.Cy, secy)
	}
	if m.Error != nil {
		seerror := xml.StartElement{Name: xml.Name{Local: "vt:error"}}
		gooxml.AddPreserveSpaceAttr(&seerror, *m.Error)
		e.EncodeElement(m.Error, seerror)
	}
	if m.Stream != nil {
		sestream := xml.StartElement{Name: xml.Name{Local: "vt:stream"}}
		gooxml.AddPreserveSpaceAttr(&sestream, *m.Stream)
		e.EncodeElement(m.Stream, sestream)
	}
	if m.Ostream != nil {
		seostream := xml.StartElement{Name: xml.Name{Local: "vt:ostream"}}
		gooxml.AddPreserveSpaceAttr(&seostream, *m.Ostream)
		e.EncodeElement(m.Ostream, seostream)
	}
	if m.Storage != nil {
		sestorage := xml.StartElement{Name: xml.Name{Local: "vt:storage"}}
		gooxml.AddPreserveSpaceAttr(&sestorage, *m.Storage)
		e.EncodeElement(m.Storage, sestorage)
	}
	if m.Ostorage != nil {
		seostorage := xml.StartElement{Name: xml.Name{Local: "vt:ostorage"}}
		gooxml.AddPreserveSpaceAttr(&seostorage, *m.Ostorage)
		e.EncodeElement(m.Ostorage, seostorage)
	}
	if m.Vstream != nil {
		sevstream := xml.StartElement{Name: xml.Name{Local: "vt:vstream"}}
		e.EncodeElement(m.Vstream, sevstream)
	}
	if m.Clsid != nil {
		seclsid := xml.StartElement{Name: xml.Name{Local: "vt:clsid"}}
		gooxml.AddPreserveSpaceAttr(&seclsid, *m.Clsid)
		e.EncodeElement(m.Clsid, seclsid)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Variant) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Variant:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "variant":
				m.Variant = NewVariant()
				if err := d.DecodeElement(m.Variant, &el); err != nil {
					return err
				}
			case "vector":
				m.Vector = NewVector()
				if err := d.DecodeElement(m.Vector, &el); err != nil {
					return err
				}
			case "array":
				m.Array = NewArray()
				if err := d.DecodeElement(m.Array, &el); err != nil {
					return err
				}
			case "blob":
				m.Blob = new(string)
				if err := d.DecodeElement(m.Blob, &el); err != nil {
					return err
				}
			case "oblob":
				m.Oblob = new(string)
				if err := d.DecodeElement(m.Oblob, &el); err != nil {
					return err
				}
			case "empty":
				m.Empty = NewEmpty()
				if err := d.DecodeElement(m.Empty, &el); err != nil {
					return err
				}
			case "null":
				m.Null = NewNull()
				if err := d.DecodeElement(m.Null, &el); err != nil {
					return err
				}
			case "i1":
				m.I1 = new(int8)
				if err := d.DecodeElement(m.I1, &el); err != nil {
					return err
				}
			case "i2":
				m.I2 = new(int16)
				if err := d.DecodeElement(m.I2, &el); err != nil {
					return err
				}
			case "i4":
				m.I4 = new(int32)
				if err := d.DecodeElement(m.I4, &el); err != nil {
					return err
				}
			case "i8":
				m.I8 = new(int64)
				if err := d.DecodeElement(m.I8, &el); err != nil {
					return err
				}
			case "int":
				m.Int = new(int32)
				if err := d.DecodeElement(m.Int, &el); err != nil {
					return err
				}
			case "ui1":
				m.Ui1 = new(uint8)
				if err := d.DecodeElement(m.Ui1, &el); err != nil {
					return err
				}
			case "ui2":
				m.Ui2 = new(uint16)
				if err := d.DecodeElement(m.Ui2, &el); err != nil {
					return err
				}
			case "ui4":
				m.Ui4 = new(uint32)
				if err := d.DecodeElement(m.Ui4, &el); err != nil {
					return err
				}
			case "ui8":
				m.Ui8 = new(uint64)
				if err := d.DecodeElement(m.Ui8, &el); err != nil {
					return err
				}
			case "uint":
				m.Uint = new(uint32)
				if err := d.DecodeElement(m.Uint, &el); err != nil {
					return err
				}
			case "r4":
				m.R4 = new(float32)
				if err := d.DecodeElement(m.R4, &el); err != nil {
					return err
				}
			case "r8":
				m.R8 = new(float64)
				if err := d.DecodeElement(m.R8, &el); err != nil {
					return err
				}
			case "decimal":
				m.Decimal = new(float64)
				if err := d.DecodeElement(m.Decimal, &el); err != nil {
					return err
				}
			case "lpstr":
				m.Lpstr = new(string)
				if err := d.DecodeElement(m.Lpstr, &el); err != nil {
					return err
				}
			case "lpwstr":
				m.Lpwstr = new(string)
				if err := d.DecodeElement(m.Lpwstr, &el); err != nil {
					return err
				}
			case "bstr":
				m.Bstr = new(string)
				if err := d.DecodeElement(m.Bstr, &el); err != nil {
					return err
				}
			case "date":
				m.Date = new(time.Time)
				if err := d.DecodeElement(m.Date, &el); err != nil {
					return err
				}
			case "filetime":
				m.Filetime = new(time.Time)
				if err := d.DecodeElement(m.Filetime, &el); err != nil {
					return err
				}
			case "bool":
				m.Bool = new(bool)
				if err := d.DecodeElement(m.Bool, &el); err != nil {
					return err
				}
			case "cy":
				m.Cy = new(string)
				if err := d.DecodeElement(m.Cy, &el); err != nil {
					return err
				}
			case "error":
				m.Error = new(string)
				if err := d.DecodeElement(m.Error, &el); err != nil {
					return err
				}
			case "stream":
				m.Stream = new(string)
				if err := d.DecodeElement(m.Stream, &el); err != nil {
					return err
				}
			case "ostream":
				m.Ostream = new(string)
				if err := d.DecodeElement(m.Ostream, &el); err != nil {
					return err
				}
			case "storage":
				m.Storage = new(string)
				if err := d.DecodeElement(m.Storage, &el); err != nil {
					return err
				}
			case "ostorage":
				m.Ostorage = new(string)
				if err := d.DecodeElement(m.Ostorage, &el); err != nil {
					return err
				}
			case "vstream":
				m.Vstream = NewVstream()
				if err := d.DecodeElement(m.Vstream, &el); err != nil {
					return err
				}
			case "clsid":
				m.Clsid = new(string)
				if err := d.DecodeElement(m.Clsid, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Variant
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Variant) Validate() error {
	return m.ValidateWithPath("CT_Variant")
}
func (m *CT_Variant) ValidateWithPath(path string) error {
	if m.Variant != nil {
		if err := m.Variant.ValidateWithPath(path + "/Variant"); err != nil {
			return err
		}
	}
	if m.Vector != nil {
		if err := m.Vector.ValidateWithPath(path + "/Vector"); err != nil {
			return err
		}
	}
	if m.Array != nil {
		if err := m.Array.ValidateWithPath(path + "/Array"); err != nil {
			return err
		}
	}
	if m.Empty != nil {
		if err := m.Empty.ValidateWithPath(path + "/Empty"); err != nil {
			return err
		}
	}
	if m.Null != nil {
		if err := m.Null.ValidateWithPath(path + "/Null"); err != nil {
			return err
		}
	}
	if m.Cy != nil {
		if !ST_CyPatternRe.MatchString(*m.Cy) {
			return fmt.Errorf(`%s/m.Cy must match '%s' (have %v)`, path, ST_CyPatternRe, *m.Cy)
		}
	}
	if m.Error != nil {
		if !ST_ErrorPatternRe.MatchString(*m.Error) {
			return fmt.Errorf(`%s/m.Error must match '%s' (have %v)`, path, ST_ErrorPatternRe, *m.Error)
		}
	}
	if m.Vstream != nil {
		if err := m.Vstream.ValidateWithPath(path + "/Vstream"); err != nil {
			return err
		}
	}
	if m.Clsid != nil {
		if !sharedTypes.ST_GuidPatternRe.MatchString(*m.Clsid) {
			return fmt.Errorf(`%s/m.Clsid must match '%s' (have %v)`, path, sharedTypes.ST_GuidPatternRe, *m.Clsid)
		}
	}
	return nil
}
