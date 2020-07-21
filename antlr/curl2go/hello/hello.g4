grammar hello;

hello
 : 'hello' NAME
 ;

NAME: [a-z]+ ;

WS
 : [ \t\n\r] + -> skip
 ;