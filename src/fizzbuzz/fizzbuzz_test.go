package fizzbuzz

import "testing"
//import "fmt"

func Test_IsFizz(t *testing.T) {

    cases := []struct {
        name string
        in  int
        out bool 
    }{
        {name:"Fizz True", in: 3, out: true,},
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
