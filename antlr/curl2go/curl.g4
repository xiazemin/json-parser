grammar curl;

curl
  : 'curl' (shortflag|longflag)*  Url
  | 'curl' '-X' 'POST' (shortflag | longflag)* Url
  | 'curl' '-X' 'POST' Url
  | 'curl' Url
  ;

shortflag
  : SHORTCMD str
  ;

longflag
  : LONGCMD str
  ;

SHORTCMD
  : '-'[a-zA-Z]+
  ;

LONGCMD
  : '--'[a-zA-Z]+'='
  ;

str
  : DoubleQuoteString
  | BackQuoteString
  | SingleQuoteString
  ;

BackQuoteString:  '`' ~'`'* '`';
SingleQuoteString: '\'' (~('\'' | '\\') | '\\' . )* '\'';
DoubleQuoteString: '"'('\\"'|~'"')* '"'; // quote-quote is an escaped quote

Url
 : '"' ('http://' | 'https://') [a-zA-Z0-9.?/]+ '"'
 | '\'' Url '\''
 | ('http://' | 'https://') [a-zA-Z0-9.?/]+
 ;

WS
  : [ ]+ -> skip
  ;
/*
WS
   : [ \t\n\r] + -> skip
   ;
 */