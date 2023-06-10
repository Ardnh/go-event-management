package domain

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type WebResponseWithPagination struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Page   int         `json:"page"`
	Limit  int         `json:"limit"`
	Sort   string      `json:"sort"`
	Data   interface{} `json:"data"`
}
