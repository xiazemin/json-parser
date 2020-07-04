%{
#include <string.h>
char* yylval;
%}
def [_a-zA-Z][_a-zA-Z0-9]*
str \"(\\\"|[^\"])*\"
num [0-9]+(\.[0-9]+)?
arrs \[
arre \]
objs \{
obje \}
split ,
desc :
%%
{str} { yylval = strdup(yytext);
printf("[STR:%s]",yylval);
free(yylval);}
{def} { yylval = strdup(yytext);
printf("[DEF:%s]",yylval);
free(yylval);}
{num} { yylval = strdup(yytext);
printf("[NUM:%s]",yylval);
free(yylval);}
{arrs} {printf("[ARRS]");}
{arre} {printf("[ARRE]");}
{objs} { printf("[OBJS]"); }
{obje} { printf("[OBJE]"); }
{desc} { printf("[DESC]"); }
{split} { printf("[SPLIT]");}
%%