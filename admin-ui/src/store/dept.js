import { ElMessage } from 'element-plus';
import { defineStore } from 'pinia'
import { Adddept, Deptlist, Updatedept, Deldept } from '../api/dept';

export const useDeptStore = defineStore('deptStore', {
    state: () => {
        return {
            collapse: false,
            deptList: [],
            deptTotal: 0,
            deptInfo: {
                "id": 0,
                "name": ''
            }
        }
    },
    getters: {
        
    },
    actions: {
        handleCollapse() {
            this.collapse = !this.collapse;
        },
        async handleDeptList(){
            await Deptlist().then( res=>{
                if (res.code === 0) {
                    this.$patch(state => {
                        state.deptList = res.data.data;
                        state.deptTotal = res.data.total;
                    })
                }
            }).catch(err => {
                console.log(err)
            })
        },
        async handleDelDept(data){
            await Deldept(data).then( res=>{
                if (res.code === 0) {
                    // this.$patch(state => {
                    //     state.deptList = res.data.data;
                    //     state.deptTotal = res.data.total;
                    // })
                    this.handleDeptList();
                    ElMessage.success(res.msg)
                }else{
                    ElMessage.warning(res.msg)
                }
            }).catch(err => {
                console.log(err)
            })
        },
        async handleEditDept(data){
            await Updatedept(data).then( res=>{
                if (res.code === 0) {
                    this.handleDeptList();
                    ElMessage.success(res.msg)
                }else{
                    ElMessage.warning(res.msg)
                }
            }).catch(err => {
                console.log(err)
            })
        },
        async handleAddDept(data){
            await Adddept(data).then( res=>{
                if (res.code === 0) {
                    // this.$patch(state => {
                    //     state.deptList = res.data.data;
                    //     state.deptTotal = res.data.total;
                    // })
                    this.handleDeptList();
                    ElMessage.success(res.msg)
                }else{
                    ElMessage.warning(res.msg)
                }
            }).catch(err => {
                console.log(err)
            })
        }
    }
})