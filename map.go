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

// MergeStrIFaceMaps merge 2 map[string]interface{} into one
func MergeStrIFaceMaps(to, from map[string]interface{}) map[string]interface{} {
	out := make(map[string]interface{}, len(to))
	for k, v := range to {
		out[k] = v
	}
	for k, v := range from {
		if v, ok := v.(map[string]interface{}); ok {
			if bv, ok := out[k]; ok {
				if bv, ok := bv.(map[string]interface{}); ok {
					out[k] = MergeStrIFaceMaps(bv, v)
					continue
				}
			}
		}
		out[k] = v
	}
	return out
}
