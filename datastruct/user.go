package datastruct

type User struct {
	Id       int
	Username string
	Age      int
	Email    string
	AddTime  int
}

// 重构操作数据库表名称
func (receiver User) TableName() string {
	return "user"
}
