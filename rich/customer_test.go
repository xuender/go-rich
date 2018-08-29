package rich

import (
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCustomer(t *testing.T) {
	Convey("Customer", t, func() {
		Convey("LoadXsl", func() {
			ext := []Ext{
				Ext{Key: "Name", Value: "0"},
				Ext{Key: "R球镜", Value: "1"},
				Ext{Key: "R柱镜", Value: "2"},
				Ext{Key: "R轴位", Value: "3"},
				Ext{Key: "L球镜", Value: "4"},
				Ext{Key: "L柱镜", Value: "5"},
				Ext{Key: "L轴位", Value: "6"},
				Ext{Key: "轴位", Value: "7"},
				Ext{Key: "镜架", Value: "8"},
				Ext{Key: "镜片", Value: "9"},
				Ext{Key: "金额", Value: "10"},
				Ext{Key: "Phone", Value: "11"},
				Ext{Key: "Note", Value: "12"},
			}
			cs, err := readXlsx("./客户档案.xlsx", ext)
			b, _ := json.Marshal(cs)
			fmt.Println(string(b))
			So(err, ShouldBeNil)
			So(len(cs), ShouldEqual, 1008)
		})
	})
}
