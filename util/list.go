package util

type ListValue interface{}

type CircularList struct {
	values []ListValue
	pos    int
}

type ListCallback func(ListValue)

func (cl *CircularList) Add(v ListValue) {
	cl.values = append(cl.values, v)
}

func (cl *CircularList) PeekN(n int) ListValue {
	pos := cl.pos + n - 1
	if pos >= len(cl.values) {
		pos = pos % len(cl.values)
	}
	return cl.values[pos]
}

func (cl *CircularList) Peek() ListValue {
	return cl.PeekN(1)
}

func (cl *CircularList) Next() ListValue {
	i := cl.values[cl.pos]
	cl.pos++
	if cl.pos >= len(cl.values) {
		cl.pos = 0
	}
	return i
}

func (cl *CircularList) Iterate(callback ListCallback) {
	for _, value := range cl.values {
		callback(value)
	}
}

type IntListCallback func(int)

type CircularIntList struct {
	CircularList
}

func (cl *CircularIntList) Add(i int)       { cl.CircularList.Add(ListValue(i)) }
func (cl *CircularIntList) Peek() int       { return cl.CircularList.Peek().(int) }
func (cl *CircularIntList) PeekN(n int) int { return cl.CircularList.PeekN(n).(int) }
func (cl *CircularIntList) Next() int       { return cl.CircularList.Next().(int) }
func (cl *CircularIntList) Iterate(cb IntListCallback) {
	cl.CircularList.Iterate(func(v ListValue) { cb(v.(int)) })
}
