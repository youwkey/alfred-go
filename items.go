package alfred

type Items []*Item

func (i *Items) Length() int {
	return len(*i)
}

func (i *Items) IsEmpty() bool {
	return i.Length() == 0
}

func (i *Items) Append(items ...*Item) {
	*i = append(*i, items...)
}
