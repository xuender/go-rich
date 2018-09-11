package rich

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCustomer(t *testing.T) {
	Convey("Customer", t, func() {
		Convey("LoadXsl", func() {
      ext:=make(map[int]string)
      ext[0]="Name"
			cs, err := readXlsx("./客户档案.xlsx", ext)
			// b, _ := json.Marshal(cs)
			// fmt.Println(string(b))
			So(err, ShouldBeNil)
			So(len(cs), ShouldEqual, 1008)
		})
	})
}
