import request from '@/utils/request'

/**
 * 菜系列表
 * @param {Object} params
 */
export function food_type_list(params) {
    return request.get('/foodType/list', params);
}


