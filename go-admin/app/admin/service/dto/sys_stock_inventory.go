package dto

import (
	"time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SysStockInventoryGetPageReq struct {
	dto.Pagination `search:"-"`
	ProductId      string    `form:"productId"  search:"type:exact;column:product_id;table:sys_stock_inventory" comment:"商品编码"`
	Quantity       string    `form:"quantity"  search:"type:exact;column:quantity;table:sys_stock_inventory" comment:"数量"`
	LastUpdated    time.Time `form:"lastUpdated"  search:"type:exact;column:last_updated;table:sys_stock_inventory" comment:"最后更新时间"`
	SysStockInventoryOrder
}

type SysStockInventoryOrder struct {
	Id          string `form:"idOrder"  search:"type:order;column:id;table:sys_stock_inventory"`
	ProductId   string `form:"productIdOrder"  search:"type:order;column:product_id;table:sys_stock_inventory"`
	Quantity    string `form:"quantityOrder"  search:"type:order;column:quantity;table:sys_stock_inventory"`
	LastUpdated string `form:"lastUpdatedOrder"  search:"type:order;column:last_updated;table:sys_stock_inventory"`
	CreatedAt   string `form:"createdAtOrder"  search:"type:order;column:created_at;table:sys_stock_inventory"`
	UpdatedAt   string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:sys_stock_inventory"`
	DeletedAt   string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:sys_stock_inventory"`
	CreateBy    string `form:"createByOrder"  search:"type:order;column:create_by;table:sys_stock_inventory"`
	UpdateBy    string `form:"updateByOrder"  search:"type:order;column:update_by;table:sys_stock_inventory"`
}

func (m *SysStockInventoryGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysStockInventoryInsertReq struct {
	Id          int       `json:"-" comment:"库存编码"` // 库存编码
	ProductId   string    `json:"productId" comment:"商品编码"`
	Quantity    string    `json:"quantity" comment:"数量"`
	LastUpdated time.Time `json:"lastUpdated" comment:"最后更新时间"`
	common.ControlBy
}

func (s *SysStockInventoryInsertReq) Generate(model *models.SysStockInventory) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ProductId = s.ProductId
	model.Quantity = s.Quantity
	model.LastUpdated = s.LastUpdated
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *SysStockInventoryInsertReq) GetId() interface{} {
	return s.Id
}

type SysStockInventoryUpdateReq struct {
	Id          int       `uri:"id" comment:"库存编码"` // 库存编码
	ProductId   string    `json:"productId" comment:"商品编码"`
	Quantity    string    `json:"quantity" comment:"数量"`
	LastUpdated time.Time `json:"lastUpdated" comment:"最后更新时间"`
	common.ControlBy
}

func (s *SysStockInventoryUpdateReq) Generate(model *models.SysStockInventory) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ProductId = s.ProductId
	model.Quantity = s.Quantity
	model.LastUpdated = s.LastUpdated
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *SysStockInventoryUpdateReq) GetId() interface{} {
	return s.Id
}

// SysStockInventoryGetReq 功能获取请求参数
type SysStockInventoryGetReq struct {
	Id int `uri:"id"`
}

func (s *SysStockInventoryGetReq) GetId() interface{} {
	return s.Id
}

// SysStockInventoryDeleteReq 功能删除请求参数
type SysStockInventoryDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysStockInventoryDeleteReq) GetId() interface{} {
	return s.Ids
}
