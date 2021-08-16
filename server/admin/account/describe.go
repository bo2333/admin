package account

// ActionLogDescribe  日志描述 第一个key = opm
var ActionLogDescribe = map[uint16]interface{} {
	0: map[string]interface{} { "name": "全部"},
	1: map[string]interface{} {
		"name": "帐户操作",
		"ops": map[uint]interface{} {
			0: map[string]interface{} { "name": "全部ops"},
			1: map[string]interface{} {
				"name": "添加帐号",
				"func": 1,
			},
			2: map[string]interface{} {
				"name": "修改帐号",
				"func": 1,
				"par1": map[uint16]interface{} {
					0: "全部par1",
					1: "修改名字",
				},
				"par2": map[uint16]interface{} {
					0: "全部par1",
					1: "修改valid=(val)",
				},
				"par3": map[uint16]interface{} {
					0: "全部par1",
					1: "修改access=(chg)",
				},
			},
			3: map[string]interface{} {
				"name": "更换皮肤",
				"func": 1,
				"val": "当前皮肤",
				"chg": "原来皮肤",
				"par1": map[uint16]interface{} {
					0: "全部par1",
					1: "val 当前皮肤, chg 原来皮肤",
				},
			},
			4: map[string]interface{} {
				"name": "重置密码",
				"func": 1,
			},
			5: map[string]interface{} {
				"name": "修改密码",
				"func": 1,
			},
		},
	},
}