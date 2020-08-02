package petrovich

import (
	"strings"
)

type Name struct {
	First  string
	Middle string
	Last   string
}

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

func checkGenderExceptions(rg genderRulesGroup, str string) []Gender {
	result := []Gender{}
	for _, n := range rg.Exceptions.Male {
		if n == str {
			result = append(result, Male)
			break
		}
	}
	for _, n := range rg.Exceptions.Female {
		if n == str {
			result = append(result, Female)
			break
		}
	}
	for _, n := range rg.Exceptions.Androgynous {
		if n == str {
			result = append(result, Androgynous)
			break
		}
	}
	return result
}

func checkGenderSuffixes(rg genderRulesGroup, str string) []Gender {
	result := []Gender{}
	for _, suffix := range rg.Suffixes.Male {
		if strings.HasSuffix(str, suffix) {
			result = append(result, Male)
			break
		}
	}
	for _, suffix := range rg.Suffixes.Female {
		if strings.HasSuffix(str, suffix) {

			result = append(result, Female)
			break
		}
	}
	for _, suffix := range rg.Suffixes.Androgynous {
		if strings.HasSuffix(str, suffix) {
			result = append(result, Androgynous)
			break
		}
	}
	return result
}

func DetectGender(name Name) Gender {

	resFirst, resMiddle, resLast := []Gender{}, []Gender{}, []Gender{}

	if name.Middle != "" {
		resMiddle = checkGenderExceptions(gender.MiddleName, name.Middle)
		resMiddle = append(resMiddle, checkGenderSuffixes(gender.MiddleName, name.Middle)...)

		if len(resMiddle) > 0 && resMiddle[0] != Androgynous {
			return resMiddle[0]
		}
	}
	if name.First != "" {
		resFirst = checkGenderExceptions(gender.FirstName, name.First)
		resFirst = append(resFirst, checkGenderSuffixes(gender.FirstName, name.First)...)
	}
	if name.Last != "" {
		resLast = checkGenderExceptions(gender.LastName, name.Last)
		resLast = append(resLast, checkGenderSuffixes(gender.LastName, name.Last)...)
	}

	if len(resFirst) > 0 && len(resLast) > 0 {
		fn, ln := resFirst[0], resLast[0]
		if fn != Androgynous && ln == Androgynous {
			return fn
		}
		if ln != Androgynous && fn == Androgynous {
			return ln
		}
	}

	joined := append(append(resFirst, resMiddle...), resLast...)

	if len(joined) > 0 {
		return joined[0]
	}

	return Androgynous
}
