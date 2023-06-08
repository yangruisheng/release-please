package releaseplease/foo

type Widget struct {
}

func (w Widget) ID() string {
	return "Foo"
}

func (w Widget) Weight() int {
	return 0
}
