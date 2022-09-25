package main

import (
	"sort"
)

type bundle struct {
	index  int
	source string
	target string
}

func findReplaceString(S string, indexes []int, sources []string, targets []string) string {
	bundles := make([]bundle, len(indexes))
	for i := 0; i < len(indexes); i++ {
		bundles[i] = bundle{
			index:  indexes[i],
			source: sources[i],
			target: targets[i],
		}
	}

	sort.Slice(
		bundles,
		func(i, j int) bool {
			return bundles[i].index > bundles[j].index
		},
	)

	for i := 0; i < len(bundles); i++ {
		bundle := bundles[i]
		if len(S) < bundle.index+len(bundle.source) || S[bundle.index:bundle.index+len(bundle.source)] != bundle.source {
			continue
		}
		S = S[:bundle.index] + bundle.target + S[bundle.index+len(bundle.source):]
	}

	return S
}

func main() {
	findReplaceString("abcd", []int{0, 2}, []string{"a", "cd"}, []string{"eee", "ffff"})
	findReplaceString("abcd", []int{0, 2}, []string{"ab", "ec"}, []string{"eee", "ffff"})
	findReplaceString("vmokgggqzp", []int{3, 5, 1}, []string{"kg", "ggq", "mo"}, []string{"s", "so", "bfr"})
}
