package models

type JSONResult struct {
	Code  int         `json:"code"`
	Count int         `json:"count"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}
