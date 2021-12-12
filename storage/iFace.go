package storage

type DataProvider interface {
	GetData(resource string, data Data) (error)
}

type Data interface {
	Append(el string) error
	Next() bool
}
