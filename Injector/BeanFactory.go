package Injector

var BeanFactory *BeanFactoryImpl

func init() {
	BeanFactory = NewBeanFactory()
}

type BeanFactoryImpl struct {
	beanMapper BeanMapper
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

func NewBeanFactory() *BeanFactoryImpl {
	return &BeanFactoryImpl{beanMapper: make(BeanMapper)}
}
