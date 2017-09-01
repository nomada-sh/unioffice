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

type CT_WebSettings struct {
	// Root Frameset Definition
	Frameset *CT_Frameset
	// Information about HTML div Elements
	Divs *CT_Divs
	// Output Encoding When Saving as Web Page
	Encoding *CT_String
	// Disable Features Not Supported by Target Web Browser
	OptimizeForBrowser *CT_OptimizeForBrowser
	// Utilize VML When Saving as Web Page
	RelyOnVML *CT_OnOff
	// Allow PNG as Graphic Format
	AllowPNG *CT_OnOff
	// Do Not Rely on CSS for Font Face Formatting
	DoNotRelyOnCSS *CT_OnOff
	// Recommend Web Page Format over Single File Web Page Format
	DoNotSaveAsSingleFile *CT_OnOff
	// Do Not Place Supporting Files in Subdirectory
	DoNotOrganizeInFolder *CT_OnOff
	// Do Not Use File Names Longer than 8.3 Characters
	DoNotUseLongFileNames *CT_OnOff
	// Pixels per Inch for Graphics/Images
	PixelsPerInch *CT_DecimalNumber
	// Target Screen Size for Web Page
	TargetScreenSz *CT_TargetScreenSz
	// Save Smart Tag Data in XML Property Bag
	SaveSmartTagsAsXml *CT_OnOff
}

func NewCT_WebSettings() *CT_WebSettings {
	ret := &CT_WebSettings{}
	return ret
}
func (m *CT_WebSettings) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.Frameset != nil {
		seframeset := xml.StartElement{Name: xml.Name{Local: "w:frameset"}}
		e.EncodeElement(m.Frameset, seframeset)
	}
	if m.Divs != nil {
		sedivs := xml.StartElement{Name: xml.Name{Local: "w:divs"}}
		e.EncodeElement(m.Divs, sedivs)
	}
	if m.Encoding != nil {
		seencoding := xml.StartElement{Name: xml.Name{Local: "w:encoding"}}
		e.EncodeElement(m.Encoding, seencoding)
	}
	if m.OptimizeForBrowser != nil {
		seoptimizeForBrowser := xml.StartElement{Name: xml.Name{Local: "w:optimizeForBrowser"}}
		e.EncodeElement(m.OptimizeForBrowser, seoptimizeForBrowser)
	}
	if m.RelyOnVML != nil {
		serelyOnVML := xml.StartElement{Name: xml.Name{Local: "w:relyOnVML"}}
		e.EncodeElement(m.RelyOnVML, serelyOnVML)
	}
	if m.AllowPNG != nil {
		seallowPNG := xml.StartElement{Name: xml.Name{Local: "w:allowPNG"}}
		e.EncodeElement(m.AllowPNG, seallowPNG)
	}
	if m.DoNotRelyOnCSS != nil {
		sedoNotRelyOnCSS := xml.StartElement{Name: xml.Name{Local: "w:doNotRelyOnCSS"}}
		e.EncodeElement(m.DoNotRelyOnCSS, sedoNotRelyOnCSS)
	}
	if m.DoNotSaveAsSingleFile != nil {
		sedoNotSaveAsSingleFile := xml.StartElement{Name: xml.Name{Local: "w:doNotSaveAsSingleFile"}}
		e.EncodeElement(m.DoNotSaveAsSingleFile, sedoNotSaveAsSingleFile)
	}
	if m.DoNotOrganizeInFolder != nil {
		sedoNotOrganizeInFolder := xml.StartElement{Name: xml.Name{Local: "w:doNotOrganizeInFolder"}}
		e.EncodeElement(m.DoNotOrganizeInFolder, sedoNotOrganizeInFolder)
	}
	if m.DoNotUseLongFileNames != nil {
		sedoNotUseLongFileNames := xml.StartElement{Name: xml.Name{Local: "w:doNotUseLongFileNames"}}
		e.EncodeElement(m.DoNotUseLongFileNames, sedoNotUseLongFileNames)
	}
	if m.PixelsPerInch != nil {
		sepixelsPerInch := xml.StartElement{Name: xml.Name{Local: "w:pixelsPerInch"}}
		e.EncodeElement(m.PixelsPerInch, sepixelsPerInch)
	}
	if m.TargetScreenSz != nil {
		setargetScreenSz := xml.StartElement{Name: xml.Name{Local: "w:targetScreenSz"}}
		e.EncodeElement(m.TargetScreenSz, setargetScreenSz)
	}
	if m.SaveSmartTagsAsXml != nil {
		sesaveSmartTagsAsXml := xml.StartElement{Name: xml.Name{Local: "w:saveSmartTagsAsXml"}}
		e.EncodeElement(m.SaveSmartTagsAsXml, sesaveSmartTagsAsXml)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_WebSettings) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_WebSettings:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "frameset":
				m.Frameset = NewCT_Frameset()
				if err := d.DecodeElement(m.Frameset, &el); err != nil {
					return err
				}
			case "divs":
				m.Divs = NewCT_Divs()
				if err := d.DecodeElement(m.Divs, &el); err != nil {
					return err
				}
			case "encoding":
				m.Encoding = NewCT_String()
				if err := d.DecodeElement(m.Encoding, &el); err != nil {
					return err
				}
			case "optimizeForBrowser":
				m.OptimizeForBrowser = NewCT_OptimizeForBrowser()
				if err := d.DecodeElement(m.OptimizeForBrowser, &el); err != nil {
					return err
				}
			case "relyOnVML":
				m.RelyOnVML = NewCT_OnOff()
				if err := d.DecodeElement(m.RelyOnVML, &el); err != nil {
					return err
				}
			case "allowPNG":
				m.AllowPNG = NewCT_OnOff()
				if err := d.DecodeElement(m.AllowPNG, &el); err != nil {
					return err
				}
			case "doNotRelyOnCSS":
				m.DoNotRelyOnCSS = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotRelyOnCSS, &el); err != nil {
					return err
				}
			case "doNotSaveAsSingleFile":
				m.DoNotSaveAsSingleFile = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotSaveAsSingleFile, &el); err != nil {
					return err
				}
			case "doNotOrganizeInFolder":
				m.DoNotOrganizeInFolder = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotOrganizeInFolder, &el); err != nil {
					return err
				}
			case "doNotUseLongFileNames":
				m.DoNotUseLongFileNames = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotUseLongFileNames, &el); err != nil {
					return err
				}
			case "pixelsPerInch":
				m.PixelsPerInch = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.PixelsPerInch, &el); err != nil {
					return err
				}
			case "targetScreenSz":
				m.TargetScreenSz = NewCT_TargetScreenSz()
				if err := d.DecodeElement(m.TargetScreenSz, &el); err != nil {
					return err
				}
			case "saveSmartTagsAsXml":
				m.SaveSmartTagsAsXml = NewCT_OnOff()
				if err := d.DecodeElement(m.SaveSmartTagsAsXml, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_WebSettings
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_WebSettings) Validate() error {
	return m.ValidateWithPath("CT_WebSettings")
}
func (m *CT_WebSettings) ValidateWithPath(path string) error {
	if m.Frameset != nil {
		if err := m.Frameset.ValidateWithPath(path + "/Frameset"); err != nil {
			return err
		}
	}
	if m.Divs != nil {
		if err := m.Divs.ValidateWithPath(path + "/Divs"); err != nil {
			return err
		}
	}
	if m.Encoding != nil {
		if err := m.Encoding.ValidateWithPath(path + "/Encoding"); err != nil {
			return err
		}
	}
	if m.OptimizeForBrowser != nil {
		if err := m.OptimizeForBrowser.ValidateWithPath(path + "/OptimizeForBrowser"); err != nil {
			return err
		}
	}
	if m.RelyOnVML != nil {
		if err := m.RelyOnVML.ValidateWithPath(path + "/RelyOnVML"); err != nil {
			return err
		}
	}
	if m.AllowPNG != nil {
		if err := m.AllowPNG.ValidateWithPath(path + "/AllowPNG"); err != nil {
			return err
		}
	}
	if m.DoNotRelyOnCSS != nil {
		if err := m.DoNotRelyOnCSS.ValidateWithPath(path + "/DoNotRelyOnCSS"); err != nil {
			return err
		}
	}
	if m.DoNotSaveAsSingleFile != nil {
		if err := m.DoNotSaveAsSingleFile.ValidateWithPath(path + "/DoNotSaveAsSingleFile"); err != nil {
			return err
		}
	}
	if m.DoNotOrganizeInFolder != nil {
		if err := m.DoNotOrganizeInFolder.ValidateWithPath(path + "/DoNotOrganizeInFolder"); err != nil {
			return err
		}
	}
	if m.DoNotUseLongFileNames != nil {
		if err := m.DoNotUseLongFileNames.ValidateWithPath(path + "/DoNotUseLongFileNames"); err != nil {
			return err
		}
	}
	if m.PixelsPerInch != nil {
		if err := m.PixelsPerInch.ValidateWithPath(path + "/PixelsPerInch"); err != nil {
			return err
		}
	}
	if m.TargetScreenSz != nil {
		if err := m.TargetScreenSz.ValidateWithPath(path + "/TargetScreenSz"); err != nil {
			return err
		}
	}
	if m.SaveSmartTagsAsXml != nil {
		if err := m.SaveSmartTagsAsXml.ValidateWithPath(path + "/SaveSmartTagsAsXml"); err != nil {
			return err
		}
	}
	return nil
}
