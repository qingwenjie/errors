package errors

type Fields map[string]interface{}

func New() Fields {
	return make(Fields)
}

func (fs Fields) Clone() Fields {
	n := make(Fields, len(fs))
	for k, v := range fs {
		n[k] = v
	}
	return n
}

func (fs Fields) Merge(fs2 Fields) Fields {
	for k, v := range fs2 {
		fs[k] = v
	}
	return fs
}

func (fs Fields) Set(key string, val interface{}) Fields {
	fs[key] = val
	return fs
}