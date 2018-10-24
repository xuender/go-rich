package rich

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUtils(t *testing.T) {
	Convey("pinyin", t, func() {
		Convey("PY", func() {
			So(PY("中国"), ShouldEqual, "zhong guo")
			So(PY("中国WTo"), ShouldEqual, "zhong guo ; zhong guo wto")
			So(PY("重瞳"), ShouldContainSubstring, "zhong tong")
			So(PY("11"), ShouldContainSubstring, "11")
		})
	})
	Convey("Axis", t, func() {
		Convey("坐标", func() {
			So(Axis(0, 0), ShouldEqual, "A1")
			So(Axis(1, 2), ShouldEqual, "B3")
			So(Axis(3, 2), ShouldEqual, "D3")
		})
	})
	Convey("PC", t, func() {
		Convey("排列组合", func() {
			a := [][]string{
				{"A", "a"},
				{"b"},
				{"c", "C"},
			}
			So(PC(a), ShouldContain, "A b c")
		})
	})
	Convey("Initial", t, func() {
		Convey("多音字", func() {
			So(Initial("长三"), ShouldContain, "Z")
			So(Initial("长三"), ShouldContain, "C")
			So(Initial(" 长三"), ShouldContain, "C")
		})
	})
}
func ExamplePY() {
	fmt.Println(PY("一沓"))
	fmt.Println(PY("1你"))
	fmt.Println(PY(" 沓 "))

	// Output:
	// yi da ; yi ta
	// 1ni ; ni
	// da ; ta
}
func ExampleInitial() {
	fmt.Println(Initial("一沓"))
	fmt.Println(Initial("1沓"))

	// Output:
	// [Y]
	// [1 D T]
}
