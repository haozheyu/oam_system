import request from '../utils/request'

/**
 * 添加主机
 */
export const AddHost = data => {
    return request({
        url: '/api/host/addHost',
        method: 'POST',
        data: {
            "addr": data.addr,
            "user": data.user,
            "port": data.port,
            "password": data.password
        }
    })
}

/**
 * 编辑主机
 */
export const EditHost = data => {
    return request({
        url: '/api/host/editHost',
        method: 'POST',
        data: {
            "user": data.user,
            "addr": data.addr,
            "port": data.port,
            "password": data.password,
            "isDel": data.isDel
        }
    })
}

/** 主机列表
 */
export const HostList = data => {
    return request({
        url: '/api/host/hosts',
        method: 'POST',
        data: {
            "page": data.page,
            "pageSize": data.pageSize
        }
    })
}

