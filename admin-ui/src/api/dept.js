import request from '../utils/request'

/**
 * 部门列表
 */
export const Deptlist = data => {
    return request({
        url: '/api/user/dept/list',
        method: 'POST'
    })
}

/**
 * 添加部门
 */
 export const Adddept = data => {
    return request({
        url: '/api/user/dept/add',
        method: 'POST',
        data: {
            "name": data.name
        }
    })
}

/**
 * 删除部门
 */
 export const Deldept = data => {
    return request({
        url: '/api/user/dept/delete',
        method: 'POST',
        data: {
            "id": data.id
        }
    })
}

/**
 * 部门更新
 */
 export const Updatedept = data => {
    return request({
        url: '/api/user/dept/update',
        method: 'POST',
        data: {
            "id": data.id,
            "name": data.name
        }
    })
}