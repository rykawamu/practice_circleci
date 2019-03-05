package fizzbuzz

import (
//    "fmt"
    "strconv"
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

func FizzBuzz(val int) string {
    isFizz := IsFizz(val)
    isBuzz := IsBuzz(val)

    s := strconv.Itoa(val)
    switch {
    case isFizz && isBuzz: s = "FizzBuzz"
    case isFizz : s = "Fizz"
    case isBuzz : s = "Buzz"
    }
    return s
}
