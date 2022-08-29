/**
 * 封装对本地存储的操作
 */

const publicChar = 'wk_'
export const get = key => {
  return uni.getStorageSync(publicChar + key)
}

export const set = (key, value) => {
  return uni.setStorageSync(publicChar + key, value)
}

export const remove = key => {
  uni.removeStorageSync(publicChar + key)
}

export const clear = () => {
  const { keys } = uni.getStorageInfoSync()
  keys.forEach(key => {
    if (key.includes(publicChar)) {
      uni.removeStorageSync(key)
    }
  })
}
export const clearWith = (excludeKeys = []) => {
  const { keys } = uni.getStorageInfoSync()
  keys.forEach(key => {
    if (key.includes(publicChar)) {
      excludeKeys.forEach(excludeKey => {
        excludeKey = publicChar + excludeKey
        if (key !== excludeKey) {
          uni.removeStorageSync(key)
        }
      })
    }
  })
}

export default {
  get,
  set,
  remove,
  clear,
  clearWith
}
