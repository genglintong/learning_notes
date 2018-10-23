#!/bin/bash

# begin end

awk '
BEGIN {
    print "NAME RATE    HOURS"
    print ""
}
$2 > maxrate{
    maxrate = $2
    maxemp = $1
}
{
    if ($3>15)
        emp = emp + 1
    pay = pay + $2 * $3
    print $0
}
END {
    print emp, "雇员大于15小时"
    if (NR > 0)
        print NR, "雇员总数"
        print pay, "工资总数"
        print pay/NR, "平均工资"
    print "最高时薪", maxrate, "姓名", maxemp
}
' ./emp.data > ./emp_begin.data
