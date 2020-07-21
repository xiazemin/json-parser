// Code generated from curl.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // curl

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 12, 47, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 2, 7, 2, 14,
	10, 2, 12, 2, 14, 2, 17, 11, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7,
	2, 25, 10, 2, 12, 2, 14, 2, 28, 11, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 5, 2, 37, 10, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 5, 2, 2, 6, 2, 4, 6, 8, 2, 3, 3, 2, 8, 10, 2, 49, 2, 36, 3, 2, 2,
	2, 4, 38, 3, 2, 2, 2, 6, 41, 3, 2, 2, 2, 8, 44, 3, 2, 2, 2, 10, 15, 7,
	3, 2, 2, 11, 14, 5, 4, 3, 2, 12, 14, 5, 6, 4, 2, 13, 11, 3, 2, 2, 2, 13,
	12, 3, 2, 2, 2, 14, 17, 3, 2, 2, 2, 15, 13, 3, 2, 2, 2, 15, 16, 3, 2, 2,
	2, 16, 18, 3, 2, 2, 2, 17, 15, 3, 2, 2, 2, 18, 37, 7, 11, 2, 2, 19, 20,
	7, 3, 2, 2, 20, 21, 7, 4, 2, 2, 21, 26, 7, 5, 2, 2, 22, 25, 5, 4, 3, 2,
	23, 25, 5, 6, 4, 2, 24, 22, 3, 2, 2, 2, 24, 23, 3, 2, 2, 2, 25, 28, 3,
	2, 2, 2, 26, 24, 3, 2, 2, 2, 26, 27, 3, 2, 2, 2, 27, 29, 3, 2, 2, 2, 28,
	26, 3, 2, 2, 2, 29, 37, 7, 11, 2, 2, 30, 31, 7, 3, 2, 2, 31, 32, 7, 4,
	2, 2, 32, 33, 7, 5, 2, 2, 33, 37, 7, 11, 2, 2, 34, 35, 7, 3, 2, 2, 35,
	37, 7, 11, 2, 2, 36, 10, 3, 2, 2, 2, 36, 19, 3, 2, 2, 2, 36, 30, 3, 2,
	2, 2, 36, 34, 3, 2, 2, 2, 37, 3, 3, 2, 2, 2, 38, 39, 7, 6, 2, 2, 39, 40,
	5, 8, 5, 2, 40, 5, 3, 2, 2, 2, 41, 42, 7, 7, 2, 2, 42, 43, 5, 8, 5, 2,
	43, 7, 3, 2, 2, 2, 44, 45, 9, 2, 2, 2, 45, 9, 3, 2, 2, 2, 7, 13, 15, 24,
	26, 36,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'curl'", "'-X'", "'POST'",
}
var symbolicNames = []string{
	"", "", "", "", "SHORTCMD", "LONGCMD", "BackQuoteString", "SingleQuoteString",
	"DoubleQuoteString", "Url", "WS",
}

var ruleNames = []string{
	"curl", "shortflag", "longflag", "str",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type curlParser struct {
	*antlr.BaseParser
}

func NewcurlParser(input antlr.TokenStream) *curlParser {
	this := new(curlParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "curl.g4"

	return this
}

// curlParser tokens.
const (
	curlParserEOF               = antlr.TokenEOF
	curlParserT__0              = 1
	curlParserT__1              = 2
	curlParserT__2              = 3
	curlParserSHORTCMD          = 4
	curlParserLONGCMD           = 5
	curlParserBackQuoteString   = 6
	curlParserSingleQuoteString = 7
	curlParserDoubleQuoteString = 8
	curlParserUrl               = 9
	curlParserWS                = 10
)

// curlParser rules.
const (
	curlParserRULE_curl      = 0
	curlParserRULE_shortflag = 1
	curlParserRULE_longflag  = 2
	curlParserRULE_str       = 3
)

// ICurlContext is an interface to support dynamic dispatch.
type ICurlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCurlContext differentiates from other interfaces.
	IsCurlContext()
}

type CurlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCurlContext() *CurlContext {
	var p = new(CurlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = curlParserRULE_curl
	return p
}

func (*CurlContext) IsCurlContext() {}

func NewCurlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CurlContext {
	var p = new(CurlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = curlParserRULE_curl

	return p
}

func (s *CurlContext) GetParser() antlr.Parser { return s.parser }

func (s *CurlContext) Url() antlr.TerminalNode {
	return s.GetToken(curlParserUrl, 0)
}

func (s *CurlContext) AllShortflag() []IShortflagContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IShortflagContext)(nil)).Elem())
	var tst = make([]IShortflagContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IShortflagContext)
		}
	}

	return tst
}

func (s *CurlContext) Shortflag(i int) IShortflagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShortflagContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IShortflagContext)
}

func (s *CurlContext) AllLongflag() []ILongflagContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILongflagContext)(nil)).Elem())
	var tst = make([]ILongflagContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILongflagContext)
		}
	}

	return tst
}

func (s *CurlContext) Longflag(i int) ILongflagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILongflagContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILongflagContext)
}

func (s *CurlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CurlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CurlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(curlListener); ok {
		listenerT.EnterCurl(s)
	}
}

func (s *CurlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(curlListener); ok {
		listenerT.ExitCurl(s)
	}
}

func (p *curlParser) Curl() (localctx ICurlContext) {
	localctx = NewCurlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, curlParserRULE_curl)
	var _la int

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

	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(8)
			p.Match(curlParserT__0)
		}
		p.SetState(13)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == curlParserSHORTCMD || _la == curlParserLONGCMD {
			p.SetState(11)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case curlParserSHORTCMD:
				{
					p.SetState(9)
					p.Shortflag()
				}

			case curlParserLONGCMD:
				{
					p.SetState(10)
					p.Longflag()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(15)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(16)
			p.Match(curlParserUrl)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(17)
			p.Match(curlParserT__0)
		}
		{
			p.SetState(18)
			p.Match(curlParserT__1)
		}
		{
			p.SetState(19)
			p.Match(curlParserT__2)
		}
		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == curlParserSHORTCMD || _la == curlParserLONGCMD {
			p.SetState(22)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case curlParserSHORTCMD:
				{
					p.SetState(20)
					p.Shortflag()
				}

			case curlParserLONGCMD:
				{
					p.SetState(21)
					p.Longflag()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(26)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(27)
			p.Match(curlParserUrl)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(28)
			p.Match(curlParserT__0)
		}
		{
			p.SetState(29)
			p.Match(curlParserT__1)
		}
		{
			p.SetState(30)
			p.Match(curlParserT__2)
		}
		{
			p.SetState(31)
			p.Match(curlParserUrl)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(32)
			p.Match(curlParserT__0)
		}
		{
			p.SetState(33)
			p.Match(curlParserUrl)
		}

	}

	return localctx
}

// IShortflagContext is an interface to support dynamic dispatch.
type IShortflagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShortflagContext differentiates from other interfaces.
	IsShortflagContext()
}

type ShortflagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShortflagContext() *ShortflagContext {
	var p = new(ShortflagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = curlParserRULE_shortflag
	return p
}

func (*ShortflagContext) IsShortflagContext() {}

func NewShortflagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShortflagContext {
	var p = new(ShortflagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = curlParserRULE_shortflag

	return p
}

func (s *ShortflagContext) GetParser() antlr.Parser { return s.parser }

func (s *ShortflagContext) SHORTCMD() antlr.TerminalNode {
	return s.GetToken(curlParserSHORTCMD, 0)
}

func (s *ShortflagContext) Str() IStrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStrContext)
}

func (s *ShortflagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShortflagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShortflagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(curlListener); ok {
		listenerT.EnterShortflag(s)
	}
}

func (s *ShortflagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(curlListener); ok {
		listenerT.ExitShortflag(s)
	}
}

func (p *curlParser) Shortflag() (localctx IShortflagContext) {
	localctx = NewShortflagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, curlParserRULE_shortflag)

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
		p.SetState(36)
		p.Match(curlParserSHORTCMD)
	}
	{
		p.SetState(37)
		p.Str()
	}

	return localctx
}

// ILongflagContext is an interface to support dynamic dispatch.
type ILongflagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLongflagContext differentiates from other interfaces.
	IsLongflagContext()
}

type LongflagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLongflagContext() *LongflagContext {
	var p = new(LongflagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = curlParserRULE_longflag
	return p
}

func (*LongflagContext) IsLongflagContext() {}

func NewLongflagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LongflagContext {
	var p = new(LongflagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = curlParserRULE_longflag

	return p
}

func (s *LongflagContext) GetParser() antlr.Parser { return s.parser }

func (s *LongflagContext) LONGCMD() antlr.TerminalNode {
	return s.GetToken(curlParserLONGCMD, 0)
}

func (s *LongflagContext) Str() IStrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStrContext)
}

func (s *LongflagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LongflagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LongflagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(curlListener); ok {
		listenerT.EnterLongflag(s)
	}
}

func (s *LongflagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(curlListener); ok {
		listenerT.ExitLongflag(s)
	}
}

func (p *curlParser) Longflag() (localctx ILongflagContext) {
	localctx = NewLongflagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, curlParserRULE_longflag)

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
		p.SetState(39)
		p.Match(curlParserLONGCMD)
	}
	{
		p.SetState(40)
		p.Str()
	}

	return localctx
}

// IStrContext is an interface to support dynamic dispatch.
type IStrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStrContext differentiates from other interfaces.
	IsStrContext()
}

type StrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrContext() *StrContext {
	var p = new(StrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = curlParserRULE_str
	return p
}

func (*StrContext) IsStrContext() {}

func NewStrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrContext {
	var p = new(StrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = curlParserRULE_str

	return p
}

func (s *StrContext) GetParser() antlr.Parser { return s.parser }

func (s *StrContext) DoubleQuoteString() antlr.TerminalNode {
	return s.GetToken(curlParserDoubleQuoteString, 0)
}

func (s *StrContext) BackQuoteString() antlr.TerminalNode {
	return s.GetToken(curlParserBackQuoteString, 0)
}

func (s *StrContext) SingleQuoteString() antlr.TerminalNode {
	return s.GetToken(curlParserSingleQuoteString, 0)
}

func (s *StrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(curlListener); ok {
		listenerT.EnterStr(s)
	}
}

func (s *StrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(curlListener); ok {
		listenerT.ExitStr(s)
	}
}

func (p *curlParser) Str() (localctx IStrContext) {
	localctx = NewStrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, curlParserRULE_str)
	var _la int

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
		p.SetState(42)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<curlParserBackQuoteString)|(1<<curlParserSingleQuoteString)|(1<<curlParserDoubleQuoteString))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
