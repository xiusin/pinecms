package search


type ISearch interface {
	Search(index string, query SearchParams) (any, error)
	Index(index string, id string, doc map[string]any) (string, error)
	Update(index, id string, doc map[string]any) error
	Delete(index string, id string) error
}

type SearchParams struct {
	Index string
	Query map[string]any
	From  int
	Size  int
	Sort  []string
	Aggs  map[string]any
}
