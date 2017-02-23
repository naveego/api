package utils

type MapReader struct {
	Map map[string]interface{}
}

func NewMapReader(m map[string]interface{}) *MapReader {
	return &MapReader{
		Map: m,
	}
}

func (mr *MapReader) ReadString(key string) (string, bool) {
	if raw, ok := mr.Map[key]; ok {
		if s, ok := raw.(string); ok {
			return s, true
		}
	}

	return "", false
}

func (mr *MapReader) ReadInt(key string) (int, bool) {
	if raw, ok := mr.Map[key]; ok {
		if i, ok := raw.(int); ok {
			return i, true
		}

		// JSON deserialization always uses float64
		if f, ok := raw.(float64); ok {
			return int(f), true
		}
	}
	return 0, false
}

func (mr *MapReader) ReadBool(key string) (bool, bool) {
	raw, ok := mr.Map[key]
	if ok {
		b, ok := raw.(bool)
		if ok {
			return b, true
		}
	}

	return false, false
}
