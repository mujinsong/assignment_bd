package config

var (
	// DSN 数据库地址
	UseDSN = DSN_local
	// Server 服务器地址 或者本机地址(需要端口号) 用于生成视频流url
	Server = "http://192.168.1.5"
	// 监听端口
	Port = ":8888"

	DSN       = "root:1234567890@tcp(182.92.131.42:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	DSN_local = "root:Zsm@20020609@tcp(localhost:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	PankerDSN = "root:123456@tcp(127.0.0.1:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"

	KM911LocalDSN  = "root:@Dzg15484@tcp(localhost:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	KM911RemoteDSN = "root:@Dzg15484@tcp(81.68.91.70)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
)
