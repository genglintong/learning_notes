## chapter2

### 2.1 模式

- 模式汇总
    1. BEGIN{ statements }<br> 
    在输入被读取之前，statements执行一次
    2. END{ statements } <br>
    当所有输入读取完毕后，statements被执行一次
    3. expression{ statements } <br>
    当expression为真时，statements被执行一次
    4. /regular expression/ or compound pattern or pattern1,pattern2 { statements } <br>
    当正则被匹配或者复合表达式为真或者范围匹配时，执行后面statements

- 将表达式用作模式
    
    |运算符|意义|
    |:------:|:-----:|
    |< |小于|
    |<=|小于或等于|
    |==|等于|
    |!=|不等于|
    |>=|大于或等于|
    |>|大于|
    |～|匹配|
    |!~|不匹配|

- 字符串匹配模式
    1. /regrxpr/ <br>
    当当前输入行包含一段被regrxpr匹配的子字符串时，该模式被匹配
    2. expression ~ /regexpr/ <br>
    expression 包含一段被regexpr匹配的子字符串时，该模式被匹配
    3. expression !~ /regexpr/ <br>
    不包含被regexpr匹配的子字符串时，模式被匹配

- 正则表达式
    1. 元字符<br>
    \ ^ $ . [  ] | (  ) * + ?
    2. 基本正则表达式
        1. 一个不是元字符的字符 eg. A
        2. 一个匹配特殊符号的转义字符 eg. \t
        3. 一个被引用的元字符 eg. \*
        4. ^ 匹配一行的开始
        5. $ 匹配一行的结束
        6. . 匹配任意一个字符
        7. 一个字符类[ABC] 匹配字符A B 或 C
        8. 字符类的缩写形式 [A-Za-z] 匹配单个字母
        9. 一个互补的字符类 [^0-9] 匹配任意一个不是数字的字符
    3. 运算符组合
        1. 选择:A|B 匹配A或B
        2. 拼接:AB 匹配后面紧跟B 的A
        3. 闭包:A* 匹配0个或者多个A
        4. 正闭包: A+ 匹配1个或者多个A
        5. 零或者一: A? 匹配空字符串或者A
        6. 括号:被(r)匹配的字符串，与r所匹配的字符串相同

- 复合模式
- 范围模式
由第一个匹配成功时开始，到第二个匹配成功时结束<br>
eg. FNR == 1, FNR == 5 {print FILRNAME ": " $0}
- 模式总结

|模式|例子|匹配|
|:------:|:------:|:------:|
|BEGIN|BEGIN|输入被读取之前|
|END|END|输入被读取完之后|
|expression|$3 < 100|第三个字段小于100的行|
|string-matching|/Asia/|含有Asia的行|
|compound|$3 < 100 && $4 == "Asia"|第三个字段小于100 且 第四个字段为Asia|
|range|NR == 10, NR == 20|输入的第10行至第20行|

### 2.2 动作
- 动作
expression 包括常量 变量 赋值 函数调用等等
```
print expression-list

printf(format, expression-list)

if (expression) 
    statements

if (expression) 
    statements 
else 
    statements

while (expression)
    statements

for (expression; expression; expression)
    statements

for (expression in array)
    statements

do 
    statements 
while (expression)

break 
continue
```
- 内建变量

|变量|意义|默认值|
|:------:|:-------:|:------:|
|ARGC|命令行参数个数|-|
|ARGV|命令行参数数组|-|
|FILENAME|当前输入文件名|-|
|FNR|当前输入文件记录个数|-|
|FS|控制输入行的字段分割符|" "|
|NF|当前记录的字段个数|-|
|NR|到目前位置读的记录数量|-|
|OFMT|数值的输出格式|"%.6g"|
|OFS|输出字段分隔符|" "|
|ORS|输出记录的分隔符|"\n"|
|RLENGTH|被函数match匹配的字符串长度|-|
|RS|控制输入行的记录分隔符|"\n"|
|RSTART|被函数match匹配的字符串的开始|-|
|SUBSEP|下标分隔符|"\034"|
