package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xiazemin/curl2go/listen"
	"github.com/xiazemin/curl2go/parser"
)

func main()  {
	//curl:=`curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer b7d03a6947b217efb6f3ec3bd3504582" -d '{"type":"A","name":"www","data":"162.10.66.0","priority":null,"port":null,"weight":null}' "https://api.digitalocean.com/v2/domains/example.com/records"`
	curl:=`curl "https://api.digitalocean.com/v2/domains/example.com/records"`

	in:=antlr.NewInputStream(curl)
  lex:=parser.NewcurlLexer(in)
  ins:=antlr.NewCommonTokenStream(lex,antlr.TokenDefaultChannel)
  yac:=parser.NewcurlParser(ins)
  var l listen.Listen
  antlr.ParseTreeWalkerDefault.Walk(l,yac.Curl())
}


