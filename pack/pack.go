package pack

type RString struct {
	S string
}

// String implements the stringer interface
func (s RString) String() string {
	return s.S
}

// NewRString whatever
func NewRString(s string) RString {
	return RString{S: s}
}

func (rstr RString) Reverse() RString {
	var buffer []rune
	runed := []rune(rstr.S)
	for i := len(runed) - 1; i >= 0; i-- {
		buffer = append(buffer, runed[i])
	}
	return NewRString(string(buffer))
}

type SortableRString []RString

// implement the sort interface
func (s SortableRString) Len() int {
	return len(s)
}
func (s SortableRString) Less(i, j int) bool {
	return s[i].S < s[j].S
}
func (s SortableRString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
