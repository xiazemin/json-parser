#!/bin/bash

#@echo off
#color 1E
#echo 当前目录:%~dp0
#cd "%~dp0"
#echo flex编译json.l生成lex.yy.c...
#flex json.l
#echo bison编译json.y生成json.tab.h、json.tab.c...
#bison -d json.y
#echo gcc编译json.tab.c输出json.exe...
#gcc json.tab.c -o json.exe
#echo 运行json.exe输入json.txt作为测试文件...
#echo 输出结果:
#json.exe<json.txt
#pause

flex json.l
#lex.yy.c

bison -d json.y

#json.tab.c	json.tab.h

gcc json.tab.c -o json.exe

./json.exe<../lex/json.json>json.yac



