import request from '@/utils/request'

/**
 * 用户注册
 * @param {Object} params
 */
export function register(params) {
    return request.post('/user/register', params);
}

/**
 * 用户登录
 * @param {Object} params
 */
export function login(params) {
    return request.post('/user/loginMobile', params);
}

/**
 * 获取用户信息
 * @param {*} params 
 */
export function user_info(params) {
    return request.getLoading('/user/getUserInfo', params);
}
/**
 * 修改用户信息
 * @param {} params 
 */
export function update_user(params) {
    return request.postLoading('/user/updateUserInfo', params);
}


