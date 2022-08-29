// 判断传入参数是否一个数组
export function isArray(arr) {
  return Object.prototype.toString.apply(arr) === '[object Array]'
}

// 对手机号脱敏
export function hidePhone(phone) {
  const reg = /(\d{3})\d*(\d{4})/
  const res = phone.replace(reg, '$1****$2')
  return res
}

export function isEmpty(value) {
  return value === undefined || value === null || value === ''
}
