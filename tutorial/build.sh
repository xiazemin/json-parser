#!/bin/bash
cd json-tutorial/
cd tutorial01
mkdir build
cd build/
#cmake -DCMAKE_BUILD_TYPE=Debug ..
/Applications/CMake.app/Contents/bin/cmake  -DCMAKE_BUILD_TYPE=Debug ..
make
/Applications/CMake.app/Contents/bin/cmake  -DCMAKE_BUILD_TYPE=Release ..
make
./leptjson_test