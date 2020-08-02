package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type genderRules struct {
	FirstName  genderRulesGroup
	MiddleName genderRulesGroup
	LastName   genderRulesGroup
}

type genderRulesGroup struct {
	Exceptions genderRule `json:"exceptions,omitempty"`
	Suffixes   genderRule `json:"suffixes,omitempty"`
}

type genderRule struct {
	Androgynous []string `json:"androgynous,omitempty"`
	Male        []string `json:"male,omitempty"`
	Female      []string `json:"female,omitempty"`
}

type gender struct {
	Gender genderRules `json:"gender"`
}

func main() {
	b, err := ioutil.ReadFile("rules/gender.json")
	if err != nil {
		panic(err)
	}

	var r gender

	err = json.Unmarshal(b, &r)
	if err != nil {
		panic(err)
	}

	o, err := os.Create("gender_generated.go")
	if err != nil {
		panic(err)
	}

	defer o.Close()

	fmt.Fprint(o, "// DO NOT EDIT!\n// Code generated from rules/gender.json\n\npackage petrovich\n\n")

	fmt.Fprint(o, "var gender = genderRules{\n\tFirstName: genderRulesGroup{\n")
	printRulesGroup(o, r.Gender.FirstName)

	fmt.Fprint(o, "\t},\n\tMiddleName: genderRulesGroup{\n")
	printRulesGroup(o, r.Gender.MiddleName)

	fmt.Fprint(o, "\t},\n\tLastName: genderRulesGroup{\n")
	printRulesGroup(o, r.Gender.LastName)

	fmt.Fprint(o, "\t},\n}\n")
}

func printRulesGroup(o io.Writer, g genderRulesGroup) {
	fmt.Fprint(o, "\t\tExceptions: genderRule{\n")
	printRule(o, g.Exceptions)
	fmt.Fprint(o, "\t\t},\n")

	fmt.Fprint(o, "\t\tSuffixes: genderRule{\n")
	printRule(o, g.Suffixes)
	fmt.Fprint(o, "\t\t},\n")
}

func printRule(o io.Writer, s genderRule) {
	fmt.Fprint(o, "\t\t\tAndrogynous: []string{\n")
	for _, t := range s.Androgynous {
		fmt.Fprintf(o, "\t\t\t\t\"%s\",\n", t)
	}
	fmt.Fprint(o, "\t\t\t\t},\n\t\t\tFemale: []string{\n")
	for _, t := range s.Female {
		fmt.Fprintf(o, "\t\t\t\t\"%s\",\n", t)
	}
	fmt.Fprint(o, "\t\t\t\t},\n\t\t\tMale: []string{\n")
	for _, t := range s.Male {
		fmt.Fprintf(o, "\t\t\t\t\"%s\",\n", t)
	}
	fmt.Fprint(o, "\t\t\t\t},\n")
}
