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

# beign 和 end
awk '
BEGIN {
    FS = "\t"   # 设置分隔符
    printf("%10s %6s %5s    %s\n\n","COUNTURY","AREA","POP","COUNTINENT")
}
{
    printf("%10s %6d %5d    %s\n", $1, $2 , $3, $4) # 按照格式输出
    area = area + $2
    pop = pop + $3
}
END {
    printf("\n%10s %6d %5d\n", "TOTAL", area , pop)
}
' ./countries.data > ./countries_beginEnd.data
