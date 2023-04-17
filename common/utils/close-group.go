package utils

type CloseGroup struct {
	workArr    []chan struct{}
	closeGroup chan struct{}
}

func (c *CloseGroup) Start() {
	if c.workArr == nil {
		c.workArr = make([]chan struct{}, 0)
	}
	newChan := make(chan struct{})
	c.workArr = append(c.workArr, newChan)
	<-newChan
}

func (c *CloseGroup) Close() {
	c.closeGroup <- struct{}{}
}

func (c *CloseGroup) Wait() {
	c.closeGroup = make(chan struct{})
	<-c.closeGroup
	if c.workArr == nil {
		return
	}
	for _, v := range c.workArr {
		v <- struct{}{}
	}

	c.workArr = nil
}
