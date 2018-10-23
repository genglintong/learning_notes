#!/bin/bash

# 输出 程序格式
awk '
{
    print \
        $1, # 国家名称
        $2, # 地域大小
        $3  # 国家人口
}
' ./countries.data > ./countries_output.data
