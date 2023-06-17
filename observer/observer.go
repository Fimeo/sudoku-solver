package observer

type IObservable interface {
	Notify()
	Subscribe(observer Observer)
}

type IObserver interface {
	Update(interface{})
}

type Observer struct {
	callbackFunc func(interface{})
}

func NewObserver(callback func(interface{})) Observer {
	return Observer{
		callbackFunc: callback,
	}
}

func (o Observer) Update(subject interface{}) {
	o.callbackFunc(subject)
}
