import request from '../utils/request'

/**
 * 角色列表
 */
export const Rolelist = data => {
    return request({
        url: '/api/user/role/list',
        method: 'POST'
    })
}

/**
 * 添加角色
 */
 export const AddRole = data => {
    return request({
        url: '/api/user/role/add',
        method: 'POST',
        data: {
            "name": data.name
        }
    })
}

/**
 * 删除角色
 */
 export const DelRole = data => {
    return request({
        url: '/api/user/role/delete',
        method: 'POST',
        data: {
            "id": data.id
        }
    })
}

/**
 * 角色更新
 */
 export const UpdateRole = data => {
    return request({
        url: '/api/user/role/update',
        method: 'POST',
        data: {
            "id": data.id,
            "name": data.name
        }
    })
}