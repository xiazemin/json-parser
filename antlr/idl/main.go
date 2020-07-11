package main

import (
	"github.com/xiazemin/json-parser/antlr/idl/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)
func main()  {
	in:=antlr.NewInputStream(`
struct StructA {
  long a;
};

struct StructB : StructA {
  long b;
};

struct StructC : StructB {
};`)
lex:=parser.NewIDLLexer(in)
stream:=antlr.NewCommonTokenStream(lex,antlr.TokenDefaultChannel)
par:=parser.NewIDLParser(stream)
antlr.ParseTreeWalkerDefault.Walk(&parser.BaseIDLListener{},par.Specification())
}
