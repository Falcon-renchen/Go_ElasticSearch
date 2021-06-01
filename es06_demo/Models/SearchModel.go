package Models

type SearchModel struct {
	BookName string `json:"book_name"`
}

func NewSearchModel() *SearchModel {
	return &SearchModel{}
}
