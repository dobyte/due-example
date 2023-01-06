package route

const (
	Register      int32 = iota + 1 // 注册账号
	Login                          // 登录账号
	CreateRoom                     // 创建房间
	JoinRoom                       // 加入房间
	NotifyMessage                  // 通知消息
)
