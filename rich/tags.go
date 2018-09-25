package rich

import (
	"sort"

	"github.com/xuender/goutils"
)

// Tags 标签
type Tags []string

// Add 增加新标签
func (t *Tags) Add(tag string) {
	for _, i := range *t {
		if i == tag {
			return
		}
	}
	*t = append(*t, tag)
	t.Sort()
}

// DelPy 删除拼音标签
func (t *Tags) DelPy() {
	goutils.Remove(t, func(s string) bool { return len(s) == 1 })
}

// Sort 标签排序
func (t Tags) Sort() {
	sort.Slice(t, func(i int, j int) bool {
		return t[i] < t[j]
	})
}

// Includes 包含
func (t Tags) Includes(tags []string) bool {
	for _, tag := range tags {
		for _, o := range t {
			if tag == o {
				return true
			}
		}
	}
	return false
}
