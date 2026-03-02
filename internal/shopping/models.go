package shopping

// SearchResponse는 네이버 쇼핑 API의 응답 구조체입니다.
type SearchResponse struct {
	Total   int       `json:"total"`
	Start   int       `json:"start"`
	Display int       `json:"display"`
	Items   []Product `json:"items"`
}

// Product는 개별 상품 정보 구조체입니다.
type Product struct {
	Title     string `json:"title"` // <b> 태그 포함 가능
	Link      string `json:"link"`
	Image     string `json:"image"`
	Lprice    string `json:"lprice"` // 최저가 (문자열)
	Hprice    string `json:"hprice"` // 최고가
	MallName  string `json:"mallName"`
	ProductId string `json:"productId"`
	Brand     string `json:"brand"`
	Maker     string `json:"maker"`
	Category1 string `json:"category1"`
}
