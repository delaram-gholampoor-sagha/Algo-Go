package main

import (
	"sort"
	"strconv"
)

func groupTitles(strs []string) [][]string {
	var output [][]string
	if len(strs) == 0 {
		return output
	}
	res := make(map[string][]string)

	for _, s := range strs {
		count := make([]int, 26)
		for _, c := range s {
			index := c - 'a'
			count[index]++
		}
		key := ""
		for i := 0; i < 26; i++ {
			key += "#"
			key += strconv.Itoa(count[i])
		}
		res[key] = append(res[key], s)
	}
	for _, v := range res {
		output = append(output, v)
	}
	return output
}

func main() {
	// Driver code
	titles := []string{"duel", "dule", "speed", "spede", "deul", "cars"}
	query := "spede"
	output := groupTitles(titles)
	for _, o := range output {
		sort.Strings(o)
		i := sort.Search(len(o), func(i int) bool {
			return o[i] >= query
		})
		if i < len(o) && o[i] == query {
			print(o)
		}
	}
}
