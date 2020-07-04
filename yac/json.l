%{
#include <string.h>
#include "json.tab.h"
extern char* yylval;
%}
space [ \t\n]+
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
return STR; }
{def} {
	if(strcmp(yytext,"true")==0){
		return TRUE;
	}else if(strcmp(yytext,"false")==0){
		return FALSE;
	}else if(strcmp(yytext,"null")==0){
		return NIL;
	}else{
		yylval = strdup(yytext);
		return STR;
	}
}
{num} { yylval = strdup(yytext);
return NUM; }
{arrs} { return ARRS; }
{arre} { return ARRE; }
{objs} { return OBJS; }
{obje} { return OBJE; }
{desc} { return DESC; }
{split} { return SPLIT; }
{space} {}
%%