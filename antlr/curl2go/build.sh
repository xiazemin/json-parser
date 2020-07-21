http 协议

https://github.com/antlr/grammars-v4/blob/master/http/http.g4

java -jar ../calc/antlr-4.8-complete.jar -Dlanguage=Go -o parser curl.g4

https://github.com/mholt/curl-to-go/blob/gh-pages/resources/js/curl-to-go.js

https://github.com/moul/http2curl

/** Optional javadoc style comment */
grammar Name;
options {...}
import ... ;

tokens {...}
channels {...} // lexer only
@actionName {...}

rule1 // parser and lexer rules, possibly intermingled
...
ruleN


https://www.cnblogs.com/clonen/p/9083359.html


syntax error: missing COLON at 'longflag' while matching a rule

https://stackoverflow.com/questions/26405713/antlr4-lexer-rule-with-init-block

https://www.crifan.com/antlr_v3_syntax_fragment/

 ANTLR里，lexer和Token的名字须以大写字母开头；
 而非lexer的规则的名字则以小写开头。除了第一个字母外，
 标识符的其余字母则大小写均可，但是却只能使用ASCII字符。
https://blog.csdn.net/tangtao_xp/article/details/19004995

error(50): /Users/didi/PhpstormProjects/c/json-parser/antlr/curl2go/curl.g4:14:2: syntax error: '|' came as a complete surprise to me while matching rule preamble

options
  : options shortflag
  | options longflag
  | longflag
  | shortflag
  ;

  没有终结符

修改掉问题解决了


$java -jar ../calc/antlr-4.8-complete.jar -Dlanguage=Go -o parser curl.g4

$go mod init github.com/xiazemin/curl2go
go: creating new go.mod: module github.com/xiazemin/curl2go

$go mod edit -replace github.com/xiazemin/curl2go/parser=./parser

export GOPROXY=https://goproxy.cn

go run ./


STRING
   : '"' (ESC | SAFECODEPOINT)* '"'
   | '\'' (ESC | SAFECODEPOINT)* '\''
   ;

fragment ESC
   : '\\' (["\\/bfnrt] | UNICODE)
   ;
fragment UNICODE
   : 'u' HEX HEX HEX HEX
   ;
fragment HEX
   : [0-9a-fA-F]
   ;
fragment SAFECODEPOINT
   : ~ ["\\\u0000-\u001F]
   ;


https://www.it1352.com/1732989.html

go mod edit -replace github.com/xiazemin/curl2go/listen=./listen

https://www.jianshu.com/p/38487a311e8e

https://blog.csdn.net/yangguosb/article/details/85624640

error(50): /Users/didi/PhpstormProjects/c/json-parser/antlr/curl2go/curl.g4::: syntax error: mismatched character '<EOF>' expecting '''
error(50): /Users/didi/PhpstormProjects/c/json-parser/antlr/curl2go/curl.g4:36:4: syntax error: '<EOF>' came as a complete surprise to me while looking for rule element

原因终结符号没有大写开头

//https://github.com/antlr/grammars-v4/blob/master/php/PhpLexer.g4

