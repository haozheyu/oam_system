import { defineStore } from 'pinia'
import { login, userList, updateUser, AddUser } from "../api/user";
import { ElMessage } from "element-plus";
import { setItem, getItem, removeAllItem } from "../utils/storage"
import router from '../router';

export const useUserStore = defineStore('userStore', {
    state: () => {
        return {
            token: getItem("token") || '',
            roleName: getItem("roleName") || '',
            userName: getItem("userName") || '',
            deptName: getItem("deptName") || '',
            accessExpire: getItem("accessExpire") || '',
            refreshAfter: getItem("refreshAfter") || '',
            collapse: false,
            storageUserList: [],
            storageUserTotole: 0
        }
    },
    getters: {

    },
    actions: {
        handleCollapse() {
            this.collapse = !this.collapse;
        },
        async handleLogin(param){
            await login(param).then(data => {
                if (data.code === 0 ) {
                    setItem("token", data.data.token);
                    setItem("roleName", data.data.roleName);
                    setItem("userName", data.data.userName);
                    setItem("deptName", data.data.deptName);
                    setItem("accessExpire", data.data.accessExpire);
                    setItem("refreshAfter", data.data.refreshAfter)
                    ElMessage.success("登录成功");
                    router.push('/')
                }else{
                    ElMessage.warning(data.msg);
                }
            })
        },
        async handleLogout(){
            removeAllItem();
            router.push('/login')
        },
        async handleUserInfo(param){
            // const res = await login(param)
            // console.log(res)
            // todo
        },
        async handleUserList(param){
            await userList(param).then((res)=>{
                if (res.code === 0) {
                    this.$patch(state => {
                        state.storageUserList =  res.data.data;  
                        state.storageUserTotole = res.data.total;
                    })
                }else{
                    this.$patch(state => {
                        state.storageUserList =  [];  
                        state.storageUserTotole = 0;
                    })
                    ElMessage.warning(res.msg)
                }
                }).catch((err)=>{
                    console.log(err)
                    ElMessage.warning("后端接口异常")
            })
        },
        async handleUpdateUser(param1,param2){
            console.log(param1, param2)
            await updateUser(param1).then((res)=>{
                // 更新用户列表
                this.handleUserList(param2)
                ElMessage.success(res.msg)
            }).catch((err)=>{
                ElMessage.warning(err)
            })
        },
        async handleAddUser(param){
            await AddUser(param).then((res)=>{
                // 更新用户列表
                ElMessage.success(res.msg)
            }).catch((err) => {
                ElMessage.warning(err)
            })
        }
    }
})
