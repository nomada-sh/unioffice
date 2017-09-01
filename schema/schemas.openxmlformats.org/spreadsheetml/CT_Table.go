// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_Table struct {
	// Table Id
	IdAttr uint32
	// Name
	NameAttr *string
	// Table Name
	DisplayNameAttr string
	// Table Comment
	CommentAttr *string
	// Reference
	RefAttr string
	// Table Type
	TableTypeAttr ST_TableType
	// Header Row Count
	HeaderRowCountAttr *uint32
	// Insert Row Showing
	InsertRowAttr *bool
	// Insert Row Shift
	InsertRowShiftAttr *bool
	// Totals Row Count
	TotalsRowCountAttr *uint32
	// Totals Row Shown
	TotalsRowShownAttr *bool
	// Published
	PublishedAttr *bool
	// Header Row Format Id
	HeaderRowDxfIdAttr *uint32
	// Data Area Format Id
	DataDxfIdAttr *uint32
	// Totals Row Format Id
	TotalsRowDxfIdAttr *uint32
	// Header Row Border Format Id
	HeaderRowBorderDxfIdAttr *uint32
	// Table Border Format Id
	TableBorderDxfIdAttr *uint32
	// Totals Row Border Format Id
	TotalsRowBorderDxfIdAttr *uint32
	// Header Row Style
	HeaderRowCellStyleAttr *string
	// Data Style Name
	DataCellStyleAttr *string
	// Totals Row Style
	TotalsRowCellStyleAttr *string
	// Connection ID
	ConnectionIdAttr *uint32
	// Table AutoFilter
	AutoFilter *CT_AutoFilter
	// Sort State
	SortState *CT_SortState
	// Table Columns
	TableColumns *CT_TableColumns
	// Table Style
	TableStyleInfo *CT_TableStyleInfo
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_Table() *CT_Table {
	ret := &CT_Table{}
	ret.TableColumns = NewCT_TableColumns()
	return ret
}
func (m *CT_Table) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "displayName"},
		Value: fmt.Sprintf("%v", m.DisplayNameAttr)})
	if m.CommentAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "comment"},
			Value: fmt.Sprintf("%v", *m.CommentAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ref"},
		Value: fmt.Sprintf("%v", m.RefAttr)})
	if m.TableTypeAttr != ST_TableTypeUnset {
		attr, err := m.TableTypeAttr.MarshalXMLAttr(xml.Name{Local: "tableType"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.HeaderRowCountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "headerRowCount"},
			Value: fmt.Sprintf("%v", *m.HeaderRowCountAttr)})
	}
	if m.InsertRowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "insertRow"},
			Value: fmt.Sprintf("%v", *m.InsertRowAttr)})
	}
	if m.InsertRowShiftAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "insertRowShift"},
			Value: fmt.Sprintf("%v", *m.InsertRowShiftAttr)})
	}
	if m.TotalsRowCountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "totalsRowCount"},
			Value: fmt.Sprintf("%v", *m.TotalsRowCountAttr)})
	}
	if m.TotalsRowShownAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "totalsRowShown"},
			Value: fmt.Sprintf("%v", *m.TotalsRowShownAttr)})
	}
	if m.PublishedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "published"},
			Value: fmt.Sprintf("%v", *m.PublishedAttr)})
	}
	if m.HeaderRowDxfIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "headerRowDxfId"},
			Value: fmt.Sprintf("%v", *m.HeaderRowDxfIdAttr)})
	}
	if m.DataDxfIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dataDxfId"},
			Value: fmt.Sprintf("%v", *m.DataDxfIdAttr)})
	}
	if m.TotalsRowDxfIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "totalsRowDxfId"},
			Value: fmt.Sprintf("%v", *m.TotalsRowDxfIdAttr)})
	}
	if m.HeaderRowBorderDxfIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "headerRowBorderDxfId"},
			Value: fmt.Sprintf("%v", *m.HeaderRowBorderDxfIdAttr)})
	}
	if m.TableBorderDxfIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "tableBorderDxfId"},
			Value: fmt.Sprintf("%v", *m.TableBorderDxfIdAttr)})
	}
	if m.TotalsRowBorderDxfIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "totalsRowBorderDxfId"},
			Value: fmt.Sprintf("%v", *m.TotalsRowBorderDxfIdAttr)})
	}
	if m.HeaderRowCellStyleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "headerRowCellStyle"},
			Value: fmt.Sprintf("%v", *m.HeaderRowCellStyleAttr)})
	}
	if m.DataCellStyleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dataCellStyle"},
			Value: fmt.Sprintf("%v", *m.DataCellStyleAttr)})
	}
	if m.TotalsRowCellStyleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "totalsRowCellStyle"},
			Value: fmt.Sprintf("%v", *m.TotalsRowCellStyleAttr)})
	}
	if m.ConnectionIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "connectionId"},
			Value: fmt.Sprintf("%v", *m.ConnectionIdAttr)})
	}
	e.EncodeToken(start)
	if m.AutoFilter != nil {
		seautoFilter := xml.StartElement{Name: xml.Name{Local: "x:autoFilter"}}
		e.EncodeElement(m.AutoFilter, seautoFilter)
	}
	if m.SortState != nil {
		sesortState := xml.StartElement{Name: xml.Name{Local: "x:sortState"}}
		e.EncodeElement(m.SortState, sesortState)
	}
	setableColumns := xml.StartElement{Name: xml.Name{Local: "x:tableColumns"}}
	e.EncodeElement(m.TableColumns, setableColumns)
	if m.TableStyleInfo != nil {
		setableStyleInfo := xml.StartElement{Name: xml.Name{Local: "x:tableStyleInfo"}}
		e.EncodeElement(m.TableStyleInfo, setableStyleInfo)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Table) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TableColumns = NewCT_TableColumns()
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.IdAttr = uint32(parsed)
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
		}
		if attr.Name.Local == "displayName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DisplayNameAttr = parsed
		}
		if attr.Name.Local == "comment" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CommentAttr = &parsed
		}
		if attr.Name.Local == "ref" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RefAttr = parsed
		}
		if attr.Name.Local == "tableType" {
			m.TableTypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "headerRowCount" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.HeaderRowCountAttr = &pt
		}
		if attr.Name.Local == "insertRow" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.InsertRowAttr = &parsed
		}
		if attr.Name.Local == "insertRowShift" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.InsertRowShiftAttr = &parsed
		}
		if attr.Name.Local == "totalsRowCount" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.TotalsRowCountAttr = &pt
		}
		if attr.Name.Local == "totalsRowShown" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.TotalsRowShownAttr = &parsed
		}
		if attr.Name.Local == "published" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PublishedAttr = &parsed
		}
		if attr.Name.Local == "headerRowDxfId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.HeaderRowDxfIdAttr = &pt
		}
		if attr.Name.Local == "dataDxfId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DataDxfIdAttr = &pt
		}
		if attr.Name.Local == "totalsRowDxfId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.TotalsRowDxfIdAttr = &pt
		}
		if attr.Name.Local == "headerRowBorderDxfId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.HeaderRowBorderDxfIdAttr = &pt
		}
		if attr.Name.Local == "tableBorderDxfId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.TableBorderDxfIdAttr = &pt
		}
		if attr.Name.Local == "totalsRowBorderDxfId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.TotalsRowBorderDxfIdAttr = &pt
		}
		if attr.Name.Local == "headerRowCellStyle" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.HeaderRowCellStyleAttr = &parsed
		}
		if attr.Name.Local == "dataCellStyle" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DataCellStyleAttr = &parsed
		}
		if attr.Name.Local == "totalsRowCellStyle" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TotalsRowCellStyleAttr = &parsed
		}
		if attr.Name.Local == "connectionId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ConnectionIdAttr = &pt
		}
	}
lCT_Table:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "autoFilter":
				m.AutoFilter = NewCT_AutoFilter()
				if err := d.DecodeElement(m.AutoFilter, &el); err != nil {
					return err
				}
			case "sortState":
				m.SortState = NewCT_SortState()
				if err := d.DecodeElement(m.SortState, &el); err != nil {
					return err
				}
			case "tableColumns":
				if err := d.DecodeElement(m.TableColumns, &el); err != nil {
					return err
				}
			case "tableStyleInfo":
				m.TableStyleInfo = NewCT_TableStyleInfo()
				if err := d.DecodeElement(m.TableStyleInfo, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Table
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Table) Validate() error {
	return m.ValidateWithPath("CT_Table")
}
func (m *CT_Table) ValidateWithPath(path string) error {
	if err := m.TableTypeAttr.ValidateWithPath(path + "/TableTypeAttr"); err != nil {
		return err
	}
	if m.AutoFilter != nil {
		if err := m.AutoFilter.ValidateWithPath(path + "/AutoFilter"); err != nil {
			return err
		}
	}
	if m.SortState != nil {
		if err := m.SortState.ValidateWithPath(path + "/SortState"); err != nil {
			return err
		}
	}
	if err := m.TableColumns.ValidateWithPath(path + "/TableColumns"); err != nil {
		return err
	}
	if m.TableStyleInfo != nil {
		if err := m.TableStyleInfo.ValidateWithPath(path + "/TableStyleInfo"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
