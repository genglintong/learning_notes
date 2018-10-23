#!/bin/bash

# 计算雇员工资

# 模式-动作 时长大于0时
awk '$3 > 0 {
    print $1, $2 * $3
}' emp.data > emp_paybill.data

# 精美输出 printf
awk '$3 > 0 {
printf("total pay for %s is $%.2f\n", $1, $2 * $3)
 }' emp.data > emp_paybill_printf.data

# 精美输出 printf
 awk '$3 > 0 {
 printf("%-8s $%6.2f\n", $1, $2 * $3)
  }' emp.data > emp_paybill_printf2.data

# 精美输出 printf
  awk '$3 > 0 {
  printf("%-8s $%6.2f %s\n", $1, $2 * $3,$2 * $3)
   }' emp.data | sort -k 3 -n > emp_paybill_sort.data

 # 打印工资大于50的员工
  awk '$3 * $2 > 50 {
  printf("%-8s $%6.2f\n", $1, $2 * $3)
   }' emp.data > emp_paybill_greater_50.data
