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
			cs, err := LoadXlsx("./客户档案.xlsx")
			b, _ := json.Marshal(cs)
			fmt.Println(string(b))
			So(err, ShouldBeNil)
			So(len(cs), ShouldEqual, 7)
		})
	})
}
