package main

import (
	"fmt"
)

func main() {

	fz, err := fuzzy.NewSearcher(nil)
	if err != nil {
		panic(err)
	}

	fz.SetTerm("test")
	fz.SetTerms([]string{"test", "test2", "test3"})

	res := fz.Search(fz.GetTerm())

	for _, r := range res {
		fmt.Println(r)
	}

}
