package reflection

import (
	"fmt"
	"reflect"
)

func Test() {
	var x float64 = 3.4
	var empty interface{} = reflect.ValueOf(x)
	fmt.Println("type:", reflect.TypeOf(x), empty)
}
