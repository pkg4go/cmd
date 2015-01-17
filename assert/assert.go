package assert

import "reflect"
import "fmt"

func Equal(a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		panic(fmt.Errorf("expect %v to equal %v", a, b))
	}
}
