package balance

import (
	"fmt"

	"github.com/openthe88/baidu-marketing/model"
)

// GetAccountTransferHistoryRequest 查询转账记录 API Request
type GetAccountTransferHistoryRequest struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

func (r GetAccountTransferHistoryRequest) Url() string {
	return fmt.Sprintf("%sMccHistoryService/getAccountTransferHistory", model.BASE_URL_SMS)
}
