package rich

import (
	"errors"
	"sort"

	"github.com/xuender/go-utils"
)

// Days 日期
type Days []string

// Clean 清空
func (d *Days) Clean() {
	*d = append([]string{})
}

// Add 增加
func (d *Days) Add(day string) bool {
	if d.Includes(day) {
		return false
	}
	*d = append(*d, day)
	sort.Sort(sort.Reverse(sort.StringSlice(*d)))
	return true
}

// Includes 包含
func (d *Days) Includes(day string) bool {
	for _, v := range *d {
		if v == day {
			return true
		}
	}
	return false
}

// 日帐增加
func (w *Web) dayAdd(t Trade) error {
	day := t.Ca.Format(DayFormat)
	w.days.Add(day)
	w.Put(DaysKey, w.days)
	idsKey := []byte(day)
	ids := []utils.ID{}
	w.Get(idsKey, &ids)
	ids = append(ids, t.ID)
	return w.Put(idsKey, ids)
}

// Reset 重置帐目
func (w *Web) Reset() error {
	for _, day := range w.days {
		w.Delete([]byte(day))
	}
	w.days.Clean()
	// 删除的客户
	delc := []Customer{}
	w.Iterator([]byte{CustomerIDPrefix, '-'}, func(key, data []byte) {
		c := Customer{}
		utils.Decode(data, &c)
		if c.IsDelete() {
			delc = append(delc, c)
		}
	})
	// 删除的商品
	deli := []Item{}
	w.Iterator([]byte{ItemIDPrefix, '-'}, func(key, data []byte) {
		i := Item{}
		utils.Decode(data, &i)
		if i.IsDelete() {
			deli = append(deli, i)
		}
	})
	w.Iterator([]byte{TradeIDPrefix, '-'}, func(key, data []byte) {
		t := Trade{}
		utils.Decode(data, &t)
		day := t.Ca.Format(DayFormat)
		w.days.Add(day)
		idsKey := []byte(day)
		ids := []utils.ID{}
		w.Get(idsKey, &ids)
		ids = append(ids, t.ID)
		w.Put(idsKey, ids)
		if !t.CID.IsNew() {
			utils.Remove(delc, func(c Customer) bool { return c.ID.Equal(t.CID) })
		}
		utils.Remove(deli, func(i Item) bool { return t.HasItem(i.ID) })
	})
	// 删除无用的商品和客户
	for _, i := range deli {
		w.Delete(i.ID[:])
	}
	for _, c := range delc {
		w.Delete(c.ID[:])
	}
	return w.Put(DaysKey, w.days)
}

// 日帐减少
func (w *Web) dayDel(t Trade) error {
	day := t.Ca.Format("2006-01-02")
	idsKey := []byte(day)
	ids := []utils.ID{}
	w.Get(idsKey, &ids)
	for i, id := range ids {
		if id.Equal(t.ID) {
			ids = append(ids[:i], ids[i+1:]...)
			w.Put(idsKey, ids)
			if len(ids) == 0 {
				for f, d := range w.days {
					if d == day {
						w.days = append(w.days[:f], w.days[f+1:]...)
						w.Put(DaysKey, w.days)
					}
				}
			}
		}
	}
	return errors.New("帐目未找到")
}
