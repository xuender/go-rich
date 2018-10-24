package rich

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCustomer(t *testing.T) {
	Convey("Customer", t, func() {
		Convey("LoadXsl", func() {
			ext := make(map[int]string)
			ext[0] = "Name"
			cs := []Customer{}
			err := ReadXlsx("../doc/客户档案.xlsx", func(row []string) {
				if c, e := newCustomer(row, ext); e == nil {
					cs = append(cs, c)
				}
			})
			So(err, ShouldBeNil)
			So(len(cs), ShouldEqual, 1008)
		})
	})
}
