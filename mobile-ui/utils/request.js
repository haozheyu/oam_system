import {
	get,
	clearWith
} from './local'
import {
	UN_AUTHORIZED,
	UNAUTHORIZED_HEADER_IS_EMPTY,
	INVALID_TOKEN,
	TOKEN_EXPIRED,
	TOKEN_PASS_OUT
} from './statusCode.js'
import log from './log'

import {
	getCurrentPage
} from './getPage'


export const baseUrl = "http://localhost:8085/mars"
//export let baseUrl = "http://47.98.194.238:8085/mars"
// export const baseUrl = "https://apibyte.wkhelpme.com/mars"

// 跳转白名单，即使登录过期，也不跳转
const whiteList = ['pages/home/home']
let showLoading
let isShowToast = false

const request = (url, data, method, header = {}) => {
	const token = get('token')
	if (token) {
		header.token = `${token}`
	}
	return new Promise((resolve, reject) => {
		uni.request({
			url: `${baseUrl}${url}`,
			data,
			header: header,
			method,
			success(res) {
				console.log(res)
				if (res.data.status === 200) {
					resolve(res.data)
				} else if (res.data.status === -1) {
					uni.showToast({
						icon: 'none',
						title: res.data.msg,
						duration: 2000
					})
				} else if (res.data.status === 401) {
					uni.navigateTo({
						url: '/pages/login/login.vue'
					})
				}
			},
			fail(err) {
				console.log(err)
				showToast(err)
				reject(err)
			},
			complete() {
				uni.hideLoading()
			}
		})
	})
}


const request_loading = (url, data, method, header = {}) => {
	const token = get('token')
	if (token) {
		header.token = `${token}`
	}
	uni.showLoading({
		title: "加载中..."
	})

	return new Promise((resolve, reject) => {
		uni.request({
			url: `${baseUrl}${url}`,
			data,
			header: header,
			method,
			success(res) {
				if (res.data.status === 200) {
					resolve(res.data)
				} else if (res.data.status === -1) {
					uni.showToast({
						icon: 'none',
						title: res.data.msg,
						duration: 2000
					})
				} else if (res.data.status === 401) {
					uni.navigateTo({
						url: '/pages/login/login'
					})
				}
			},
			fail(err) {
				console.log(err)
				showToast(err)
				reject(err)
			},
			complete() {

				uni.hideLoading()
			}
		})
	})
}

const showToast = msg => {
	isShowToast = true
	uni.showToast({
		icon: 'none',
		title: msg,
		duration: 2000
	})


}


const getReq = (url, data, header) => {
	return request(url, data, 'GET', header, false)
}

const getLoadingReq = (url, data, header) => {
	return request_loading(url, data, 'GET', header, false)
}

const postReq = (url, data, header) => {
	return request(url, data, 'POST', header, false)
}

const postLoadingReq = (url, data, header) => {
	return request_loading(url, data, 'POST', header, false)
}
const putReq = (url, data, header) => {
	return request(url, data, 'PUT', header)
}
const delReq = (url, data, header) => {
	return request(url, data, 'DELETE', header)
}

export default {
	get: getReq,
	getLoading: getLoadingReq,
	postLoading: postLoadingReq,
	post: postReq,
	put: putReq,
	delete: delReq
}
