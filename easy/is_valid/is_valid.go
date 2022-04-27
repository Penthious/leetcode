package is_valid

type stack []int32

func isValid(input string) bool {
	if len(input)%2 != 0 {
		return false
	}

	data := map[int32]int32{
		40:  41,
		123: 125,
		91:  93,
	}

	st := stack{}
	for _, rune := range input {
		if st.isEmpty() {
			st.push(rune)
			continue
		}
		el := st[len(st)-1]
		if rune != data[el] {
			st.push(rune)
		}
		if rune == data[el] {
			st.pop()
		}
	}

	return len(st) == 0
}

func (s *stack) push(num int32) {
	*s = append(*s, num)
}

func (s *stack) pop() {
	i := len(*s) - 1
	*s = (*s)[:i]

}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}
