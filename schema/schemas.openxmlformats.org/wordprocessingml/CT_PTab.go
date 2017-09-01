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
)

type CT_PTab struct {
	// Positional Tab Stop Alignment
	AlignmentAttr ST_PTabAlignment
	// Positional Tab Base
	RelativeToAttr ST_PTabRelativeTo
	// Tab Leader Character
	LeaderAttr ST_PTabLeader
}

func NewCT_PTab() *CT_PTab {
	ret := &CT_PTab{}
	ret.AlignmentAttr = ST_PTabAlignment(1)
	ret.RelativeToAttr = ST_PTabRelativeTo(1)
	ret.LeaderAttr = ST_PTabLeader(1)
	return ret
}
func (m *CT_PTab) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	attr, err := m.AlignmentAttr.MarshalXMLAttr(xml.Name{Local: "w:alignment"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	attr, err = m.RelativeToAttr.MarshalXMLAttr(xml.Name{Local: "w:relativeTo"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	attr, err = m.LeaderAttr.MarshalXMLAttr(xml.Name{Local: "w:leader"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_PTab) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.AlignmentAttr = ST_PTabAlignment(1)
	m.RelativeToAttr = ST_PTabRelativeTo(1)
	m.LeaderAttr = ST_PTabLeader(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "alignment" {
			m.AlignmentAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "relativeTo" {
			m.RelativeToAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "leader" {
			m.LeaderAttr.UnmarshalXMLAttr(attr)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_PTab: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_PTab) Validate() error {
	return m.ValidateWithPath("CT_PTab")
}
func (m *CT_PTab) ValidateWithPath(path string) error {
	if m.AlignmentAttr == ST_PTabAlignmentUnset {
		return fmt.Errorf("%s/AlignmentAttr is a mandatory field", path)
	}
	if err := m.AlignmentAttr.ValidateWithPath(path + "/AlignmentAttr"); err != nil {
		return err
	}
	if m.RelativeToAttr == ST_PTabRelativeToUnset {
		return fmt.Errorf("%s/RelativeToAttr is a mandatory field", path)
	}
	if err := m.RelativeToAttr.ValidateWithPath(path + "/RelativeToAttr"); err != nil {
		return err
	}
	if m.LeaderAttr == ST_PTabLeaderUnset {
		return fmt.Errorf("%s/LeaderAttr is a mandatory field", path)
	}
	if err := m.LeaderAttr.ValidateWithPath(path + "/LeaderAttr"); err != nil {
		return err
	}
	return nil
}
