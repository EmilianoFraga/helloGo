package main

import (
	"fmt"
)

func testMaps() {
	var personAge map[string]int

	fmt.Printf("len:%d %v\n", len(personAge), personAge)

	personAge = make(map[string]int)

	fmt.Printf("len:%d %v\n", len(personAge), personAge)

	personAge["Mark"] = 10
	personAge["Angela"] = 6
	personAge["John"] = 23
	personAge["Fabiana"] = 34

	fmt.Printf("len:%d %v\n", len(personAge), personAge)

	countryPopulation := map[string]int{
		"Angola":        31825295,
		"Rwanda":        12626950,
		"Côte d'Ivoire": 26008328,
		"Morocco":       36471769,
	}

	key := "Non existing"

	testRetrieveMapValue(countryPopulation, key)

	countryPopulation[key] = -1

	testRetrieveMapValue(countryPopulation, key)

	delete(countryPopulation, key)

	testRetrieveMapValue(countryPopulation, key)

}

func testRetrieveMapValue(m map[string]int, key string) {
	fmt.Printf("len:%d %v\n", len(m), m)

	var ok bool

	_, ok = m[key]

	if ok {
		fmt.Printf("Value for %s is %d\n", key, m[key])
	} else {
		fmt.Printf("Key \"%s\" not found\n", key)
	}
}
