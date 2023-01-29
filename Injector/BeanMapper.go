package Injector

import "reflect"

type BeanMapper map[reflect.Type]reflect.Value

func (bm BeanMapper) add(bean interface{}) {
	t := reflect.TypeOf(bean)
	if t.Kind() != reflect.Ptr {
		panic("need ptr")
	}
	bm[t] = reflect.ValueOf(bean)
}
func (bm BeanMapper) get(bean interface{}) reflect.Value {
	t := reflect.TypeOf(bean)
	if v, ok := bm[t]; ok {
		return v
	}
	return reflect.Value{}
}
