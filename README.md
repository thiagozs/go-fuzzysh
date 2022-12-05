# go-fuzzysh -  wrapper for algorithm fuzzy search

Start with search

```golang
fz, err := fuzzy.NewSearcher(nil)
if err != nil {
        panic(err)
}

fz.SetTerm("test")
fz.SetTerms([]string{"test", "test2", "test3"})

res := fz.Search(fz.GetTerm())

for _, r := range res {
        fmt.Printf("%+v\n", r)
}
```
## Versioning and license

Our version numbers follow the [semantic versioning specification](http://semver.org/). You can see the available versions by checking the [tags on this repository](https://github.com/thiagozs/go-fuzzysh/tags). For more details about our license model, please take a look at the [LICENSE](LICENSE) file.

**2022**, thiagozs.