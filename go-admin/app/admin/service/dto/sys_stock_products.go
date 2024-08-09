package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SysStockProductsGetPageReq struct {
	dto.Pagination     `search:"-"`
    Name string `form:"name"  search:"type:exact;column:name;table:sys_stock_products" comment:"商品名称"`
    Description string `form:"description"  search:"type:exact;column:description;table:sys_stock_products" comment:"商品描述"`
    PriceInCents string `form:"priceInCents"  search:"type:exact;column:price_in_cents;table:sys_stock_products" comment:"价格（分）"`
    SysStockProductsOrder
}

type SysStockProductsOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:sys_stock_products"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:sys_stock_products"`
    Description string `form:"descriptionOrder"  search:"type:order;column:description;table:sys_stock_products"`
    PriceInCents string `form:"priceInCentsOrder"  search:"type:order;column:price_in_cents;table:sys_stock_products"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:sys_stock_products"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:sys_stock_products"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:sys_stock_products"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:sys_stock_products"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:sys_stock_products"`
    
}

func (m *SysStockProductsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysStockProductsInsertReq struct {
    Id int `json:"-" comment:"商品编码"` // 商品编码
    Name string `json:"name" comment:"商品名称"`
    Description string `json:"description" comment:"商品描述"`
    PriceInCents string `json:"priceInCents" comment:"价格（分）"`
    common.ControlBy
}

func (s *SysStockProductsInsertReq) Generate(model *models.SysStockProducts)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Description = s.Description
    model.PriceInCents = s.PriceInCents
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *SysStockProductsInsertReq) GetId() interface{} {
	return s.Id
}

type SysStockProductsUpdateReq struct {
    Id int `uri:"id" comment:"商品编码"` // 商品编码
    Name string `json:"name" comment:"商品名称"`
    Description string `json:"description" comment:"商品描述"`
    PriceInCents string `json:"priceInCents" comment:"价格（分）"`
    common.ControlBy
}

func (s *SysStockProductsUpdateReq) Generate(model *models.SysStockProducts)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Description = s.Description
    model.PriceInCents = s.PriceInCents
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *SysStockProductsUpdateReq) GetId() interface{} {
	return s.Id
}

// SysStockProductsGetReq 功能获取请求参数
type SysStockProductsGetReq struct {
     Id int `uri:"id"`
}
func (s *SysStockProductsGetReq) GetId() interface{} {
	return s.Id
}

// SysStockProductsDeleteReq 功能删除请求参数
type SysStockProductsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysStockProductsDeleteReq) GetId() interface{} {
	return s.Ids
}
