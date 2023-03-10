package Injector

import (
	"log"
	"reflect"
)

type BeanMapper map[reflect.Type]reflect.Value

func (bm BeanMapper) add(bean interface{}) {
	t := reflect.TypeOf(bean)
	if t.Kind() != reflect.Ptr {
		panic("need ptr")
	}
	bm[t] = reflect.ValueOf(bean)
}
func (bm BeanMapper) get(bean interface{}) reflect.Value {
	var t reflect.Type
	if bt, ok := bean.(reflect.Type); ok {
		t = bt
	} else {
		t = reflect.TypeOf(bean)
	}

	if v, ok := bm[t]; ok {
		return v
	}
	// 处理接口 继承
	for k, v := range bm {
		log.Printf(k.String(), t.String())
		if k.Implements(t) {
			return v
		}
	}
	return reflect.Value{}
}
