package consts

// 权限类型
const (
	PERMISSION_TYPE_DIR  = 1 // 目录
	PERMISSION_TYPE_MENU = 2 // 菜单
	PERMISSION_TYPE_ABLE = 3 // 功能
)

// 系统必须权限名称
const PERMISSION_SYSTEM_REQUIRED_NAME = "system-required"

// 主机终端会话文件默认保存目录，可通过配置 host.terminal.sessionFileDir 修改
const HOST_TERMINAL_SESSION_SAVE_DIRECTORY = "host-sessions"
