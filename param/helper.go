package param

func Bool(in Params, key any) (bool, bool) {
	data, ok := in.Param(key)
	if !ok {
		return false, false
	}

	res, ok := data.(bool)

	return res, ok
}

func String(in Params, key any) (string, bool) {
	data, ok := in.Param(key)
	if !ok {
		return "", false
	}

	res, ok := data.(string)

	return res, ok
}
