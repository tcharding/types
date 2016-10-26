// bag of integers, implements Bag interface
//
// dynamic container allowing non-unique items
package bag

type count int

type Bag map[int]count

// Len: total number of items in bag
func (b *Bag) Len() int {
	sum := 0
	for _, v := range *b {
		sum += int(v)
	}
	return sum
}

// Add an item to bag
func (b *Bag) Add(x int) {
	(*b)[x]++
}

// Delete an item from bag
func (b *Bag) Delete(x int) {
	_, ok := (*b)[x]
	if ok {
		(*b)[x]--
	}
}

// Find: check if item is in bag, return item and true if so, false if not
func (b *Bag) Find(x int) (int, bool) {
	count, ok := (*b)[x]
	return int(count), ok
}

// FindAll: finds all the items in a bag (not useful for this implementation)
func (b *Bag) FindAll(x int) (int, bool) {
	return b.Find(x) // no implementation
}
