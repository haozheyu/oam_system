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
            <el-table :data="deptStore.deptList" border class="table" ref="multipleTable" header-cell-class-name="table-header">
                <el-table-column prop="id" label="ID" width="55" align="center"></el-table-column>
                <el-table-column prop="name" label="name" align="center"></el-table-column>
                <el-table-column prop="createBy" label="创建人" align="center"></el-table-column>
                <el-table-column prop="lastUpdateBy" label="最后更新人" align="center"></el-table-column>
                <el-table-column label="状态" align="center">
                    <template #default="delFlag">
                        <el-tag :type="
                                delFlag.row.delFlag === '成功'
                                    ? 'success'
                                    : delFlag.row.delFlag === '失败'
                                    ? 'danger'
                                    : ''
                            ">
                            <div v-if="delFlag.row.delFlag === 1">
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
            
        </div>

        <!-- 编辑弹出框 -->
        <el-dialog title="编辑" v-model="editVisible" width="30%">
            <el-form label-width="70px">
                 <el-form-item label="ID" disabled>
                    <el-input v-model="form.id" disabled></el-input>
                </el-form-item>
                <el-form-item label="用户名">
                    <el-input v-model="form.name"></el-input>
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <el-button @click="editVisible = false">取 消</el-button>
                    <el-button type="primary" @click="saveEdit(form)">确 定</el-button>
                </span>
            </template>
        </el-dialog>



         <el-form
                ref="formRef"
                :model="addDeptValidateForm"
                label-width="100px"
                class="demo-ruleForm"
            >
                <el-form-item
                label="添加部门"
                prop="dept_name"
                :rules="[
                    { required: true, message: 'dept_name is required' },
                ]"
                >
                <el-input
                    v-model.number="addDeptValidateForm.dept_name"
                    type="text"
                    autocomplete="off"
                />
                </el-form-item>
                <el-form-item>
                <el-button type="primary" @click="submitForm(formRef)">Submit</el-button>
                <el-button @click="resetForm(formRef)">Reset</el-button>
                </el-form-item>
            </el-form>
    </div>
</template>

<script>
import { ref, reactive } from "vue";
import { ElMessage } from "element-plus";
import { useDeptStore } from "../store/dept";

export default {
    name: "deptList",
    setup() {
        const deptStore = useDeptStore();
        deptStore.handleDeptList();
        const form = reactive({
            id: '',
            name: ''
        })

        const formRef = ref(null)

        const addDeptValidateForm = reactive({
            dept_name: '',
        })

        const submitForm = (formEl) => {
            if (!formEl) return
            formEl.validate((valid) => {
                if (valid) {
                    deptStore.$patch(state => {
                        state.deptInfo.name = addDeptValidateForm.dept_name
                    })
                    deptStore.handleAddDept(deptStore.deptInfo)
                    // console.log('submit!',addDeptValidateForm)
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
            deptStore.$patch(state => {
                state.deptInfo.id = form.id
                state.deptInfo.name = form.name
            })
            deptStore.handleEditDept(deptStore.deptInfo)
        };
        const handleDelete = (index, row)=>{
            editVisible.value = false;
            Object.keys(form).forEach((item) => {
                form[item] = row[item];
            });
            // console.log(row.id, row.name, "------")
            deptStore.$patch(state => {
                state.deptInfo.id = row.id
                state.deptInfo.name = row.name
            })
            deptStore.handleDelDept(deptStore.deptInfo)
        }

        return {
            editVisible,
            form,
            submitFormData,
            handleEdit,
            handleDelete,
            saveEdit,
            deptStore,
            addDeptValidateForm,
            formRef,
            submitForm,
            resetForm
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
