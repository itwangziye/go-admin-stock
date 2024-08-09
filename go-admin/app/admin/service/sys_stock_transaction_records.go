package service

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type SysStockTransactionRecords struct {
	service.Service
}

// GetPage 获取SysStockTransactionRecords列表
func (e *SysStockTransactionRecords) GetPage(c *dto.SysStockTransactionRecordsGetPageReq, p *actions.DataPermission, list *[]models.SysStockTransactionRecords, count *int64) error {
	var err error
	var data models.SysStockTransactionRecords

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Order("created_at DESC").
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SysStockTransactionRecordsService GetPage error:%s \r\n", err)
		return err

	}
	return nil
}

// Get 获取SysStockTransactionRecords对象
func (e *SysStockTransactionRecords) Get(d *dto.SysStockTransactionRecordsGetReq, p *actions.DataPermission, model *models.SysStockTransactionRecords) error {
	var data models.SysStockTransactionRecords

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSysStockTransactionRecords error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SysStockTransactionRecords对象
func (e *SysStockTransactionRecords) Insert(c *dto.SysStockTransactionRecordsInsertReq) error {
	return e.Orm.Transaction(func(tx *gorm.DB) error {
		var err error
		var data models.SysStockTransactionRecords
		var inventory models.SysStockInventory

		// 查找库存记录
		result := tx.Find(&inventory, "product_id = ?", c.ProductId)
		if result.Error != nil {
			return fmt.Errorf("查找库存记录失败: %w", result.Error)
		}

		if result.RowsAffected == 0 {
			// 如果库存不存在，则创建库存
			SysStockInventoryInsertReq := dto.SysStockInventoryInsertReq{
				ProductId:   c.ProductId,
				Quantity:    "0",
				LastUpdated: time.Now(),
			}
			SysStockInventoryInsertReq.SetCreateBy(c.CreateBy)
			SysStockInventoryInsertReq.Generate(&inventory)

			if err := tx.Create(&inventory).Error; err != nil {
				return fmt.Errorf("创建库存记录失败: %w", err)
			}
		}

		// 计算变更量
		quantityChange, err := strconv.Atoi(c.Quantity)
		if err != nil {
			return fmt.Errorf("无效的数量: %w", err)
		}

		if c.TransactionType == "OUT" {
			quantityChange = -quantityChange
		}

		// 更新库存
		originInventoryInt, err := strconv.Atoi(inventory.Quantity)
		if err != nil {
			return fmt.Errorf("无效的库存数量: %w", err)
		}

		newInventoryQuantity := originInventoryInt + quantityChange
		if newInventoryQuantity < 0 {
			return errors.New("库存不足")
		}

		inventory.Quantity = strconv.Itoa(newInventoryQuantity)

		if err := tx.Save(&inventory).Error; err != nil {
			return fmt.Errorf("更新库存失败: %w", err)
		}

		// 出入库记录
		c.Generate(&data)
		err = e.Orm.Create(&data).Error
		if err != nil {
			e.Log.Errorf("SysStockTransactionRecordsService Insert error: %s \r\n", err)
			return fmt.Errorf("插入出入库记录失败: %w", err)
		}
		return nil
	})
}

// Update 修改SysStockTransactionRecords对象
func (e *SysStockTransactionRecords) Update(c *dto.SysStockTransactionRecordsUpdateReq, p *actions.DataPermission) error {

	return e.Orm.Transaction(func(tx *gorm.DB) error {
		var err error
		var data = models.SysStockTransactionRecords{}
		e.Orm.Scopes(
			actions.Permission(data.TableName(), p),
		).First(&data, c.GetId())
		c.Generate(&data)

		db := e.Orm.Save(&data)

		if err = db.Error; err != nil {
			e.Log.Errorf("SysStockTransactionRecordsService Save error:%s \r\n", err)
			return err
		}
		if db.RowsAffected == 0 {
			return errors.New("无权更新该数据~")
		}
		return nil
	})

}

func (e *SysStockTransactionRecords) UpdateInventory(tx *gorm.DB, productId string, quantityChange int) error {
	var inventory models.SysStockInventory

	// 查找库存记录
	if err := tx.First(&inventory, "product_id = ?", productId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果库存记录不存在
			return errors.New("库存记录不存在")
		}
		return fmt.Errorf("查找库存记录失败: %w", err)
	}

	// 更新现有库存记录
	quantity, _ := strconv.Atoi(inventory.Quantity)

	inventory.Quantity = strconv.Itoa(quantity + quantityChange)

	inventory.UpdatedAt = time.Now()
	if err := tx.Save(&inventory).Error; err != nil {
		return fmt.Errorf("更新库存记录失败: %w", err)
	}

	return nil
}

// Remove 删除SysStockTransactionRecords
func (e *SysStockTransactionRecords) Remove(d *dto.SysStockTransactionRecordsDeleteReq, p *actions.DataPermission) error {

	return e.Orm.Transaction(func(tx *gorm.DB) error {
		var records []models.SysStockTransactionRecords
		var data models.SysStockTransactionRecords

		// 查找所有需要删除的库存交易记录
		if err := tx.Model(&models.SysStockTransactionRecords{}).
			Scopes(
				actions.Permission(models.SysStockTransactionRecords{}.TableName(), p),
			).
			Where("id IN ?", d.Ids).Find(&records).Error; err != nil {
			e.Log.Errorf("Service RemoveSysStockTransactionRecords error:%s \r\n", err)
			return err
		}

		deleteResult := tx.Model(&data).
			Scopes(
				actions.Permission(data.TableName(), p),
			).Delete(&data, d.GetId())

		if deleteResult.Error != nil {
			e.Log.Errorf("Service RemoveSysStockTransactionRecords error:%s \r\n", deleteResult.Error)
			return deleteResult.Error
		}

		if deleteResult.RowsAffected == 0 {
			return errors.New("无权删除该数据")
		}

		// 更新库存表
		for _, record := range records {
			quantity, _ := strconv.Atoi(record.Quantity)
			if record.TransactionType == "IN" {
				quantity = -quantity
			}
			if err := e.UpdateInventory(tx, record.ProductId, quantity); err != nil {
				e.Log.Errorf("UpdateInventory error: %s \r\n", err)
				return err
			}
		}

		return nil
	})

}
