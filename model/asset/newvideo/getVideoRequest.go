package newvideo

import (
	"fmt"
	"github.com/openthe88/baidu-marketing/model"
)

// GetVideoRequest 视频拉取
type GetVideoRequest struct {
	Ids       []int64 `json:"ids" dc:"视频id集合"`
	PageSize  int     `json:"pageSize" dc:"每页记录的数量"`
	PageNo    int     `json:"pageNo" dc:"第几页"`
	VideoName string  `json:"videoName,omitempty" dc:"视频名称"`
	//取值范围：枚举值，列表如下
	//ALL - 全部
	//FC - 搜索推广
	//FEED - 信息流推广
	//JMY - 基木鱼平台
	//SHOPPING - 电商推广
	Channel string `json:"channel,omitempty" dc:"渠道筛选，字符串类型，默认值为ALL"`
	//取值范围：枚举值，列表如下
	//ALL - 全部
	//LOCAL - 本地上传
	//HUIHE - 慧合制作
	//TOOL - 视频工具
	Source    string `json:"source,omitempty" dc:"来源筛选，字符串类型，默认值为ALL"`
	StartDate string `json:"startDate,omitempty" dc:"开始日期（修改时间） 格式：yyyy-MM-dd"`
	EndDate   string `json:"endDate,omitempty" dc:"结束日期日期（修改时间） 格式：yyyy-MM-dd"`
	//取值范围：枚举值，列表如下
	//asc - 按照修改时间降序排序
	//desc - 按照修改时间升序排序
	Order string `json:"order,omitempty" dc:"asc-按照修改时间倒序排列，desc-按照修改时间降序排列，默认为desc"`
}

func (r GetVideoRequest) Url() string {
	return fmt.Sprintf("%sVideoService/getVideoInfos", model.BASE_URL_SMS)
}
