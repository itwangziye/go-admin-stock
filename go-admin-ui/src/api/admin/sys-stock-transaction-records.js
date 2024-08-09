import request from '@/utils/request'

// 查询SysStockTransactionRecords列表
export function listSysStockTransactionRecords(query) {
    return request({
        url: '/api/v1/sys-stock-transaction-records',
        method: 'get',
        params: query
    })
}

// 查询SysStockTransactionRecords详细
export function getSysStockTransactionRecords (id) {
    return request({
        url: '/api/v1/sys-stock-transaction-records/' + id,
        method: 'get'
    })
}


// 新增SysStockTransactionRecords
export function addSysStockTransactionRecords(data) {
    return request({
        url: '/api/v1/sys-stock-transaction-records',
        method: 'post',
        data: data
    })
}

// 修改SysStockTransactionRecords
export function updateSysStockTransactionRecords(data) {
    return request({
        url: '/api/v1/sys-stock-transaction-records/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除SysStockTransactionRecords
export function delSysStockTransactionRecords(data) {
    return request({
        url: '/api/v1/sys-stock-transaction-records',
        method: 'delete',
        data: data
    })
}

