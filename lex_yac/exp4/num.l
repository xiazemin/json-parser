%{
	#include "num.tab.h"
	#include <math.h>
	extern double vbltable[26];
%}
%option noyywrap
%%
([0-9]+|([0-9]*\.[0-9]+)([eE][-+]?[0-9]+)?)	{
		yylval.dval = atof(yytext);
		return NUMBER;
	}
[ \t]	;		/*忽略空白*/
[a-z]	{	yylval.vblno = yytext[0] - 'a';
			return NAME;
		}
"$"		{	return 0; /*输入结束*/ }
\n		|
.		return yytext[0];
%%