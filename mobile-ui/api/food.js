import request from '@/utils/request'

/**
 * 推荐列表
 * @param {Object} params
 */
export function food_recommend(params) {
    return request.post('/food/recommend', params);
}
/**
 * 发布食谱
 * @param {Object} params
 */
export function publish_food(params) {
    return request.post('/food/publish', params);
}

/**
 * 食谱列表广场
 * @param {Object} params
 */
export function food_list(params) {
    return request.post('/food/foodList', params);
}

/**
 * 发布笔记
 * @param {Object} params
 */
export function create_notes(params) {
    return request.post('/food/notes/createNotes', params);
}

/**
 * 笔记列表
 * @param {*} params 
 */
export function note_list(params) {
    return request.get('/food/notes/list', params);
}
/**
 * 获取分类
 * @param {*} params 
 */
export function food_categroy_detail(params) {
    return request.get('/food/getFoodsListByCategoryId', params);
}

/**
 * 首页热门菜品
 * @param {*} params 
 */
export function food_hot_list(params) {
    return request.post('/food/foodHotList', params);
}

export function food_detail(params) {
    return request.get('/food/detail', params);
}

/**
 * 收藏
 * @param {*} params 
 */
export function food_collect(params) {
    return request.getLoading('/food/collect', params);
}


/**
 * 我发布的食谱
 * @param {*} params 
 */
export function mine_publish_food(params) {
    return request.getLoading('/food/mineFoods', params);
}

/**
 * 我发布的笔记
 * @param {} params 
 */
export function mine_publish_notes(params) {
    return request.getLoading('/food/notes/minePublishList', params);
}

/**
 * 食谱收藏列表
 * @param {*} params 
 */
export function mine_collect_list(params) {
    return request.getLoading('/food/collectList', params);
}

export function food_comm_detail(params) {
    return request.getLoading('/food/foodCommDetail', params);
}
/**
 * 查询窗口列表根据餐厅ID
 * @param {Object} params
 */
export function get_windows_by_restaurant(params) {
    return request.getLoading('/food/getWindowsByRestaurantId', params);
}

/**
 * 创建订单
 * @param {Object} params
 */
export function food_order_create(params) {
    return request.postLoading('/food/order/create', params);
}

/**
 * 订单列表
 * @param {Object} params
 */
export function food_order_list(params) {
    return request.getLoading('/food/order/orderList', params);
}

/**
 * 修改订单状态
 * @param {Object} params
 */
export function food_order_update_status(params) {
    return request.getLoading('/food/order/updateStatus', params);
}




