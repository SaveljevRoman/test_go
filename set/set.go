package set

type IntSet struct {
	//elems *[]int
	elems *map[int]int
}

func MakeIntSet() IntSet {
	//var elems []int
	elems := make(map[int]int)
	return IntSet{&elems}
}

func (s IntSet) Contains(elem int) bool {
	_, ok := (*s.elems)[elem]
	if ok {
		return true
	}
	return false
}

func (s IntSet) Add(elem int) bool {
	if s.Contains(elem) {
		return false
	}
	(*s.elems)[elem] = elem

	return true
}
