package model

import "fmt"

type QiyuResult struct {
	Code    int
	Message string
}

func (qr *QiyuResult) String() string {
	return fmt.Sprintf("Code: %d, Message: %s", qr.Code, qr.Message)
}
