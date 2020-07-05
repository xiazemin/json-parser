#!/bin/bash
bison -d num.y
flex num.l
gcc num.tab.c lex.yy.c  -o n