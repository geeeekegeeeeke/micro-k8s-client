package dto

type PageResult struct {
	Total int64       `json:"total"`
	Rows  interface{} `json:"rows"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Options struct {
	Option string `json:"option"`
}
