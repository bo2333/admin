package account

import (
	"boadmin/server/util"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type ActionLogRes struct {
	util.Res
}

type ActionLogInfo struct {
	Aid  uint        `gorm:"size:11;comment:'用户id'" json:"aid"`
	Opm uint        `gorm:"size:11;comment:'操作-类'" json:"opm"`
	Ops uint        `gorm:"size:11;comment:'操作-子类'" json:"ops"`
	Val uint64      `gorm:"size:16;comment:'当前值'" json:"val"`
	Chg int64       `gorm:"size:16;comment:'改变值'" json:"chg"`
	Par1 uint64     `gorm:"size:16;comment:'参数1'" json:"par1"`
	Par2 uint64     `gorm:"size:16;comment:'参数2'" json:"par2"`
	Par3 int64      `gorm:"size:16;comment:'参数3'" json:"par3"`
	Par4 string		`gorm:"size:256;comment:'参数4'" json:"par4"`
}

type ActionLogData struct {
	Id int       	`gorm:"primary_key;auto_increment" json:"id"`
	ActionLogInfo
	Ctm int64 		`gorm:"size:16;comment:'权限'" json:"ctm"`
}


// TableName 为UserData 绑定表名
func (ActionLogData) TableName() string {
	return "action_log"
}


// Index 页面
func (ret *ActionLogRes)Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		//var users []UserBase
		//util.DB.Table("account").Find(&users)
		ret.Html(c, "action-log.html", gin.H{})
	}
}

func (ret *ActionLogRes)Data() gin.HandlerFunc {
	return func(c *gin.Context) {
		ret.Json(c, util.Success, gin.H{
			"describe" : ActionLogDescribe,
			"accounts" : GetIdNames(),
		})
	}
}

// Page 按而读取页数据
func (ret *ActionLogRes)Page() gin.HandlerFunc {
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

		db := util.DB.Table("action_log")
		whereInt := func(key string) int {
			tmp := c.Query(key)
			v, e := strconv.Atoi(tmp)
			if e == nil {
				if v > 0 {
					db = db.Where(key + " = ?", v)
					return v
				}
			}
			return 0
		}

		whereTm := func(key string, compared string) int64 {
			tmp := c.Query(key)
			v, e := time.ParseInLocation("2006-01-02 15:04:05" , tmp, time.Local)
			if e == nil {
				if v.Unix() > 0 {
					db = db.Where("ctm " + compared + " ?", v.Unix())
					return v.Unix()
				}
			}
			return 0
		}

		whereInt("aid")
		whereInt("opm")
		whereInt("ops")
		whereInt("par1")
		whereInt("par2")
		whereInt("par3")
		stm := whereTm("stm", ">=")
		etm := whereTm("etm", "<=")
		if stm > etm && etm != 0 {
			ret.Json(c, util.Acct, "param err stm > etm")
		}
		var totalSize int64
		var infos []ActionLogData
		offset := (page - 1) * limit
		db.Count(&totalSize).Order("id ASC").Offset(offset).Limit(limit).Find(&infos)

		ret.Table(c, util.Success, totalSize, "", infos)
	}
}

func AddActionLog(info ActionLogInfo) int {
	cell := ActionLogData{
		Id: 0,
		ActionLogInfo: info,
	}
	cell.Ctm = time.Now().Unix()
	saveRet := util.DB.Save(&cell)
	if saveRet.Error != nil {
		return util.DbSave
	}
	return util.Success
}
