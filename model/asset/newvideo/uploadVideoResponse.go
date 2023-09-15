package newvideo

// UploadVideoResponse 上传视频返回参数
type UploadVideoResponse struct {
	AddTime        string  `json:"addTime,omitempty" dc:"添加时间"`
	AudioRate      float64 `json:"audioRate,omitempty" dc:"音频码率 kbps"`
	AuditContent   string  `json:"auditContent,omitempty" dc:"待审核视频信息"`
	BitRate        float64 `json:"bitRate,omitempty" dc:"比特率 kbps"`
	Capacity       float64 `json:"capacity,omitempty" dc:"视频大小 MB"`
	Content        string  `json:"content,omitempty" dc:"审核通过的视频等信息"`
	Duration       int     `json:"duration,omitempty" dc:"视频时长 单位-秒"`
	Format         string  `json:"format,omitempty" dc:"视频格式"`
	FrameRate      float64 `json:"frameRate,omitempty" dc:"帧率 fps"`
	Height         int     `json:"height,omitempty" dc:"视频尺寸-高 单位-像素"`
	Width          int     `json:"width,omitempty" dc:"视频尺寸-宽 单位-像素"`
	ModTime        string  `json:"modTime,omitempty" dc:"修改时间"`
	Source         int     `json:"source,omitempty" dc:"产品线"`
	Thumbnail      string  `json:"thumbnail,omitempty" dc:"封面图 url"`
	Uploaded       bool    `json:"uploaded,omitempty" dc:"true:该视频已上传过 null/false:该视频未上传过"`
	Url            string  `json:"url,omitempty" dc:"视频 url"`
	UserId         int64   `json:"userId,omitempty" dc:"用户id"`
	VideoCodec     string  `json:"videoCodec,omitempty" dc:"视频编码容器 举例 h264等"`
	VideContentMd5 string  `json:"videContentMd5,omitempty" dc:"视频元信息 md5"`
	VideoId        int64   `json:"videoId,omitempty" dc:"视频 id"`
	VideoMd5       string  `json:"videoMd5,omitempty" dc:"视频 md5"`
	VideoName      string  `json:"videoName,omitempty" dc:"视频名称"`
	VideoRate      float64 `json:"videoRate,omitempty" dc:"视频码率 kbps"`
}
