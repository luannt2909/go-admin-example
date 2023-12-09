package user

type GetListParams struct {
	Filter           map[string]interface{}
	Limit, Offset    int
	SortBy, SortType string
}
