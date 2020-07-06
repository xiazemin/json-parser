#!/bin/bash
flex --prefix=yy --header-file=calc.lex.h -o calc.lex.c calc.l
goyacc -o calc.y.go -p "calc" calc.y