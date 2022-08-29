import { ElMessage } from 'element-plus';
import { defineStore } from 'pinia'
import { AddHost, EditHost, HostList } from '../api/ssh';

export const hostsStore = defineStore('hostsStore', {
    state: () => {
        return {
            collapse: false,
            HostItem: {
                "addr": '',
                "user": '',
                "port": '',
                "password": ''
            },
            hostListData: {
                page: 0,
                pageSise: 0,
                dataItems: [],
                total: 0
            },
            editData: {
                "user": "",
                "addr": "",
                "port": "",
                "password": "",
                "isDel": ""
            },
            addData: {
                "user": "",
                "addr": "",
                "port": "",
                "password": ""
            }
        }
    },
    getters: {

    },
    actions: {
        handleCollapse() {
            this.collapse = !this.collapse;
        },
        async handleHosts(data){
            await HostList(data).then((res)=>{
                // console.log(res)
                if (res.code === 0) {
                    this.$patch(state=>{
                        state.hostListData.page = res.data.page,
                        state.hostListData.pageSise = res.data.pageSise,
                        state.hostListData.dataItems = res.data.data,
                        state.hostListData.total = res.data.total
                    })
                }else{
                    ElMessage.warning(res.msg)
                }
            }).catch(err=>{
                console.log(err)
                ElMessage.warning(err)
            })
        },
        async handleEditHost(data){
            await EditHost(data).then((res)=>{
                console.log(res)
                if (res.code === 0) {
                    ElMessage.success(res.msg)
                }else{
                    ElMessage.warning(res.msg)
                }
            }).catch(err=>{
                console.log(err)
                ElMessage.warning(err)
            })
        },
        async handleAddHost(data){
            await AddHost(data).then((res)=>{
                console.log(res)
                if (res.code === 0) {
                    ElMessage.success(res.msg)
                }else{
                    ElMessage.warning(res.msg)
                }
            }).catch(err=>{
                console.log(err)
                ElMessage.warning(err)
            })
        }
    }
})
