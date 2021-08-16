package home

import (
	"boadmin/server/admin/account"
)

type Nav struct {
	Title string       `json:"title"`
	Access uint        `json:"access"`
	Icon string        `json:"icon"`
	Spread bool        `json:"spread"`
	Href string        `json:"href"`
	Children    []Node `json:"children"`
}

type Node struct {
	Title  string `json:"title"`
	Access  uint `json:"access"`
	Icon string `json:"icon"`
	Href string `json:"href"`
}



var Navs []Nav = []Nav{
	{
		Title:  "系统",
		Access: account.AccountManage | account.AccountRead,
		Icon:   "fa fa-cog",
		Spread: false, //是否默认展开
		Children: []Node{
			{
				Title:  "帐号管理",
				Href:   "/account",
				Access: account.AccountManage | account.AccountRead,
				Icon:   "fa fa-cog",
			},
			{
				Title:  "layui-border-orange", // 分隔线的色彩
				Href:   "",
				Access: account.AccountManage | account.AccountRead,
				Icon:   "",
			},
			{
				Title:  "帐号行为",
				Href:   "/action",
				Access: account.AccountManage | account.AccountRead,
				Icon:   "fa fa-cog",
			},
		},
	},
	{
		Title:  "示例",
		Access: account.AccountManage | account.AccountRead,
		Icon:   "fa fa-cog",
		Children: []Node{
			{
				Title:  "附件管理",
				Href:   "/admin/demo/show",
				Access: account.AccountManage | account.AccountRead,
				Icon:   "fa fa-cog",
			},
		},
	},
	{
		Title:  "示例",
		Access: account.AccountManage | account.AccountRead,
		Icon:   "fa fa-cog",
		Children: []Node{
			{
				Title:  "附件管理",
				Href:   "/admin/demo/show",
				Access: account.AccountManage | account.AccountRead,
				Icon:   "fa fa-cog",
			},
		},
	},
	{
		Title:  "示例",
		Access: account.AccountManage | account.AccountRead,
		Icon:   "fa fa-cog",
		Children: []Node{
			{
				Title:  "附件管理",
				Href:   "/admin/demo/show",
				Access: account.AccountManage | account.AccountRead,
				Icon:   "fa fa-cog",
			},
		},
	},
	{
		Title:  "示例",
		Access: account.AccountManage | account.AccountRead,
		Icon:   "fa fa-cog",
		Children: []Node{
			{
				Title:  "附件管理",
				Href:   "/admin/demo/show",
				Access: account.AccountManage | account.AccountRead,
				Icon:   "fa fa-cog",
			},
		},
	},
}

func GetNavs(uAccess uint) []Nav {
	var tmp []Nav
	for _, nav := range Navs {
		if account.IsAccess(uAccess, nav.Access) {
			var m Nav
			m.Title = nav.Title
			m.Access = nav.Access
			m.Access = nav.Access
			m.Icon =  nav.Icon

			for _, node := range nav.Children {
				if account.IsAccess(uAccess, node.Access) {
					m.Children = append(m.Children, node)
				}
			}
			tmp = append(tmp, m)
		}
	}
	return tmp
}

