package account

var AccountManage uint = 1 << 0   // 帐号管理 1
var AccountRead uint = 1 << 1     // 帐号查询 2
var DeveloperManage uint = 1 << 2 // 开发管理 4
var DeveloperRead uint = 1 << 3   // 开发查询 8
var DevOpsManage uint = 1 << 4    // 运维管理 16
var DevOpsRead uint = 1 << 5      // 运维查询 32
var OperationManage uint = 1 << 6 // 运营管理 64
var OperationRead uint = 1 << 7   // 运营查询 128
var BusinessManage uint = 1 << 8  // 业务管理 256
var BusinessRead uint = 1 << 9    // 业务查询 512

var Names = map[uint]string {
	AccountManage:   "帐号管理",
	AccountRead:     "帐号查询",
	DeveloperManage: "开发管理",
	DeveloperRead:   "开发查询",
	DevOpsManage:    "运维管理",
	DevOpsRead:      "运维查询",
	OperationManage: "运营管理",
	OperationRead:   "运营查询",
	BusinessManage:  "业务管理",
	BusinessRead:    "业务查询",
}

var Group = []map[uint]string {
	{
		AccountManage: "帐号管理",
		AccountRead:   "帐号查询",
	},
	{
		DeveloperManage: "开发管理",
		DeveloperRead:   "开发查询",
	},
	{
		DevOpsManage: "运维管理",
		DevOpsRead:   "运维查询",
	},
	{
		OperationManage: "运营管理",
		OperationRead:   "运营查询",
	},
	{
		BusinessManage: "业务管理",
		BusinessRead:   "业务查询",
	},
}

// IsAccess 是否有权限 uAccess：用户自身的权限码 hAccess:页面要求的权限码
func IsAccess(uAccess uint, hAccess uint) bool {
	if hAccess == 0 {
		return true
	}

	if uAccess <= 0 {
		return false
	}

	return uAccess & hAccess != 0
}

// GetNames 获取当前权限下的各个名称
func GetNames(uAccess uint) map[uint]string {
	return Names
}