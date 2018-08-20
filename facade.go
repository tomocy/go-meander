package meander

type facade interface {
	public() interface{}
}

func MakePublic(data []interface{}) []interface{} {
	publicData := make([]interface{}, len(data))
	for i, d := range data {
		publicData[i] = public(d)
	}

	return publicData
}

func public(o interface{}) interface{} {
	if f, ok := o.(facade); ok {
		return f.public()
	}

	return o
}
