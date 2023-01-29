package Injector

import (
	"github.com/shenyisyn/goft-expr/src/expr"
	"reflect"
)

var BeanFactory *BeanFactoryImpl

func init() {
	BeanFactory = NewBeanFactory()
}

type BeanFactoryImpl struct {
	beanMapper BeanMapper
	ExprMap    map[string]interface{}
}

func NewBeanFactory() *BeanFactoryImpl {
	return &BeanFactoryImpl{beanMapper: make(BeanMapper), ExprMap: make(map[string]interface{})}
}

func (b *BeanFactoryImpl) Get(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	getV := b.beanMapper.get(v)
	if getV.IsValid() {
		return getV.Interface()
	}
	return nil
}

func (b *BeanFactoryImpl) Set(vlist ...interface{}) {
	if vlist == nil || len(vlist) == 0 {
		return
	}
	for _, v := range vlist {
		b.beanMapper.add(v)
	}
}

// Apply 处理依赖注入
func (b *BeanFactoryImpl) Apply(bean interface{}) {
	if bean == nil {
		return
	}
	v := reflect.ValueOf(bean)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		if v.Field(i).CanSet() && field.Tag.Get("inject") != "" {
			//先看看容器里有没有
			if getV := b.Get(field.Type); getV != nil {
				v.Field(i).Set(reflect.ValueOf(getV))
				continue
			}
			//如果没用，看看表达式
			if field.Tag.Get("inject") != "-" {
				ret := expr.BeanExpr(field.Tag.Get("inject"), b.ExprMap)
				if ret != nil && !ret.IsEmpty() {
					retVal := ret[0]
					if retVal != nil {
						v.Field(i).Set(reflect.ValueOf(retVal))
						b.Set(retVal)
					}
				}
			}
		}
	}
}
