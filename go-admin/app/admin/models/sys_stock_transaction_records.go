package models

import (
	"time"

	"go-admin/common/models"
)

type SysStockTransactionRecords struct {
	models.Model

	ProductId       string    `json:"productId" gorm:"type:int unsigned;comment:商品编码"`
	TransactionType string    `json:"transactionType" gorm:"type:enum('IN','OUT');comment:交易类型"`
	Quantity        string    `json:"quantity" gorm:"type:int;comment:数量"`
	TransactionDate time.Time `json:"transactionDate" gorm:"type:timestamp;comment:交易日期"`
	Remarks         string    `json:"remarks" gorm:"type:varchar(255);comment:备注"`
	models.ModelTime
	models.ControlBy
}

func (SysStockTransactionRecords) TableName() string {
	return "sys_stock_transaction_records"
}

func (e *SysStockTransactionRecords) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysStockTransactionRecords) GetId() interface{} {
	return e.Id
}
