package setting

func Bool(in Setting, key any) (bool, bool) {
	data, ok := in.Param(key)
	if !ok {
		return false, false
	}

	res, ok := data.(bool)

	return res, ok
}

func String(in Setting, key any) (string, bool) {
	data, ok := in.Param(key)
	if !ok {
		return "", false
	}

	res, ok := data.(string)

	return res, ok
}
