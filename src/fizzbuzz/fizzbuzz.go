package fizzbuzz

import (
//    "fmt"
)

// Is Fizz is check function.
func IsFizz(val int) bool {
    isFizz := val % 3 == 0
    //isFizz := true
    return isFizz
}

// Is Buzz is check function.
func IsBuzz(val int) bool {
    isBuzz := val % 5 == 0
    //isBuzz := true
    return isBuzz
}
