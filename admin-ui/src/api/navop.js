import request from '../utils/request'

/**
 * 导航列表
 */
 export const Navlist = (data) => {
    return request({
        url: '/api/nav/list_body',
        method: 'POST',
        data: {
            "page": data.page,
            "pageSize": data.pageSize,
        }
    })
}

/**
 * 添加导航标题内容
 */
 export const AddNavTopBody = (data) => {
    return request({
        url: '/api/nav/add_body',
        method: 'POST',
        data: {
            "top_name_id": data.top_name_id,
            "link_addr": data.link_addr,
            "link_name": data.link_name,
            "link_desc": data.link_desc,
            "link_icon": data.link_icon
        }
    })
}

/**
 * 添加导航标题
 */
 export const AddNavTop = (data) => {
    return request({
        url: '/api/nav/add_top',
        method: 'POST',
        data: {
            "name": data.name
        }
    })
}

/**
 * 删除导航标题内容
 */
 export const DelNavTopBody = (data) => {
    return request({
        url: '/api/nav/delete_body',
        method: 'POST',
        data: {
            "link_name": data.link_name
        }
    })
}

/**
 * 删除导航标题
 */
 export const DelNavTop = (data) => {
    return request({
        url: '/api/nav/delete_top',
        method: 'POST',
        data: {
            "name": data.name
        }
    })
}

/**
 * 根据导航名获取导航内容
 */
 export const GetNavBody = (data) => {
    return request({
        url: '/api/nav/get_top_body',
        method: 'POST',
        data: {
            "name": data.name,
            "page": data.page,
            "pageSize": data.pageSize
        }
    })
}

/**
 * 获取导航标题
 */
 export const GetNavTopList = () => {
    return request({
        url: '/api/nav/list_top',
        method: 'POST'
    })
}

/**
 * 更新导航标题内容
 */
 export const UpdateNavTopList = (data) => {
    return request({
        url: '/api/nav/update',
        method: 'POST',
        data: {
            "link_addr": data.link_addr,
            "link_name": data.link_name,
            "link_desc": data.link_desc,
            "link_icon": data.link_icon
        }
    })
}