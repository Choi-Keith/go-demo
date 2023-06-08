package errors

type CodeMessage struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func ModuleError(baseCode int) int {
	return Module*100000 + baseCode
}

var (
	ActionSuccess        = &CodeMessage{0, "ok.ActionSuccess"}
	ErrorNotFound        = &CodeMessage{ModuleError(00001), "err.ResourceNotFound"}
	ErrorMissingJwtToken = &CodeMessage{ModuleError(00002), "err.ErrorMissingJwtToken"}
	ErrorCreateUser      = &CodeMessage{ModuleError(00003), "创建用户失败"}
	ErrorParamsLogin     = &CodeMessage{ModuleError(00004), "用户名或密码错误"}
	ErrorNotExistsLogin  = &CodeMessage{ModuleError(00005), "用户不存在"}
	ErrLoginOauthParams  = &CodeMessage{ModuleError(00006), "oauth参数错误"}
	ErrCreateCategory    = &CodeMessage{ModuleError(10001), "创建分类失败"}
	ErrGetCategory       = &CodeMessage{ModuleError(10002), "获取分类失败"}
	ErrListCategory      = &CodeMessage{ModuleError(10003), "查询分类失败"}
	ErrUpdateCategory    = &CodeMessage{ModuleError(10004), "修改分类失败"}
	ErrCreateArticle     = &CodeMessage{ModuleError(20001), "创建文章失败"}
	ErrGetArticle        = &CodeMessage{ModuleError(20002), "获取文章失败"}
	ErrListArticle       = &CodeMessage{ModuleError(20003), "查询文章失败"}
	ErrUpdateArticle     = &CodeMessage{ModuleError(20004), "修改文章失败"}
)
