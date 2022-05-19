package http

// ResponseEntity 响应数据体
type ResponseEntity struct {
	// 错误码
	ErrorCode int `json:"errcode"`
	// 响应消息
	Message string `json:"errmsg"`
	// 响应数据
	Data interface{} `json:"data"`
}

// PaginationEntity 分页数据体
type PaginationEntity struct {
	// 是否有下一页
	HasMore bool `json:"more"`
	// 下一页开始数据
	Start int32 `json:"start"`
	// 数据列表
	List interface{} `json:"list"`
}
