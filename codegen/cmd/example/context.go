package main

// Context can be used to callback
type Context struct {
	clickCallback func(int)
	moveCallback func(uint32)
}

// OnClick register a callback function
func (c *Context) OnClick(callback func(arg int)) {
	c.clickCallback = callback
}

// EmitClick emit a callback event
func (c *Context) EmitClick(arg int) {
	c.clickCallback(arg)
}

// OnMove register a callback function
func (c *Context) OnMove(callback func(arg uint32)) {
	c.moveCallback = callback
}

// EmitMove emit a callback event
func (c *Context) EmitMove(arg uint32) {
	c.moveCallback(arg)
}
