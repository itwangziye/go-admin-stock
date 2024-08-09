package dto

import (
	"time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SysStockTransactionRecordsGetPageReq struct {
	dto.Pagination  `search:"-"`
	TransactionType string    `form:"transactionType"  search:"type:exact;column:transaction_type;table:sys_stock_transaction_records" comment:"交易类型"`
	TransactionDate time.Time `form:"transactionDate"  search:"type:exact;column:transaction_date;table:sys_stock_transaction_records" comment:"交易日期"`
	Remarks         string    `form:"remarks"  search:"type:contains;column:remarks;table:sys_stock_transaction_records" comment:"备注"`
	SysStockTransactionRecordsOrder
}

type SysStockTransactionRecordsOrder struct {
	Id              string `form:"idOrder"  search:"type:order;column:id;table:sys_stock_transaction_records"`
	ProductId       string `form:"productIdOrder"  search:"type:order;column:product_id;table:sys_stock_transaction_records"`
	TransactionType string `form:"transactionTypeOrder"  search:"type:order;column:transaction_type;table:sys_stock_transaction_records"`
	Quantity        string `form:"quantityOrder"  search:"type:order;column:quantity;table:sys_stock_transaction_records"`
	TransactionDate string `form:"transactionDateOrder"  search:"type:order;column:transaction_date;table:sys_stock_transaction_records"`
	Remarks         string `form:"remarksOrder"  search:"type:order;column:remarks;table:sys_stock_transaction_records"`
	CreatedAt       string `form:"createdAtOrder"  search:"type:order;column:created_at;table:sys_stock_transaction_records"`
	UpdatedAt       string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:sys_stock_transaction_records"`
	DeletedAt       string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:sys_stock_transaction_records"`
	CreateBy        string `form:"createByOrder"  search:"type:order;column:create_by;table:sys_stock_transaction_records"`
	UpdateBy        string `form:"updateByOrder"  search:"type:order;column:update_by;table:sys_stock_transaction_records"`
}

func (m *SysStockTransactionRecordsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysStockTransactionRecordsInsertReq struct {
	Id              int       `json:"-" comment:"交易记录编码"` // 交易记录编码
	ProductId       string    `json:"productId" comment:"商品编码"`
	TransactionType string    `json:"transactionType" comment:"交易类型"`
	Quantity        string    `json:"quantity" comment:"数量"`
	TransactionDate time.Time `json:"transactionDate" comment:"交易日期"`
	Remarks         string    `json:"remarks" comment:"备注"`
	common.ControlBy
}

func (s *SysStockTransactionRecordsInsertReq) Generate(model *models.SysStockTransactionRecords) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ProductId = s.ProductId
	model.TransactionType = s.TransactionType
	model.Quantity = s.Quantity
	model.TransactionDate = s.TransactionDate
	model.Remarks = s.Remarks
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *SysStockTransactionRecordsInsertReq) GetId() interface{} {
	return s.Id
}

type SysStockTransactionRecordsUpdateReq struct {
	Id              int       `uri:"id" comment:"交易记录编码"` // 交易记录编码
	ProductId       string    `json:"productId" comment:"商品编码"`
	TransactionType string    `json:"transactionType" comment:"交易类型"`
	Quantity        string    `json:"quantity" comment:"数量"`
	TransactionDate time.Time `json:"transactionDate" comment:"交易日期"`
	Remarks         string    `json:"remarks" comment:"备注"`
	common.ControlBy
}

func (s *SysStockTransactionRecordsUpdateReq) Generate(model *models.SysStockTransactionRecords) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ProductId = s.ProductId
	model.TransactionType = s.TransactionType
	model.Quantity = s.Quantity
	model.TransactionDate = s.TransactionDate
	model.Remarks = s.Remarks
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *SysStockTransactionRecordsUpdateReq) GetId() interface{} {
	return s.Id
}

// SysStockTransactionRecordsGetReq 功能获取请求参数
type SysStockTransactionRecordsGetReq struct {
	Id int `uri:"id"`
}

func (s *SysStockTransactionRecordsGetReq) GetId() interface{} {
	return s.Id
}

// SysStockTransactionRecordsDeleteReq 功能删除请求参数
type SysStockTransactionRecordsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysStockTransactionRecordsDeleteReq) GetId() interface{} {
	return s.Ids
}
