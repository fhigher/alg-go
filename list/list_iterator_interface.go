package list

type Iterator interface {
	FrontNext() interface{}
	BackNext() interface{}
	FrontHasNext() bool
	BackHasNext() bool
}

type Iterable interface {
	Iterator() Iterator
}
