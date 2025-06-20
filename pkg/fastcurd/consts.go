package fastcurd

type (
	Code int
)

const (
	CodeOk Code = iota
	CodeDefaultError
	CodeNoAuth
	CodeBadReq
	CodeValidError
	CodeNoLogin
	CodeServerError
	CodeRateLimitError
)

const (
	PrimaryField    = "id"
	CreateTimeField = "ctime"
	UpdateTimeField = "utime"
)
