package gftenant

var (
	ModeNone   = "none"   //关闭
	ModeColumn = "column" //列模式
	ModeDB     = "db"     //数据源模式
)

type Model struct {
	Debug     bool   `json:"debug"`     // 是否开启调试模式
	DbName    string `json:"dbName"`    // 数据库名称
	DbAddr    string `json:"dbAddr"`    // 数据库地址
	DbPort    string `json:"dbPort"`    // 数据库端口
	DbAccount string `json:"dbAccount"` // 数据库账号
	DbPass    string `json:"dbPass"`    // 数据库密码
}
