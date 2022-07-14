import { ElMessage } from 'element-plus';
import { defineStore } from 'pinia'
import { Navlist, GetNavTopList, AddNavTopBody, AddNavTop, DelNavTopBody, DelNavTop, GetNavBody, UpdateNavTopList } from '../api/navop';

export const useNavStore = defineStore('navStore', {
    state: () => {
        return {
            collapse: false,
            NavBodyData: {
                "top_name_id": '',
                "link_addr": '',
                "link_name": '',
                "link_desc": '',
                "link_icon": ''
            },
            navListData: {
                page: 0,
                pageSise: 0,
                dataItems: [],
                total: 0
            },
            navTopList:[],
            navTopBodyData: {
                page: 0,
                pageSise: 0,
                dataItems: [],
                total: 0
            },
            NavTopData: {
                "name": ''
            },
            updateNavData: {
                "link_addr": '',
                "link_name": '',
                "link_desc": '',
                "link_icon": ''
            }
        }
    },
    getters: {
        
    },
    actions: {
        handleCollapse() {
            this.collapse = !this.collapse;
        },
        async handleNavList(data){
            await Navlist(data).then((res)=>{
                // console.log(res)
                if (res.code === 0) {
                    this.$patch(state=>{
                        state.navListData.page = res.data.page,
                        state.navListData.pageSise = res.data.pageSise,
                        state.navListData.dataItems = res.data.data,
                        state.navListData.total = res.data.total
                    })
                }else{
                    ElMessage.warning(res.msg)
                }
            }).catch(err=>{
                console.log(err)
                ElMessage.warning(err)
            })
        },
        async handleAddNavTopBody(data){
            await AddNavTopBody(data).then((res)=>{
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
        async handleAddNavTop(data){
            await AddNavTop(data).then((res)=>{
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
        async handleDelNavTopBody(data){
            await DelNavTopBody(data).then((res)=>{
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
        async handleDelNavTop(data){
            await DelNavTop(data).then((res)=>{
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
        async handleGetNavBody(data){
            await GetNavBody(data).then((res)=>{
                // console.log(res)
                if (res.code === 0) {
                    this.$patch(state=>{
                        state.navTopBodyData.page = res.data.page,
                        state.navTopBodyData.pageSise = res.data.pageSise,
                        state.navTopBodyData.dataItems = res.data.data,
                        state.navTopBodyData.total = res.data.total
                    })
                    
                }else{
                    ElMessage.warning(res.msg)
                }
            }).catch(err=>{
                console.log(err)
                ElMessage.warning(err)
            })
        },
        async handleGetNavTopList(){
            await GetNavTopList().then((res)=>{
                // console.log(res)
                if (res.code === 0) {
                    this.$patch(state=>{
                        state.navTopList = res.data.data
                    })
                    
                }else{
                    ElMessage.warning(res.msg)
                }
            }).catch(err=>{
                console.log(err)
                ElMessage.warning(err)
            })
        },
        async handleUpdateNavTopList(data){
            await UpdateNavTopList(data).then((res)=>{
                // console.log(res)
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