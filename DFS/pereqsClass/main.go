package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{

	"algorithms": {"data structures"}, "calculus": {"linear algebra"}, "compilers": {
		"data structures",
		"formal languages",
		"computer organization"},
	"data structures": {"discrete math"}, "databases": {"data structures"},
	"discrete math":    {"intro to programming"},
	"formal languages": {"discrete math"}, "networks": {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

var seen = make(map[string]bool)
var order []string

func visitAll(m map[string][]string, items []string) {
	for _, item := range items {
		if !seen[item] {
			seen[item] = true
			visitAll(m, m[item])
			order = append(order, item)
		}
	}
}

func topoSort(m map[string][]string) []string {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(m, keys)
	return order
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
