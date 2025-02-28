package balance

import (
	"fmt"

	"github.com/openthe88/baidu-marketing/model"
)

// GetPaymentRecordRequest 查询付款信息与待加款信息 API Request
type GetPaymentRecordRequest struct {
	// FundType 查询付款记录类型，不同类型返回的参数不同，详见返回结果：1：直销客户实际资金流动; 2：直销客户优惠资金流动;21：KA实际资金流动;22：KA优惠资金流动;23：KA待加款记录
	FundType int `json:"fundType,omitempty"`
	// Condition 查询条件对象
	Condition *model.Condition `json:"condition,omitempty"`
	// ChunkSize 单次请求数据行数，默认500
	ChunkSize int64 `json:"chunkSize,omitempty"`
	// Offset 返回数据行数偏移值，只支持id的gt（大于）、lt（小于），如：{ "id": { "gt": 0 } }
	Offset map[string]map[string]int64 `json:"offset,omitempty"`
	// Sort 排序字段和排序方向列表，只支持id的asc或desc，如：{"in": [ "id asc",  ]}
	Sort map[string][]string
}

func (r GetPaymentRecordRequest) Url() string {
	return fmt.Sprintf("%sPaymentService/getPaymentRecord", model.BASE_URL_SMS)
}
