package newvideo

// GetVideoResponse 视频拉取
type GetVideoResponse struct {
	Data   []GetVideoData   `json:"data,omitempty" dc:"data信息"`
	Expand []GetVideoExpand `json:"expand,omitempty" dc:"expand信息"`
}

type GetVideoData struct {
	Videoid   int64  `json:"videoid,omitempty" dc:"视频id"`
	Userid    int64  `json:"userid,omitempty" dc:"账户id"`
	Url       string `json:"url,omitempty" dc:"视频url"`
	VideoName string `json:"videoName,omitempty" dc:"视频名称"`
	//视频来源：2-信息流，10-创意中心自提，13-搜索，20-基木鱼自提，19-慧合，18,23,32-视频工具
	Source    int     `json:"source,omitempty" dc:"视频来源"`
	Capacity  float64 `json:"capacity,omitempty" dc:"视频大小，单位MB"`
	Format    string  `json:"format,omitempty" dc:"视频格式"`
	Width     int     `json:"width,omitempty" dc:"视频宽度，单位px"`
	Height    int     `json:"height,omitempty" dc:"视频高度，单位px	"`
	Duration  int     `json:"duration,omitempty" dc:"时长，单位s"`
	AddTime   string  `json:"addTime,omitempty" dc:"添加时间，格式：yyyy-mm-dd hh:mm:ss"`
	ModTime   string  `json:"modTime,omitempty" dc:"修改时间，格式：yyyy-mm-dd hh:mm:ss"`
	Thumbnail string  `json:"thumbnail,omitempty" dc:"封面图片url，仅做预览使用。为空代表该视频没有封面图。"`
	//转码状态：枚举值，列举如下
	//0-转码成功：可以用于投放视频样式，也可以分享给其他账户。
	//30-转码中：转码中的视频暂时不可以投放与分享，需等待转码成功。
	//40-转码失败：转码失败视频的视频不可以投放与分享
	Istranscode int    `json:"istranscode,omitempty" dc:"转码状态"`
	FromUserId  int64  `json:"fromUserId,omitempty" dc:"由谁分享,为0代码非分享视频"`
	VideoMd5    string `json:"videoMd5,omitempty" dc:"视频文件md5"`
}

type GetVideoExpand struct {
	TotalCount int64 `json:"totalCount,omitempty" dc:"视频总数量"`
}
