package mongo

type SingleResult interface {
	Decode(any) error
	Err() error
}
