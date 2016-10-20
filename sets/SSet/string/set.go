// ordered set of strings, implements the SSet interface
package sset

type SSet []string

// size of set
func (set *SSet) Size() int {
	return len(*set)
}

// add s to set, return true if added
func (set *SSet) Add(s string) bool {
	for i, cur := range *set {
		if cur == s {
			return false
		}
		if s < cur {
			set.insertBefore(s, i)
			return true
		}
	}
	*set = append(*set, s)
	return true
}

func (set *SSet) insertBefore(s string, i int) {
	new := copy((*set)[:i])
	new = append(new, s)
	new = append(new, (*set)[i:]...)
	*set = new
}

func copy(v []string) []string {
	new := make([]string, len(v))
	for i := range new {
		new[i] = v[i]
	}
	return new
}

// delete s from set, return true if deleted
func (set *SSet) Delete(s string) bool {
	for i, cur := range *set {
		if cur == s {
			set.deleteIndex(i)
			return true
		}
	}
	return false
}

func (set *SSet) deleteIndex(i int) {
	new := (*set)[:i]
	new = append(new, (*set)[i+1:]...)
	*set = new
}

// Find an item in set, returns item if found and true, if not found
// returns successor item and true. If no successor returns false
func (set *SSet) Find(s string) (string, bool) {
	for _, cur := range *set {
		if cur == s {
			return s, true
		}
		if s < cur {
			return cur, true
		}
	}
	return "", false
}
