import request from '@/utils/request'

// 查询SysStockInventory列表
export function listSysStockInventory(query) {
    return request({
        url: '/api/v1/sys-stock-inventory',
        method: 'get',
        params: query
    })
}

// 查询SysStockInventory详细
export function getSysStockInventory (id) {
    return request({
        url: '/api/v1/sys-stock-inventory/' + id,
        method: 'get'
    })
}


// 新增SysStockInventory
export function addSysStockInventory(data) {
    return request({
        url: '/api/v1/sys-stock-inventory',
        method: 'post',
        data: data
    })
}

// 修改SysStockInventory
export function updateSysStockInventory(data) {
    return request({
        url: '/api/v1/sys-stock-inventory/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除SysStockInventory
export function delSysStockInventory(data) {
    return request({
        url: '/api/v1/sys-stock-inventory',
        method: 'delete',
        data: data
    })
}

