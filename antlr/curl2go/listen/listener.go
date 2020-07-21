package listen

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xiazemin/curl2go/parser"
)

type Listen struct{
	* parser.BasecurlListener
}

func (l Listen) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println("EnterEveryRule",ctx.GetText(),ctx.GetChildCount())
}


func (l*Listen) EnterCurl(ctx *parser.CurlContext) {
	fmt.Println("EnterCurl",ctx.GetText())
}
// ExitCurl is called when production curl is exited.
func (l*Listen) ExitCurl(ctx *parser.CurlContext) {
	fmt.Println("ExitCurl",ctx.GetText())
}
// ExitContent is called when production content is exited.
//func (l*Listen) ExitContent(ctx *parser.ContentContext) {
//	fmt.Println("ExitContent",ctx.GetText())
//}
// ExitShortflag is called when production shortflag is exited.
func (l*Listen) ExitShortflag(ctx *parser.ShortflagContext) {
	fmt.Println("short",ctx.GetText())
}
// ExitLongflag is called when production longflag is exited.
func (l*Listen) ExitLongflag(ctx *parser.LongflagContext) {
	fmt.Println("long",ctx.GetText())
}

func (l*Listen) EnterStr(ctx *parser.StrContext) {
	fmt.Println("EnterStr",ctx.GetText())
}

// ExitStr is called when production str is exited.
func (l*Listen) ExitStr(ctx *parser.StrContext) {
	fmt.Println("ExitStr",ctx.GetText())
}
