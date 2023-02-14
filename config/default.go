package config

var (
	// DSN 数据库地址(只需要修改这个就可以了,不必修改main.go中的数据库地址 这样就可以只修改一个地方)
	UseDSN = PankerDSN
	// Server 服务器地址 或者本机地址 用于生成视频的播放地址
	UseServer = PankerServer
	// 监听端口
	Port = ":8888"

	// 成员数据库地址和服务器地址 用于修改上方的 UseDSN 和 UseServer
	DSN       = "root:1234567890@tcp(182.92.131.42:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	DSN_local = "root:Zsm@20020609@tcp(localhost:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"

	PankerDSN    = "root:123456@tcp(127.0.0.1:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	PankerServer = "http://192.168.43.6"

	KM911Server    = "http://192.168.1.3"
	KM911LocalDSN  = "root:@Dzg15484@tcp(localhost:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	KM911RemoteDSN = "root:@Dzg15484@tcp(81.68.91.70)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
)
