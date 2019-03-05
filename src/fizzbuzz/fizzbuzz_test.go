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
