#!/bin/bash

# while 语句
awk '{
    print $0
    i = 1
    while (i <= $3) {
        printf("\t%.2f\n", $1 * (1 + $2) ^ i)
        i = i + 1
    }
}' ./rate.data > ./rate_while.data

# for 语句
 awk '{
     print $0
     for (i = 1 ; i <= $3; i = i + 1) {
         printf("\t%.2f\n", $1 * (1 + $2) ^ i)
     }
 }' ./rate.data > ./rate_for.data
