package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	//"testing"

	"github.com/xiazemin/json-parser/antlr/json2xml/go/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type j2xConvert struct {
	*parser.BaseJSONListener
	xml map[antlr.Tree]string
}

func NewJ2xConvert() *j2xConvert {
	return &j2xConvert{
		&parser.BaseJSONListener{},
		make(map[antlr.Tree]string),
	}
}

func (j *j2xConvert) setXML(ctx antlr.Tree, s string) {
	j.xml[ctx] = s
}

func (j *j2xConvert) getXML(ctx antlr.Tree) string {
	return j.xml[ctx]
}

// j2xConvert methods
func (j *j2xConvert) ExitJson(ctx *parser.JsonContext) {
	j.setXML(ctx, j.getXML(ctx.GetChild(0)));
}

func (j *j2xConvert) stripQuotes(s string) string {
	if s == "" || ! strings.Contains(s, "\"") {
		return s
	}
	return s[1 : len(s)-1]
}

func (j *j2xConvert) ExitAnObject(ctx *parser.AnObjectContext) {
	sb := strings.Builder{}
	sb.WriteString("\n")
	for _, p := range ctx.AllPair() {
		sb.WriteString(j.getXML(p))
	}
	j.setXML(ctx, sb.String())
}

func (j *j2xConvert) ExitNullObject(ctx *parser.NullObjectContext) {
	j.setXML(ctx, "")
}

func (j *j2xConvert) ExitArrayOfValues(ctx *parser.ArrayOfValuesContext) {
	sb := strings.Builder{}
	sb.WriteString("\n")
	for _, p := range ctx.AllValue() {
		sb.WriteString("<element>")
		sb.WriteString(j.getXML(p))
		sb.WriteString("</element>")
		sb.WriteString("\n")
	}
	j.setXML(ctx, sb.String())
}

func (j *j2xConvert) ExitNullArray(ctx *parser.NullArrayContext) {
	j.setXML(ctx, "")
}

func (j *j2xConvert) ExitPair(ctx *parser.PairContext) {
	tag := j.stripQuotes(ctx.STRING().GetText())
	v := ctx.Value()
	r := fmt.Sprintf("<%s>%s</%s>\n", tag, j.getXML(v), tag)
	j.setXML(ctx, r)
}

func (j *j2xConvert) ExitObjectValue(ctx *parser.ObjectValueContext) {
	j.setXML(ctx, j.getXML(ctx.Object()))
}

func (j *j2xConvert) ExitArrayValue(ctx *parser.ArrayValueContext) {
	j.setXML(ctx, j.getXML(ctx.Array()))
}

func (j *j2xConvert) ExitAtom(ctx *parser.AtomContext) {
	j.setXML(ctx, ctx.GetText())
}

func (j *j2xConvert) ExitString(ctx *parser.StringContext) {
	j.setXML(ctx, j.stripQuotes(ctx.GetText()))
}

func main(){
	f, err := os.Open("./t.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	// Setup the input
	is := antlr.NewInputStream(string(content))

	// Create lexter
	lexer := parser.NewJSONLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)

	// Create parser and tree
	p := parser.NewJSONParser(stream)
	p.BuildParseTrees = true
	tree := p.Json()

	// Finally AST tree
	j2x := NewJ2xConvert()
	antlr.ParseTreeWalkerDefault.Walk(j2x, tree)
	log.Println(j2x.getXML(tree))
}