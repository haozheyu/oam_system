import request from '@/utils/request'



/**
 * 创建搜索
 * @param {Object} params
 */
export function create_search_content(params) {
    return request.getLoading('/historySearch/create', params);
}

/**
 * 获取用户信息
 * @param {*} params 
 */
export function search_list(params) {
    return request.getLoading('/historySearch/list', params);
}



