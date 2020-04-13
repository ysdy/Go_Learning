package entity

type Pet struct {
	Id      string `json:"id"`
	Species string `json:"species"`
	Name    string `json:"name"`
	Age     int32  `json:"age"`
	StoreId string `json:"store_id"`
}
