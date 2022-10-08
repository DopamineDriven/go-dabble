package main

import (
	"fmt"
)

type Maths interface {
	Println(x float64, y float64) float64
}

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type Make interface {
	make(map[string]int)
};



func main() {
// var p =go3d(new()) UnitXYZ[int,int,int];
	b := make(map[string]int)

	f := fib()

	b["k1"] = 7
	b["k2"] = 13

	fmt.Println("map:", b)

	v1 := b["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(b))

	delete(b, "k2")
	fmt.Println("map:", b)

	_, prs := b["k2"]
	fmt.Println("prs:", prs)

	n := map[int]int{f():0, f():1, f():2,f():3,f():4,f():5,f():6, f():7, f():8};

	fmt.Println("map:", n)
	// Function calls are evaluated left-to-right.
	fmt.Println(f(), f(), f(), f(), f(), f(), f(), f(), f(), f(), f(), f())

}


// func main() (
//   vvv := make(map[string]int)

//    type Make interface {
// 	typeof(make(map[string]int))
//    }

//     mm["k1"] = 7,
//     mm["k2"] = 13

//     fmt.Println("map:", mm)

//     v1 := mm["k1"]
//     fmt.Println("v1: ", v1)

//     fmt.Println("len:", len(mm))

//     delete(m, "k2")
//     fmt.Println("map:", m)

//     _, prs := m["k2"]
//     fmt.Println("prs:", prs)

//     n := map[string]int{"foo": 1, "bar": 2}
//     fmt.Println("map:", n)
// }
// func main() {

// 	fmt.Println(math.Pow(2,-2.242123213))
// }

// fib returns a function that returns
// successive Fibonacci numbers.