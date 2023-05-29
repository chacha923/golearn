package structure

type Memo struct {
	cache map[int]map[int]int
}

func NewMemo() *Memo {
	return &Memo{
		cache: make(map[int]map[int]int),
	}
}

func (m *Memo) Get(i, j int) (int, bool) {
	if mi, ok := m.cache[i]; ok {
		if val, ok := mi[j]; ok {
			return val, true
		}
	}
	return -1, false
}

func (m *Memo) Put(i, j, value int) {
	if mi, ok := m.cache[i]; !ok {
		mi = make(map[int]int)
		mi[j] = value
		m.cache[i] = mi
		return
	} else {
		m.cache[i][j] = value
	}
}
