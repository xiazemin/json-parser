// Code generated from curl.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // curl

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasecurlListener is a complete listener for a parse tree produced by curlParser.
type BasecurlListener struct{}

var _ curlListener = &BasecurlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasecurlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasecurlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasecurlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasecurlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCurl is called when production curl is entered.
func (s *BasecurlListener) EnterCurl(ctx *CurlContext) {}

// ExitCurl is called when production curl is exited.
func (s *BasecurlListener) ExitCurl(ctx *CurlContext) {}

// EnterShortflag is called when production shortflag is entered.
func (s *BasecurlListener) EnterShortflag(ctx *ShortflagContext) {}

// ExitShortflag is called when production shortflag is exited.
func (s *BasecurlListener) ExitShortflag(ctx *ShortflagContext) {}

// EnterLongflag is called when production longflag is entered.
func (s *BasecurlListener) EnterLongflag(ctx *LongflagContext) {}

// ExitLongflag is called when production longflag is exited.
func (s *BasecurlListener) ExitLongflag(ctx *LongflagContext) {}

// EnterStr is called when production str is entered.
func (s *BasecurlListener) EnterStr(ctx *StrContext) {}

// ExitStr is called when production str is exited.
func (s *BasecurlListener) ExitStr(ctx *StrContext) {}
