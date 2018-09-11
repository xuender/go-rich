package rich

import "sort"

// Days 日期
type Days []string

// Add 增加
func (d *Days) Add(day string) bool {
	for _, v := range *d {
		if v == day {
			return false
		}
	}
	*d = append(*d, day)
	sort.Sort(sort.Reverse(sort.StringSlice(*d)))
	return true
}
