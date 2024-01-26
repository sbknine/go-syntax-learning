package customer

//Expose var use Pascal case
var Name = "Customer"

type Person struct {
	name string
	age int
}

// Method
func (p Person) Hello() string {
	return "Hello " + p.name
}

func (p Person) GetName() string {
	return p.name
}

func (p *Person) SetNameAge(name string, age int) {
	p.name = name
	p.age = age
}


func Sum(a, b int) int {
	return a + b
}