package universe

type Id uint64

type IdOrder []Id

func (ids IdOrder) Len() int {
	return len(ids)
}

func (ids IdOrder) Swap(i, j int) {
	ids[i], ids[j] = ids[j], ids[i]
}

func (ids IdOrder) Less(i, j int) bool {
	return ids[i] < ids[j]
}
