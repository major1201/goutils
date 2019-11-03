package goutils

// CopyMapSS makes a shallow copy of a map.
func CopyMapSS(m map[string]string) map[string]string {
	if m == nil {
		return nil
	}
	cp := make(map[string]string, len(m))
	for k, v := range m {
		cp[k] = v
	}
	return cp
}

// MergeMapSS merges multiple maps into one
func MergeMapSS(base map[string]string, overrides ...map[string]string) map[string]string {
	if base == nil && len(overrides) == 0 {
		return nil
	}

	if base == nil {
		base = make(map[string]string)
	}

	for _, override := range overrides {
		if override == nil {
			continue
		}
		for k, v := range override {
			base[k] = v
		}
	}

	return base
}
