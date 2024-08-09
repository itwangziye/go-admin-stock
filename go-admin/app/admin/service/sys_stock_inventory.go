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

type SysStockInventory struct {
	service.Service
}

// GetPage 获取SysStockInventory列表
func (e *SysStockInventory) GetPage(c *dto.SysStockInventoryGetPageReq, p *actions.DataPermission, list *[]models.SysStockInventory, count *int64) error {
	var err error
	var data models.SysStockInventory

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SysStockInventoryService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取SysStockInventory对象
func (e *SysStockInventory) Get(d *dto.SysStockInventoryGetReq, p *actions.DataPermission, model *models.SysStockInventory) error {
	var data models.SysStockInventory

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSysStockInventory error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SysStockInventory对象
func (e *SysStockInventory) Insert(c *dto.SysStockInventoryInsertReq) error {
    var err error
    var data models.SysStockInventory
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SysStockInventoryService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改SysStockInventory对象
func (e *SysStockInventory) Update(c *dto.SysStockInventoryUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.SysStockInventory{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("SysStockInventoryService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除SysStockInventory
func (e *SysStockInventory) Remove(d *dto.SysStockInventoryDeleteReq, p *actions.DataPermission) error {
	var data models.SysStockInventory

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveSysStockInventory error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
