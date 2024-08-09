package models

import (
	"time"

	"go-admin/common/models"
)

type SysStockInventory struct {
	models.Model

	ProductId   string    `json:"productId" gorm:"type:int unsigned;comment:商品编码"`
	Quantity    string    `json:"quantity" gorm:"type:int;comment:数量"`
	LastUpdated time.Time `json:"lastUpdated" gorm:"type:timestamp;comment:最后更新时间"`
	models.ModelTime
	models.ControlBy
}

func (SysStockInventory) TableName() string {
	return "sys_stock_inventory"
}

func (e *SysStockInventory) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysStockInventory) GetId() interface{} {
	return e.Id
}
