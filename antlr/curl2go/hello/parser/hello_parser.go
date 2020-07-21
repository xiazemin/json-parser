// Code generated from hello.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // hello

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 5, 8, 4,
	2, 9, 2, 3, 2, 3, 2, 3, 2, 3, 2, 2, 2, 3, 2, 2, 2, 2, 6, 2, 4, 3, 2, 2,
	2, 4, 5, 7, 3, 2, 2, 5, 6, 7, 4, 2, 2, 6, 3, 3, 2, 2, 2, 2,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'hello'",
}
var symbolicNames = []string{
	"", "", "NAME", "WS",
}

var ruleNames = []string{
	"hello",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type helloParser struct {
	*antlr.BaseParser
}

func NewhelloParser(input antlr.TokenStream) *helloParser {
	this := new(helloParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "hello.g4"

	return this
}

// helloParser tokens.
const (
	helloParserEOF  = antlr.TokenEOF
	helloParserT__0 = 1
	helloParserNAME = 2
	helloParserWS   = 3
)

// helloParserRULE_hello is the helloParser rule.
const helloParserRULE_hello = 0

// IHelloContext is an interface to support dynamic dispatch.
type IHelloContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHelloContext differentiates from other interfaces.
	IsHelloContext()
}

type HelloContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHelloContext() *HelloContext {
	var p = new(HelloContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = helloParserRULE_hello
	return p
}

func (*HelloContext) IsHelloContext() {}

func NewHelloContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HelloContext {
	var p = new(HelloContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = helloParserRULE_hello

	return p
}

func (s *HelloContext) GetParser() antlr.Parser { return s.parser }

func (s *HelloContext) NAME() antlr.TerminalNode {
	return s.GetToken(helloParserNAME, 0)
}

func (s *HelloContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HelloContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HelloContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(helloListener); ok {
		listenerT.EnterHello(s)
	}
}

func (s *HelloContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(helloListener); ok {
		listenerT.ExitHello(s)
	}
}

func (p *helloParser) Hello() (localctx IHelloContext) {
	localctx = NewHelloContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, helloParserRULE_hello)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(2)
		p.Match(helloParserT__0)
	}
	{
		p.SetState(3)
		p.Match(helloParserNAME)
	}

	return localctx
}
