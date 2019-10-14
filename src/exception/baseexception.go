package exception

import "fmt"

type BaseException struct {
	Code        int
	Description string
}

// 实现error接口
func (be *BaseException) Error() string {
	exceptionContent := `
		code:%d
		description:%s
`
	return fmt.Sprintf(exceptionContent, be.Code, be.Description)
}
