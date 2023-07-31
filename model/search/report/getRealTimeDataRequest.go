package report

import (
	"fmt"

	"github.com/openthe88/baidu-marketing/model"
)

// GetRealTimeDataRequest 推广报告 API Request
type GetRealTimeDataRequest struct {
	RealTimeRequestType *RealTimeRequest `json:"realTimeRequestType"`
}

func (r GetRealTimeDataRequest) Url() string {
	return fmt.Sprintf("%sReportService/getRealTimeData", model.BASE_URL_SMS)
}
