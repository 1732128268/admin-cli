package system

import (
	"admin-cli/global"
	"admin-cli/model"
	"admin-cli/model/gen"
	"admin-cli/model/request"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

type OperationRecordApi struct{}

func CreateSysOperationRecord(sysOperationRecord model.SysOperationRecord) error {
	return global.Db.Create(&sysOperationRecord).Error
}

// GetSysOperationRecordList SysOperationRecord列表
func (s *OperationRecordApi) GetSysOperationRecordList(c *gin.Context) {
	size := c.DefaultQuery("size", "10")      //每页数
	current := c.DefaultQuery("current", "1") //当前页
	//总页数
	var total int64
	global.Db.Model(&model.SysOperationRecord{}).Count(&total)
	var sysOperationRecordList []model.SysOperationRecord
	err := global.Db.Scopes(gen.Paginate(cast.ToInt(current), cast.ToInt(size))).Preload("User").Find(&sysOperationRecordList).Error
	if err != nil {
		logrus.Errorf("get sysOperationRecord list error: %v", err)
		global.Response(c, nil, err)
		return
	}
	global.Response(c, gin.H{
		"total":   total,   //总条数
		"current": current, //当前页
		"size":    size,    //每页数
		"list":    sysOperationRecordList,
	}, nil)
}

// FindSysOperationRecord 用id查询SysOperationRecord
func (s *OperationRecordApi) FindSysOperationRecord(c *gin.Context) {
	var id request.GetById
	if err := c.ShouldBind(&id); err != nil {
		logrus.Errorf("查询SysOperationRecord参数错误:%v", err)
		global.ValidatorResponse(c, err)
		return
	}
	var sysOperationRecord model.SysOperationRecord
	if err := global.Db.Where("id = ?", id.ID).First(&sysOperationRecord).Error; err != nil {
		logrus.Errorf("查询SysOperationRecord失败:%v", err)
		global.Response(c, nil, err)
		return
	}
	global.Response(c, gin.H{
		"sysOperationRecord": sysOperationRecord,
	}, nil)

}

// CreateSysOperationRecord 创建SysOperationRecord
func (s *OperationRecordApi) CreateSysOperationRecord(c *gin.Context) {
	var sysOperationRecord model.SysOperationRecord
	if err := c.ShouldBind(&sysOperationRecord); err != nil {
		logrus.Errorf("创建SysOperationRecord参数错误:%v", err)
		global.ValidatorResponse(c, err)
		return
	}
	if err := CreateSysOperationRecord(sysOperationRecord); err != nil {
		logrus.Errorf("创建SysOperationRecord失败:%v", err)
		global.Response(c, nil, err)
		return
	}
	global.Response(c, gin.H{
		"sysOperationRecord": sysOperationRecord,
	}, nil)
}

// DeleteSysOperationRecord 删除SysOperationRecord
func (s *OperationRecordApi) DeleteSysOperationRecord(c *gin.Context) {
	var id request.GetById
	if err := c.ShouldBind(&id); err != nil {
		logrus.Errorf("删除SysOperationRecord参数错误:%v", err)
		global.ValidatorResponse(c, err)
		return
	}
	err := global.Db.Delete(&model.SysOperationRecord{}, "id = ?", id.ID).Error
	if err != nil {
		logrus.Errorf("删除SysOperationRecord失败:%v", err)
		global.Response(c, nil, err)
		return
	}
	global.Response(c, nil, nil)
}

// DeleteSysOperationRecordByIds 批量删除SysOperationRecord
func (s *OperationRecordApi) DeleteSysOperationRecordByIds(c *gin.Context) {
	var ids request.IdsReq
	if err := c.ShouldBind(&ids); err != nil {
		logrus.Errorf("删除SysOperationRecord参数错误:%v", err)
		global.ValidatorResponse(c, err)
		return
	}
	if err := global.Db.Where("id in (?)", ids.Ids).Delete(&model.SysOperationRecord{}).Error; err != nil {
		logrus.Errorf("删除SysOperationRecord失败:%v", err)
		global.Response(c, nil, err)
		return
	}
	global.Response(c, nil, nil)
}
