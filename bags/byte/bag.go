// bag of bytes, dynamic container storing non-unique byte values
package bag

type Bag map[byte]int

// make a bag from an ASCII string
func makeBag(s string) Bag {
	bag := make(Bag)
	for i := 0; i < len(s); i++ {
		bag.Add(s[i])
	}
	return bag
}

// Len: total number of items in bag
func (b *Bag) Len() int {
	sum := 0
	for _, v := range *b {
		sum += int(v)
	}
	return sum
}

// Add an item to bag
func (b *Bag) Add(x byte) {
	(*b)[x]++
}

// bagEq: check contents of bag b against contents of bag c
func bagEq(b, c Bag) bool {
	if len(b) != len(c) {
		return false
	}
	for k, v := range b {
		if v != c[k] {
			return false
		}
	}
	return true
}

// Delete an item from bag
func (b *Bag) Delete(x byte) {
	_, ok := (*b)[x]
	if ok {
		(*b)[x]--
	}
}

// Find: check if item is in bag, return item and true if so, false if not
func (b *Bag) Find(x byte) (int, bool) {
	count, ok := (*b)[x]
	return count, ok
}

// intersection: return bag containing items present in both bags b and c
func (b *Bag) intersection(c Bag) Bag {
	bag := make(Bag)
	for k, vb := range *b {
		vc, ok := c[k]
		if ok {
			if vb < vc {
				bag[k] = vb
			} else {
				bag[k] = vc
			}
		}
	}
	return bag
}

// union: return bag containing combined items from bag b and c
func (b *Bag) union(c Bag) Bag {
	bag := make(Bag)
	for k, v := range *b {
		bag[k] += v
	}
	for k, v := range c {
		bag[k] += v
	}
	return bag
}

// difference: return bag containing items in bag b that are not in bag c
func (b *Bag) difference(c Bag) Bag {
	bag := make(Bag)
	for k, vb := range *b {
		vc, ok := c[k]
		if ok {
			if vb > vc {
				bag[k] = vb - vc
			}
		} else {
			bag[k] = vb
		}
	}
	return bag
}
