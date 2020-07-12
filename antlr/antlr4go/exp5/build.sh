go mod init github.com/xiazemin/json-parser/antlr/antlr4go/exp5
go mod edit -replace github.com/xiazemin/json-parser/antlr/antlr4go/parser=../parser
export GOPROXY=https://goproxy.cn
go run todot.go

#https://github.com/awalterschulze/gographviz