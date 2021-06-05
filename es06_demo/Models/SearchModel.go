package Models

type SearchModel struct {
	BookName        string  `json:"book_name" binding:"omitempty"`
	BookPress       string  `json:"book_press" binding:"omitempty"`
	BookPrice1Start float32 `json:"book_price1_start" binding:"gte=0,lt=10000"`
	BookPrice1End   float32 `json:"book_price1_end" binding:"required,gte=0,lt=10000,gtefield=BookPrice1Start"`
	//gtefield=BookPrice1Start gte=大于等于
}

func NewSearchModel() *SearchModel {
	return &SearchModel{}
}
