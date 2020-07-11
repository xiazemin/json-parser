// Code generated from MarkdownParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // MarkdownParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMarkdownParserListener is a complete listener for a parse tree produced by MarkdownParser.
type BaseMarkdownParserListener struct{}

var _ MarkdownParserListener = &BaseMarkdownParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMarkdownParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMarkdownParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMarkdownParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMarkdownParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDocument is called when production document is entered.
func (s *BaseMarkdownParserListener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *BaseMarkdownParserListener) ExitDocument(ctx *DocumentContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseMarkdownParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseMarkdownParserListener) ExitBlock(ctx *BlockContext) {}

// EnterHtmlBlockTags is called when production htmlBlockTags is entered.
func (s *BaseMarkdownParserListener) EnterHtmlBlockTags(ctx *HtmlBlockTagsContext) {}

// ExitHtmlBlockTags is called when production htmlBlockTags is exited.
func (s *BaseMarkdownParserListener) ExitHtmlBlockTags(ctx *HtmlBlockTagsContext) {}

// EnterHtmlBlockSelfClosing is called when production htmlBlockSelfClosing is entered.
func (s *BaseMarkdownParserListener) EnterHtmlBlockSelfClosing(ctx *HtmlBlockSelfClosingContext) {}

// ExitHtmlBlockSelfClosing is called when production htmlBlockSelfClosing is exited.
func (s *BaseMarkdownParserListener) ExitHtmlBlockSelfClosing(ctx *HtmlBlockSelfClosingContext) {}

// EnterHeading is called when production heading is entered.
func (s *BaseMarkdownParserListener) EnterHeading(ctx *HeadingContext) {}

// ExitHeading is called when production heading is exited.
func (s *BaseMarkdownParserListener) ExitHeading(ctx *HeadingContext) {}

// EnterSetextHeading is called when production setextHeading is entered.
func (s *BaseMarkdownParserListener) EnterSetextHeading(ctx *SetextHeadingContext) {}

// ExitSetextHeading is called when production setextHeading is exited.
func (s *BaseMarkdownParserListener) ExitSetextHeading(ctx *SetextHeadingContext) {}

// EnterSetextHeading1 is called when production setextHeading1 is entered.
func (s *BaseMarkdownParserListener) EnterSetextHeading1(ctx *SetextHeading1Context) {}

// ExitSetextHeading1 is called when production setextHeading1 is exited.
func (s *BaseMarkdownParserListener) ExitSetextHeading1(ctx *SetextHeading1Context) {}

// EnterSetextHeading2 is called when production setextHeading2 is entered.
func (s *BaseMarkdownParserListener) EnterSetextHeading2(ctx *SetextHeading2Context) {}

// ExitSetextHeading2 is called when production setextHeading2 is exited.
func (s *BaseMarkdownParserListener) ExitSetextHeading2(ctx *SetextHeading2Context) {}

// EnterAtxHeading is called when production atxHeading is entered.
func (s *BaseMarkdownParserListener) EnterAtxHeading(ctx *AtxHeadingContext) {}

// ExitAtxHeading is called when production atxHeading is exited.
func (s *BaseMarkdownParserListener) ExitAtxHeading(ctx *AtxHeadingContext) {}

// EnterRawLine is called when production rawLine is entered.
func (s *BaseMarkdownParserListener) EnterRawLine(ctx *RawLineContext) {}

// ExitRawLine is called when production rawLine is exited.
func (s *BaseMarkdownParserListener) ExitRawLine(ctx *RawLineContext) {}

// EnterNonIndentSpace is called when production nonIndentSpace is entered.
func (s *BaseMarkdownParserListener) EnterNonIndentSpace(ctx *NonIndentSpaceContext) {}

// ExitNonIndentSpace is called when production nonIndentSpace is exited.
func (s *BaseMarkdownParserListener) ExitNonIndentSpace(ctx *NonIndentSpaceContext) {}

// EnterBlockQuote is called when production blockQuote is entered.
func (s *BaseMarkdownParserListener) EnterBlockQuote(ctx *BlockQuoteContext) {}

// ExitBlockQuote is called when production blockQuote is exited.
func (s *BaseMarkdownParserListener) ExitBlockQuote(ctx *BlockQuoteContext) {}

// EnterBlockQuoteBlankLine is called when production blockQuoteBlankLine is entered.
func (s *BaseMarkdownParserListener) EnterBlockQuoteBlankLine(ctx *BlockQuoteBlankLineContext) {}

// ExitBlockQuoteBlankLine is called when production blockQuoteBlankLine is exited.
func (s *BaseMarkdownParserListener) ExitBlockQuoteBlankLine(ctx *BlockQuoteBlankLineContext) {}

// EnterVerbatim is called when production verbatim is entered.
func (s *BaseMarkdownParserListener) EnterVerbatim(ctx *VerbatimContext) {}

// ExitVerbatim is called when production verbatim is exited.
func (s *BaseMarkdownParserListener) ExitVerbatim(ctx *VerbatimContext) {}

// EnterVerbatimBlankLine is called when production verbatimBlankLine is entered.
func (s *BaseMarkdownParserListener) EnterVerbatimBlankLine(ctx *VerbatimBlankLineContext) {}

// ExitVerbatimBlankLine is called when production verbatimBlankLine is exited.
func (s *BaseMarkdownParserListener) ExitVerbatimBlankLine(ctx *VerbatimBlankLineContext) {}

// EnterHorizontalRule is called when production horizontalRule is entered.
func (s *BaseMarkdownParserListener) EnterHorizontalRule(ctx *HorizontalRuleContext) {}

// ExitHorizontalRule is called when production horizontalRule is exited.
func (s *BaseMarkdownParserListener) ExitHorizontalRule(ctx *HorizontalRuleContext) {}

// EnterReferences is called when production references is entered.
func (s *BaseMarkdownParserListener) EnterReferences(ctx *ReferencesContext) {}

// ExitReferences is called when production references is exited.
func (s *BaseMarkdownParserListener) ExitReferences(ctx *ReferencesContext) {}

// EnterReference is called when production reference is entered.
func (s *BaseMarkdownParserListener) EnterReference(ctx *ReferenceContext) {}

// ExitReference is called when production reference is exited.
func (s *BaseMarkdownParserListener) ExitReference(ctx *ReferenceContext) {}

// EnterReferenceLabel is called when production referenceLabel is entered.
func (s *BaseMarkdownParserListener) EnterReferenceLabel(ctx *ReferenceLabelContext) {}

// ExitReferenceLabel is called when production referenceLabel is exited.
func (s *BaseMarkdownParserListener) ExitReferenceLabel(ctx *ReferenceLabelContext) {}

// EnterReferenceId is called when production referenceId is entered.
func (s *BaseMarkdownParserListener) EnterReferenceId(ctx *ReferenceIdContext) {}

// ExitReferenceId is called when production referenceId is exited.
func (s *BaseMarkdownParserListener) ExitReferenceId(ctx *ReferenceIdContext) {}

// EnterReferenceUrl is called when production referenceUrl is entered.
func (s *BaseMarkdownParserListener) EnterReferenceUrl(ctx *ReferenceUrlContext) {}

// ExitReferenceUrl is called when production referenceUrl is exited.
func (s *BaseMarkdownParserListener) ExitReferenceUrl(ctx *ReferenceUrlContext) {}

// EnterReferenceTitle is called when production referenceTitle is entered.
func (s *BaseMarkdownParserListener) EnterReferenceTitle(ctx *ReferenceTitleContext) {}

// ExitReferenceTitle is called when production referenceTitle is exited.
func (s *BaseMarkdownParserListener) ExitReferenceTitle(ctx *ReferenceTitleContext) {}

// EnterReferenceTitleSingle is called when production referenceTitleSingle is entered.
func (s *BaseMarkdownParserListener) EnterReferenceTitleSingle(ctx *ReferenceTitleSingleContext) {}

// ExitReferenceTitleSingle is called when production referenceTitleSingle is exited.
func (s *BaseMarkdownParserListener) ExitReferenceTitleSingle(ctx *ReferenceTitleSingleContext) {}

// EnterReferenceTitleDouble is called when production referenceTitleDouble is entered.
func (s *BaseMarkdownParserListener) EnterReferenceTitleDouble(ctx *ReferenceTitleDoubleContext) {}

// ExitReferenceTitleDouble is called when production referenceTitleDouble is exited.
func (s *BaseMarkdownParserListener) ExitReferenceTitleDouble(ctx *ReferenceTitleDoubleContext) {}

// EnterReferenceTitleParens is called when production referenceTitleParens is entered.
func (s *BaseMarkdownParserListener) EnterReferenceTitleParens(ctx *ReferenceTitleParensContext) {}

// ExitReferenceTitleParens is called when production referenceTitleParens is exited.
func (s *BaseMarkdownParserListener) ExitReferenceTitleParens(ctx *ReferenceTitleParensContext) {}

// EnterOrderedList is called when production orderedList is entered.
func (s *BaseMarkdownParserListener) EnterOrderedList(ctx *OrderedListContext) {}

// ExitOrderedList is called when production orderedList is exited.
func (s *BaseMarkdownParserListener) ExitOrderedList(ctx *OrderedListContext) {}

// EnterBulletList is called when production bulletList is entered.
func (s *BaseMarkdownParserListener) EnterBulletList(ctx *BulletListContext) {}

// ExitBulletList is called when production bulletList is exited.
func (s *BaseMarkdownParserListener) ExitBulletList(ctx *BulletListContext) {}

// EnterOrderedListItem is called when production orderedListItem is entered.
func (s *BaseMarkdownParserListener) EnterOrderedListItem(ctx *OrderedListItemContext) {}

// ExitOrderedListItem is called when production orderedListItem is exited.
func (s *BaseMarkdownParserListener) ExitOrderedListItem(ctx *OrderedListItemContext) {}

// EnterBulletListItem is called when production bulletListItem is entered.
func (s *BaseMarkdownParserListener) EnterBulletListItem(ctx *BulletListItemContext) {}

// ExitBulletListItem is called when production bulletListItem is exited.
func (s *BaseMarkdownParserListener) ExitBulletListItem(ctx *BulletListItemContext) {}

// EnterInlineListItem is called when production inlineListItem is entered.
func (s *BaseMarkdownParserListener) EnterInlineListItem(ctx *InlineListItemContext) {}

// ExitInlineListItem is called when production inlineListItem is exited.
func (s *BaseMarkdownParserListener) ExitInlineListItem(ctx *InlineListItemContext) {}

// EnterListItemBlankLine is called when production listItemBlankLine is entered.
func (s *BaseMarkdownParserListener) EnterListItemBlankLine(ctx *ListItemBlankLineContext) {}

// ExitListItemBlankLine is called when production listItemBlankLine is exited.
func (s *BaseMarkdownParserListener) ExitListItemBlankLine(ctx *ListItemBlankLineContext) {}

// EnterPara is called when production para is entered.
func (s *BaseMarkdownParserListener) EnterPara(ctx *ParaContext) {}

// ExitPara is called when production para is exited.
func (s *BaseMarkdownParserListener) ExitPara(ctx *ParaContext) {}

// EnterInline is called when production inline is entered.
func (s *BaseMarkdownParserListener) EnterInline(ctx *InlineContext) {}

// ExitInline is called when production inline is exited.
func (s *BaseMarkdownParserListener) ExitInline(ctx *InlineContext) {}

// EnterSpan is called when production span is entered.
func (s *BaseMarkdownParserListener) EnterSpan(ctx *SpanContext) {}

// ExitSpan is called when production span is exited.
func (s *BaseMarkdownParserListener) ExitSpan(ctx *SpanContext) {}

// EnterEmph is called when production emph is entered.
func (s *BaseMarkdownParserListener) EnterEmph(ctx *EmphContext) {}

// ExitEmph is called when production emph is exited.
func (s *BaseMarkdownParserListener) ExitEmph(ctx *EmphContext) {}

// EnterEmphStar is called when production emphStar is entered.
func (s *BaseMarkdownParserListener) EnterEmphStar(ctx *EmphStarContext) {}

// ExitEmphStar is called when production emphStar is exited.
func (s *BaseMarkdownParserListener) ExitEmphStar(ctx *EmphStarContext) {}

// EnterEmphUl is called when production emphUl is entered.
func (s *BaseMarkdownParserListener) EnterEmphUl(ctx *EmphUlContext) {}

// ExitEmphUl is called when production emphUl is exited.
func (s *BaseMarkdownParserListener) ExitEmphUl(ctx *EmphUlContext) {}

// EnterStrong is called when production strong is entered.
func (s *BaseMarkdownParserListener) EnterStrong(ctx *StrongContext) {}

// ExitStrong is called when production strong is exited.
func (s *BaseMarkdownParserListener) ExitStrong(ctx *StrongContext) {}

// EnterStrongStar is called when production strongStar is entered.
func (s *BaseMarkdownParserListener) EnterStrongStar(ctx *StrongStarContext) {}

// ExitStrongStar is called when production strongStar is exited.
func (s *BaseMarkdownParserListener) ExitStrongStar(ctx *StrongStarContext) {}

// EnterStrongUl is called when production strongUl is entered.
func (s *BaseMarkdownParserListener) EnterStrongUl(ctx *StrongUlContext) {}

// ExitStrongUl is called when production strongUl is exited.
func (s *BaseMarkdownParserListener) ExitStrongUl(ctx *StrongUlContext) {}

// EnterImage is called when production image is entered.
func (s *BaseMarkdownParserListener) EnterImage(ctx *ImageContext) {}

// ExitImage is called when production image is exited.
func (s *BaseMarkdownParserListener) ExitImage(ctx *ImageContext) {}

// EnterImageLink is called when production imageLink is entered.
func (s *BaseMarkdownParserListener) EnterImageLink(ctx *ImageLinkContext) {}

// ExitImageLink is called when production imageLink is exited.
func (s *BaseMarkdownParserListener) ExitImageLink(ctx *ImageLinkContext) {}

// EnterExplicitImageLink is called when production explicitImageLink is entered.
func (s *BaseMarkdownParserListener) EnterExplicitImageLink(ctx *ExplicitImageLinkContext) {}

// ExitExplicitImageLink is called when production explicitImageLink is exited.
func (s *BaseMarkdownParserListener) ExitExplicitImageLink(ctx *ExplicitImageLinkContext) {}

// EnterImageAlt is called when production imageAlt is entered.
func (s *BaseMarkdownParserListener) EnterImageAlt(ctx *ImageAltContext) {}

// ExitImageAlt is called when production imageAlt is exited.
func (s *BaseMarkdownParserListener) ExitImageAlt(ctx *ImageAltContext) {}

// EnterReferenceImageLink is called when production referenceImageLink is entered.
func (s *BaseMarkdownParserListener) EnterReferenceImageLink(ctx *ReferenceImageLinkContext) {}

// ExitReferenceImageLink is called when production referenceImageLink is exited.
func (s *BaseMarkdownParserListener) ExitReferenceImageLink(ctx *ReferenceImageLinkContext) {}

// EnterLink is called when production link is entered.
func (s *BaseMarkdownParserListener) EnterLink(ctx *LinkContext) {}

// ExitLink is called when production link is exited.
func (s *BaseMarkdownParserListener) ExitLink(ctx *LinkContext) {}

// EnterLinkContent is called when production linkContent is entered.
func (s *BaseMarkdownParserListener) EnterLinkContent(ctx *LinkContentContext) {}

// ExitLinkContent is called when production linkContent is exited.
func (s *BaseMarkdownParserListener) ExitLinkContent(ctx *LinkContentContext) {}

// EnterExplicitLink is called when production explicitLink is entered.
func (s *BaseMarkdownParserListener) EnterExplicitLink(ctx *ExplicitLinkContext) {}

// ExitExplicitLink is called when production explicitLink is exited.
func (s *BaseMarkdownParserListener) ExitExplicitLink(ctx *ExplicitLinkContext) {}

// EnterLinkUrl is called when production linkUrl is entered.
func (s *BaseMarkdownParserListener) EnterLinkUrl(ctx *LinkUrlContext) {}

// ExitLinkUrl is called when production linkUrl is exited.
func (s *BaseMarkdownParserListener) ExitLinkUrl(ctx *LinkUrlContext) {}

// EnterLinkTitle is called when production linkTitle is entered.
func (s *BaseMarkdownParserListener) EnterLinkTitle(ctx *LinkTitleContext) {}

// ExitLinkTitle is called when production linkTitle is exited.
func (s *BaseMarkdownParserListener) ExitLinkTitle(ctx *LinkTitleContext) {}

// EnterLinkTitleSingle is called when production linkTitleSingle is entered.
func (s *BaseMarkdownParserListener) EnterLinkTitleSingle(ctx *LinkTitleSingleContext) {}

// ExitLinkTitleSingle is called when production linkTitleSingle is exited.
func (s *BaseMarkdownParserListener) ExitLinkTitleSingle(ctx *LinkTitleSingleContext) {}

// EnterLinkTitleDouble is called when production linkTitleDouble is entered.
func (s *BaseMarkdownParserListener) EnterLinkTitleDouble(ctx *LinkTitleDoubleContext) {}

// ExitLinkTitleDouble is called when production linkTitleDouble is exited.
func (s *BaseMarkdownParserListener) ExitLinkTitleDouble(ctx *LinkTitleDoubleContext) {}

// EnterReferenceLink is called when production referenceLink is entered.
func (s *BaseMarkdownParserListener) EnterReferenceLink(ctx *ReferenceLinkContext) {}

// ExitReferenceLink is called when production referenceLink is exited.
func (s *BaseMarkdownParserListener) ExitReferenceLink(ctx *ReferenceLinkContext) {}

// EnterEntity is called when production entity is entered.
func (s *BaseMarkdownParserListener) EnterEntity(ctx *EntityContext) {}

// ExitEntity is called when production entity is exited.
func (s *BaseMarkdownParserListener) ExitEntity(ctx *EntityContext) {}

// EnterHexEntityName is called when production hexEntityName is entered.
func (s *BaseMarkdownParserListener) EnterHexEntityName(ctx *HexEntityNameContext) {}

// ExitHexEntityName is called when production hexEntityName is exited.
func (s *BaseMarkdownParserListener) ExitHexEntityName(ctx *HexEntityNameContext) {}

// EnterDecEntityName is called when production decEntityName is entered.
func (s *BaseMarkdownParserListener) EnterDecEntityName(ctx *DecEntityNameContext) {}

// ExitDecEntityName is called when production decEntityName is exited.
func (s *BaseMarkdownParserListener) ExitDecEntityName(ctx *DecEntityNameContext) {}

// EnterCharEntityName is called when production charEntityName is entered.
func (s *BaseMarkdownParserListener) EnterCharEntityName(ctx *CharEntityNameContext) {}

// ExitCharEntityName is called when production charEntityName is exited.
func (s *BaseMarkdownParserListener) ExitCharEntityName(ctx *CharEntityNameContext) {}

// EnterEscapedChar is called when production escapedChar is entered.
func (s *BaseMarkdownParserListener) EnterEscapedChar(ctx *EscapedCharContext) {}

// ExitEscapedChar is called when production escapedChar is exited.
func (s *BaseMarkdownParserListener) ExitEscapedChar(ctx *EscapedCharContext) {}

// EnterCode is called when production code is entered.
func (s *BaseMarkdownParserListener) EnterCode(ctx *CodeContext) {}

// ExitCode is called when production code is exited.
func (s *BaseMarkdownParserListener) ExitCode(ctx *CodeContext) {}

// EnterBacktickCode is called when production backtickCode is entered.
func (s *BaseMarkdownParserListener) EnterBacktickCode(ctx *BacktickCodeContext) {}

// ExitBacktickCode is called when production backtickCode is exited.
func (s *BaseMarkdownParserListener) ExitBacktickCode(ctx *BacktickCodeContext) {}

// EnterDoubleBacktickCode is called when production doubleBacktickCode is entered.
func (s *BaseMarkdownParserListener) EnterDoubleBacktickCode(ctx *DoubleBacktickCodeContext) {}

// ExitDoubleBacktickCode is called when production doubleBacktickCode is exited.
func (s *BaseMarkdownParserListener) ExitDoubleBacktickCode(ctx *DoubleBacktickCodeContext) {}

// EnterSpaceBacktickCode is called when production spaceBacktickCode is entered.
func (s *BaseMarkdownParserListener) EnterSpaceBacktickCode(ctx *SpaceBacktickCodeContext) {}

// ExitSpaceBacktickCode is called when production spaceBacktickCode is exited.
func (s *BaseMarkdownParserListener) ExitSpaceBacktickCode(ctx *SpaceBacktickCodeContext) {}

// EnterAttributeName is called when production attributeName is entered.
func (s *BaseMarkdownParserListener) EnterAttributeName(ctx *AttributeNameContext) {}

// ExitAttributeName is called when production attributeName is exited.
func (s *BaseMarkdownParserListener) ExitAttributeName(ctx *AttributeNameContext) {}

// EnterHtmlAttributeS is called when production htmlAttributeS is entered.
func (s *BaseMarkdownParserListener) EnterHtmlAttributeS(ctx *HtmlAttributeSContext) {}

// ExitHtmlAttributeS is called when production htmlAttributeS is exited.
func (s *BaseMarkdownParserListener) ExitHtmlAttributeS(ctx *HtmlAttributeSContext) {}

// EnterHtmlAttributeD is called when production htmlAttributeD is entered.
func (s *BaseMarkdownParserListener) EnterHtmlAttributeD(ctx *HtmlAttributeDContext) {}

// ExitHtmlAttributeD is called when production htmlAttributeD is exited.
func (s *BaseMarkdownParserListener) ExitHtmlAttributeD(ctx *HtmlAttributeDContext) {}

// EnterHtmlBlockOpenDiv is called when production htmlBlockOpenDiv is entered.
func (s *BaseMarkdownParserListener) EnterHtmlBlockOpenDiv(ctx *HtmlBlockOpenDivContext) {}

// ExitHtmlBlockOpenDiv is called when production htmlBlockOpenDiv is exited.
func (s *BaseMarkdownParserListener) ExitHtmlBlockOpenDiv(ctx *HtmlBlockOpenDivContext) {}

// EnterHtmlBlockCloseDiv is called when production htmlBlockCloseDiv is entered.
func (s *BaseMarkdownParserListener) EnterHtmlBlockCloseDiv(ctx *HtmlBlockCloseDivContext) {}

// ExitHtmlBlockCloseDiv is called when production htmlBlockCloseDiv is exited.
func (s *BaseMarkdownParserListener) ExitHtmlBlockCloseDiv(ctx *HtmlBlockCloseDivContext) {}

// EnterHtmlBlockDiv is called when production htmlBlockDiv is entered.
func (s *BaseMarkdownParserListener) EnterHtmlBlockDiv(ctx *HtmlBlockDivContext) {}

// ExitHtmlBlockDiv is called when production htmlBlockDiv is exited.
func (s *BaseMarkdownParserListener) ExitHtmlBlockDiv(ctx *HtmlBlockDivContext) {}

// EnterHtmlBlockOpenSpan is called when production htmlBlockOpenSpan is entered.
func (s *BaseMarkdownParserListener) EnterHtmlBlockOpenSpan(ctx *HtmlBlockOpenSpanContext) {}

// ExitHtmlBlockOpenSpan is called when production htmlBlockOpenSpan is exited.
func (s *BaseMarkdownParserListener) ExitHtmlBlockOpenSpan(ctx *HtmlBlockOpenSpanContext) {}

// EnterHtmlBlockCloseSpan is called when production htmlBlockCloseSpan is entered.
func (s *BaseMarkdownParserListener) EnterHtmlBlockCloseSpan(ctx *HtmlBlockCloseSpanContext) {}

// ExitHtmlBlockCloseSpan is called when production htmlBlockCloseSpan is exited.
func (s *BaseMarkdownParserListener) ExitHtmlBlockCloseSpan(ctx *HtmlBlockCloseSpanContext) {}

// EnterHtmlBlockSpan is called when production htmlBlockSpan is entered.
func (s *BaseMarkdownParserListener) EnterHtmlBlockSpan(ctx *HtmlBlockSpanContext) {}

// ExitHtmlBlockSpan is called when production htmlBlockSpan is exited.
func (s *BaseMarkdownParserListener) ExitHtmlBlockSpan(ctx *HtmlBlockSpanContext) {}

// EnterHtmlBlockHr is called when production htmlBlockHr is entered.
func (s *BaseMarkdownParserListener) EnterHtmlBlockHr(ctx *HtmlBlockHrContext) {}

// ExitHtmlBlockHr is called when production htmlBlockHr is exited.
func (s *BaseMarkdownParserListener) ExitHtmlBlockHr(ctx *HtmlBlockHrContext) {}

// EnterHtmlBlockInTags is called when production htmlBlockInTags is entered.
func (s *BaseMarkdownParserListener) EnterHtmlBlockInTags(ctx *HtmlBlockInTagsContext) {}

// ExitHtmlBlockInTags is called when production htmlBlockInTags is exited.
func (s *BaseMarkdownParserListener) ExitHtmlBlockInTags(ctx *HtmlBlockInTagsContext) {}

// EnterHtmlBlockInSelfClosing is called when production htmlBlockInSelfClosing is entered.
func (s *BaseMarkdownParserListener) EnterHtmlBlockInSelfClosing(ctx *HtmlBlockInSelfClosingContext) {}

// ExitHtmlBlockInSelfClosing is called when production htmlBlockInSelfClosing is exited.
func (s *BaseMarkdownParserListener) ExitHtmlBlockInSelfClosing(ctx *HtmlBlockInSelfClosingContext) {}

// EnterHtmlComment is called when production htmlComment is entered.
func (s *BaseMarkdownParserListener) EnterHtmlComment(ctx *HtmlCommentContext) {}

// ExitHtmlComment is called when production htmlComment is exited.
func (s *BaseMarkdownParserListener) ExitHtmlComment(ctx *HtmlCommentContext) {}

// EnterAutolink is called when production autolink is entered.
func (s *BaseMarkdownParserListener) EnterAutolink(ctx *AutolinkContext) {}

// ExitAutolink is called when production autolink is exited.
func (s *BaseMarkdownParserListener) ExitAutolink(ctx *AutolinkContext) {}
