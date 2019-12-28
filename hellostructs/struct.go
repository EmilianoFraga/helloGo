package hellostructs

// Person is a sample struct
type Person struct {
	name string
	age  uint8
}

// GetName returns the person name
func (p *Person) GetName() string {
	return p.name
}

// GetAge returns the person name
func (p *Person) GetAge() uint8 {
	return p.age
}

// MakePersonOlder inevitable
func (p *Person) MakePersonOlder() {
	p.age++
}

// MakePersonYounger miracle
func (p *Person) MakePersonYounger() {
	p.age--
}

// CreateNewPerson is a person instance creator
func CreateNewPerson(name string, age uint8) Person {
	return Person {name: name, age: age}
}

// MakePersonOlderByRef inevitable
func MakePersonOlderByRef(p *Person) {
	p.age++
}

// MakePersonOlderByValue this doesn't work
func MakePersonOlderByValue(p Person) {
	p.age++
}
