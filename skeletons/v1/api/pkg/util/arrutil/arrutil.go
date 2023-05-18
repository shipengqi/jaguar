package arrutil

func DistinguishStringSlice(old, new []string) (removed, added []string) {
	if len(old) == 0 {
		return nil, new
	}
	tmp := make(map[string]struct{})
	for i := 0; i < len(old); i++ {
		tmp[old[i]] = struct{}{}
	}
	for i := 0; i < len(new); i++ {
		if _, ok := tmp[new[i]]; !ok {
			added = append(added, new[i])
		} else {
			delete(tmp, new[i])
		}
	}
	for k := range tmp {
		removed = append(removed, k)
	}
	return
}

func DistinguishUint64Slice(old, new []uint64) (removed, added []uint64) {
	if len(old) == 0 {
		return nil, new
	}

	tmp := make(map[uint64]struct{})
	for i := 0; i < len(old); i++ {
		tmp[old[i]] = struct{}{}
	}
	for i := 0; i < len(new); i++ {
		if _, ok := tmp[new[i]]; !ok {
			added = append(added, new[i])
		} else {
			delete(tmp, new[i])
		}
	}
	for k := range tmp {
		removed = append(removed, k)
	}
	return
}
