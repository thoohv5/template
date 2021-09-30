package server

// GetExampleReq Example请求
type GetExampleReq struct {
	// Id
	Id int64 `form:"id" binding:"required,gte=1"`
	// 名称
	Name string `form:"name" binding:"required" label:"名称"`
}

// GetExampleResp Example返回值
type GetExampleResp struct {
	// Id
	Id int64
	// 名称
	Name string
}

// PostExampleReq Example请求
type PostExampleReq struct {
	// Id
	Id int64 `json:"id" binding:"required,gte=1"`
	// 名称
	Name string `json:"name" binding:"omitempty,gte=2"`
}

// PostExampleResp Example返回值
type PostExampleResp struct {
	// Id
	Id int64
	// 名称
	Name string
}
