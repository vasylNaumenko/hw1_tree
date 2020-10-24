package main

import (
	"github.com/antonholmquist/jason"
	"log"
)

func main() {
	exampleJSON := `{
    "name": "Walter White",
    "age": 51,
    "children": [
      "junior",
      "holly"
    ],
    "other": {
      "occupation": "chemist",
      "years": 23
    }
  }`

	v, _ := jason.NewObjectFromBytes([]byte(exampleJSON))

	name, _ := v.GetString("name")
	age, _ := v.GetNumber("age")
	occupation, _ := v.GetString("other", "occupation")
	years, _ := v.GetNumber("other", "years")

	log.Println("age:", age)
	log.Println("name:", name)
	log.Println("occupation:", occupation)
	log.Println("years:", years)

	children, _ := v.GetStringArray("children")
	for i, child := range children {
		log.Printf("child %d: %s", i, child)
	}

	others, _ := v.GetObject("other")

	for _, value := range others.Map() {

		s, sErr := value.String()
		n, nErr := value.Number()

		if sErr == nil {
			log.Println("string value: ", s)
		} else if nErr == nil {
			log.Println("number value: ", n)
		}
	}
}
