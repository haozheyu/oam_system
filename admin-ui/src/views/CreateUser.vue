<template>
    <div>
        <div class="crumbs">
            <el-breadcrumb separator="/">
                <el-breadcrumb-item>
                    <i class="el-icon-lx-calendar"></i> 表单
                </el-breadcrumb-item>
                <el-breadcrumb-item>创建用户</el-breadcrumb-item>
            </el-breadcrumb>
        </div>
        <div class="container">
            <div class="form-box">
                <el-form ref="formRef" :rules="rules" :model="form" label-width="80px">
                    <el-form-item label="登录用户" prop="name">
                        <el-input v-model="form.name"></el-input>
                    </el-form-item>
                    <el-form-item label="显示用户" prop="nickName">
                        <el-input v-model="form.nickName"></el-input>
                    </el-form-item>
                    <el-form-item label="邮箱" prop="email">
                        <el-input v-model="form.email"></el-input>
                    </el-form-item>
                    <el-form-item label="电话" prop="mobile">
                        <el-input v-model="form.mobile"></el-input>
                    </el-form-item>
                    <el-form-item label="密码" prop="password">
                        <el-input v-model="form.password"></el-input>
                    </el-form-item>
                    <el-form-item label="年龄" prop="age">
                        <el-input v-model="form.age"></el-input>
                    </el-form-item>

                    <el-form-item label="性别" prop="sex">
                        <el-select v-model="form.sex" placeholder="男">
                            <el-option label="男" value=1 />
                            <el-option label="女" value=0 />
                        </el-select>
                    </el-form-item>  
                    <el-form-item label="部门" prop="deptId">
                    <el-select v-model="form.deptId" placeholder="选择部门">
                        <el-option v-for="(item,index) in DeptListResp" :key="index" :label="item.name" :value="item.id"></el-option>
                    </el-select>
                    </el-form-item>
                    <el-form-item label="角色" prop="roleId">
                        <el-select v-model="form.roleId" placeholder="选择角色">
                            <el-option v-for="(item,index) in RoleListResp" :key="index" :label="item.name" :value="item.id"></el-option>
                        </el-select>
                    </el-form-item>

                    <el-form-item>
                        <el-button type="primary" @click="onSubmit">表单提交</el-button>
                        <el-button @click="onReset">重置表单</el-button>
                    </el-form-item>
                </el-form>
            </div>
        </div>
    </div>
</template>

<script>
import { reactive, ref } from "vue";
import { ElMessage } from "element-plus";
import { getItem } from "../utils/storage";
import { setTimeStamp } from "../utils/time";
import { Rolelist } from "../api/role";
import { Deptlist } from "../api/dept";
import { useUserStore } from "../store/user";
export default {
    name: "createUser",
    setup() {
        const userStore = useUserStore();
        const DeptListResp = ref([]);
        const RoleListResp = ref([]);
        const rules = {
            name: [
                { required: true, message: "请输入登录用户名", trigger: "blur" },
            ],
            email: [
                { required: true, message: "请输入邮箱", trigger: "blur" },
            ],
            password: [
                { required: true, message: "请输入密码", trigger: "blur" },
            ],
            age: [
                { required: true, message: "请输入年龄", trigger: "blur" },
            ],
            sex: [
                { required: true, message: "请输入性别", trigger: "blur" },
            ],
            nickName: [
                { required: true, message: "请输入名称", trigger: "blur" },
            ],
            mobile: [
                { required: true, message: "请输入电话号", trigger: "blur" },
            ],
        };
        const formRef = ref(null);
        const form = reactive({
            email: "",
            name: "",
            deptId: 1,
            roleId: 8,
            password: "",
            lastUpdateBy: getItem("userName"),
            lastUpdateTime: setTimeStamp(),
            age: 0,
            sex: 1,
            nickName: "",
            mobile: "",
            avatar: "http://dummyimage.com/100x100"
        });
        // 提交
        const onSubmit = () => {
            // 表单校验
            formRef.value.validate((valid) => {
                if (valid) {
                    console.log(form);
                     let user = {
                        age: parseInt(form.age),
                        avatar: form.avatar,
                        deptId: parseInt(form.deptId),
                        email: form.email,
                        lastUpdateBy: form.lastUpdateBy,
                        lastUpdateTime: parseInt(form.lastUpdateTime),
                        mobile: form.mobile,
                        name: form.name,
                        nickName: form.nickName,
                        password: form.password,
                        roleId: parseInt(form.roleId),
                        sex: parseInt(form.sex),
                    }
                    userStore.handleAddUser(user)
                } else {
                    return false;
                }
            });
        };
        // 重置
        const onReset = () => {
            formRef.value.resetFields();
        };
        // 获取部门列表
        const getDeptList = () => {
            Deptlist().then((res)=>{
                if (res.code !== 0) {
                    ElMessage.warning(res.msg)
                }else{
                    DeptListResp.value = res.data.data
                }
            }).catch((err)=>{
                console.log(err)
                ElMessage.warning("后端接口异常")
            })
        };
        getDeptList();
        // 获取角色列表
        const getRoleList = () => {
             Rolelist().then((res)=>{
                if (res.code !== 0) {
                    ElMessage.warning(res.msg)
                }else{
                    RoleListResp.value = res.data.data
                }
            }).catch((err)=>{
                console.log(err)
                ElMessage.warning("后端接口异常")
            })
        };
        getRoleList();
        

        return {
            DeptListResp,
            RoleListResp,
            rules,
            formRef,
            form,
            onSubmit,
            onReset,
            userStore
        };
    },
};
</script>