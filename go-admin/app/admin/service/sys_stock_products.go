package service

import (
	"errors"

    "github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type SysStockProducts struct {
	service.Service
}

// GetPage 获取SysStockProducts列表
func (e *SysStockProducts) GetPage(c *dto.SysStockProductsGetPageReq, p *actions.DataPermission, list *[]models.SysStockProducts, count *int64) error {
	var err error
	var data models.SysStockProducts

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SysStockProductsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取SysStockProducts对象
func (e *SysStockProducts) Get(d *dto.SysStockProductsGetReq, p *actions.DataPermission, model *models.SysStockProducts) error {
	var data models.SysStockProducts

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSysStockProducts error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SysStockProducts对象
func (e *SysStockProducts) Insert(c *dto.SysStockProductsInsertReq) error {
    var err error
    var data models.SysStockProducts
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SysStockProductsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改SysStockProducts对象
func (e *SysStockProducts) Update(c *dto.SysStockProductsUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.SysStockProducts{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("SysStockProductsService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除SysStockProducts
func (e *SysStockProducts) Remove(d *dto.SysStockProductsDeleteReq, p *actions.DataPermission) error {
	var data models.SysStockProducts

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveSysStockProducts error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
