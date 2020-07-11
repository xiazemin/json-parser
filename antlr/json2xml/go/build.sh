$go mod init github.com/xiazemin/json-parser/antlr/json/parser
go: creating new go.mod: module github.com/xiazemin/json-parser/antlr/json/parser

$go mod init github.com/xiazemin/json-parser/antlr/json2xml/go
go: creating new go.mod: module github.com/xiazemin/json-parser/antlr/json2xml/go

$go mod edit -replace github.com/xiazemin/json-parser/antlr/json/parser=../../json/parser/

$go run main.go


注意文法和标准json相比多了  # AnObject
直接编译是没法通过的

$java -jar ../../calc/antlr-4.8-complete.jar -Dlanguage=Go -o parser JSON.g4

$go mod edit -replace github.com/xiazemin/json-parser/antlr/json2xml/go/parser=./parser

$go run main.go >result.xml


line 1:1 token recognition error at: '
'
line 2:0 token recognition error at: ' '
line 2:1 token recognition error at: ' '
line 2:15 token recognition error at: ' '
line 2:17 token recognition error at: ' '
line 2:52 token recognition error at: '
'
line 3:0 token recognition error at: ' '
line 3:1 token recognition error at: ' '
line 3:8 token recognition error at: ' '
line 3:10 token recognition error at: ' '
line 3:30 token recognition error at: ' '
line 3:49 token recognition error at: '
'
line 4:0 token recognition error at: ' '
line 4:1 token recognition error at: ' '
line 4:8 token recognition error at: ' '
line 4:10 token recognition error at: ' '
line 4:23 token recognition error at: '
'
line 5:0 token recognition error at: ' '
line 5:1 token recognition error at: ' '
line 5:10 token recognition error at: ' '
line 5:20 token recognition error at: ' '
line 5:30 token recognition error at: '
'
line 6:0 token recognition error at: ' '
line 6:1 token recognition error at: ' '
line 6:12 token recognition error at: ' '
line 6:15 token recognition error at: '
'
2020/07/11 12:05:53 main.go:131:
<description>An imaginary server config file</description>
<logs>
<level>verbose</level>
<dir>/var/log</dir>
</logs>
<host>antlr.org</host>
<admin>
<element>parrt</element>
<element>tombu</element>
</admin>
<aliases></aliases>

{
  "description" : "An imaginary server config file",
  "logs" : {"level":"verbose", "dir":"/var/log"},
  "host" : "antlr.org",
  "admin": ["parrt", "tombu"],
  "aliases": []
}