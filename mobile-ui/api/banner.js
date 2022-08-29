import request from '@/utils/request'

/**
 * 广告列表
 * @param {Object} params
 */
export function banner_list(params) {
    return request.get('/banner/list', params);
}


