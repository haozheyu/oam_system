import request from '../utils/request'
import { setTimeStamp } from '../utils/time'

/**
 * 登录
 */
export const login = data => {
    return request({
        url: '/api/user/login',
        method: 'POST',
        data: {
            "userName": data.username,
            "password": data.password
        }
    })
}

/**
 * 获取用户列表
 */
 export const userList = data => {
    return request({
        url: '/api/user/list',
        method: 'POST',
        data: {
            "current": data.pageIndex,
            "pageSize": data.pageSize
        }
    })
}

/**
 * 用户更新
 * 		Email:          req.Email,
		Mobile:         req.Mobile,
		Status:         req.Status,
		DeptId:         req.DeptId,
        RoleId:         req.RoleId,
		Sex:            req.Sex,
		Age:            req.Age,
        NickName:       req.NickName,
 */
 export const updateUser = data => {
    return request({
        url: '/api/user/update',
        method: 'POST',
        data: {
            "name": data.name,
            "email": data.email,
            "mobile": data.mobile,
            "nickName": data.nick_name,
            "deptId": data.deptId,
            "roleId": data.roleId,
            "status": data.status,
            "sex": data.sex,
            "age": data.age
        }
    })
}

/**
 * 新增用户
 */
 export const AddUser = data => {
    return request({
        url: '/api/user/add',
        method: 'POST',
        data: {
            "email": data.email,
            "name": data.name,
            "deptId": data.deptId,
            "roleId": data.roleId,
            "password": data.password,
            "lastUpdateBy": data.lastUpdateBy,
            "lastUpdateTime": setTimeStamp(),
            "avatar": "http://dummyimage.com/100x100",
            "mobile": data.mobile,
            "nickName": data.nickName
        }
    })
}

/**
 * 删除用户
 */
 export const DelUser = data => {
    return request({
        url: '/api/user/delete',
        method: 'POST',
        data: {
            "userName": data.userName
        }
    })
}

/**
 * 重置用户密码
 */
 export const SetUserPassword = data => {
    return request({
        url: '/api/user/reSetPassword',
        method: 'POST',
        data: {
            "userName": data.userName,
            "password": data.password
        }
    })
}



/**
 * 获取用户信息
 */
export const getUserInfo = data => {
    return request({
        url: '/api/user/userInfo',
        method: 'POST',
        data: {
            "userName": data.userName
        }
    })
}
