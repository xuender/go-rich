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

			customerProMap := make(map[int]string)
			customerProMap[0] = "Name"
			customerProMap[1] = "R球镜"
			customerProMap[2] = "R柱镜"
			customerProMap[3] = "R轴位"
			customerProMap[4] = "L球镜"
			customerProMap[5] = "L柱镜"
			customerProMap[6] = "L轴位"
			customerProMap[7] = "轴位"
			customerProMap[8] = "镜架"
			customerProMap[9] = "镜片"
			customerProMap[10] = "金额"
			customerProMap[11] = "Phone"
			customerProMap[12] = "Note"
			cs, err := readXlsx("./客户档案.xlsx", customerProMap)
			b, _ := json.Marshal(cs)
			fmt.Println(string(b))
			So(err, ShouldBeNil)
			So(len(cs), ShouldEqual, 7)
		})
	})
}
