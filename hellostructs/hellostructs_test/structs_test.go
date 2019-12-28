package hellostructs_test

import (
	"github.com/EmilianoFraga/helloGo/hellostructs"
	"testing"
)

func TestNewInstance(t *testing.T) {
	// this gives visibility error
	// p := hellostructs.Person {name: "Maria", age: expectedAge}
	var expectedName string = "Maria"
	var expectedAge uint8 = 32

	var p hellostructs.Person
	p = hellostructs.CreateNewPerson(expectedName, expectedAge)

	if p.GetAge() != expectedAge {
		t.Errorf("Invalid age %d expected as %d", p.GetAge(), expectedAge)
	}

	if p.GetName() != expectedName {
		t.Errorf("Invalid age %s expected as %s", p.GetName(), expectedName)
	}
}

func TestMakeOlder(t *testing.T) {
	var expectedName string = "Maria"
	var originalAge uint8 = 32

	var p hellostructs.Person
	p = hellostructs.CreateNewPerson(expectedName, originalAge)

	hellostructs.MakePersonOlderByValue(p)

	if p.GetAge() != originalAge {
		t.Errorf("Invalid age %d expected as %d", p.GetAge(), originalAge)
	}

	hellostructs.MakePersonOlderByRef(&p)

	if p.GetAge() != originalAge+1 {
		t.Errorf("Invalid age %d expected as %d", p.GetAge(), originalAge+1)
	}

	p.MakePersonOlder()

	if p.GetAge() != originalAge+2 {
		t.Errorf("Invalid age %d expected as %d", p.GetAge(), originalAge+2)
	}

	p.MakePersonYounger()
	p.MakePersonYounger()

	if p.GetAge() != originalAge {
		t.Errorf("Invalid age %d expected as %d", p.GetAge(), originalAge)
	}
}
