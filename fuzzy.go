package fuzzy

import (
	fzf "go.deanishe.net/fuzzy"
)

type Result struct {
	Match   bool
	Query   string
	Score   float64
	SortKey string
}

type Searcher struct {
	terms []string
	term  string
}

func NewSearcher(opts []FuzzyshOptions) (*Searcher, error) {

	c := &OptionsParams{}
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return &Searcher{
		terms: c.terms,
		term:  c.term,
	}, nil
}

func (f *Searcher) SetTerms(terms []string) {
	f.terms = terms
}

func (f *Searcher) SetTerm(term string) {
	f.term = term
}

func (f *Searcher) GetTerm() string {
	return f.term
}

func (f *Searcher) GetTerms() []string {
	return f.terms
}

func (f *Searcher) Search(term string) []Result {
	rr := []Result{}
	result := fzf.SortStrings(f.terms, term)

	for _, r := range result {
		rr = append(rr, Result{
			Match:   r.Match,
			Query:   r.Query,
			Score:   r.Score,
			SortKey: r.SortKey,
		})
	}

	return rr
}

func (f *Searcher) SearchWithTerms(terms []string, term string) []Result {
	rr := []Result{}
	result := fzf.SortStrings(terms, term)

	for _, r := range result {
		rr = append(rr, Result{
			Match:   r.Match,
			Query:   r.Query,
			Score:   r.Score,
			SortKey: r.SortKey,
		})
	}

	return rr
}
