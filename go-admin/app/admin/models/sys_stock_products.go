package models

import (
	"go-admin/common/models"
)

type SysStockProducts struct {
	models.Model

	Name         string `json:"name" gorm:"type:varchar(128);comment:商品名称"`
	Description  string `json:"description" gorm:"type:varchar(255);comment:商品描述"`
	PriceInCents string `json:"priceInCents" gorm:"type:int unsigned;comment:价格（分）"`
	models.ModelTime
	models.ControlBy
}

func (SysStockProducts) TableName() string {
	return "sys_stock_products"
}

func (e *SysStockProducts) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysStockProducts) GetId() interface{} {
	return e.Id
}
