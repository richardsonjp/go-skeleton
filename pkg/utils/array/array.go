package array

func DeleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}

	return r
}

func ContainsUint(items []uint, search uint) bool {
	for _, item := range items {
		if item == search {
			return true
		}
	}
	return false
}
