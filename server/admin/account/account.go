package account

import (
	"boadmin/server/middleware/session"
	"boadmin/server/util"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type UserName struct {
	Id   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"size:100;comment:'名字'" json:"name"`
}

type UserBase struct {
	//Id			uint   `gorm:"primary_key;auto_increment"`
	//Name		string `gorm:"size:100;comment:'名字'"`
	UserName
	Acct		string	`gorm:"size:64;comment:'帐号'" json:"acct"`
	Access   	uint	`gorm:"size:11;comment:'权限'" json:"access"`
	Last 		string	`gorm:"size:30;comment:'最后登录ip地址'" json:"last"`
	Valid		uint	`gorm:"size:11;comment:'是否有效0.无效 1有效'" json:"valid"`//
	Ctm 		int64	`gorm:"size:11;comment:'创建时间'" json:"ctm"`
	Utm 		int64	`gorm:"size:11;comment:'更新时间'" json:"utm"`
	Skin 		uint8	`gorm:"size:6;comment:'更新时间'" json:"skin"`
}

type UserData struct {
	UserBase
	Password  	string `gorm:"size:200;comment:'密码'"`
}



type UserInfo struct {
	UserData
	Token string // "用户登录凭证"
}

// TableName 为UserData 绑定表名
func (UserData) TableName() string {
	return "account"
}

type AccountRes struct {
	util.Res
}

type ChangePasswordInfo struct {
	Id    uint `json:"id"`
	Old   string `json:"old_password"`
	New    uint `json:"new_password"`
}

// Index 页面
func (ret *AccountRes)Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		// var users []UserBase
		// util.DB.Table("account").Find(&users)
		ret.Html(c, "account.html", gin.H{})
	}
}

// ChangePassword 更新密码
func (ret *AccountRes)ChangePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		// old := c.PostForm("old_password")
		// println(old)
		old := c.PostForm("old")
		chg := c.PostForm("chg")
		cl := len(chg)
		if cl < 6 || cl > 12 {
			ret.Json(c, util.PassLen, "密码长度错误")
			return
		}

		ssInfo, er := session.GetInfo(c)
		if er != nil {
			ret.Json(c, util.Login, "没有登录")
			return
		}

		old = util.Encryption(ssInfo.Acct, old, util.Ini.SBase.PasswordSalt)
		if old != ssInfo.Pass {
			ret.Json(c, util.Pass, "错误的密码")
			return
		}

		err := InnerChangePassword(c, ssInfo.Acct, chg, ssInfo.Id)
		if err != nil {
			ret.Json(c, util.Fail, "修改失败")
			return
		}

		AddActionLog(ActionLogInfo{
			Aid: ssInfo.Id,
			Opm: 1, // 帐户操作
			Ops: 5, // 修改密码
			Par4: ssInfo.Name,
		})

		ret.Json(c, util.Success, "修改成功")
	}
}

// ResetPassword 重置密码
func (ret *AccountRes)ResetPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid := c.PostForm("id")
		id := 0
		if sid != "" {
			var err error
			id, err = strconv.Atoi(sid)
			if err != nil {
				ret.Json(c, util.Acct, "param err id")
				return
			}
			if id < 0 {
				ret.Json(c, util.Acct, "param err id")
				return
			}
		}

		user := UserData{}
		util.DB.Where("id = ?", id).Take(&user)

		err := InnerChangePassword(c, user.Acct, util.Ini.SBase.DefPass, user.Id)
		if err != nil {
			ret.Json(c, util.Fail, "修改失败")
			return
		}

		AddActionLog(ActionLogInfo{
			Aid: user.Id,
			Opm: 1, // 帐户操作
			Ops: 4, // 重置密码
			Par4: user.Name,
		})

		ret.Json(c, util.Success, "修改成功")
	}
}

// InnerChangePassword 密码修改
func InnerChangePassword(c *gin.Context, acct string, pass string, id uint) error {
	if pass == "" {
		return nil
	}
	passSalt := util.Encryption(acct, pass, util.Ini.SBase.PasswordSalt)
	err := util.DB.Model(&UserData{}).Where("id = ?", id).Update("password", passSalt).Error
	if err == nil {
		session.Mod(c, "pass", passSalt)
	}
	return err
}

// ChangeSkin 更换皮肤
func (ret *AccountRes)ChangeSkin() gin.HandlerFunc {
	return func(c *gin.Context) {
		par := c.PostForm("skin")
		skin, err := strconv.Atoi(par)
		if err != nil {
			ret.Json(c, util.Fail, "修改失败:参数错误")
			return
		}
		id, er := session.Get(c, "id")
		if er != nil {
			ret.Json(c, util.Login, "没有登录")
			return
		}
		old, _ := session.Get(c, "skin")
		// old := ssInfo.Skin
		errSql := util.DB.Model(&UserData{}).Where("id = ?", id).Update("skin", skin).Error
		if errSql != nil {
			ret.Json(c, util.Fail, "修改失败")
			return
		}

		if session.Mod(c, "skin", uint8(skin)) {
			AddActionLog(ActionLogInfo{
				Aid: id.(uint),
				Opm: 1, // 帐户操作
				Ops: 3, // 更换皮肤
				Val: uint64(skin),
				Chg: old.(int64),
				Par1: 1,
			})
			ret.Json(c, util.Success, skin)
		} else {
			ret.Json(c, util.Fail, skin)
		}
	}
}

// Page 按而读取页数据
func (ret *AccountRes)Page() gin.HandlerFunc {
	return func(c *gin.Context) {
		page := 1
		limit := 10
		pageStr := c.Query("page")
		if pageStr != "" {
			tmp, err := strconv.Atoi(pageStr)
			if err == nil {
				page = tmp
			}
		}
		limitStr := c.Query("limit")
		if limitStr != "" {
			tmp, err := strconv.Atoi(limitStr)
			if err == nil {
				limit = tmp
			}
		}

		var users []UserBase
		offset := (page - 1) * limit
		var totalSize int64
		util.DB.Table("account").Count(&totalSize).Order("id ASC").Offset(offset).Limit(limit).Find(&users)

		ret.Table(c, util.Success, totalSize, "", users)
	}
}

func GetIdNames() []UserName {
	var users []UserName
	util.DB.Table("account").Find(&users)
	return users
}

// Group 返回按分组的权限
func (ret *AccountRes)Group() gin.HandlerFunc {
	return func(c *gin.Context) {
		ret.Json(c, util.Success, gin.H{
			"group":   Group,
			"defPass": util.Ini.SBase.DefPass,
		})
	}
}

// Modify 新增/修改属性
func (ret *AccountRes)Modify() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid := c.PostForm("id")
		id := 0
		if sid != "" {
			var err error
			id, err = strconv.Atoi(sid)
			if err != nil {
				ret.Json(c, util.Acct, "param err id")
				return
			}
			if id < 0 {
				ret.Json(c, util.Acct, "param err id")
				return
			}
		}

		acct := c.PostForm("acct")
		if acct == "" {
			ret.Json(c, util.Param, "param err acct")
			return
		}

		name := c.PostForm("name")
		if name == "" {
			ret.Json(c, util.Param, "param err name")
			return
		}

		valid, err := strconv.Atoi(c.PostForm("valid"))
		if err != nil {
			ret.Json(c, util.Param, "param err valid")
			return
		}

		access, err := strconv.Atoi( c.PostForm("access"))
		if err != nil {
			ret.Json(c, util.Param, "param err access")
			return
		}

		user := UserData{}
		util.DB.Where("id = ?", id).Take(&user)

		user.Acct = acct
		user.Name = name
		user.Valid = uint(valid)
		user.Access = uint(access)


		if id == 0 /* 新增 */{
			user.Ctm = time.Now().Unix()
			user.Password = util.Encryption(user.Acct, util.Ini.SBase.DefPass, util.Ini.SBase.PasswordSalt)
		}

		saveRet := util.DB.Save(&user)
		if saveRet.Error != nil {
			ret.Json(c, util.DbSave, "save err")
			return
		}

		sessionInfo, sessionErr := session.GetInfo(c)
		var logInfo ActionLogInfo
		var userBase UserBase
		if id == 0 {
			saveRet.Last(&userBase)
			if sessionErr == nil {
				logInfo.Aid = sessionInfo.Id
				logInfo.Opm = 1
				logInfo.Ops = 1
				logInfo.Par4 = name
				AddActionLog(logInfo)
			}
		} else {
			userBase = user.UserBase
			logInfo.Aid = sessionInfo.Id
			logInfo.Opm = 1
			logInfo.Ops = 2
			logInfo.Par4 = name
			if sessionInfo.Name != name {
				logInfo.Par1 = 1
			}
			if valid > 0 {
				logInfo.Par2 = 1
				logInfo.Val = uint64(valid)
			}
			if sessionInfo.Access != uint(access) {
				logInfo.Par3 = 1
				logInfo.Chg = int64(access)
			}
			AddActionLog(logInfo)

		}
		ret.Json(c, util.Success, userBase)
	}
}


// UpdateLoginHost 更新登录ip
func UpdateLoginHost(id uint, lastHost string) error {
	return util.DB.Model(&UserData{}).Where("id = ? AND last <> ?", id, lastHost).Update("last", lastHost).Error
}


// Info 取出帐号信息
func Info(id uint) (UserInfo, error) {
	var user UserInfo
	err := util.DB.Where("id = ?", id).First(&user).Error
	return user, err
}

// InfoByAcct 根据acct读取信息
func InfoByAcct(acct string) (UserInfo, error) {
	var user UserInfo
	err := util.DB.Where("acct = ?", acct).First(&user).Error
	return user, err
}
