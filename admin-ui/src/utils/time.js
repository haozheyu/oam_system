import { getItem } from "./storage";

/**
 * 时间戳转换
 */
 export function getTimeStamp(value) {
    return new Date(parseInt(value) * 1000).toLocaleString().replace(/:\d  {1,2}$/,' ');
  }
  /**
   * 设置时间戳
   */
  export function setTimeStamp() {
    return Date.parse(new Date());
  }
  /**
   * 是否超时
   */
  export function isCheckTimeout() {
    // 当前时间戳
    var currentTime = Date.parse(new Date());
    // 缓存时间戳
    var timeStamp = getItem("refreshAfter")
    return currentTime < timeStamp 
  }
  