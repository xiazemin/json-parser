https://github.com/antlr/grammars-v4/blob/master/thrift/Thrift.g4
$java -jar ../calc/antlr-4.8-complete.jar -Dlanguage=Go -o parser Thrift.g4
error(134): Thrift.g4:51:0: symbol struct conflicts with generated code in target language or runtime
error(134): Thrift.g4:28:49: symbol struct conflicts with generated code in target language or runtime


https://tech.liuchao.me/2019/05/parse-thrift-in-go-the-antlr-way/

The struct definition from the canonical Thrift.g4 may cause exceptions when generating Golang runtimes,
 this can be fixed by changing struct to goStruct (or any other name you desire).

 $java -jar ../calc/antlr-4.8-complete.jar -Dlanguage=Go -o parser Thrift.g4

 成功

 $go mod init github.com/xiazemin/json-parser/antlr/thrift
go: creating new go.mod: module github.com/xiazemin/json-parser/antlr/thrift
$go mod edit -replace github.com/xiazemin/json-parser/antlr/thrift/parser=./parser

$go run example.go

//https://github.com/antlr/grammars-v4/tree/master/thrift