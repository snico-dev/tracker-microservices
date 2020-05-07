package dtos

//ResponseDTO response base
type ResponseDTO struct {
	Result  bool        `json:"result"`
	Content interface{} `json:"content"`
}
