import request from '@/utils/request'

/**
 * 需求列表
 * @param {Object} params
 */
export function demand_list(params) {
    return request.post('/demand/page', params);
}

/**
 * 需求详情
 * @param {Object} params
 */
export function demand_detail(params) {
    return request.getLoading('/demand/detail', params);
}


/**
 * 申请接单
 * @param {Object} params
 */
export function apply_order_taking(params) {
    return request.getLoading('/demand/applyOrderTaking', params);
}

/*
	下拉技术类型
 * @param {Object} params
 */
export function techType_pull(params) {
    return request.get('/tech/type/listPull', params);
}

/*
	下拉预算
 * @param {Object} params
 */
export function estimate_list(params) {
    return request.get('/estimate/amount/list', params);
}

/**
 * 更新浏览量
 * @param {Object} params
 */
export function update_view_count(params) {
    return request.get('/demand/updateViewCount', params);
}


/**
 * 发布需求
 * @param {Object} params
 */
export function publish_demand(params) {
    return request.post('/demand/publish', params);
}

