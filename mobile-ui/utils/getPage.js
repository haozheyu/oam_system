import { isEmpty } from './index.js'

export const getCurrentPage = () => {
  // 获取所有页面
  const pages = getCurrentPages()

  const currentPage = !isEmpty(pages) ? (pages[pages.length - 1]).route : null

  return currentPage
}

export const getPrevPage = () => {
  // 获取所有页面
  const pages = getCurrentPages()

  const PrevPage = !isEmpty(pages) ? (pages[pages.length - 2]).route : null

  return PrevPage
}
