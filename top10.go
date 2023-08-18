package analysis

import (
	"regexp"
	"sort"
	"strings"
)

type Top struct {
	Word  string
	Count int
}

func Top10(str string) []string {
	if len(str) == 0 {
		return nil
	}

	var re = regexp.MustCompile(`[[:punct:]]`)
	lit := make(map[string]int)
	var res []string

	str45 := re.ReplaceAllString(strings.ToLower(str), "")
	an := strings.Fields(str45)
	for _, val := range an {
		lit[val] += 1
	}
	sortMap := make([]Top, 0, len(lit))
	for key, val := range lit {
		sortMap = append(sortMap, Top{key, val})
	}
	sort.SliceStable(sortMap, func(i, j int) bool {
		return sortMap[i].Word < sortMap[j].Word
	})
	sort.SliceStable(sortMap, func(i, j int) bool {
		return sortMap[i].Count > sortMap[j].Count
	})
	for i := 0; i < 10; i++ {
		res = append(res, sortMap[i].Word)
	}
	return res
}
