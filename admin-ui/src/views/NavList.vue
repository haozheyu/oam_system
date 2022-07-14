<template>
    <div>
        <div class="crumbs">
            <el-breadcrumb separator="/">
                <el-breadcrumb-item>
                    <i class="el-icon-lx-cascades"></i> 部门列表
                </el-breadcrumb-item>
            </el-breadcrumb>
        </div>
        <div class="container">
            <el-tabs type="border-card" v-model="bindNavModel" :tab-click="getNavTopBodyData(navStore.navTopList, bindNavModel)">
                <el-tab-pane label="所有标签">
                    <el-table :data="navStore.navListData.dataItems" border class="table" ref="multipleTable" header-cell-class-name="table-header">
                        <el-table-column prop="top_id" label="id" width="55" align="center"></el-table-column>
                        <el-table-column prop="top_name" label="导航标题" width="55" align="center"></el-table-column>
                        <el-table-column prop="link_name" label="导航名称" align="center"></el-table-column>
                        <el-table-column prop="link_addr" label="导航地址" align="center"></el-table-column>
                        <el-table-column prop="link_desc" label="描述" align="center"></el-table-column>
                        <el-table-column prop="link_icon" label="icon" align="center"></el-table-column>
                        <el-table-column prop="create_by" label="创建人" align="center"></el-table-column>
                        <el-table-column label="状态" align="center">
                            <template #default="status">
                                <el-tag :type="
                                        status.row.status === '成功'
                                            ? 'success'
                                            : status.row.status === '失败'
                                            ? 'danger'
                                            : ''
                                    ">
                                    <div v-if="status.row.status === 1">
                                    激活
                                    </div>
                                    <div v-else>
                                    未激活
                                    </div>
                                    </el-tag>
                            </template>
                        </el-table-column>

                        <el-table-column label="操作" width="180" align="center">
                            <template #default="scope">
                                <el-button type="text" icon="el-icon-edit" @click="handleEdit(scope.$index, scope.row)">编辑
                                </el-button>
                                <el-button type="text" icon="el-icon-delete" class="red"
                                    @click="handleDelete(scope.$index, scope.row)">删除</el-button>
                            </template>
                        </el-table-column>
                    </el-table>  
                     <div class="pagination">
                        <el-pagination background layout="total, prev, pager, next" :current-page="navStore.navListData.page"
                            :page-size="navStore.navListData.pageSise" :total="navStore.navListData.total" @current-change="handlePageChange1">
                        </el-pagination>
                    </div>                 
                </el-tab-pane>
                
                <el-tab-pane
                v-for="(item,index) in navStore.navTopList"
                :key="index"
                :label="item.name"
                
                >

                    <el-table :data="navStore.navTopBodyData.dataItems" border class="table" ref="multipleTable" header-cell-class-name="table-header">
                        <el-table-column prop="top_id" label="id" width="55" align="center"></el-table-column>
                        <el-table-column prop="top_name" label="导航标题" width="55" align="center"></el-table-column>
                        <el-table-column prop="link_name" label="导航名称" align="center"></el-table-column>
                        <el-table-column prop="link_addr" label="导航地址" align="center"></el-table-column>
                        <el-table-column prop="link_desc" label="描述" align="center"></el-table-column>
                        <el-table-column prop="link_icon" label="icon" align="center"></el-table-column>
                        <el-table-column prop="create_by" label="创建人" align="center"></el-table-column>
                        <el-table-column label="状态" align="center">
                            <template #default="status">
                                <el-tag :type="
                                        status.row.status === '成功'
                                            ? 'success'
                                            : status.row.status === '失败'
                                            ? 'danger'
                                            : ''
                                    ">
                                    <div v-if="status.row.status === 1">
                                    激活
                                    </div>
                                    <div v-else>
                                    未激活
                                    </div>
                                    </el-tag>
                            </template>
                        </el-table-column>

                        <el-table-column label="操作" width="180" align="center">
                            <template #default="scope">
                                <el-button type="text" icon="el-icon-edit" @click="handleEdit(scope.$index, scope.row)">编辑
                                </el-button>
                                <el-button type="text" icon="el-icon-delete" class="red"
                                    @click="handleDelete(scope.$index, scope.row)">删除</el-button>
                            </template>
                        </el-table-column>
                    </el-table>  
                     <div class="pagination">
                        <el-pagination background layout="total, prev, pager, next" :current-page="navStore.navTopBodyData.page"
                            :page-size="navStore.navTopBodyData.pageSise" :total="navStore.navTopBodyData.total" @current-change="handlePageChange">
                        </el-pagination>
                    </div> 

                </el-tab-pane>

            </el-tabs>    

        </div>

        <!-- 编辑弹出框 -->
        <el-dialog title="编辑" v-model="editVisible" width="30%">
            <el-form label-width="70px">
                 <el-form-item label="链接地址">
                    <el-input v-model="form.link_addr" ></el-input>
                </el-form-item>
                <el-form-item label="链接名称">
                    <el-input v-model="form.link_name" disabled></el-input>
                </el-form-item>
                <el-form-item label="链接描述">
                    <el-input v-model="form.link_desc"></el-input>
                </el-form-item>
                <el-form-item label="链接图标">
                    <el-input v-model="form.link_icon"></el-input>
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <el-button @click="editVisible = false">取 消</el-button>
                    <el-button type="primary" @click="saveEdit(form)">确 定</el-button>
                </span>
            </template>
        </el-dialog>

    </div>
</template>

<script>
import { ref, reactive } from "vue";
import { ElMessage } from "element-plus";
import { useNavStore } from "../store/nav";

export default {
    name: "navList",
    setup() {
        const navStore = useNavStore();
         const query = reactive({
            page: 1,
            pageSize: 10,
        });
        navStore.handleNavList(query);
        navStore.handleGetNavTopList();
        const bindNavModel = ref('')

        const form = reactive({
            "link_addr": '',
            "link_name": '',
            "link_desc": '',
            "link_icon": ''
        })

         // 分页导航
        const handlePageChange = (val) => {
            query.page = val;
            navStore.handleNavList(query);
        };
         // 分页导航1
        const handlePageChange1 = (val) => {
            query.page = val;
            navStore.handleNavList(query);
        };

        const getNavTopBodyData = (toplist, index) => {
            // 1234-index 0123-toplist
            let newindex =  parseInt(index) - 1
            let data = {
                "name": ''
            } & query
            toplist.forEach((item,inx)=>{
                if (newindex === inx) {
                    navStore.handleGetNavBody(item);
                }
            })
            
        };

        const formRef = ref(null)

        const addDeptValidateForm = reactive({
            dept_name: '',
        })

        const submitForm = (formEl) => {
            if (!formEl) return
            formEl.validate((valid) => {
                if (valid) {
                    // navStore.$patch(state => {
                    //     state.deptInfo.name = addDeptValidateForm.dept_name
                    // })
                    // deptStore.handleAddDept(deptStore.deptInfo)
                    console.log('submit!',addDeptValidateForm)
                } else {
                    ElMessage.warning('error submit!')
                    return false
                }
            })
        }

        const resetForm = (formEl) => {
            if (!formEl) return
            formEl.resetFields()
        }
       
        // // 表格编辑时弹窗和保存
        const editVisible = ref(false);
    
        const handleEdit = (index, row) => {
            Object.keys(form).forEach((item) => {
                form[item] = row[item];
            });
            editVisible.value = true;
        };
        const submitFormData =  reactive({
            id: '',
            name: ''
        })
        const saveEdit = (form) => {
            editVisible.value = false;          
            navStore.handleUpdateNavTopList(form)
            // 重新加载一次数据
            navStore.handleNavList(query)
        };
        const handleDelete = (index, row)=>{
            editVisible.value = false;
            navStore.handleDelNavTopBody(row)
            // 重新加载一次数据
            navStore.handleNavList(query)
        }

        return {
            query,
            editVisible,
            form,
            submitFormData,
            handleEdit,
            handleDelete,
            handlePageChange,
            saveEdit,
            navStore,
            addDeptValidateForm,
            handlePageChange1,
            getNavTopBodyData,
            formRef,
            submitForm,
            resetForm,
            bindNavModel
        };
    },
};
</script>

<style scoped>
.handle-box {
    margin-bottom: 20px;
}

.handle-select {
    width: 120px;
}

.handle-input {
    width: 300px;
    display: inline-block;
}
.table {
    width: 100%;
    font-size: 14px;
}
.red {
    color: #ff0000;
}
.mr10 {
    margin-right: 10px;
}
.table-td-thumb {
    display: block;
    margin: auto;
    width: 40px;
    height: 40px;
}
</style>
