package meander

type Facade interface {
	Public() interface{}
}

func Public(o interface{}) interface{} {
	if f, ok := o.(Facade); ok {
		return f.Public()
	}

	return o
}
