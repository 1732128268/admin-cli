package router

import (
	"admin-cli/serve/api/v1/system"
	"github.com/gin-gonic/gin"
)

func SysOperationRecordRouter(router *gin.RouterGroup) {
	operationRecordRouter := router.Group("sysOperationRecord")
	operationRecordApi := system.OperationRecordApi{}
	{
		operationRecordRouter.GET("getSysOperationRecordList", operationRecordApi.GetSysOperationRecordList)          // 获取SysOperationRecord列表
		operationRecordRouter.POST("findSysOperationRecord", operationRecordApi.FindSysOperationRecord)               // 根据ID获取SysOperationRecord
		operationRecordRouter.POST("createSysOperationRecord", operationRecordApi.CreateSysOperationRecord)           // 新建SysOperationRecord
		operationRecordRouter.POST("deleteSysOperationRecord", operationRecordApi.DeleteSysOperationRecord)           // 删除SysOperationRecord
		operationRecordRouter.POST("deleteSysOperationRecordByIds", operationRecordApi.DeleteSysOperationRecordByIds) // 批量删除SysOperationRecord
	}
}
