import request from '@/utils/request'

// 查询SysStockProducts列表
export function listSysStockProducts(query) {
    return request({
        url: '/api/v1/sys-stock-products',
        method: 'get',
        params: query
    })
}

// 查询SysStockProducts详细
export function getSysStockProducts (id) {
    return request({
        url: '/api/v1/sys-stock-products/' + id,
        method: 'get'
    })
}


// 新增SysStockProducts
export function addSysStockProducts(data) {
    return request({
        url: '/api/v1/sys-stock-products',
        method: 'post',
        data: data
    })
}

// 修改SysStockProducts
export function updateSysStockProducts(data) {
    return request({
        url: '/api/v1/sys-stock-products/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除SysStockProducts
export function delSysStockProducts(data) {
    return request({
        url: '/api/v1/sys-stock-products',
        method: 'delete',
        data: data
    })
}

