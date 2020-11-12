package foo

type dbMock struct{}

func newDbMock() *dbMock {
	return &dbMock{}
}

func (d *dbMock) First(out interface{}) {
	out.(*User).Age = 1
}
