#!/bin/bash
lex -o ./json_lex.c ./json.l

#gcc -o ./json_lex.exe ./json_lex.c -lfl
#ld: library not found for -lfl

#https://stackoverflow.com/questions/21298097/library-not-found-for-lfl
#gcc -o ./json_lex.exe ./json_lex.c -ll

gcc -o ./json_lex.exe ./json_lex.c -ll

#@echo off
# color 1E
# echo 当前目录:%~dp0
# cd "%~dp0"
# echo flex编译testjson.l生成testjson.c...
# flex -o testjson.c testjson.l
# echo gcc编译testjson.c链接libfl.a输出testjson.exe...
# gcc -o testjson.exe testjson.c -L. -lfl
# echo 运行testjson.exe输入json.txt作为测试文件...
# echo 输出结果:
# testjson.exe<json.txt
# pause

./json_lex.exe<./json.json >./json.token
