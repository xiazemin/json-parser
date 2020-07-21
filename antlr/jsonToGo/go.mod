module github.com/xiazemin/json-parser/antlr/jsonToGo

go 1.13

replace github.com/xiazemin/json-parser/antlr/jsonToGo/parser => ./parser

require github.com/antlr/antlr4 v0.0.0-20200712162734-eb1adaa8a7a6

replace github.com/xiazemin/json-parser/antlr/jsonToGo/listener => ./listener

replace github.com/xiazemin/json-parser/antlr/jsonToGo/file => ./file
