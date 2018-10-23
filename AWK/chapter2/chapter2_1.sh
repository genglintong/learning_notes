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

# 计算亚洲总人口
awk '
    $4 == "Asia" {
        pop += $3
        n ++
    }
END {
    print "亚洲国家总数:",n,"人口总数:",pop
}
' ./countries.data

 # 计算最大人口和国家
 awk '
     $3 > maxpop {
         maxpop = $3
         country = $1
     }
 END {
     print "最大人口:",maxpop,"所属国家:",country
 }
 ' ./countries.data
