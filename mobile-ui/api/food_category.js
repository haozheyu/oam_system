import request from '@/utils/request'

/**
 * 菜品列表
 * @param {Object} params
 */
export function food_category_list(params) {
    return request.get('/food/category/list', params);
}


