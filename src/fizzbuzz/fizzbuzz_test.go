package fizzbuzz

import "testing"
//import "fmt"

func Test_IsFizz(t *testing.T) {

    cases := []struct {
        name string
        in  int
        out bool
    }{
        {name:"Fizz False", in: 1, out: false,},
        {name:"Fizz True", in: 3, out: true,},
        {name:"Fizz False", in: 5, out: false,},
        {name:"Fizz True", in: 6, out: true,},
        {name:"Fizz False", in: 10, out: false,},
        {name:"Fizz True", in: 15, out: true,},
    }

    for _, tt := range cases {
       st := tt
       t.Run(st.name, func(t *testing.T){
           s := IsFizz(st.in)
           if s != st.out {
               t.Fatal("出力値が異なる[", st.in, "]=>[", s, "] [" , st.out , "]")
           }
       })
    }
}

func Test_IsBuzz(t *testing.T) {

    cases := []struct {
        name string
        in  int
        out bool
    }{
        {name:"Fizz False", in: 1, out: false,},
        {name:"Fizz False", in: 3, out: false,},
        {name:"Fizz True", in: 5, out: true,},
        {name:"Fizz False", in: 6, out: false,},
        {name:"Fizz True", in: 10, out: true,},
        {name:"Fizz True", in: 15, out: true,},
    }

    for _, tt := range cases {
       st := tt
       t.Run(st.name, func(t *testing.T){
           s := IsBuzz(st.in)
           if s != st.out {
               t.Fatal("出力値が異なる[", st.in, "]=>[", s, "] [" , st.out , "]")
           }
       })
    }
}

func Test_FizzBuzz(t *testing.T) {

    cases := []struct {
        name string
        in  int
        out string
    }{
        {name:"1 is [1]", in: 1, out: "1",},
        {name:"3 is [Fizz]", in: 3, out: "Fizz",},
        {name:"5 is [Buzz]", in: 5, out: "Buzz",},
        {name:"15 is [FizzBuzz]", in: 15, out: "FizzBuzz",},
        {name:"18 is [Fizz]", in: 18, out: "Fizz",},
        {name:"19 is [19]", in: 19, out: "19",},
        {name:"20 is [Buzz]", in: 20, out: "Buzz",},
        {name:"30 is [FizzBuzz]", in: 30, out: "FizzBuzz",},
    }

    for _, tt := range cases {
       st := tt
       t.Run(st.name, func(t *testing.T){
           s := FizzBuzz(st.in)
           if s != st.out {
               t.Fatal("出力値が異なる[", st.in, "]=>[", s, "] [" , st.out , "]")
           }
       })
    }
}
