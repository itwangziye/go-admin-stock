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

type SysStockInventory struct {
	api.Api
}

// GetPage 获取库存列表
// @Summary 获取库存列表
// @Description 获取库存列表
// @Tags 库存
// @Param productId query string false "商品编码"
// @Param quantity query string false "数量"
// @Param lastUpdated query time.Time false "最后更新时间"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SysStockInventory}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-stock-inventory [get]
// @Security Bearer
func (e SysStockInventory) GetPage(c *gin.Context) {
    req := dto.SysStockInventoryGetPageReq{}
    s := service.SysStockInventory{}
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
	list := make([]models.SysStockInventory, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取库存失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取库存
// @Summary 获取库存
// @Description 获取库存
// @Tags 库存
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.SysStockInventory} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-stock-inventory/{id} [get]
// @Security Bearer
func (e SysStockInventory) Get(c *gin.Context) {
	req := dto.SysStockInventoryGetReq{}
	s := service.SysStockInventory{}
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
	var object models.SysStockInventory

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取库存失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建库存
// @Summary 创建库存
// @Description 创建库存
// @Tags 库存
// @Accept application/json
// @Product application/json
// @Param data body dto.SysStockInventoryInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/sys-stock-inventory [post]
// @Security Bearer
func (e SysStockInventory) Insert(c *gin.Context) {
    req := dto.SysStockInventoryInsertReq{}
    s := service.SysStockInventory{}
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
		e.Error(500, err, fmt.Sprintf("创建库存失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改库存
// @Summary 修改库存
// @Description 修改库存
// @Tags 库存
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SysStockInventoryUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/sys-stock-inventory/{id} [put]
// @Security Bearer
func (e SysStockInventory) Update(c *gin.Context) {
    req := dto.SysStockInventoryUpdateReq{}
    s := service.SysStockInventory{}
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
		e.Error(500, err, fmt.Sprintf("修改库存失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除库存
// @Summary 删除库存
// @Description 删除库存
// @Tags 库存
// @Param data body dto.SysStockInventoryDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/sys-stock-inventory [delete]
// @Security Bearer
func (e SysStockInventory) Delete(c *gin.Context) {
    s := service.SysStockInventory{}
    req := dto.SysStockInventoryDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除库存失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
