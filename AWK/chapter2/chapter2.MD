## chapter2
- 模式汇总
    1. BEGIN{ statements }<br> 
    在输入被读取之前，statements执行一次
    2. END{ statements } <br>
    当所有输入读取完毕后，statements被执行一次
    3. expression{ statements } <br>
    当expression为真时，statements被执行一次
    4. /regular expression/ or compound pattern or pattern1,pattern2 { statements } <br>
    当正则被匹配或者复合表达式为真或者多个匹配为真时，执行后面statements
