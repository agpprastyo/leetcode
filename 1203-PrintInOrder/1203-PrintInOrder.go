// Last updated: 2/11/2026, 8:30:18 PM
type Foo struct {
    doneFirst chan struct{}
    doneSecond chan struct{}

}

func NewFoo() *Foo {
	return &Foo{
        doneFirst : make(chan struct{}),
        doneSecond: make(chan struct{}),
	}
}

func (f *Foo) First(printFirst func()) {
	// Do not change this line
	printFirst()

    close(f.doneFirst)
}

func (f *Foo) Second(printSecond func()) {
    <-f.doneFirst
	/// Do not change this line
	printSecond()

    close(f.doneSecond)
}

func (f *Foo) Third(printThird func()) {
    <-f.doneSecond
	// Do not change this line
	printThird()
}