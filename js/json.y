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
command : value {printf("\treturn curValue;\n");}
value: STR {printf("\tcurValue=%s\n",$1);free($1);}
	|NUM {printf("\tcurValue=%s\n",$1);free($1);}
	|FALSE {printf("\tcurValue=false;\n");}
	|TRUE {printf("\tcurValue=true;\n");}
	|NIL {printf("\tcurValue=null;\n");}
	|object {printf("\tcurValue=curObj;\n\tcurObj=stack[stack.length-1];\n\n");}
	|array {printf("\tcurValue=curObj;\n\tcurObj=stack[stack.length-1];\n\n");}
;
arrs: ARRS {printf("\tcurObj=[];\n\tstack.push(curObj);\n");};
array : arrs ARRE {printf("\tcurObj=stack.pop();\n"); }
	| arrs values ARRE {printf("\tcurObj=stack.pop();\n"); }
;
objs: OBJS {printf("\tcurObj={};\n\tstack.push(curObj);\n");};
object : objs OBJE {printf("\tcurObj=stack.pop();\n");}
	| objs pairs OBJE {printf("\tcurObj=stack.pop();\n");}
;
values : values SPLIT value {printf("\tcurObj.push(curValue);\n");}
	|value {printf("\tcurObj.push(curValue);\n");}
	|values SPLIT
;
pairs : pairs SPLIT STR DESC value {printf("\tcurObj[%s]=curValue;\n",$3);}
	|STR DESC value {printf("\tcurObj[%s]=curValue;\n",$1);}
	|pairs SPLIT
;
;
%%
int main()
{
    printf("(function(){\n\tvar stack = [];\n\tvar curObj = null;\n\tvar curValue=null;\n\n");
    int r = yyparse();
    printf("})()");
    return r;
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