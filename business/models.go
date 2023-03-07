package business

import (
	"bytes"
	"fmt"
	"math"

	"golang.org/x/exp/slices"
)

type Solar struct {
	Name  string
	Netto float64
}

type SolarSlice []Solar

type Wind struct {
	Name  string
	Netto float64
}

func (s *Solar) Print() string {
	return fmt.Sprintf("%s - %v\n", startduskPoint, *s)
}

func (t *Wind) Print() string {
	return fmt.Sprintf("%s - %v\n", startduskPoint, *t)
}

var startduskPoint string = "startdusk: "

func PrintGeneric[T Energy](t T) string {
	return fmt.Sprintf("%s - %v\n", startduskPoint, t)
}

func PrintSlice[T Energy](tt []T) string {
	fmt.Printf("type of slice: %T\n", tt)
	var buf bytes.Buffer
	buf.WriteString("[\n")
	for i := range tt {
		buf.WriteString("\t" + PrintGeneric(tt[i]))
	}
	buf.WriteString("]")
	return buf.String()
}

// Go新增了一个符号 ~ ，
// 在类型约束中使用类似 ~int 这种写法的话，就代表着不光是 int ，所有以 int 为底层类型的类型也都可用于实例化
// 限制：使用 ~ 时有一定的限制：
// ~后面的类型不能为接口
// ~后面的类型必须为基本类型
// type MyInt int

// type _ interface {
// 	~[]byte // 正确
// 	~MyInt  // 错误，~后的类型必须为基本类型
// 	~error  // 错误，~后的类型不能为接口
// }

// 这里的S 就能识别出SolarSlice类型
func PrintSlice2[T Energy, S ~[]T](tt S) string {
	fmt.Printf("type of slice2: %T\n", tt)
	var buf bytes.Buffer
	buf.WriteString("[\n")
	for i := range tt {
		buf.WriteString("\t" + PrintGeneric(tt[i]))
	}
	buf.WriteString("]")
	return buf.String()
}

type Energy interface {
	Wind | Solar
	Cost() float64
}

func SortByCost[T Energy](a []T) {
	slices.SortFunc(a, func(a, b T) bool {
		return a.Cost() < b.Cost() || math.IsNaN(a.Cost()) && !math.IsNaN(b.Cost())
	})
}
