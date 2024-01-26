package main

import (
	"fmt"
	"golang-syntax-learning/customer"
	"unicode/utf8"
)

//GO have zero value not default value
var test int

func main() {
	declare()
	arrayList()
	dictMap()
	loop()
	c,d := add(10, 20)
	println("sum =", c,"?=", d)
	fmt.Println(sub(2, 5))
	fmt.Println(hello("k"))

	// Anonymous func
	sum := func(a,b int)int {
		return a + b
	}
	println("sum=",sum(5, 5))

	cal(sub)

	println(customer.Sum(10, 20))

	// Pointer
	x := 10
	y := &x
	fmt.Println(&x, y)
	fmt.Println("*y = ",*y)
	// var y *int
	*y = 20
	fmt.Println("x =",x)
	
	sumPointer(&x)
	println("sum pointer = ",x)

	// p1 := customer.Person{Name: "Bond", Age: 18}
	// p2 := customer.Person{
	// 	Age: 18, 
	// 	Name: "Mond",
	// }
	// fmt.Println("Person 1:",p1)
	// fmt.Println("Person 2:",p2)
	p1 := customer.Person{}
	p1.SetNameAge("Bond", 18)
	println(p1.GetName())
}

func sumPointer(result *int) {
	a := 10
	b := 20
	*result = a +b
}

func cal (f func(int,int) int) {
	fmt.Println(f(10, 20))
}

func declare() {
	println("Hello World!")
	//Value
	fmt.Printf("Hello %v\n",10)
	//Type
	fmt.Printf("%T\n", 10)

	//Short declared is only accept in {}
	y := 10
	println(test,y)
	
	if y >= 50 {
		test = 20
		println(test)
	} else {
		test = 25
	}
}

func arrayList() {
	// var x [3]int = [3]int{1}
	// flexible range array use [...] instead of [3]
	x := [3]int{}
	x[1] = 10
	fmt.Printf("%#v\n", x)

	// List call slice in GO
	slice := []int{1,2,3,4}
	t := append(slice,4)
	fmt.Println(t)
	fmt.Println(slice)
	println(len(t))

	// List slice
	fmt.Println(slice[1:3])

	// String count should use runecountinstring
	name := "test ฟ"
	count := utf8.RuneCountInString(name)
	println(count)
}

func dictMap() {
	// Map or dictionary
	dict := map[string]string{}
	dict["th"]="Thailand"
	dict["en"]="USA"
	
	a, b := 10,20
	println(a,b)

	dic, ok := dict["jp"]
	if ok {
		println(dic)
	} else {
		println("no value")
	}
}

func loop() {

	// For loop 
	list := []int{1,2,3,4,5}

	for i := 0; i < len(list); i++ {
		println("a",list[i])
	}

	for n,m := range list {
		println(n, m)
	}
}
// (int, int) int, int
func add(a int, b int) (int, int) {
	return a + b, 20
}
//(int,int)int
func sub(a, b int)int {
	return a -b 
}
//(string)string
func hello(name string) string  {
	return "Hello" +name
}