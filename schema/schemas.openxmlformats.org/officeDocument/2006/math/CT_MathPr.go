// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"
	"log"
)

type CT_MathPr struct {
	MathFont  *CT_String
	BrkBin    *CT_BreakBin
	BrkBinSub *CT_BreakBinSub
	SmallFrac *CT_OnOff
	DispDef   *CT_OnOff
	LMargin   *CT_TwipsMeasure
	RMargin   *CT_TwipsMeasure
	DefJc     *CT_OMathJc
	PreSp     *CT_TwipsMeasure
	PostSp    *CT_TwipsMeasure
	InterSp   *CT_TwipsMeasure
	IntraSp   *CT_TwipsMeasure
	Choice    *CT_MathPrChoice
	IntLim    *CT_LimLoc
	NaryLim   *CT_LimLoc
}

func NewCT_MathPr() *CT_MathPr {
	ret := &CT_MathPr{}
	return ret
}
func (m *CT_MathPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.MathFont != nil {
		semathFont := xml.StartElement{Name: xml.Name{Local: "m:mathFont"}}
		e.EncodeElement(m.MathFont, semathFont)
	}
	if m.BrkBin != nil {
		sebrkBin := xml.StartElement{Name: xml.Name{Local: "m:brkBin"}}
		e.EncodeElement(m.BrkBin, sebrkBin)
	}
	if m.BrkBinSub != nil {
		sebrkBinSub := xml.StartElement{Name: xml.Name{Local: "m:brkBinSub"}}
		e.EncodeElement(m.BrkBinSub, sebrkBinSub)
	}
	if m.SmallFrac != nil {
		sesmallFrac := xml.StartElement{Name: xml.Name{Local: "m:smallFrac"}}
		e.EncodeElement(m.SmallFrac, sesmallFrac)
	}
	if m.DispDef != nil {
		sedispDef := xml.StartElement{Name: xml.Name{Local: "m:dispDef"}}
		e.EncodeElement(m.DispDef, sedispDef)
	}
	if m.LMargin != nil {
		selMargin := xml.StartElement{Name: xml.Name{Local: "m:lMargin"}}
		e.EncodeElement(m.LMargin, selMargin)
	}
	if m.RMargin != nil {
		serMargin := xml.StartElement{Name: xml.Name{Local: "m:rMargin"}}
		e.EncodeElement(m.RMargin, serMargin)
	}
	if m.DefJc != nil {
		sedefJc := xml.StartElement{Name: xml.Name{Local: "m:defJc"}}
		e.EncodeElement(m.DefJc, sedefJc)
	}
	if m.PreSp != nil {
		sepreSp := xml.StartElement{Name: xml.Name{Local: "m:preSp"}}
		e.EncodeElement(m.PreSp, sepreSp)
	}
	if m.PostSp != nil {
		sepostSp := xml.StartElement{Name: xml.Name{Local: "m:postSp"}}
		e.EncodeElement(m.PostSp, sepostSp)
	}
	if m.InterSp != nil {
		seinterSp := xml.StartElement{Name: xml.Name{Local: "m:interSp"}}
		e.EncodeElement(m.InterSp, seinterSp)
	}
	if m.IntraSp != nil {
		seintraSp := xml.StartElement{Name: xml.Name{Local: "m:intraSp"}}
		e.EncodeElement(m.IntraSp, seintraSp)
	}
	if m.Choice != nil {
		m.Choice.MarshalXML(e, start)
	}
	if m.IntLim != nil {
		seintLim := xml.StartElement{Name: xml.Name{Local: "m:intLim"}}
		e.EncodeElement(m.IntLim, seintLim)
	}
	if m.NaryLim != nil {
		senaryLim := xml.StartElement{Name: xml.Name{Local: "m:naryLim"}}
		e.EncodeElement(m.NaryLim, senaryLim)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_MathPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_MathPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "mathFont":
				m.MathFont = NewCT_String()
				if err := d.DecodeElement(m.MathFont, &el); err != nil {
					return err
				}
			case "brkBin":
				m.BrkBin = NewCT_BreakBin()
				if err := d.DecodeElement(m.BrkBin, &el); err != nil {
					return err
				}
			case "brkBinSub":
				m.BrkBinSub = NewCT_BreakBinSub()
				if err := d.DecodeElement(m.BrkBinSub, &el); err != nil {
					return err
				}
			case "smallFrac":
				m.SmallFrac = NewCT_OnOff()
				if err := d.DecodeElement(m.SmallFrac, &el); err != nil {
					return err
				}
			case "dispDef":
				m.DispDef = NewCT_OnOff()
				if err := d.DecodeElement(m.DispDef, &el); err != nil {
					return err
				}
			case "lMargin":
				m.LMargin = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.LMargin, &el); err != nil {
					return err
				}
			case "rMargin":
				m.RMargin = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.RMargin, &el); err != nil {
					return err
				}
			case "defJc":
				m.DefJc = NewCT_OMathJc()
				if err := d.DecodeElement(m.DefJc, &el); err != nil {
					return err
				}
			case "preSp":
				m.PreSp = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.PreSp, &el); err != nil {
					return err
				}
			case "postSp":
				m.PostSp = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.PostSp, &el); err != nil {
					return err
				}
			case "interSp":
				m.InterSp = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.InterSp, &el); err != nil {
					return err
				}
			case "intraSp":
				m.IntraSp = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.IntraSp, &el); err != nil {
					return err
				}
			case "wrapIndent":
				m.Choice = NewCT_MathPrChoice()
				if err := d.DecodeElement(&m.Choice.WrapIndent, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "wrapRight":
				m.Choice = NewCT_MathPrChoice()
				if err := d.DecodeElement(&m.Choice.WrapRight, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "intLim":
				m.IntLim = NewCT_LimLoc()
				if err := d.DecodeElement(m.IntLim, &el); err != nil {
					return err
				}
			case "naryLim":
				m.NaryLim = NewCT_LimLoc()
				if err := d.DecodeElement(m.NaryLim, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_MathPr
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_MathPr) Validate() error {
	return m.ValidateWithPath("CT_MathPr")
}
func (m *CT_MathPr) ValidateWithPath(path string) error {
	if m.MathFont != nil {
		if err := m.MathFont.ValidateWithPath(path + "/MathFont"); err != nil {
			return err
		}
	}
	if m.BrkBin != nil {
		if err := m.BrkBin.ValidateWithPath(path + "/BrkBin"); err != nil {
			return err
		}
	}
	if m.BrkBinSub != nil {
		if err := m.BrkBinSub.ValidateWithPath(path + "/BrkBinSub"); err != nil {
			return err
		}
	}
	if m.SmallFrac != nil {
		if err := m.SmallFrac.ValidateWithPath(path + "/SmallFrac"); err != nil {
			return err
		}
	}
	if m.DispDef != nil {
		if err := m.DispDef.ValidateWithPath(path + "/DispDef"); err != nil {
			return err
		}
	}
	if m.LMargin != nil {
		if err := m.LMargin.ValidateWithPath(path + "/LMargin"); err != nil {
			return err
		}
	}
	if m.RMargin != nil {
		if err := m.RMargin.ValidateWithPath(path + "/RMargin"); err != nil {
			return err
		}
	}
	if m.DefJc != nil {
		if err := m.DefJc.ValidateWithPath(path + "/DefJc"); err != nil {
			return err
		}
	}
	if m.PreSp != nil {
		if err := m.PreSp.ValidateWithPath(path + "/PreSp"); err != nil {
			return err
		}
	}
	if m.PostSp != nil {
		if err := m.PostSp.ValidateWithPath(path + "/PostSp"); err != nil {
			return err
		}
	}
	if m.InterSp != nil {
		if err := m.InterSp.ValidateWithPath(path + "/InterSp"); err != nil {
			return err
		}
	}
	if m.IntraSp != nil {
		if err := m.IntraSp.ValidateWithPath(path + "/IntraSp"); err != nil {
			return err
		}
	}
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	if m.IntLim != nil {
		if err := m.IntLim.ValidateWithPath(path + "/IntLim"); err != nil {
			return err
		}
	}
	if m.NaryLim != nil {
		if err := m.NaryLim.ValidateWithPath(path + "/NaryLim"); err != nil {
			return err
		}
	}
	return nil
}
