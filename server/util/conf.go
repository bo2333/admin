package util

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type SIni struct {
	SMysql `ini:"mysql"`
	SBase  `ini:"base"`
	SLogo  `ini:"logo"`
}

type SMysql struct {
	Host        string `ini:"host"`
	Port        string `ini:"port"`
	UserName    string `ini:"username"`
	Password    string `ini:"password"`
	Database    string `ini:"database"`
	MaxOpenConn int    `ini:"max_open_conn"`
	MaxIdleConn int    `ini:"max_idle_conn"`
}

type SBase struct {
	Port     		string `ini:"port"`
	Host     		string `ini:"host"`
	FillData 		bool   `ini:"fill_data"`
	PasswordSalt    string `ini:"password_salt"`
	LogoIco    		string `ini:"logo_ico"`
	LogoTitle    	string `ini:"logo_title"`
	Website			string `ini:"website"`
	Email			string `ini:"email"`
	Version			string `ini:"version"`
	DefPass 		string `ini:"def_pass"`
}

type SLogo struct {
	Icon    	string `ini:"icon"`
	Title    	string `ini:"title"`
}

var Ini = new(SIni)

//初始化配置文件
func init() {
	path, err := RootPath()
	if err != nil {
		fmt.Printf("get root path err:%v", err)
	}
	fmt.Println("rootPath:" + path)
	err = ini.MapTo(Ini, path + "/resources/conf/config.ini")
	if err != nil {
		fmt.Printf("load ini err:%v", err)
	}
}
