package setting

//nolint:gochecknoglobals
var eparam = empty{}

func New(opts ...Option) Setting {
	var param Setting

	param = eparam
	for _, opt := range opts {
		param = opt(param)
	}

	return param
}

type Setting interface {
	Param(key any) (any, bool)
	With(key, val any) Setting
}

type Option func(Setting) Setting

type empty struct{}

func (e empty) Param(any) (any, bool) {
	return nil, false
}

func (e empty) With(key, val any) Setting {
	return data{
		parent: e,
		key:    key,
		val:    val,
	}
}

type data struct {
	parent   Setting
	key, val any
}

func (d data) Param(key any) (any, bool) {
	if d.key == key {
		return d.val, true
	}

	return d.parent.Param(key)
}

func (d data) With(key, val any) Setting {
	return data{
		parent: d,
		key:    key,
		val:    val,
	}
}
