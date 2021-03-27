package main

import (
	"fmt"
	"sort"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")}, {"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")}, {"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")}}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (c customSort) Len() int {
	return len(c.t)
}
func (c customSort) Less(i, j int) bool {
	return c.less(c.t[i], c.t[j])
}

func (c customSort) Swap(i, j int) {
	c.t[i], c.t[j] = c.t[j], c.t[i]
}

func main() {
	//sort.Sort(customSort{
	//	t: tracks,
	//	less: func(x, y *Track) bool {
	//		if x.Title != y.Title {
	//			return x.Title < y.Title
	//		}
	//		if x.Year != y.Year {
	//			return x.Year < y.Year
	//		}
	//		if x.Length != y.Length {
	//			return x.Length < y.Length
	//		}
	//		return false
	//	},
	//})
	//const format = "\t%v\t%v\t%v\t%v\t%v\t\n"
	//for _, t := range tracks {
	//	fmt.Printf(format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	//}
	s := [][]int{{1, 2}, {12, 13}, {23, 22}}
	sort.SliceIsSorted(s, func(i, j int) bool {
		if s[i][0] != s[j][0] {
			return s[i][0] < s[j][0]
		}
		return s[i][0] > s[j][0]
	})
	fmt.Println(s)
}

func name() {

}
