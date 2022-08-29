import request from '@/utils/request'

/**
 * 用户注册
 * @param {Object} params
 */
export function register(params) {
    return request.post('/user/register', params);
}


/**
 * 获取用户信息
 * @param {*} params 
 */
export function food_comment(params) {
    return request.getLoading('/food/comment/createComment', params);
}



