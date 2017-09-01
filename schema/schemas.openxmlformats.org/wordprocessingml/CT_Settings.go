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
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/math"
	"baliance.com/gooxml/schema/schemas.openxmlformats.org/schemaLibrary"
)

type CT_Settings struct {
	// Write Protection
	WriteProtection *CT_WriteProtection
	// Document View Setting
	View *CT_View
	// Magnification Setting
	Zoom *CT_Zoom
	// Remove Personal Information from Document Properties
	RemovePersonalInformation *CT_OnOff
	// Remove Date and Time from Annotations
	RemoveDateAndTime *CT_OnOff
	// Do Not Display Visual Boundary For Header/Footer or Between Pages
	DoNotDisplayPageBoundaries *CT_OnOff
	// Display Background Objects When Displaying Document
	DisplayBackgroundShape *CT_OnOff
	// Print PostScript Codes With Document Text
	PrintPostScriptOverText *CT_OnOff
	// Print Fractional Character Widths
	PrintFractionalCharacterWidth *CT_OnOff
	// Only Print Form Field Content
	PrintFormsData *CT_OnOff
	// Embed TrueType Fonts
	EmbedTrueTypeFonts *CT_OnOff
	// Embed Common System Fonts
	EmbedSystemFonts *CT_OnOff
	// Subset Fonts When Embedding
	SaveSubsetFonts *CT_OnOff
	// Only Save Form Field Content
	SaveFormsData *CT_OnOff
	// Mirror Page Margins
	MirrorMargins *CT_OnOff
	// Align Paragraph and Table Borders with Page Border
	AlignBordersAndEdges *CT_OnOff
	// Page Border Excludes Header
	BordersDoNotSurroundHeader *CT_OnOff
	// Page Border Excludes Footer
	BordersDoNotSurroundFooter *CT_OnOff
	// Position Gutter At Top of Page
	GutterAtTop *CT_OnOff
	// Do Not Display Visual Indication of Spelling Errors
	HideSpellingErrors *CT_OnOff
	// Do Not Display Visual Indication of Grammatical Errors
	HideGrammaticalErrors *CT_OnOff
	// Grammar Checking Settings
	ActiveWritingStyle []*CT_WritingStyle
	// Spelling and Grammatical Checking State
	ProofState *CT_Proof
	// Structured Document Tag Placeholder Text Should be Resaved
	FormsDesign *CT_OnOff
	// Attached Document Template
	AttachedTemplate *CT_Rel
	// Automatically Update Styles From Document Template
	LinkStyles *CT_OnOff
	// Suggested Filtering for List of Document Styles
	StylePaneFormatFilter *CT_StylePaneFilter
	// Suggested Sorting for List of Document Styles
	StylePaneSortMethod *CT_StyleSort
	// Document Classification
	DocumentType *CT_DocType
	// Mail Merge Settings
	MailMerge *CT_MailMerge
	// Visibility of Annotation Types
	RevisionView *CT_TrackChangesView
	// Track Revisions to Document
	TrackRevisions *CT_OnOff
	// Do Not Use Move Syntax When Tracking Revisions
	DoNotTrackMoves *CT_OnOff
	// Do Not Track Formatting Revisions When Tracking Revisions
	DoNotTrackFormatting *CT_OnOff
	// Document Editing Restrictions
	DocumentProtection *CT_DocProtect
	// Allow Automatic Formatting to Override Formatting Protection Settings
	AutoFormatOverride *CT_OnOff
	// Prevent Modification of Themes Part
	StyleLockTheme *CT_OnOff
	// Prevent Replacement of Styles Part
	StyleLockQFSet *CT_OnOff
	// Distance Between Automatic Tab Stops
	DefaultTabStop *CT_TwipsMeasure
	// Automatically Hyphenate Document Contents When Displayed
	AutoHyphenation *CT_OnOff
	// Maximum Number of Consecutively Hyphenated Lines
	ConsecutiveHyphenLimit *CT_DecimalNumber
	// Hyphenation Zone
	HyphenationZone *CT_TwipsMeasure
	// Do Not Hyphenate Words in ALL CAPITAL LETTERS
	DoNotHyphenateCaps *CT_OnOff
	// Show E-Mail Message Header
	ShowEnvelope *CT_OnOff
	// Percentage of Document to Use When Generating Summary
	SummaryLength *CT_DecimalNumberOrPrecent
	// Paragraph Style Applied to Automatically Generated Paragraphs
	ClickAndTypeStyle *CT_String
	// Default Table Style for Newly Inserted Tables
	DefaultTableStyle *CT_String
	// Different Even/Odd Page Headers and Footers
	EvenAndOddHeaders *CT_OnOff
	// Reverse Book Fold Printing
	BookFoldRevPrinting *CT_OnOff
	// Book Fold Printing
	BookFoldPrinting *CT_OnOff
	// Number of Pages Per Booklet
	BookFoldPrintingSheets *CT_DecimalNumber
	// Drawing Grid Horizontal Grid Unit Size
	DrawingGridHorizontalSpacing *CT_TwipsMeasure
	// Drawing Grid Vertical Grid Unit Size
	DrawingGridVerticalSpacing *CT_TwipsMeasure
	// Distance between Horizontal Gridlines
	DisplayHorizontalDrawingGridEvery *CT_DecimalNumber
	// Distance between Vertical Gridlines
	DisplayVerticalDrawingGridEvery *CT_DecimalNumber
	// Do Not Use Margins for Drawing Grid Origin
	DoNotUseMarginsForDrawingGridOrigin *CT_OnOff
	// Drawing Grid Horizontal Origin Point
	DrawingGridHorizontalOrigin *CT_TwipsMeasure
	// Drawing Grid Vertical Origin Point
	DrawingGridVerticalOrigin *CT_TwipsMeasure
	// Do Not Show Visual Indicator For Form Fields
	DoNotShadeFormData *CT_OnOff
	// Never Kern Punctuation Characters
	NoPunctuationKerning *CT_OnOff
	// Character-Level Whitespace Compression
	CharacterSpacingControl *CT_CharacterSpacing
	// Print Two Pages Per Sheet
	PrintTwoOnOne *CT_OnOff
	// Use Strict Kinsoku Rules for Japanese Text
	StrictFirstAndLastChars *CT_OnOff
	// Custom Set of Characters Which Cannot End a Line
	NoLineBreaksAfter *CT_Kinsoku
	// Custom Set Of Characters Which Cannot Begin A Line
	NoLineBreaksBefore *CT_Kinsoku
	// Generate Thumbnail For Document On Save
	SavePreviewPicture *CT_OnOff
	// Do Not Validate Custom XML Markup Against Schemas
	DoNotValidateAgainstSchema *CT_OnOff
	// Allow Saving Document As XML File When Custom XML Markup Is Invalid
	SaveInvalidXml *CT_OnOff
	// Ignore Mixed Content When Validating Custom XML Markup
	IgnoreMixedContent *CT_OnOff
	// Use Custom XML Element Names as Default Placeholder Text
	AlwaysShowPlaceholderText *CT_OnOff
	// Do Not Show Visual Indicator For Invalid Custom XML Markup
	DoNotDemarcateInvalidXml *CT_OnOff
	// Only Save Custom XML Markup
	SaveXmlDataOnly *CT_OnOff
	// Save Document as XML File through Custom XSL Transform
	UseXSLTWhenSaving *CT_OnOff
	// Custom XSL Transform To Use When Saving As XML File
	SaveThroughXslt *CT_SaveThroughXslt
	// Show Visual Indicators for Custom XML Markup Start/End Locations
	ShowXMLTags *CT_OnOff
	// Do Not Mark Custom XML Elements With No Namespace As Invalid
	AlwaysMergeEmptyNamespace *CT_OnOff
	// Automatically Recalculate Fields on Open
	UpdateFields *CT_OnOff
	// Default Properties for VML Objects in Header and Footer
	HdrShapeDefaults *CT_ShapeDefaults
	// Document-Wide Footnote Properties
	FootnotePr *CT_FtnDocProps
	// Document-Wide Endnote Properties
	EndnotePr *CT_EdnDocProps
	// Compatibility Settings
	Compat *CT_Compat
	// Document Variables
	DocVars *CT_DocVars
	// Listing of All Revision Save ID Values
	Rsids  *CT_DocRsids
	MathPr *math.MathPr
	// Attached Custom XML Schema
	AttachedSchema []*CT_String
	// Theme Font Languages
	ThemeFontLang *CT_Language
	// Theme Color Mappings
	ClrSchemeMapping *CT_ColorSchemeMapping
	// Do Not Include Content in Text Boxes, Footnotes, and Endnotes in Document Statistics
	DoNotIncludeSubdocsInStats *CT_OnOff
	// Do Not Automatically Compress Images
	DoNotAutoCompressPictures *CT_OnOff
	// Upgrade Document on Open
	ForceUpgrade *CT_Empty
	// Caption Settings
	Captions *CT_Captions
	// Freeze Document Layout
	ReadModeInkLockDown *CT_ReadingModeInkLockDown
	// Supplementary Smart Tag Information
	SmartTagType  []*CT_SmartTagType
	SchemaLibrary *schemaLibrary.SchemaLibrary
	// Default Properties for VML Objects in Main Document
	ShapeDefaults *CT_ShapeDefaults
	// Remove Smart Tags When Saving
	DoNotEmbedSmartTags *CT_OnOff
	// Radix Point for Field Code Evaluation
	DecimalSymbol *CT_String
	// List Separator for Field Code Evaluation
	ListSeparator *CT_String
}

func NewCT_Settings() *CT_Settings {
	ret := &CT_Settings{}
	return ret
}
func (m *CT_Settings) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.WriteProtection != nil {
		sewriteProtection := xml.StartElement{Name: xml.Name{Local: "w:writeProtection"}}
		e.EncodeElement(m.WriteProtection, sewriteProtection)
	}
	if m.View != nil {
		seview := xml.StartElement{Name: xml.Name{Local: "w:view"}}
		e.EncodeElement(m.View, seview)
	}
	if m.Zoom != nil {
		sezoom := xml.StartElement{Name: xml.Name{Local: "w:zoom"}}
		e.EncodeElement(m.Zoom, sezoom)
	}
	if m.RemovePersonalInformation != nil {
		seremovePersonalInformation := xml.StartElement{Name: xml.Name{Local: "w:removePersonalInformation"}}
		e.EncodeElement(m.RemovePersonalInformation, seremovePersonalInformation)
	}
	if m.RemoveDateAndTime != nil {
		seremoveDateAndTime := xml.StartElement{Name: xml.Name{Local: "w:removeDateAndTime"}}
		e.EncodeElement(m.RemoveDateAndTime, seremoveDateAndTime)
	}
	if m.DoNotDisplayPageBoundaries != nil {
		sedoNotDisplayPageBoundaries := xml.StartElement{Name: xml.Name{Local: "w:doNotDisplayPageBoundaries"}}
		e.EncodeElement(m.DoNotDisplayPageBoundaries, sedoNotDisplayPageBoundaries)
	}
	if m.DisplayBackgroundShape != nil {
		sedisplayBackgroundShape := xml.StartElement{Name: xml.Name{Local: "w:displayBackgroundShape"}}
		e.EncodeElement(m.DisplayBackgroundShape, sedisplayBackgroundShape)
	}
	if m.PrintPostScriptOverText != nil {
		seprintPostScriptOverText := xml.StartElement{Name: xml.Name{Local: "w:printPostScriptOverText"}}
		e.EncodeElement(m.PrintPostScriptOverText, seprintPostScriptOverText)
	}
	if m.PrintFractionalCharacterWidth != nil {
		seprintFractionalCharacterWidth := xml.StartElement{Name: xml.Name{Local: "w:printFractionalCharacterWidth"}}
		e.EncodeElement(m.PrintFractionalCharacterWidth, seprintFractionalCharacterWidth)
	}
	if m.PrintFormsData != nil {
		seprintFormsData := xml.StartElement{Name: xml.Name{Local: "w:printFormsData"}}
		e.EncodeElement(m.PrintFormsData, seprintFormsData)
	}
	if m.EmbedTrueTypeFonts != nil {
		seembedTrueTypeFonts := xml.StartElement{Name: xml.Name{Local: "w:embedTrueTypeFonts"}}
		e.EncodeElement(m.EmbedTrueTypeFonts, seembedTrueTypeFonts)
	}
	if m.EmbedSystemFonts != nil {
		seembedSystemFonts := xml.StartElement{Name: xml.Name{Local: "w:embedSystemFonts"}}
		e.EncodeElement(m.EmbedSystemFonts, seembedSystemFonts)
	}
	if m.SaveSubsetFonts != nil {
		sesaveSubsetFonts := xml.StartElement{Name: xml.Name{Local: "w:saveSubsetFonts"}}
		e.EncodeElement(m.SaveSubsetFonts, sesaveSubsetFonts)
	}
	if m.SaveFormsData != nil {
		sesaveFormsData := xml.StartElement{Name: xml.Name{Local: "w:saveFormsData"}}
		e.EncodeElement(m.SaveFormsData, sesaveFormsData)
	}
	if m.MirrorMargins != nil {
		semirrorMargins := xml.StartElement{Name: xml.Name{Local: "w:mirrorMargins"}}
		e.EncodeElement(m.MirrorMargins, semirrorMargins)
	}
	if m.AlignBordersAndEdges != nil {
		sealignBordersAndEdges := xml.StartElement{Name: xml.Name{Local: "w:alignBordersAndEdges"}}
		e.EncodeElement(m.AlignBordersAndEdges, sealignBordersAndEdges)
	}
	if m.BordersDoNotSurroundHeader != nil {
		sebordersDoNotSurroundHeader := xml.StartElement{Name: xml.Name{Local: "w:bordersDoNotSurroundHeader"}}
		e.EncodeElement(m.BordersDoNotSurroundHeader, sebordersDoNotSurroundHeader)
	}
	if m.BordersDoNotSurroundFooter != nil {
		sebordersDoNotSurroundFooter := xml.StartElement{Name: xml.Name{Local: "w:bordersDoNotSurroundFooter"}}
		e.EncodeElement(m.BordersDoNotSurroundFooter, sebordersDoNotSurroundFooter)
	}
	if m.GutterAtTop != nil {
		segutterAtTop := xml.StartElement{Name: xml.Name{Local: "w:gutterAtTop"}}
		e.EncodeElement(m.GutterAtTop, segutterAtTop)
	}
	if m.HideSpellingErrors != nil {
		sehideSpellingErrors := xml.StartElement{Name: xml.Name{Local: "w:hideSpellingErrors"}}
		e.EncodeElement(m.HideSpellingErrors, sehideSpellingErrors)
	}
	if m.HideGrammaticalErrors != nil {
		sehideGrammaticalErrors := xml.StartElement{Name: xml.Name{Local: "w:hideGrammaticalErrors"}}
		e.EncodeElement(m.HideGrammaticalErrors, sehideGrammaticalErrors)
	}
	if m.ActiveWritingStyle != nil {
		seactiveWritingStyle := xml.StartElement{Name: xml.Name{Local: "w:activeWritingStyle"}}
		e.EncodeElement(m.ActiveWritingStyle, seactiveWritingStyle)
	}
	if m.ProofState != nil {
		seproofState := xml.StartElement{Name: xml.Name{Local: "w:proofState"}}
		e.EncodeElement(m.ProofState, seproofState)
	}
	if m.FormsDesign != nil {
		seformsDesign := xml.StartElement{Name: xml.Name{Local: "w:formsDesign"}}
		e.EncodeElement(m.FormsDesign, seformsDesign)
	}
	if m.AttachedTemplate != nil {
		seattachedTemplate := xml.StartElement{Name: xml.Name{Local: "w:attachedTemplate"}}
		e.EncodeElement(m.AttachedTemplate, seattachedTemplate)
	}
	if m.LinkStyles != nil {
		selinkStyles := xml.StartElement{Name: xml.Name{Local: "w:linkStyles"}}
		e.EncodeElement(m.LinkStyles, selinkStyles)
	}
	if m.StylePaneFormatFilter != nil {
		sestylePaneFormatFilter := xml.StartElement{Name: xml.Name{Local: "w:stylePaneFormatFilter"}}
		e.EncodeElement(m.StylePaneFormatFilter, sestylePaneFormatFilter)
	}
	if m.StylePaneSortMethod != nil {
		sestylePaneSortMethod := xml.StartElement{Name: xml.Name{Local: "w:stylePaneSortMethod"}}
		e.EncodeElement(m.StylePaneSortMethod, sestylePaneSortMethod)
	}
	if m.DocumentType != nil {
		sedocumentType := xml.StartElement{Name: xml.Name{Local: "w:documentType"}}
		e.EncodeElement(m.DocumentType, sedocumentType)
	}
	if m.MailMerge != nil {
		semailMerge := xml.StartElement{Name: xml.Name{Local: "w:mailMerge"}}
		e.EncodeElement(m.MailMerge, semailMerge)
	}
	if m.RevisionView != nil {
		serevisionView := xml.StartElement{Name: xml.Name{Local: "w:revisionView"}}
		e.EncodeElement(m.RevisionView, serevisionView)
	}
	if m.TrackRevisions != nil {
		setrackRevisions := xml.StartElement{Name: xml.Name{Local: "w:trackRevisions"}}
		e.EncodeElement(m.TrackRevisions, setrackRevisions)
	}
	if m.DoNotTrackMoves != nil {
		sedoNotTrackMoves := xml.StartElement{Name: xml.Name{Local: "w:doNotTrackMoves"}}
		e.EncodeElement(m.DoNotTrackMoves, sedoNotTrackMoves)
	}
	if m.DoNotTrackFormatting != nil {
		sedoNotTrackFormatting := xml.StartElement{Name: xml.Name{Local: "w:doNotTrackFormatting"}}
		e.EncodeElement(m.DoNotTrackFormatting, sedoNotTrackFormatting)
	}
	if m.DocumentProtection != nil {
		sedocumentProtection := xml.StartElement{Name: xml.Name{Local: "w:documentProtection"}}
		e.EncodeElement(m.DocumentProtection, sedocumentProtection)
	}
	if m.AutoFormatOverride != nil {
		seautoFormatOverride := xml.StartElement{Name: xml.Name{Local: "w:autoFormatOverride"}}
		e.EncodeElement(m.AutoFormatOverride, seautoFormatOverride)
	}
	if m.StyleLockTheme != nil {
		sestyleLockTheme := xml.StartElement{Name: xml.Name{Local: "w:styleLockTheme"}}
		e.EncodeElement(m.StyleLockTheme, sestyleLockTheme)
	}
	if m.StyleLockQFSet != nil {
		sestyleLockQFSet := xml.StartElement{Name: xml.Name{Local: "w:styleLockQFSet"}}
		e.EncodeElement(m.StyleLockQFSet, sestyleLockQFSet)
	}
	if m.DefaultTabStop != nil {
		sedefaultTabStop := xml.StartElement{Name: xml.Name{Local: "w:defaultTabStop"}}
		e.EncodeElement(m.DefaultTabStop, sedefaultTabStop)
	}
	if m.AutoHyphenation != nil {
		seautoHyphenation := xml.StartElement{Name: xml.Name{Local: "w:autoHyphenation"}}
		e.EncodeElement(m.AutoHyphenation, seautoHyphenation)
	}
	if m.ConsecutiveHyphenLimit != nil {
		seconsecutiveHyphenLimit := xml.StartElement{Name: xml.Name{Local: "w:consecutiveHyphenLimit"}}
		e.EncodeElement(m.ConsecutiveHyphenLimit, seconsecutiveHyphenLimit)
	}
	if m.HyphenationZone != nil {
		sehyphenationZone := xml.StartElement{Name: xml.Name{Local: "w:hyphenationZone"}}
		e.EncodeElement(m.HyphenationZone, sehyphenationZone)
	}
	if m.DoNotHyphenateCaps != nil {
		sedoNotHyphenateCaps := xml.StartElement{Name: xml.Name{Local: "w:doNotHyphenateCaps"}}
		e.EncodeElement(m.DoNotHyphenateCaps, sedoNotHyphenateCaps)
	}
	if m.ShowEnvelope != nil {
		seshowEnvelope := xml.StartElement{Name: xml.Name{Local: "w:showEnvelope"}}
		e.EncodeElement(m.ShowEnvelope, seshowEnvelope)
	}
	if m.SummaryLength != nil {
		sesummaryLength := xml.StartElement{Name: xml.Name{Local: "w:summaryLength"}}
		e.EncodeElement(m.SummaryLength, sesummaryLength)
	}
	if m.ClickAndTypeStyle != nil {
		seclickAndTypeStyle := xml.StartElement{Name: xml.Name{Local: "w:clickAndTypeStyle"}}
		e.EncodeElement(m.ClickAndTypeStyle, seclickAndTypeStyle)
	}
	if m.DefaultTableStyle != nil {
		sedefaultTableStyle := xml.StartElement{Name: xml.Name{Local: "w:defaultTableStyle"}}
		e.EncodeElement(m.DefaultTableStyle, sedefaultTableStyle)
	}
	if m.EvenAndOddHeaders != nil {
		seevenAndOddHeaders := xml.StartElement{Name: xml.Name{Local: "w:evenAndOddHeaders"}}
		e.EncodeElement(m.EvenAndOddHeaders, seevenAndOddHeaders)
	}
	if m.BookFoldRevPrinting != nil {
		sebookFoldRevPrinting := xml.StartElement{Name: xml.Name{Local: "w:bookFoldRevPrinting"}}
		e.EncodeElement(m.BookFoldRevPrinting, sebookFoldRevPrinting)
	}
	if m.BookFoldPrinting != nil {
		sebookFoldPrinting := xml.StartElement{Name: xml.Name{Local: "w:bookFoldPrinting"}}
		e.EncodeElement(m.BookFoldPrinting, sebookFoldPrinting)
	}
	if m.BookFoldPrintingSheets != nil {
		sebookFoldPrintingSheets := xml.StartElement{Name: xml.Name{Local: "w:bookFoldPrintingSheets"}}
		e.EncodeElement(m.BookFoldPrintingSheets, sebookFoldPrintingSheets)
	}
	if m.DrawingGridHorizontalSpacing != nil {
		sedrawingGridHorizontalSpacing := xml.StartElement{Name: xml.Name{Local: "w:drawingGridHorizontalSpacing"}}
		e.EncodeElement(m.DrawingGridHorizontalSpacing, sedrawingGridHorizontalSpacing)
	}
	if m.DrawingGridVerticalSpacing != nil {
		sedrawingGridVerticalSpacing := xml.StartElement{Name: xml.Name{Local: "w:drawingGridVerticalSpacing"}}
		e.EncodeElement(m.DrawingGridVerticalSpacing, sedrawingGridVerticalSpacing)
	}
	if m.DisplayHorizontalDrawingGridEvery != nil {
		sedisplayHorizontalDrawingGridEvery := xml.StartElement{Name: xml.Name{Local: "w:displayHorizontalDrawingGridEvery"}}
		e.EncodeElement(m.DisplayHorizontalDrawingGridEvery, sedisplayHorizontalDrawingGridEvery)
	}
	if m.DisplayVerticalDrawingGridEvery != nil {
		sedisplayVerticalDrawingGridEvery := xml.StartElement{Name: xml.Name{Local: "w:displayVerticalDrawingGridEvery"}}
		e.EncodeElement(m.DisplayVerticalDrawingGridEvery, sedisplayVerticalDrawingGridEvery)
	}
	if m.DoNotUseMarginsForDrawingGridOrigin != nil {
		sedoNotUseMarginsForDrawingGridOrigin := xml.StartElement{Name: xml.Name{Local: "w:doNotUseMarginsForDrawingGridOrigin"}}
		e.EncodeElement(m.DoNotUseMarginsForDrawingGridOrigin, sedoNotUseMarginsForDrawingGridOrigin)
	}
	if m.DrawingGridHorizontalOrigin != nil {
		sedrawingGridHorizontalOrigin := xml.StartElement{Name: xml.Name{Local: "w:drawingGridHorizontalOrigin"}}
		e.EncodeElement(m.DrawingGridHorizontalOrigin, sedrawingGridHorizontalOrigin)
	}
	if m.DrawingGridVerticalOrigin != nil {
		sedrawingGridVerticalOrigin := xml.StartElement{Name: xml.Name{Local: "w:drawingGridVerticalOrigin"}}
		e.EncodeElement(m.DrawingGridVerticalOrigin, sedrawingGridVerticalOrigin)
	}
	if m.DoNotShadeFormData != nil {
		sedoNotShadeFormData := xml.StartElement{Name: xml.Name{Local: "w:doNotShadeFormData"}}
		e.EncodeElement(m.DoNotShadeFormData, sedoNotShadeFormData)
	}
	if m.NoPunctuationKerning != nil {
		senoPunctuationKerning := xml.StartElement{Name: xml.Name{Local: "w:noPunctuationKerning"}}
		e.EncodeElement(m.NoPunctuationKerning, senoPunctuationKerning)
	}
	if m.CharacterSpacingControl != nil {
		secharacterSpacingControl := xml.StartElement{Name: xml.Name{Local: "w:characterSpacingControl"}}
		e.EncodeElement(m.CharacterSpacingControl, secharacterSpacingControl)
	}
	if m.PrintTwoOnOne != nil {
		seprintTwoOnOne := xml.StartElement{Name: xml.Name{Local: "w:printTwoOnOne"}}
		e.EncodeElement(m.PrintTwoOnOne, seprintTwoOnOne)
	}
	if m.StrictFirstAndLastChars != nil {
		sestrictFirstAndLastChars := xml.StartElement{Name: xml.Name{Local: "w:strictFirstAndLastChars"}}
		e.EncodeElement(m.StrictFirstAndLastChars, sestrictFirstAndLastChars)
	}
	if m.NoLineBreaksAfter != nil {
		senoLineBreaksAfter := xml.StartElement{Name: xml.Name{Local: "w:noLineBreaksAfter"}}
		e.EncodeElement(m.NoLineBreaksAfter, senoLineBreaksAfter)
	}
	if m.NoLineBreaksBefore != nil {
		senoLineBreaksBefore := xml.StartElement{Name: xml.Name{Local: "w:noLineBreaksBefore"}}
		e.EncodeElement(m.NoLineBreaksBefore, senoLineBreaksBefore)
	}
	if m.SavePreviewPicture != nil {
		sesavePreviewPicture := xml.StartElement{Name: xml.Name{Local: "w:savePreviewPicture"}}
		e.EncodeElement(m.SavePreviewPicture, sesavePreviewPicture)
	}
	if m.DoNotValidateAgainstSchema != nil {
		sedoNotValidateAgainstSchema := xml.StartElement{Name: xml.Name{Local: "w:doNotValidateAgainstSchema"}}
		e.EncodeElement(m.DoNotValidateAgainstSchema, sedoNotValidateAgainstSchema)
	}
	if m.SaveInvalidXml != nil {
		sesaveInvalidXml := xml.StartElement{Name: xml.Name{Local: "w:saveInvalidXml"}}
		e.EncodeElement(m.SaveInvalidXml, sesaveInvalidXml)
	}
	if m.IgnoreMixedContent != nil {
		seignoreMixedContent := xml.StartElement{Name: xml.Name{Local: "w:ignoreMixedContent"}}
		e.EncodeElement(m.IgnoreMixedContent, seignoreMixedContent)
	}
	if m.AlwaysShowPlaceholderText != nil {
		sealwaysShowPlaceholderText := xml.StartElement{Name: xml.Name{Local: "w:alwaysShowPlaceholderText"}}
		e.EncodeElement(m.AlwaysShowPlaceholderText, sealwaysShowPlaceholderText)
	}
	if m.DoNotDemarcateInvalidXml != nil {
		sedoNotDemarcateInvalidXml := xml.StartElement{Name: xml.Name{Local: "w:doNotDemarcateInvalidXml"}}
		e.EncodeElement(m.DoNotDemarcateInvalidXml, sedoNotDemarcateInvalidXml)
	}
	if m.SaveXmlDataOnly != nil {
		sesaveXmlDataOnly := xml.StartElement{Name: xml.Name{Local: "w:saveXmlDataOnly"}}
		e.EncodeElement(m.SaveXmlDataOnly, sesaveXmlDataOnly)
	}
	if m.UseXSLTWhenSaving != nil {
		seuseXSLTWhenSaving := xml.StartElement{Name: xml.Name{Local: "w:useXSLTWhenSaving"}}
		e.EncodeElement(m.UseXSLTWhenSaving, seuseXSLTWhenSaving)
	}
	if m.SaveThroughXslt != nil {
		sesaveThroughXslt := xml.StartElement{Name: xml.Name{Local: "w:saveThroughXslt"}}
		e.EncodeElement(m.SaveThroughXslt, sesaveThroughXslt)
	}
	if m.ShowXMLTags != nil {
		seshowXMLTags := xml.StartElement{Name: xml.Name{Local: "w:showXMLTags"}}
		e.EncodeElement(m.ShowXMLTags, seshowXMLTags)
	}
	if m.AlwaysMergeEmptyNamespace != nil {
		sealwaysMergeEmptyNamespace := xml.StartElement{Name: xml.Name{Local: "w:alwaysMergeEmptyNamespace"}}
		e.EncodeElement(m.AlwaysMergeEmptyNamespace, sealwaysMergeEmptyNamespace)
	}
	if m.UpdateFields != nil {
		seupdateFields := xml.StartElement{Name: xml.Name{Local: "w:updateFields"}}
		e.EncodeElement(m.UpdateFields, seupdateFields)
	}
	if m.HdrShapeDefaults != nil {
		sehdrShapeDefaults := xml.StartElement{Name: xml.Name{Local: "w:hdrShapeDefaults"}}
		e.EncodeElement(m.HdrShapeDefaults, sehdrShapeDefaults)
	}
	if m.FootnotePr != nil {
		sefootnotePr := xml.StartElement{Name: xml.Name{Local: "w:footnotePr"}}
		e.EncodeElement(m.FootnotePr, sefootnotePr)
	}
	if m.EndnotePr != nil {
		seendnotePr := xml.StartElement{Name: xml.Name{Local: "w:endnotePr"}}
		e.EncodeElement(m.EndnotePr, seendnotePr)
	}
	if m.Compat != nil {
		secompat := xml.StartElement{Name: xml.Name{Local: "w:compat"}}
		e.EncodeElement(m.Compat, secompat)
	}
	if m.DocVars != nil {
		sedocVars := xml.StartElement{Name: xml.Name{Local: "w:docVars"}}
		e.EncodeElement(m.DocVars, sedocVars)
	}
	if m.Rsids != nil {
		sersids := xml.StartElement{Name: xml.Name{Local: "w:rsids"}}
		e.EncodeElement(m.Rsids, sersids)
	}
	if m.MathPr != nil {
		semathPr := xml.StartElement{Name: xml.Name{Local: "m:mathPr"}}
		e.EncodeElement(m.MathPr, semathPr)
	}
	if m.AttachedSchema != nil {
		seattachedSchema := xml.StartElement{Name: xml.Name{Local: "w:attachedSchema"}}
		e.EncodeElement(m.AttachedSchema, seattachedSchema)
	}
	if m.ThemeFontLang != nil {
		sethemeFontLang := xml.StartElement{Name: xml.Name{Local: "w:themeFontLang"}}
		e.EncodeElement(m.ThemeFontLang, sethemeFontLang)
	}
	if m.ClrSchemeMapping != nil {
		seclrSchemeMapping := xml.StartElement{Name: xml.Name{Local: "w:clrSchemeMapping"}}
		e.EncodeElement(m.ClrSchemeMapping, seclrSchemeMapping)
	}
	if m.DoNotIncludeSubdocsInStats != nil {
		sedoNotIncludeSubdocsInStats := xml.StartElement{Name: xml.Name{Local: "w:doNotIncludeSubdocsInStats"}}
		e.EncodeElement(m.DoNotIncludeSubdocsInStats, sedoNotIncludeSubdocsInStats)
	}
	if m.DoNotAutoCompressPictures != nil {
		sedoNotAutoCompressPictures := xml.StartElement{Name: xml.Name{Local: "w:doNotAutoCompressPictures"}}
		e.EncodeElement(m.DoNotAutoCompressPictures, sedoNotAutoCompressPictures)
	}
	if m.ForceUpgrade != nil {
		seforceUpgrade := xml.StartElement{Name: xml.Name{Local: "w:forceUpgrade"}}
		e.EncodeElement(m.ForceUpgrade, seforceUpgrade)
	}
	if m.Captions != nil {
		secaptions := xml.StartElement{Name: xml.Name{Local: "w:captions"}}
		e.EncodeElement(m.Captions, secaptions)
	}
	if m.ReadModeInkLockDown != nil {
		sereadModeInkLockDown := xml.StartElement{Name: xml.Name{Local: "w:readModeInkLockDown"}}
		e.EncodeElement(m.ReadModeInkLockDown, sereadModeInkLockDown)
	}
	if m.SmartTagType != nil {
		sesmartTagType := xml.StartElement{Name: xml.Name{Local: "w:smartTagType"}}
		e.EncodeElement(m.SmartTagType, sesmartTagType)
	}
	if m.SchemaLibrary != nil {
		seschemaLibrary := xml.StartElement{Name: xml.Name{Local: "ma:schemaLibrary"}}
		e.EncodeElement(m.SchemaLibrary, seschemaLibrary)
	}
	if m.ShapeDefaults != nil {
		seshapeDefaults := xml.StartElement{Name: xml.Name{Local: "w:shapeDefaults"}}
		e.EncodeElement(m.ShapeDefaults, seshapeDefaults)
	}
	if m.DoNotEmbedSmartTags != nil {
		sedoNotEmbedSmartTags := xml.StartElement{Name: xml.Name{Local: "w:doNotEmbedSmartTags"}}
		e.EncodeElement(m.DoNotEmbedSmartTags, sedoNotEmbedSmartTags)
	}
	if m.DecimalSymbol != nil {
		sedecimalSymbol := xml.StartElement{Name: xml.Name{Local: "w:decimalSymbol"}}
		e.EncodeElement(m.DecimalSymbol, sedecimalSymbol)
	}
	if m.ListSeparator != nil {
		selistSeparator := xml.StartElement{Name: xml.Name{Local: "w:listSeparator"}}
		e.EncodeElement(m.ListSeparator, selistSeparator)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Settings) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Settings:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "writeProtection":
				m.WriteProtection = NewCT_WriteProtection()
				if err := d.DecodeElement(m.WriteProtection, &el); err != nil {
					return err
				}
			case "view":
				m.View = NewCT_View()
				if err := d.DecodeElement(m.View, &el); err != nil {
					return err
				}
			case "zoom":
				m.Zoom = NewCT_Zoom()
				if err := d.DecodeElement(m.Zoom, &el); err != nil {
					return err
				}
			case "removePersonalInformation":
				m.RemovePersonalInformation = NewCT_OnOff()
				if err := d.DecodeElement(m.RemovePersonalInformation, &el); err != nil {
					return err
				}
			case "removeDateAndTime":
				m.RemoveDateAndTime = NewCT_OnOff()
				if err := d.DecodeElement(m.RemoveDateAndTime, &el); err != nil {
					return err
				}
			case "doNotDisplayPageBoundaries":
				m.DoNotDisplayPageBoundaries = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotDisplayPageBoundaries, &el); err != nil {
					return err
				}
			case "displayBackgroundShape":
				m.DisplayBackgroundShape = NewCT_OnOff()
				if err := d.DecodeElement(m.DisplayBackgroundShape, &el); err != nil {
					return err
				}
			case "printPostScriptOverText":
				m.PrintPostScriptOverText = NewCT_OnOff()
				if err := d.DecodeElement(m.PrintPostScriptOverText, &el); err != nil {
					return err
				}
			case "printFractionalCharacterWidth":
				m.PrintFractionalCharacterWidth = NewCT_OnOff()
				if err := d.DecodeElement(m.PrintFractionalCharacterWidth, &el); err != nil {
					return err
				}
			case "printFormsData":
				m.PrintFormsData = NewCT_OnOff()
				if err := d.DecodeElement(m.PrintFormsData, &el); err != nil {
					return err
				}
			case "embedTrueTypeFonts":
				m.EmbedTrueTypeFonts = NewCT_OnOff()
				if err := d.DecodeElement(m.EmbedTrueTypeFonts, &el); err != nil {
					return err
				}
			case "embedSystemFonts":
				m.EmbedSystemFonts = NewCT_OnOff()
				if err := d.DecodeElement(m.EmbedSystemFonts, &el); err != nil {
					return err
				}
			case "saveSubsetFonts":
				m.SaveSubsetFonts = NewCT_OnOff()
				if err := d.DecodeElement(m.SaveSubsetFonts, &el); err != nil {
					return err
				}
			case "saveFormsData":
				m.SaveFormsData = NewCT_OnOff()
				if err := d.DecodeElement(m.SaveFormsData, &el); err != nil {
					return err
				}
			case "mirrorMargins":
				m.MirrorMargins = NewCT_OnOff()
				if err := d.DecodeElement(m.MirrorMargins, &el); err != nil {
					return err
				}
			case "alignBordersAndEdges":
				m.AlignBordersAndEdges = NewCT_OnOff()
				if err := d.DecodeElement(m.AlignBordersAndEdges, &el); err != nil {
					return err
				}
			case "bordersDoNotSurroundHeader":
				m.BordersDoNotSurroundHeader = NewCT_OnOff()
				if err := d.DecodeElement(m.BordersDoNotSurroundHeader, &el); err != nil {
					return err
				}
			case "bordersDoNotSurroundFooter":
				m.BordersDoNotSurroundFooter = NewCT_OnOff()
				if err := d.DecodeElement(m.BordersDoNotSurroundFooter, &el); err != nil {
					return err
				}
			case "gutterAtTop":
				m.GutterAtTop = NewCT_OnOff()
				if err := d.DecodeElement(m.GutterAtTop, &el); err != nil {
					return err
				}
			case "hideSpellingErrors":
				m.HideSpellingErrors = NewCT_OnOff()
				if err := d.DecodeElement(m.HideSpellingErrors, &el); err != nil {
					return err
				}
			case "hideGrammaticalErrors":
				m.HideGrammaticalErrors = NewCT_OnOff()
				if err := d.DecodeElement(m.HideGrammaticalErrors, &el); err != nil {
					return err
				}
			case "activeWritingStyle":
				tmp := NewCT_WritingStyle()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ActiveWritingStyle = append(m.ActiveWritingStyle, tmp)
			case "proofState":
				m.ProofState = NewCT_Proof()
				if err := d.DecodeElement(m.ProofState, &el); err != nil {
					return err
				}
			case "formsDesign":
				m.FormsDesign = NewCT_OnOff()
				if err := d.DecodeElement(m.FormsDesign, &el); err != nil {
					return err
				}
			case "attachedTemplate":
				m.AttachedTemplate = NewCT_Rel()
				if err := d.DecodeElement(m.AttachedTemplate, &el); err != nil {
					return err
				}
			case "linkStyles":
				m.LinkStyles = NewCT_OnOff()
				if err := d.DecodeElement(m.LinkStyles, &el); err != nil {
					return err
				}
			case "stylePaneFormatFilter":
				m.StylePaneFormatFilter = NewCT_StylePaneFilter()
				if err := d.DecodeElement(m.StylePaneFormatFilter, &el); err != nil {
					return err
				}
			case "stylePaneSortMethod":
				m.StylePaneSortMethod = NewCT_StyleSort()
				if err := d.DecodeElement(m.StylePaneSortMethod, &el); err != nil {
					return err
				}
			case "documentType":
				m.DocumentType = NewCT_DocType()
				if err := d.DecodeElement(m.DocumentType, &el); err != nil {
					return err
				}
			case "mailMerge":
				m.MailMerge = NewCT_MailMerge()
				if err := d.DecodeElement(m.MailMerge, &el); err != nil {
					return err
				}
			case "revisionView":
				m.RevisionView = NewCT_TrackChangesView()
				if err := d.DecodeElement(m.RevisionView, &el); err != nil {
					return err
				}
			case "trackRevisions":
				m.TrackRevisions = NewCT_OnOff()
				if err := d.DecodeElement(m.TrackRevisions, &el); err != nil {
					return err
				}
			case "doNotTrackMoves":
				m.DoNotTrackMoves = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotTrackMoves, &el); err != nil {
					return err
				}
			case "doNotTrackFormatting":
				m.DoNotTrackFormatting = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotTrackFormatting, &el); err != nil {
					return err
				}
			case "documentProtection":
				m.DocumentProtection = NewCT_DocProtect()
				if err := d.DecodeElement(m.DocumentProtection, &el); err != nil {
					return err
				}
			case "autoFormatOverride":
				m.AutoFormatOverride = NewCT_OnOff()
				if err := d.DecodeElement(m.AutoFormatOverride, &el); err != nil {
					return err
				}
			case "styleLockTheme":
				m.StyleLockTheme = NewCT_OnOff()
				if err := d.DecodeElement(m.StyleLockTheme, &el); err != nil {
					return err
				}
			case "styleLockQFSet":
				m.StyleLockQFSet = NewCT_OnOff()
				if err := d.DecodeElement(m.StyleLockQFSet, &el); err != nil {
					return err
				}
			case "defaultTabStop":
				m.DefaultTabStop = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.DefaultTabStop, &el); err != nil {
					return err
				}
			case "autoHyphenation":
				m.AutoHyphenation = NewCT_OnOff()
				if err := d.DecodeElement(m.AutoHyphenation, &el); err != nil {
					return err
				}
			case "consecutiveHyphenLimit":
				m.ConsecutiveHyphenLimit = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.ConsecutiveHyphenLimit, &el); err != nil {
					return err
				}
			case "hyphenationZone":
				m.HyphenationZone = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.HyphenationZone, &el); err != nil {
					return err
				}
			case "doNotHyphenateCaps":
				m.DoNotHyphenateCaps = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotHyphenateCaps, &el); err != nil {
					return err
				}
			case "showEnvelope":
				m.ShowEnvelope = NewCT_OnOff()
				if err := d.DecodeElement(m.ShowEnvelope, &el); err != nil {
					return err
				}
			case "summaryLength":
				m.SummaryLength = NewCT_DecimalNumberOrPrecent()
				if err := d.DecodeElement(m.SummaryLength, &el); err != nil {
					return err
				}
			case "clickAndTypeStyle":
				m.ClickAndTypeStyle = NewCT_String()
				if err := d.DecodeElement(m.ClickAndTypeStyle, &el); err != nil {
					return err
				}
			case "defaultTableStyle":
				m.DefaultTableStyle = NewCT_String()
				if err := d.DecodeElement(m.DefaultTableStyle, &el); err != nil {
					return err
				}
			case "evenAndOddHeaders":
				m.EvenAndOddHeaders = NewCT_OnOff()
				if err := d.DecodeElement(m.EvenAndOddHeaders, &el); err != nil {
					return err
				}
			case "bookFoldRevPrinting":
				m.BookFoldRevPrinting = NewCT_OnOff()
				if err := d.DecodeElement(m.BookFoldRevPrinting, &el); err != nil {
					return err
				}
			case "bookFoldPrinting":
				m.BookFoldPrinting = NewCT_OnOff()
				if err := d.DecodeElement(m.BookFoldPrinting, &el); err != nil {
					return err
				}
			case "bookFoldPrintingSheets":
				m.BookFoldPrintingSheets = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.BookFoldPrintingSheets, &el); err != nil {
					return err
				}
			case "drawingGridHorizontalSpacing":
				m.DrawingGridHorizontalSpacing = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.DrawingGridHorizontalSpacing, &el); err != nil {
					return err
				}
			case "drawingGridVerticalSpacing":
				m.DrawingGridVerticalSpacing = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.DrawingGridVerticalSpacing, &el); err != nil {
					return err
				}
			case "displayHorizontalDrawingGridEvery":
				m.DisplayHorizontalDrawingGridEvery = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.DisplayHorizontalDrawingGridEvery, &el); err != nil {
					return err
				}
			case "displayVerticalDrawingGridEvery":
				m.DisplayVerticalDrawingGridEvery = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.DisplayVerticalDrawingGridEvery, &el); err != nil {
					return err
				}
			case "doNotUseMarginsForDrawingGridOrigin":
				m.DoNotUseMarginsForDrawingGridOrigin = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotUseMarginsForDrawingGridOrigin, &el); err != nil {
					return err
				}
			case "drawingGridHorizontalOrigin":
				m.DrawingGridHorizontalOrigin = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.DrawingGridHorizontalOrigin, &el); err != nil {
					return err
				}
			case "drawingGridVerticalOrigin":
				m.DrawingGridVerticalOrigin = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.DrawingGridVerticalOrigin, &el); err != nil {
					return err
				}
			case "doNotShadeFormData":
				m.DoNotShadeFormData = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotShadeFormData, &el); err != nil {
					return err
				}
			case "noPunctuationKerning":
				m.NoPunctuationKerning = NewCT_OnOff()
				if err := d.DecodeElement(m.NoPunctuationKerning, &el); err != nil {
					return err
				}
			case "characterSpacingControl":
				m.CharacterSpacingControl = NewCT_CharacterSpacing()
				if err := d.DecodeElement(m.CharacterSpacingControl, &el); err != nil {
					return err
				}
			case "printTwoOnOne":
				m.PrintTwoOnOne = NewCT_OnOff()
				if err := d.DecodeElement(m.PrintTwoOnOne, &el); err != nil {
					return err
				}
			case "strictFirstAndLastChars":
				m.StrictFirstAndLastChars = NewCT_OnOff()
				if err := d.DecodeElement(m.StrictFirstAndLastChars, &el); err != nil {
					return err
				}
			case "noLineBreaksAfter":
				m.NoLineBreaksAfter = NewCT_Kinsoku()
				if err := d.DecodeElement(m.NoLineBreaksAfter, &el); err != nil {
					return err
				}
			case "noLineBreaksBefore":
				m.NoLineBreaksBefore = NewCT_Kinsoku()
				if err := d.DecodeElement(m.NoLineBreaksBefore, &el); err != nil {
					return err
				}
			case "savePreviewPicture":
				m.SavePreviewPicture = NewCT_OnOff()
				if err := d.DecodeElement(m.SavePreviewPicture, &el); err != nil {
					return err
				}
			case "doNotValidateAgainstSchema":
				m.DoNotValidateAgainstSchema = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotValidateAgainstSchema, &el); err != nil {
					return err
				}
			case "saveInvalidXml":
				m.SaveInvalidXml = NewCT_OnOff()
				if err := d.DecodeElement(m.SaveInvalidXml, &el); err != nil {
					return err
				}
			case "ignoreMixedContent":
				m.IgnoreMixedContent = NewCT_OnOff()
				if err := d.DecodeElement(m.IgnoreMixedContent, &el); err != nil {
					return err
				}
			case "alwaysShowPlaceholderText":
				m.AlwaysShowPlaceholderText = NewCT_OnOff()
				if err := d.DecodeElement(m.AlwaysShowPlaceholderText, &el); err != nil {
					return err
				}
			case "doNotDemarcateInvalidXml":
				m.DoNotDemarcateInvalidXml = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotDemarcateInvalidXml, &el); err != nil {
					return err
				}
			case "saveXmlDataOnly":
				m.SaveXmlDataOnly = NewCT_OnOff()
				if err := d.DecodeElement(m.SaveXmlDataOnly, &el); err != nil {
					return err
				}
			case "useXSLTWhenSaving":
				m.UseXSLTWhenSaving = NewCT_OnOff()
				if err := d.DecodeElement(m.UseXSLTWhenSaving, &el); err != nil {
					return err
				}
			case "saveThroughXslt":
				m.SaveThroughXslt = NewCT_SaveThroughXslt()
				if err := d.DecodeElement(m.SaveThroughXslt, &el); err != nil {
					return err
				}
			case "showXMLTags":
				m.ShowXMLTags = NewCT_OnOff()
				if err := d.DecodeElement(m.ShowXMLTags, &el); err != nil {
					return err
				}
			case "alwaysMergeEmptyNamespace":
				m.AlwaysMergeEmptyNamespace = NewCT_OnOff()
				if err := d.DecodeElement(m.AlwaysMergeEmptyNamespace, &el); err != nil {
					return err
				}
			case "updateFields":
				m.UpdateFields = NewCT_OnOff()
				if err := d.DecodeElement(m.UpdateFields, &el); err != nil {
					return err
				}
			case "hdrShapeDefaults":
				m.HdrShapeDefaults = NewCT_ShapeDefaults()
				if err := d.DecodeElement(m.HdrShapeDefaults, &el); err != nil {
					return err
				}
			case "footnotePr":
				m.FootnotePr = NewCT_FtnDocProps()
				if err := d.DecodeElement(m.FootnotePr, &el); err != nil {
					return err
				}
			case "endnotePr":
				m.EndnotePr = NewCT_EdnDocProps()
				if err := d.DecodeElement(m.EndnotePr, &el); err != nil {
					return err
				}
			case "compat":
				m.Compat = NewCT_Compat()
				if err := d.DecodeElement(m.Compat, &el); err != nil {
					return err
				}
			case "docVars":
				m.DocVars = NewCT_DocVars()
				if err := d.DecodeElement(m.DocVars, &el); err != nil {
					return err
				}
			case "rsids":
				m.Rsids = NewCT_DocRsids()
				if err := d.DecodeElement(m.Rsids, &el); err != nil {
					return err
				}
			case "mathPr":
				m.MathPr = math.NewMathPr()
				if err := d.DecodeElement(m.MathPr, &el); err != nil {
					return err
				}
			case "attachedSchema":
				tmp := NewCT_String()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AttachedSchema = append(m.AttachedSchema, tmp)
			case "themeFontLang":
				m.ThemeFontLang = NewCT_Language()
				if err := d.DecodeElement(m.ThemeFontLang, &el); err != nil {
					return err
				}
			case "clrSchemeMapping":
				m.ClrSchemeMapping = NewCT_ColorSchemeMapping()
				if err := d.DecodeElement(m.ClrSchemeMapping, &el); err != nil {
					return err
				}
			case "doNotIncludeSubdocsInStats":
				m.DoNotIncludeSubdocsInStats = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotIncludeSubdocsInStats, &el); err != nil {
					return err
				}
			case "doNotAutoCompressPictures":
				m.DoNotAutoCompressPictures = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotAutoCompressPictures, &el); err != nil {
					return err
				}
			case "forceUpgrade":
				m.ForceUpgrade = NewCT_Empty()
				if err := d.DecodeElement(m.ForceUpgrade, &el); err != nil {
					return err
				}
			case "captions":
				m.Captions = NewCT_Captions()
				if err := d.DecodeElement(m.Captions, &el); err != nil {
					return err
				}
			case "readModeInkLockDown":
				m.ReadModeInkLockDown = NewCT_ReadingModeInkLockDown()
				if err := d.DecodeElement(m.ReadModeInkLockDown, &el); err != nil {
					return err
				}
			case "smartTagType":
				tmp := NewCT_SmartTagType()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.SmartTagType = append(m.SmartTagType, tmp)
			case "schemaLibrary":
				m.SchemaLibrary = schemaLibrary.NewSchemaLibrary()
				if err := d.DecodeElement(m.SchemaLibrary, &el); err != nil {
					return err
				}
			case "shapeDefaults":
				m.ShapeDefaults = NewCT_ShapeDefaults()
				if err := d.DecodeElement(m.ShapeDefaults, &el); err != nil {
					return err
				}
			case "doNotEmbedSmartTags":
				m.DoNotEmbedSmartTags = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotEmbedSmartTags, &el); err != nil {
					return err
				}
			case "decimalSymbol":
				m.DecimalSymbol = NewCT_String()
				if err := d.DecodeElement(m.DecimalSymbol, &el); err != nil {
					return err
				}
			case "listSeparator":
				m.ListSeparator = NewCT_String()
				if err := d.DecodeElement(m.ListSeparator, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Settings
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Settings) Validate() error {
	return m.ValidateWithPath("CT_Settings")
}
func (m *CT_Settings) ValidateWithPath(path string) error {
	if m.WriteProtection != nil {
		if err := m.WriteProtection.ValidateWithPath(path + "/WriteProtection"); err != nil {
			return err
		}
	}
	if m.View != nil {
		if err := m.View.ValidateWithPath(path + "/View"); err != nil {
			return err
		}
	}
	if m.Zoom != nil {
		if err := m.Zoom.ValidateWithPath(path + "/Zoom"); err != nil {
			return err
		}
	}
	if m.RemovePersonalInformation != nil {
		if err := m.RemovePersonalInformation.ValidateWithPath(path + "/RemovePersonalInformation"); err != nil {
			return err
		}
	}
	if m.RemoveDateAndTime != nil {
		if err := m.RemoveDateAndTime.ValidateWithPath(path + "/RemoveDateAndTime"); err != nil {
			return err
		}
	}
	if m.DoNotDisplayPageBoundaries != nil {
		if err := m.DoNotDisplayPageBoundaries.ValidateWithPath(path + "/DoNotDisplayPageBoundaries"); err != nil {
			return err
		}
	}
	if m.DisplayBackgroundShape != nil {
		if err := m.DisplayBackgroundShape.ValidateWithPath(path + "/DisplayBackgroundShape"); err != nil {
			return err
		}
	}
	if m.PrintPostScriptOverText != nil {
		if err := m.PrintPostScriptOverText.ValidateWithPath(path + "/PrintPostScriptOverText"); err != nil {
			return err
		}
	}
	if m.PrintFractionalCharacterWidth != nil {
		if err := m.PrintFractionalCharacterWidth.ValidateWithPath(path + "/PrintFractionalCharacterWidth"); err != nil {
			return err
		}
	}
	if m.PrintFormsData != nil {
		if err := m.PrintFormsData.ValidateWithPath(path + "/PrintFormsData"); err != nil {
			return err
		}
	}
	if m.EmbedTrueTypeFonts != nil {
		if err := m.EmbedTrueTypeFonts.ValidateWithPath(path + "/EmbedTrueTypeFonts"); err != nil {
			return err
		}
	}
	if m.EmbedSystemFonts != nil {
		if err := m.EmbedSystemFonts.ValidateWithPath(path + "/EmbedSystemFonts"); err != nil {
			return err
		}
	}
	if m.SaveSubsetFonts != nil {
		if err := m.SaveSubsetFonts.ValidateWithPath(path + "/SaveSubsetFonts"); err != nil {
			return err
		}
	}
	if m.SaveFormsData != nil {
		if err := m.SaveFormsData.ValidateWithPath(path + "/SaveFormsData"); err != nil {
			return err
		}
	}
	if m.MirrorMargins != nil {
		if err := m.MirrorMargins.ValidateWithPath(path + "/MirrorMargins"); err != nil {
			return err
		}
	}
	if m.AlignBordersAndEdges != nil {
		if err := m.AlignBordersAndEdges.ValidateWithPath(path + "/AlignBordersAndEdges"); err != nil {
			return err
		}
	}
	if m.BordersDoNotSurroundHeader != nil {
		if err := m.BordersDoNotSurroundHeader.ValidateWithPath(path + "/BordersDoNotSurroundHeader"); err != nil {
			return err
		}
	}
	if m.BordersDoNotSurroundFooter != nil {
		if err := m.BordersDoNotSurroundFooter.ValidateWithPath(path + "/BordersDoNotSurroundFooter"); err != nil {
			return err
		}
	}
	if m.GutterAtTop != nil {
		if err := m.GutterAtTop.ValidateWithPath(path + "/GutterAtTop"); err != nil {
			return err
		}
	}
	if m.HideSpellingErrors != nil {
		if err := m.HideSpellingErrors.ValidateWithPath(path + "/HideSpellingErrors"); err != nil {
			return err
		}
	}
	if m.HideGrammaticalErrors != nil {
		if err := m.HideGrammaticalErrors.ValidateWithPath(path + "/HideGrammaticalErrors"); err != nil {
			return err
		}
	}
	for i, v := range m.ActiveWritingStyle {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ActiveWritingStyle[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.ProofState != nil {
		if err := m.ProofState.ValidateWithPath(path + "/ProofState"); err != nil {
			return err
		}
	}
	if m.FormsDesign != nil {
		if err := m.FormsDesign.ValidateWithPath(path + "/FormsDesign"); err != nil {
			return err
		}
	}
	if m.AttachedTemplate != nil {
		if err := m.AttachedTemplate.ValidateWithPath(path + "/AttachedTemplate"); err != nil {
			return err
		}
	}
	if m.LinkStyles != nil {
		if err := m.LinkStyles.ValidateWithPath(path + "/LinkStyles"); err != nil {
			return err
		}
	}
	if m.StylePaneFormatFilter != nil {
		if err := m.StylePaneFormatFilter.ValidateWithPath(path + "/StylePaneFormatFilter"); err != nil {
			return err
		}
	}
	if m.StylePaneSortMethod != nil {
		if err := m.StylePaneSortMethod.ValidateWithPath(path + "/StylePaneSortMethod"); err != nil {
			return err
		}
	}
	if m.DocumentType != nil {
		if err := m.DocumentType.ValidateWithPath(path + "/DocumentType"); err != nil {
			return err
		}
	}
	if m.MailMerge != nil {
		if err := m.MailMerge.ValidateWithPath(path + "/MailMerge"); err != nil {
			return err
		}
	}
	if m.RevisionView != nil {
		if err := m.RevisionView.ValidateWithPath(path + "/RevisionView"); err != nil {
			return err
		}
	}
	if m.TrackRevisions != nil {
		if err := m.TrackRevisions.ValidateWithPath(path + "/TrackRevisions"); err != nil {
			return err
		}
	}
	if m.DoNotTrackMoves != nil {
		if err := m.DoNotTrackMoves.ValidateWithPath(path + "/DoNotTrackMoves"); err != nil {
			return err
		}
	}
	if m.DoNotTrackFormatting != nil {
		if err := m.DoNotTrackFormatting.ValidateWithPath(path + "/DoNotTrackFormatting"); err != nil {
			return err
		}
	}
	if m.DocumentProtection != nil {
		if err := m.DocumentProtection.ValidateWithPath(path + "/DocumentProtection"); err != nil {
			return err
		}
	}
	if m.AutoFormatOverride != nil {
		if err := m.AutoFormatOverride.ValidateWithPath(path + "/AutoFormatOverride"); err != nil {
			return err
		}
	}
	if m.StyleLockTheme != nil {
		if err := m.StyleLockTheme.ValidateWithPath(path + "/StyleLockTheme"); err != nil {
			return err
		}
	}
	if m.StyleLockQFSet != nil {
		if err := m.StyleLockQFSet.ValidateWithPath(path + "/StyleLockQFSet"); err != nil {
			return err
		}
	}
	if m.DefaultTabStop != nil {
		if err := m.DefaultTabStop.ValidateWithPath(path + "/DefaultTabStop"); err != nil {
			return err
		}
	}
	if m.AutoHyphenation != nil {
		if err := m.AutoHyphenation.ValidateWithPath(path + "/AutoHyphenation"); err != nil {
			return err
		}
	}
	if m.ConsecutiveHyphenLimit != nil {
		if err := m.ConsecutiveHyphenLimit.ValidateWithPath(path + "/ConsecutiveHyphenLimit"); err != nil {
			return err
		}
	}
	if m.HyphenationZone != nil {
		if err := m.HyphenationZone.ValidateWithPath(path + "/HyphenationZone"); err != nil {
			return err
		}
	}
	if m.DoNotHyphenateCaps != nil {
		if err := m.DoNotHyphenateCaps.ValidateWithPath(path + "/DoNotHyphenateCaps"); err != nil {
			return err
		}
	}
	if m.ShowEnvelope != nil {
		if err := m.ShowEnvelope.ValidateWithPath(path + "/ShowEnvelope"); err != nil {
			return err
		}
	}
	if m.SummaryLength != nil {
		if err := m.SummaryLength.ValidateWithPath(path + "/SummaryLength"); err != nil {
			return err
		}
	}
	if m.ClickAndTypeStyle != nil {
		if err := m.ClickAndTypeStyle.ValidateWithPath(path + "/ClickAndTypeStyle"); err != nil {
			return err
		}
	}
	if m.DefaultTableStyle != nil {
		if err := m.DefaultTableStyle.ValidateWithPath(path + "/DefaultTableStyle"); err != nil {
			return err
		}
	}
	if m.EvenAndOddHeaders != nil {
		if err := m.EvenAndOddHeaders.ValidateWithPath(path + "/EvenAndOddHeaders"); err != nil {
			return err
		}
	}
	if m.BookFoldRevPrinting != nil {
		if err := m.BookFoldRevPrinting.ValidateWithPath(path + "/BookFoldRevPrinting"); err != nil {
			return err
		}
	}
	if m.BookFoldPrinting != nil {
		if err := m.BookFoldPrinting.ValidateWithPath(path + "/BookFoldPrinting"); err != nil {
			return err
		}
	}
	if m.BookFoldPrintingSheets != nil {
		if err := m.BookFoldPrintingSheets.ValidateWithPath(path + "/BookFoldPrintingSheets"); err != nil {
			return err
		}
	}
	if m.DrawingGridHorizontalSpacing != nil {
		if err := m.DrawingGridHorizontalSpacing.ValidateWithPath(path + "/DrawingGridHorizontalSpacing"); err != nil {
			return err
		}
	}
	if m.DrawingGridVerticalSpacing != nil {
		if err := m.DrawingGridVerticalSpacing.ValidateWithPath(path + "/DrawingGridVerticalSpacing"); err != nil {
			return err
		}
	}
	if m.DisplayHorizontalDrawingGridEvery != nil {
		if err := m.DisplayHorizontalDrawingGridEvery.ValidateWithPath(path + "/DisplayHorizontalDrawingGridEvery"); err != nil {
			return err
		}
	}
	if m.DisplayVerticalDrawingGridEvery != nil {
		if err := m.DisplayVerticalDrawingGridEvery.ValidateWithPath(path + "/DisplayVerticalDrawingGridEvery"); err != nil {
			return err
		}
	}
	if m.DoNotUseMarginsForDrawingGridOrigin != nil {
		if err := m.DoNotUseMarginsForDrawingGridOrigin.ValidateWithPath(path + "/DoNotUseMarginsForDrawingGridOrigin"); err != nil {
			return err
		}
	}
	if m.DrawingGridHorizontalOrigin != nil {
		if err := m.DrawingGridHorizontalOrigin.ValidateWithPath(path + "/DrawingGridHorizontalOrigin"); err != nil {
			return err
		}
	}
	if m.DrawingGridVerticalOrigin != nil {
		if err := m.DrawingGridVerticalOrigin.ValidateWithPath(path + "/DrawingGridVerticalOrigin"); err != nil {
			return err
		}
	}
	if m.DoNotShadeFormData != nil {
		if err := m.DoNotShadeFormData.ValidateWithPath(path + "/DoNotShadeFormData"); err != nil {
			return err
		}
	}
	if m.NoPunctuationKerning != nil {
		if err := m.NoPunctuationKerning.ValidateWithPath(path + "/NoPunctuationKerning"); err != nil {
			return err
		}
	}
	if m.CharacterSpacingControl != nil {
		if err := m.CharacterSpacingControl.ValidateWithPath(path + "/CharacterSpacingControl"); err != nil {
			return err
		}
	}
	if m.PrintTwoOnOne != nil {
		if err := m.PrintTwoOnOne.ValidateWithPath(path + "/PrintTwoOnOne"); err != nil {
			return err
		}
	}
	if m.StrictFirstAndLastChars != nil {
		if err := m.StrictFirstAndLastChars.ValidateWithPath(path + "/StrictFirstAndLastChars"); err != nil {
			return err
		}
	}
	if m.NoLineBreaksAfter != nil {
		if err := m.NoLineBreaksAfter.ValidateWithPath(path + "/NoLineBreaksAfter"); err != nil {
			return err
		}
	}
	if m.NoLineBreaksBefore != nil {
		if err := m.NoLineBreaksBefore.ValidateWithPath(path + "/NoLineBreaksBefore"); err != nil {
			return err
		}
	}
	if m.SavePreviewPicture != nil {
		if err := m.SavePreviewPicture.ValidateWithPath(path + "/SavePreviewPicture"); err != nil {
			return err
		}
	}
	if m.DoNotValidateAgainstSchema != nil {
		if err := m.DoNotValidateAgainstSchema.ValidateWithPath(path + "/DoNotValidateAgainstSchema"); err != nil {
			return err
		}
	}
	if m.SaveInvalidXml != nil {
		if err := m.SaveInvalidXml.ValidateWithPath(path + "/SaveInvalidXml"); err != nil {
			return err
		}
	}
	if m.IgnoreMixedContent != nil {
		if err := m.IgnoreMixedContent.ValidateWithPath(path + "/IgnoreMixedContent"); err != nil {
			return err
		}
	}
	if m.AlwaysShowPlaceholderText != nil {
		if err := m.AlwaysShowPlaceholderText.ValidateWithPath(path + "/AlwaysShowPlaceholderText"); err != nil {
			return err
		}
	}
	if m.DoNotDemarcateInvalidXml != nil {
		if err := m.DoNotDemarcateInvalidXml.ValidateWithPath(path + "/DoNotDemarcateInvalidXml"); err != nil {
			return err
		}
	}
	if m.SaveXmlDataOnly != nil {
		if err := m.SaveXmlDataOnly.ValidateWithPath(path + "/SaveXmlDataOnly"); err != nil {
			return err
		}
	}
	if m.UseXSLTWhenSaving != nil {
		if err := m.UseXSLTWhenSaving.ValidateWithPath(path + "/UseXSLTWhenSaving"); err != nil {
			return err
		}
	}
	if m.SaveThroughXslt != nil {
		if err := m.SaveThroughXslt.ValidateWithPath(path + "/SaveThroughXslt"); err != nil {
			return err
		}
	}
	if m.ShowXMLTags != nil {
		if err := m.ShowXMLTags.ValidateWithPath(path + "/ShowXMLTags"); err != nil {
			return err
		}
	}
	if m.AlwaysMergeEmptyNamespace != nil {
		if err := m.AlwaysMergeEmptyNamespace.ValidateWithPath(path + "/AlwaysMergeEmptyNamespace"); err != nil {
			return err
		}
	}
	if m.UpdateFields != nil {
		if err := m.UpdateFields.ValidateWithPath(path + "/UpdateFields"); err != nil {
			return err
		}
	}
	if m.HdrShapeDefaults != nil {
		if err := m.HdrShapeDefaults.ValidateWithPath(path + "/HdrShapeDefaults"); err != nil {
			return err
		}
	}
	if m.FootnotePr != nil {
		if err := m.FootnotePr.ValidateWithPath(path + "/FootnotePr"); err != nil {
			return err
		}
	}
	if m.EndnotePr != nil {
		if err := m.EndnotePr.ValidateWithPath(path + "/EndnotePr"); err != nil {
			return err
		}
	}
	if m.Compat != nil {
		if err := m.Compat.ValidateWithPath(path + "/Compat"); err != nil {
			return err
		}
	}
	if m.DocVars != nil {
		if err := m.DocVars.ValidateWithPath(path + "/DocVars"); err != nil {
			return err
		}
	}
	if m.Rsids != nil {
		if err := m.Rsids.ValidateWithPath(path + "/Rsids"); err != nil {
			return err
		}
	}
	if m.MathPr != nil {
		if err := m.MathPr.ValidateWithPath(path + "/MathPr"); err != nil {
			return err
		}
	}
	for i, v := range m.AttachedSchema {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AttachedSchema[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.ThemeFontLang != nil {
		if err := m.ThemeFontLang.ValidateWithPath(path + "/ThemeFontLang"); err != nil {
			return err
		}
	}
	if m.ClrSchemeMapping != nil {
		if err := m.ClrSchemeMapping.ValidateWithPath(path + "/ClrSchemeMapping"); err != nil {
			return err
		}
	}
	if m.DoNotIncludeSubdocsInStats != nil {
		if err := m.DoNotIncludeSubdocsInStats.ValidateWithPath(path + "/DoNotIncludeSubdocsInStats"); err != nil {
			return err
		}
	}
	if m.DoNotAutoCompressPictures != nil {
		if err := m.DoNotAutoCompressPictures.ValidateWithPath(path + "/DoNotAutoCompressPictures"); err != nil {
			return err
		}
	}
	if m.ForceUpgrade != nil {
		if err := m.ForceUpgrade.ValidateWithPath(path + "/ForceUpgrade"); err != nil {
			return err
		}
	}
	if m.Captions != nil {
		if err := m.Captions.ValidateWithPath(path + "/Captions"); err != nil {
			return err
		}
	}
	if m.ReadModeInkLockDown != nil {
		if err := m.ReadModeInkLockDown.ValidateWithPath(path + "/ReadModeInkLockDown"); err != nil {
			return err
		}
	}
	for i, v := range m.SmartTagType {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/SmartTagType[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.SchemaLibrary != nil {
		if err := m.SchemaLibrary.ValidateWithPath(path + "/SchemaLibrary"); err != nil {
			return err
		}
	}
	if m.ShapeDefaults != nil {
		if err := m.ShapeDefaults.ValidateWithPath(path + "/ShapeDefaults"); err != nil {
			return err
		}
	}
	if m.DoNotEmbedSmartTags != nil {
		if err := m.DoNotEmbedSmartTags.ValidateWithPath(path + "/DoNotEmbedSmartTags"); err != nil {
			return err
		}
	}
	if m.DecimalSymbol != nil {
		if err := m.DecimalSymbol.ValidateWithPath(path + "/DecimalSymbol"); err != nil {
			return err
		}
	}
	if m.ListSeparator != nil {
		if err := m.ListSeparator.ValidateWithPath(path + "/ListSeparator"); err != nil {
			return err
		}
	}
	return nil
}
