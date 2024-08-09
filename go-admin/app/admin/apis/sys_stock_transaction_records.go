package apis

import (
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type SysStockTransactionRecords struct {
	api.Api
}

// GetPage 获取交易记录列表
// @Summary 获取交易记录列表
// @Description 获取交易记录列表
// @Tags 交易记录
// @Param transactionType query string false "交易类型"
// @Param transactionDate query time.Time false "交易日期"
// @Param remarks query string false "备注"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SysStockTransactionRecords}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-stock-transaction-records [get]
// @Security Bearer
func (e SysStockTransactionRecords) GetPage(c *gin.Context) {
    req := dto.SysStockTransactionRecordsGetPageReq{}
    s := service.SysStockTransactionRecords{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
   	if err != nil {
   		e.Logger.Error(err)
   		e.Error(500, err, err.Error())
   		return
   	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.SysStockTransactionRecords, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取交易记录失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取交易记录
// @Summary 获取交易记录
// @Description 获取交易记录
// @Tags 交易记录
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.SysStockTransactionRecords} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-stock-transaction-records/{id} [get]
// @Security Bearer
func (e SysStockTransactionRecords) Get(c *gin.Context) {
	req := dto.SysStockTransactionRecordsGetReq{}
	s := service.SysStockTransactionRecords{}
    err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.SysStockTransactionRecords

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取交易记录失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建交易记录
// @Summary 创建交易记录
// @Description 创建交易记录
// @Tags 交易记录
// @Accept application/json
// @Product application/json
// @Param data body dto.SysStockTransactionRecordsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/sys-stock-transaction-records [post]
// @Security Bearer
func (e SysStockTransactionRecords) Insert(c *gin.Context) {
    req := dto.SysStockTransactionRecordsInsertReq{}
    s := service.SysStockTransactionRecords{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建交易记录失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改交易记录
// @Summary 修改交易记录
// @Description 修改交易记录
// @Tags 交易记录
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SysStockTransactionRecordsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/sys-stock-transaction-records/{id} [put]
// @Security Bearer
func (e SysStockTransactionRecords) Update(c *gin.Context) {
    req := dto.SysStockTransactionRecordsUpdateReq{}
    s := service.SysStockTransactionRecords{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改交易记录失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除交易记录
// @Summary 删除交易记录
// @Description 删除交易记录
// @Tags 交易记录
// @Param data body dto.SysStockTransactionRecordsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/sys-stock-transaction-records [delete]
// @Security Bearer
func (e SysStockTransactionRecords) Delete(c *gin.Context) {
    s := service.SysStockTransactionRecords{}
    req := dto.SysStockTransactionRecordsDeleteReq{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除交易记录失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
