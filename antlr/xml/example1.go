package main

import(
	"github.com/antlr/antlr4/runtime/Go/antlr"
   "github.com/xiazemin/json-parser/antlr/xml/parser"
)

type listener struct {
	*parser.BaseXMLParserListener
}
func main()  {
    in:=antlr.NewInputStream(`
 <xml>
<key>k1</key>
<value>v1</value>
</xml>
`)
lex:=parser.NewXMLLexer(in)
stream:=antlr.NewCommonTokenStream(lex,antlr.TokenDefaultChannel)
parser:=parser.NewXMLParser(stream)
antlr.ParseTreeWalkerDefault.Walk(&listener{},parser.Document())
}