package main

import (
	"fmt"
	"log"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xiazemin/json-parser/antlr/visitor/parser"
)

func main() {
	fs, err := antlr.NewFileStream("input_file")
	if err != nil {
		log.Fatal(err)
	}
	lex := parser.NewJSONLexer(fs)
	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewJSONParser(tokens)
	visitor := parser.BaseJSONVisitor{}
	tree := p.Json()
	fmt.Println(visitor.Visit(tree))
}
