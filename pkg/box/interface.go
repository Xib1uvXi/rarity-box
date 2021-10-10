package box

type dao interface {
	SaveOrUpdate(address string, tokenIDs []uint64) error
}
