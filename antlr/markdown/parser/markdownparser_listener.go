// Code generated from MarkdownParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // MarkdownParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MarkdownParserListener is a complete listener for a parse tree produced by MarkdownParser.
type MarkdownParserListener interface {
	antlr.ParseTreeListener

	// EnterDocument is called when entering the document production.
	EnterDocument(c *DocumentContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterHtmlBlockTags is called when entering the htmlBlockTags production.
	EnterHtmlBlockTags(c *HtmlBlockTagsContext)

	// EnterHtmlBlockSelfClosing is called when entering the htmlBlockSelfClosing production.
	EnterHtmlBlockSelfClosing(c *HtmlBlockSelfClosingContext)

	// EnterHeading is called when entering the heading production.
	EnterHeading(c *HeadingContext)

	// EnterSetextHeading is called when entering the setextHeading production.
	EnterSetextHeading(c *SetextHeadingContext)

	// EnterSetextHeading1 is called when entering the setextHeading1 production.
	EnterSetextHeading1(c *SetextHeading1Context)

	// EnterSetextHeading2 is called when entering the setextHeading2 production.
	EnterSetextHeading2(c *SetextHeading2Context)

	// EnterAtxHeading is called when entering the atxHeading production.
	EnterAtxHeading(c *AtxHeadingContext)

	// EnterRawLine is called when entering the rawLine production.
	EnterRawLine(c *RawLineContext)

	// EnterNonIndentSpace is called when entering the nonIndentSpace production.
	EnterNonIndentSpace(c *NonIndentSpaceContext)

	// EnterBlockQuote is called when entering the blockQuote production.
	EnterBlockQuote(c *BlockQuoteContext)

	// EnterBlockQuoteBlankLine is called when entering the blockQuoteBlankLine production.
	EnterBlockQuoteBlankLine(c *BlockQuoteBlankLineContext)

	// EnterVerbatim is called when entering the verbatim production.
	EnterVerbatim(c *VerbatimContext)

	// EnterVerbatimBlankLine is called when entering the verbatimBlankLine production.
	EnterVerbatimBlankLine(c *VerbatimBlankLineContext)

	// EnterHorizontalRule is called when entering the horizontalRule production.
	EnterHorizontalRule(c *HorizontalRuleContext)

	// EnterReferences is called when entering the references production.
	EnterReferences(c *ReferencesContext)

	// EnterReference is called when entering the reference production.
	EnterReference(c *ReferenceContext)

	// EnterReferenceLabel is called when entering the referenceLabel production.
	EnterReferenceLabel(c *ReferenceLabelContext)

	// EnterReferenceId is called when entering the referenceId production.
	EnterReferenceId(c *ReferenceIdContext)

	// EnterReferenceUrl is called when entering the referenceUrl production.
	EnterReferenceUrl(c *ReferenceUrlContext)

	// EnterReferenceTitle is called when entering the referenceTitle production.
	EnterReferenceTitle(c *ReferenceTitleContext)

	// EnterReferenceTitleSingle is called when entering the referenceTitleSingle production.
	EnterReferenceTitleSingle(c *ReferenceTitleSingleContext)

	// EnterReferenceTitleDouble is called when entering the referenceTitleDouble production.
	EnterReferenceTitleDouble(c *ReferenceTitleDoubleContext)

	// EnterReferenceTitleParens is called when entering the referenceTitleParens production.
	EnterReferenceTitleParens(c *ReferenceTitleParensContext)

	// EnterOrderedList is called when entering the orderedList production.
	EnterOrderedList(c *OrderedListContext)

	// EnterBulletList is called when entering the bulletList production.
	EnterBulletList(c *BulletListContext)

	// EnterOrderedListItem is called when entering the orderedListItem production.
	EnterOrderedListItem(c *OrderedListItemContext)

	// EnterBulletListItem is called when entering the bulletListItem production.
	EnterBulletListItem(c *BulletListItemContext)

	// EnterInlineListItem is called when entering the inlineListItem production.
	EnterInlineListItem(c *InlineListItemContext)

	// EnterListItemBlankLine is called when entering the listItemBlankLine production.
	EnterListItemBlankLine(c *ListItemBlankLineContext)

	// EnterPara is called when entering the para production.
	EnterPara(c *ParaContext)

	// EnterInline is called when entering the inline production.
	EnterInline(c *InlineContext)

	// EnterSpan is called when entering the span production.
	EnterSpan(c *SpanContext)

	// EnterEmph is called when entering the emph production.
	EnterEmph(c *EmphContext)

	// EnterEmphStar is called when entering the emphStar production.
	EnterEmphStar(c *EmphStarContext)

	// EnterEmphUl is called when entering the emphUl production.
	EnterEmphUl(c *EmphUlContext)

	// EnterStrong is called when entering the strong production.
	EnterStrong(c *StrongContext)

	// EnterStrongStar is called when entering the strongStar production.
	EnterStrongStar(c *StrongStarContext)

	// EnterStrongUl is called when entering the strongUl production.
	EnterStrongUl(c *StrongUlContext)

	// EnterImage is called when entering the image production.
	EnterImage(c *ImageContext)

	// EnterImageLink is called when entering the imageLink production.
	EnterImageLink(c *ImageLinkContext)

	// EnterExplicitImageLink is called when entering the explicitImageLink production.
	EnterExplicitImageLink(c *ExplicitImageLinkContext)

	// EnterImageAlt is called when entering the imageAlt production.
	EnterImageAlt(c *ImageAltContext)

	// EnterReferenceImageLink is called when entering the referenceImageLink production.
	EnterReferenceImageLink(c *ReferenceImageLinkContext)

	// EnterLink is called when entering the link production.
	EnterLink(c *LinkContext)

	// EnterLinkContent is called when entering the linkContent production.
	EnterLinkContent(c *LinkContentContext)

	// EnterExplicitLink is called when entering the explicitLink production.
	EnterExplicitLink(c *ExplicitLinkContext)

	// EnterLinkUrl is called when entering the linkUrl production.
	EnterLinkUrl(c *LinkUrlContext)

	// EnterLinkTitle is called when entering the linkTitle production.
	EnterLinkTitle(c *LinkTitleContext)

	// EnterLinkTitleSingle is called when entering the linkTitleSingle production.
	EnterLinkTitleSingle(c *LinkTitleSingleContext)

	// EnterLinkTitleDouble is called when entering the linkTitleDouble production.
	EnterLinkTitleDouble(c *LinkTitleDoubleContext)

	// EnterReferenceLink is called when entering the referenceLink production.
	EnterReferenceLink(c *ReferenceLinkContext)

	// EnterEntity is called when entering the entity production.
	EnterEntity(c *EntityContext)

	// EnterHexEntityName is called when entering the hexEntityName production.
	EnterHexEntityName(c *HexEntityNameContext)

	// EnterDecEntityName is called when entering the decEntityName production.
	EnterDecEntityName(c *DecEntityNameContext)

	// EnterCharEntityName is called when entering the charEntityName production.
	EnterCharEntityName(c *CharEntityNameContext)

	// EnterEscapedChar is called when entering the escapedChar production.
	EnterEscapedChar(c *EscapedCharContext)

	// EnterCode is called when entering the code production.
	EnterCode(c *CodeContext)

	// EnterBacktickCode is called when entering the backtickCode production.
	EnterBacktickCode(c *BacktickCodeContext)

	// EnterDoubleBacktickCode is called when entering the doubleBacktickCode production.
	EnterDoubleBacktickCode(c *DoubleBacktickCodeContext)

	// EnterSpaceBacktickCode is called when entering the spaceBacktickCode production.
	EnterSpaceBacktickCode(c *SpaceBacktickCodeContext)

	// EnterAttributeName is called when entering the attributeName production.
	EnterAttributeName(c *AttributeNameContext)

	// EnterHtmlAttributeS is called when entering the htmlAttributeS production.
	EnterHtmlAttributeS(c *HtmlAttributeSContext)

	// EnterHtmlAttributeD is called when entering the htmlAttributeD production.
	EnterHtmlAttributeD(c *HtmlAttributeDContext)

	// EnterHtmlBlockOpenDiv is called when entering the htmlBlockOpenDiv production.
	EnterHtmlBlockOpenDiv(c *HtmlBlockOpenDivContext)

	// EnterHtmlBlockCloseDiv is called when entering the htmlBlockCloseDiv production.
	EnterHtmlBlockCloseDiv(c *HtmlBlockCloseDivContext)

	// EnterHtmlBlockDiv is called when entering the htmlBlockDiv production.
	EnterHtmlBlockDiv(c *HtmlBlockDivContext)

	// EnterHtmlBlockOpenSpan is called when entering the htmlBlockOpenSpan production.
	EnterHtmlBlockOpenSpan(c *HtmlBlockOpenSpanContext)

	// EnterHtmlBlockCloseSpan is called when entering the htmlBlockCloseSpan production.
	EnterHtmlBlockCloseSpan(c *HtmlBlockCloseSpanContext)

	// EnterHtmlBlockSpan is called when entering the htmlBlockSpan production.
	EnterHtmlBlockSpan(c *HtmlBlockSpanContext)

	// EnterHtmlBlockHr is called when entering the htmlBlockHr production.
	EnterHtmlBlockHr(c *HtmlBlockHrContext)

	// EnterHtmlBlockInTags is called when entering the htmlBlockInTags production.
	EnterHtmlBlockInTags(c *HtmlBlockInTagsContext)

	// EnterHtmlBlockInSelfClosing is called when entering the htmlBlockInSelfClosing production.
	EnterHtmlBlockInSelfClosing(c *HtmlBlockInSelfClosingContext)

	// EnterHtmlComment is called when entering the htmlComment production.
	EnterHtmlComment(c *HtmlCommentContext)

	// EnterAutolink is called when entering the autolink production.
	EnterAutolink(c *AutolinkContext)

	// ExitDocument is called when exiting the document production.
	ExitDocument(c *DocumentContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitHtmlBlockTags is called when exiting the htmlBlockTags production.
	ExitHtmlBlockTags(c *HtmlBlockTagsContext)

	// ExitHtmlBlockSelfClosing is called when exiting the htmlBlockSelfClosing production.
	ExitHtmlBlockSelfClosing(c *HtmlBlockSelfClosingContext)

	// ExitHeading is called when exiting the heading production.
	ExitHeading(c *HeadingContext)

	// ExitSetextHeading is called when exiting the setextHeading production.
	ExitSetextHeading(c *SetextHeadingContext)

	// ExitSetextHeading1 is called when exiting the setextHeading1 production.
	ExitSetextHeading1(c *SetextHeading1Context)

	// ExitSetextHeading2 is called when exiting the setextHeading2 production.
	ExitSetextHeading2(c *SetextHeading2Context)

	// ExitAtxHeading is called when exiting the atxHeading production.
	ExitAtxHeading(c *AtxHeadingContext)

	// ExitRawLine is called when exiting the rawLine production.
	ExitRawLine(c *RawLineContext)

	// ExitNonIndentSpace is called when exiting the nonIndentSpace production.
	ExitNonIndentSpace(c *NonIndentSpaceContext)

	// ExitBlockQuote is called when exiting the blockQuote production.
	ExitBlockQuote(c *BlockQuoteContext)

	// ExitBlockQuoteBlankLine is called when exiting the blockQuoteBlankLine production.
	ExitBlockQuoteBlankLine(c *BlockQuoteBlankLineContext)

	// ExitVerbatim is called when exiting the verbatim production.
	ExitVerbatim(c *VerbatimContext)

	// ExitVerbatimBlankLine is called when exiting the verbatimBlankLine production.
	ExitVerbatimBlankLine(c *VerbatimBlankLineContext)

	// ExitHorizontalRule is called when exiting the horizontalRule production.
	ExitHorizontalRule(c *HorizontalRuleContext)

	// ExitReferences is called when exiting the references production.
	ExitReferences(c *ReferencesContext)

	// ExitReference is called when exiting the reference production.
	ExitReference(c *ReferenceContext)

	// ExitReferenceLabel is called when exiting the referenceLabel production.
	ExitReferenceLabel(c *ReferenceLabelContext)

	// ExitReferenceId is called when exiting the referenceId production.
	ExitReferenceId(c *ReferenceIdContext)

	// ExitReferenceUrl is called when exiting the referenceUrl production.
	ExitReferenceUrl(c *ReferenceUrlContext)

	// ExitReferenceTitle is called when exiting the referenceTitle production.
	ExitReferenceTitle(c *ReferenceTitleContext)

	// ExitReferenceTitleSingle is called when exiting the referenceTitleSingle production.
	ExitReferenceTitleSingle(c *ReferenceTitleSingleContext)

	// ExitReferenceTitleDouble is called when exiting the referenceTitleDouble production.
	ExitReferenceTitleDouble(c *ReferenceTitleDoubleContext)

	// ExitReferenceTitleParens is called when exiting the referenceTitleParens production.
	ExitReferenceTitleParens(c *ReferenceTitleParensContext)

	// ExitOrderedList is called when exiting the orderedList production.
	ExitOrderedList(c *OrderedListContext)

	// ExitBulletList is called when exiting the bulletList production.
	ExitBulletList(c *BulletListContext)

	// ExitOrderedListItem is called when exiting the orderedListItem production.
	ExitOrderedListItem(c *OrderedListItemContext)

	// ExitBulletListItem is called when exiting the bulletListItem production.
	ExitBulletListItem(c *BulletListItemContext)

	// ExitInlineListItem is called when exiting the inlineListItem production.
	ExitInlineListItem(c *InlineListItemContext)

	// ExitListItemBlankLine is called when exiting the listItemBlankLine production.
	ExitListItemBlankLine(c *ListItemBlankLineContext)

	// ExitPara is called when exiting the para production.
	ExitPara(c *ParaContext)

	// ExitInline is called when exiting the inline production.
	ExitInline(c *InlineContext)

	// ExitSpan is called when exiting the span production.
	ExitSpan(c *SpanContext)

	// ExitEmph is called when exiting the emph production.
	ExitEmph(c *EmphContext)

	// ExitEmphStar is called when exiting the emphStar production.
	ExitEmphStar(c *EmphStarContext)

	// ExitEmphUl is called when exiting the emphUl production.
	ExitEmphUl(c *EmphUlContext)

	// ExitStrong is called when exiting the strong production.
	ExitStrong(c *StrongContext)

	// ExitStrongStar is called when exiting the strongStar production.
	ExitStrongStar(c *StrongStarContext)

	// ExitStrongUl is called when exiting the strongUl production.
	ExitStrongUl(c *StrongUlContext)

	// ExitImage is called when exiting the image production.
	ExitImage(c *ImageContext)

	// ExitImageLink is called when exiting the imageLink production.
	ExitImageLink(c *ImageLinkContext)

	// ExitExplicitImageLink is called when exiting the explicitImageLink production.
	ExitExplicitImageLink(c *ExplicitImageLinkContext)

	// ExitImageAlt is called when exiting the imageAlt production.
	ExitImageAlt(c *ImageAltContext)

	// ExitReferenceImageLink is called when exiting the referenceImageLink production.
	ExitReferenceImageLink(c *ReferenceImageLinkContext)

	// ExitLink is called when exiting the link production.
	ExitLink(c *LinkContext)

	// ExitLinkContent is called when exiting the linkContent production.
	ExitLinkContent(c *LinkContentContext)

	// ExitExplicitLink is called when exiting the explicitLink production.
	ExitExplicitLink(c *ExplicitLinkContext)

	// ExitLinkUrl is called when exiting the linkUrl production.
	ExitLinkUrl(c *LinkUrlContext)

	// ExitLinkTitle is called when exiting the linkTitle production.
	ExitLinkTitle(c *LinkTitleContext)

	// ExitLinkTitleSingle is called when exiting the linkTitleSingle production.
	ExitLinkTitleSingle(c *LinkTitleSingleContext)

	// ExitLinkTitleDouble is called when exiting the linkTitleDouble production.
	ExitLinkTitleDouble(c *LinkTitleDoubleContext)

	// ExitReferenceLink is called when exiting the referenceLink production.
	ExitReferenceLink(c *ReferenceLinkContext)

	// ExitEntity is called when exiting the entity production.
	ExitEntity(c *EntityContext)

	// ExitHexEntityName is called when exiting the hexEntityName production.
	ExitHexEntityName(c *HexEntityNameContext)

	// ExitDecEntityName is called when exiting the decEntityName production.
	ExitDecEntityName(c *DecEntityNameContext)

	// ExitCharEntityName is called when exiting the charEntityName production.
	ExitCharEntityName(c *CharEntityNameContext)

	// ExitEscapedChar is called when exiting the escapedChar production.
	ExitEscapedChar(c *EscapedCharContext)

	// ExitCode is called when exiting the code production.
	ExitCode(c *CodeContext)

	// ExitBacktickCode is called when exiting the backtickCode production.
	ExitBacktickCode(c *BacktickCodeContext)

	// ExitDoubleBacktickCode is called when exiting the doubleBacktickCode production.
	ExitDoubleBacktickCode(c *DoubleBacktickCodeContext)

	// ExitSpaceBacktickCode is called when exiting the spaceBacktickCode production.
	ExitSpaceBacktickCode(c *SpaceBacktickCodeContext)

	// ExitAttributeName is called when exiting the attributeName production.
	ExitAttributeName(c *AttributeNameContext)

	// ExitHtmlAttributeS is called when exiting the htmlAttributeS production.
	ExitHtmlAttributeS(c *HtmlAttributeSContext)

	// ExitHtmlAttributeD is called when exiting the htmlAttributeD production.
	ExitHtmlAttributeD(c *HtmlAttributeDContext)

	// ExitHtmlBlockOpenDiv is called when exiting the htmlBlockOpenDiv production.
	ExitHtmlBlockOpenDiv(c *HtmlBlockOpenDivContext)

	// ExitHtmlBlockCloseDiv is called when exiting the htmlBlockCloseDiv production.
	ExitHtmlBlockCloseDiv(c *HtmlBlockCloseDivContext)

	// ExitHtmlBlockDiv is called when exiting the htmlBlockDiv production.
	ExitHtmlBlockDiv(c *HtmlBlockDivContext)

	// ExitHtmlBlockOpenSpan is called when exiting the htmlBlockOpenSpan production.
	ExitHtmlBlockOpenSpan(c *HtmlBlockOpenSpanContext)

	// ExitHtmlBlockCloseSpan is called when exiting the htmlBlockCloseSpan production.
	ExitHtmlBlockCloseSpan(c *HtmlBlockCloseSpanContext)

	// ExitHtmlBlockSpan is called when exiting the htmlBlockSpan production.
	ExitHtmlBlockSpan(c *HtmlBlockSpanContext)

	// ExitHtmlBlockHr is called when exiting the htmlBlockHr production.
	ExitHtmlBlockHr(c *HtmlBlockHrContext)

	// ExitHtmlBlockInTags is called when exiting the htmlBlockInTags production.
	ExitHtmlBlockInTags(c *HtmlBlockInTagsContext)

	// ExitHtmlBlockInSelfClosing is called when exiting the htmlBlockInSelfClosing production.
	ExitHtmlBlockInSelfClosing(c *HtmlBlockInSelfClosingContext)

	// ExitHtmlComment is called when exiting the htmlComment production.
	ExitHtmlComment(c *HtmlCommentContext)

	// ExitAutolink is called when exiting the autolink production.
	ExitAutolink(c *AutolinkContext)
}
