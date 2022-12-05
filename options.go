package fuzzy

type FuzzyshOptions func(s *OptionsParams) error

type OptionsParams struct {
	terms []string
	term  string
}

func OptsTerms(terms []string) FuzzyshOptions {
	return func(c *OptionsParams) error {
		c.terms = terms
		return nil
	}
}

func OptsTerm(term string) FuzzyshOptions {
	return func(c *OptionsParams) error {
		c.term = term
		return nil
	}
}
