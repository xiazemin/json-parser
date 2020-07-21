// Code generated from hello.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // hello

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasehelloListener is a complete listener for a parse tree produced by helloParser.
type BasehelloListener struct{}

var _ helloListener = &BasehelloListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasehelloListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasehelloListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasehelloListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasehelloListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterHello is called when production hello is entered.
func (s *BasehelloListener) EnterHello(ctx *HelloContext) {}

// ExitHello is called when production hello is exited.
func (s *BasehelloListener) ExitHello(ctx *HelloContext) {}
