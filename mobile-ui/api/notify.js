import request from '@/utils/request'

/**
 * 通知列表
 * @param {Object} params
 */
export function notify_list(params) {
    return request.getLoading('/notify/notifyList', params);
}
/**
 * ��ҳ֪ͨ
 * @param {Object} paramss
 */
export function sys_notify_list(params) {
    return request.getLoading('/sysNotify/getList', params);
}




