package main
import (
	"bramp.net/antlr4/json"
	"github.com/xiazemin/json-parser/antlr/thrift/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type exampleListener struct {
	// https://godoc.org/bramp.net/antlr4/json#BaseJSONListener
	*parser.BaseThriftListener
}

func (l*exampleListener)Start()  {

}

func main() {
	// Setup the input
	is := antlr.NewInputStream(`
namespace go go.antlr.thrift  
  
struct Request {  
    1: string username;        
    2: string password;             
}  
  
exception RequestException {  
    1: required i32 code;  
    2: optional string reason;  
}  
  
// 服务名  
service LoginService {  
    string doAction(1: Request request) throws (1:RequestException qe); // 可能抛出异常。  
} `)


	// Create the Thrift Lexer
	lexer := parser.NewThriftLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p:=parser.NewThriftParser(stream)
	// Finally walk the tree
	var listener exampleListener
	antlr.ParseTreeWalkerDefault.Walk(&listener{}, p.Start() )
	return listener.pop()
}
