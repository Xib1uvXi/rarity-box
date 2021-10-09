package box

type IDataSynchronizer interface {
	Sync(address string) error
}


type dao interface {
	SaveOrUpdate(address string, tokenIDs []uint64) error
}
