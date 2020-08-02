package petrovich

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
	Exceptions genderRule `json:"exceptions"`
	Suffixes   genderRule `json:"suffixes"`
}

type genderRule struct {
	Androgynous []string `json:"androgynous,omitempty"`
	Male        []string `json:"male,omitempty"`
	Female      []string `json:"female,omitempty"`
}

func DetectGender(name Name) Gender {
	return Androgynous
}
