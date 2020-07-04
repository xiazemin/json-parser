%{
#define YYSTYPE char*
#include <stdio.h>
#include "lex.yy.c"
int yyparse(void);
int yyerror(const char* msg);
int yywrap();
%}
%token STR NUM DESC SPLIT ARRS OBJS ARRE OBJE FALSE TRUE NIL
%%
command : value {printf("end\n");}
value: STR {printf("value:[string]%s\n",$1);free($1);}
	|NUM {printf("value:[number]%s\n",$1);free($1);}
	|FALSE {printf("value:FALSE\n");}
	|TRUE {printf("value:TRUE\n");}
	|NIL {printf("value:NULL\n");}
	|object {printf("value:OBJECT\n");}
	|array {printf("value:ARRAY\n");}
;
arrs: ARRS {printf("array_start\n");};
array : arrs ARRE {printf("[empty]\n"); }
	| arrs values ARRE {printf("[...]\n"); }
;
objs: OBJS {printf("object_start\n");};
object : objs OBJE {printf("{empty}\n");}
	| objs pairs OBJE {printf("{...}\n");}
;
values : values SPLIT value {printf("add %s\n",$3);}
	|value {printf("add %s\n",$1);}
	|values SPLIT
;
pairs : pairs SPLIT pair {printf("put\n");}
	|pair {printf("put\n");}
	|pairs SPLIT
;
pair : STR DESC value {printf("key %s\n",$1);}
;
%%
int main()
{
    return yyparse();
}
int yyerror( const char* msg)
{
	fprintf (stderr,"%s\n",msg);
	return 0;
}
int yywrap()
{
    return 1;
}