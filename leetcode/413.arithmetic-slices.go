package main

func numberOfArithmeticSlices(A []int) int {

    sum := 0;

    if len(A) < 3 {
        return 0
    }

    d := A[1] - A[0];
    z := 0;
    t := 0;

    for i := 2 ; i < len(A); i++ {
        t = A[i] - A[i-1];
        if  t == d {
            z += 1
            sum += z
        } else {
            z = 0
            d = t
        }
    }

    return sum
}
