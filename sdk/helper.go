package sdk

type Config map[string]any

func (cfg Config) GetString(key, fallback string) string {
	v, ok := cfg[key]
	if !ok {
		return fallback
	}
	s, ok := v.(string)
	if !ok {
		return fallback
	}
	return s
}

func (cfg Config) GetBool(key string, fallback bool) bool {
	v, ok := cfg[key]
	if !ok {
		return fallback
	}
	b, ok := v.(bool)
	if !ok {
		return fallback
	}
	return b
}

func (cfg Config) GetFloat64(key string, fallback float64) float64 {
	v, ok := cfg[key]
	if !ok {
		return fallback
	}
	f, ok := v.(float64)
	if !ok {
		return fallback
	}
	return f
}
