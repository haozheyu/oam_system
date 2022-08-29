import request from '@/utils/request'


/**
 * 餐厅列表
 * @param {Object} paramss
 */
export function restaurant_list(params) {
    return request.getLoading('/restaurant/getList', params);
}




