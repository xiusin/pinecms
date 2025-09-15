package search

type NullSearch struct {
}

func (n *NullSearch) Search(index string, query SearchParams) (any, error) {
	return nil, nil
}

func (n *NullSearch) Index(index string, id string, doc map[string]any) (string, error) {
	return "", nil
}

func (n *NullSearch) Update(index string, id string, doc map[string]any) error {
	return nil
}

func (n *NullSearch) Delete(index string, id string) error {
	return nil
}
