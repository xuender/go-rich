package rich

import (
	"time"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/xuender/goutils"
)

// Item 商品
type Item struct {
	Obj
	Name   string `json:"name"`             // 名称
	Price  int64  `json:"price"`            // 价格,单位分
	Extend []Ext  `json:"extend,omitempty"` // 扩展属性
}

// NewItem 新建商品
func NewItem() *Item {
	return &Item{
		Obj: Obj{
			ID: goutils.NewId(ItemIDPrefix),
			Ca: time.Now(),
		},
		Name:  "新商品",
		Price: 0,
	}
}

// Save 保存商品
func (i *Item) Save(db *leveldb.DB) error {
	bs, err := goutils.Encode(i)
	if err != nil {
		return err
	}
	return db.Put(i.ID[:], bs, nil)
}
