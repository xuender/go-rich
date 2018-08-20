package rich

import (
	"time"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/xuender/goutils"
)

// 商品
type Item struct {
	Id    goutils.ID `json:"id"`    // 主键
	Title string     `json:"title"` // 标题
	Ca    time.Time  `json:"ca"`    // 创建时间
	Price int64      `json:"price"` // 价格,单位分
}

func NewItem() *Item {
	return &Item{
		Id:    goutils.NewId(ItemIdPrefix),
		Title: "新商品",
		Ca:    time.Now(),
		Price: 0,
	}
}

// 保存商品
func (i *Item) Save(db *leveldb.DB) error {
	bs, err := goutils.Encode(i)
	if err != nil {
		return err
	}
	return db.Put(i.Id[:], bs, nil)
}