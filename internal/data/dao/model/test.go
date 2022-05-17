package model

// Test 测试
type Test struct {
	Base `gorm:"embedded"`
	// 名称
	Name string `gorm:"name"`
}

func (*Test) TableName() string {
	return "test"
}
