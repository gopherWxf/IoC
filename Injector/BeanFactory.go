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
			//表达式
			if field.Tag.Get("inject") != "-" {
				ret := expr.BeanExpr(field.Tag.Get("inject"), b.ExprMap)
				if ret != nil && !ret.IsEmpty() {
					retVal := ret[0]
					if retVal != nil {
						v.Field(i).Set(reflect.ValueOf(retVal))
						b.Set(retVal)
					}
				}
			} else {
				//-
				if getV := b.Get(field.Type); getV != nil {
					v.Field(i).Set(reflect.ValueOf(getV))
					continue
				}
			}
		}
	}
}
func (b *BeanFactoryImpl) Config(cfgs ...interface{}) {
	for _, cfg := range cfgs {
		t := reflect.TypeOf(cfg)
		if t.Kind() != reflect.Ptr {
			panic("need ptr")
		}
		//把config本身也加入bean
		b.Set(t)
		//自动构建ExprMap
		b.ExprMap[t.Elem().Name()] = cfg
		v := reflect.ValueOf(cfg)
		for i := 0; i < t.NumMethod(); i++ {
			method := v.Method(i)
			callRet := method.Call(nil)
			if callRet != nil && len(callRet) == 1 {
				b.Set(callRet[0].Interface())
			}
		}
	}
}
