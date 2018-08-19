package meander

// Facade is an interface which display data human friendly
type Facade interface {
	Public() interface{}
}

// Public returns the result of Facade.Public() if the passed value is Facade or returns it
func Public(o interface{}) interface{} {
	if f, ok := o.(Facade); ok {
		return f.Public()
	}

	return o
}
