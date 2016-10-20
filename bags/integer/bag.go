// bag of integers, implements Bag interface
//
// dynamic container allowing non-unique items
package bag

type count int

type Bag map[int]count

func (b *Bag) Len() int {
	sum := 0
	for _, v := range *b {
		sum += int(v)
	}
	return sum
}

func (b *Bag) Add(x int) {
	(*b)[x]++
}

func (b *Bag) Delete(x int) {
	_, ok := (*b)[x]
	if ok {
		(*b)[x]--
	}
}

func (b *Bag) Find(x int) (int, bool) {
	count, ok := (*b)[x]
	return int(count), ok
}

func (b *Bag) FindAll(x int) (int, bool) {
	return b.Find(x) // not useful for this implementation
}
