package storage

type DataProvider interface {
	GetData() (IntCollection, error)
}
