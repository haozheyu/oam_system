
/**
 * 设置token
 */
export function set_token(token) {
   uni.setStorageSync('token',token)
}

/**
 * 获取token
 */
export function get_token() {
    return uni.getStorageSync('token')
}