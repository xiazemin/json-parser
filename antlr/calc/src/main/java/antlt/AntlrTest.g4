grammar AntlrTest;
prog:stat+;

stat: expr NEWLINE                  # print
|ID '=' expr NEWLINE                # assign
|NEWLINE                            # blank
;

expr:expr (MULTIPLY|DIVIDE) expr    # MulDiv
|expr (PLUS|MINUS) expr             # AddSub
|'('expr')'                         # Private
|value                              # IdInt
;
value:INT
|ID
;

MULTIPLY:'*';
DIVIDE:'/';
PLUS:'+';
MINUS:'-';
ID:[a-z]+;
INT:[1-9]+;
NEWLINE:'\r'?'\n';
WS:[ \t\r\n] -> skip;