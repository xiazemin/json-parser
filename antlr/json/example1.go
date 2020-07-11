package main

import (
	"github.com/xiazemin/json-parser/antlr/json/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type listener struct {
	*parser.BaseJSONListener
}
func main()  {
	in:=antlr.NewInputStream(
		`{"k1":1,"k2":["v1","v2"]}`,
		)
	lex:=parser.NewJSONLexer(in)
	stream:=antlr.NewCommonTokenStream(lex,antlr.TokenDefaultChannel)
	par:=parser.NewJSONParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(&listener{},par.Json())
}
