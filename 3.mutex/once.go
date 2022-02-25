package mutex

func NewCount() count {
	return count{}
}

type count struct {
	value int
}

func (c *count) Increment() {
	c.value += 1
}

func (c *count) Value() int {
	return c.value
}
