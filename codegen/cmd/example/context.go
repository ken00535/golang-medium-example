package main

// Context can be used to callback
type Context struct {
	ClickCallback func(int)
	MoveCallback func(uint32)
}

// OnClick register a callback function
func (c *Context) OnClick(callback func(arg int)) {
	c.ClickCallback = callback
}

// EmitClick emit a callback event
func (c *Context) EmitClick(arg int) {
	c.ClickCallback(arg)
}

// OnMove register a callback function
func (c *Context) OnMove(callback func(arg uint32)) {
	c.MoveCallback = callback
}

// EmitMove emit a callback event
func (c *Context) EmitMove(arg uint32) {
	c.MoveCallback(arg)
}
