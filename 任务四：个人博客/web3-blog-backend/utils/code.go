package utils

var ok = 200

var (
	DBError            = NewError(10000, "数据库错误")
	AlreadyRegister    = NewError(10100, "用户已注册")
	SelectError        = NewError(100200, "数据查询失败")
	InserError         = NewError(100300, "注册失败！！")
	UserNotExistError  = NewError(100400, "查询用户失败")
	PasswordError      = NewError(100500, "用户名或密码错误")
	TokenGenerateError = NewError(100600, "生成token失败")
	EncryptError       = NewError(100700, "加密失败")
	ServerErr          = NewError(100800, "服务器异常")
	LoginErr           = NewError(100900, "登录失败")
	ResponseErro       = NewError(101000, "数据获取失败")
	ResponseErro_1     = NewError(102000, "创建失败")
	ResponseErro_2     = NewError(103000, "权限不足")
	ResponseErro_3     = NewError(103000, "获取数据失败")
	ResponseErro_4     = NewError(103000, "数据修改失败")
)
