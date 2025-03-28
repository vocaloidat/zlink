package sequence

type SequenceSql interface {
	Next() (uint64, error)
}
