<template>
    <div>
        <div class="crumbs">
            <el-breadcrumb separator="/">
                <el-breadcrumb-item>
                    <i class="el-icon-lx-calendar"></i> 表单
                </el-breadcrumb-item>
                <el-breadcrumb-item>添加导航</el-breadcrumb-item>
            </el-breadcrumb>
        </div>
        <div class="container">
            <div class="form-box">
  <el-row :gutter="24">
    <el-col :span="12">
      <el-card shadow="hover"> 
            <el-form ref="formRef" :rules="rules" :model="form" label-width="80px">
                    <el-form-item label="添加标题" prop="name">
                        <el-input v-model="form.name"></el-input>
                    </el-form-item>
                
                    <el-form-item>
                        <el-button type="primary" @click="onSubmit">提交</el-button>
                        <el-button @click="onReset">重置</el-button>
                    </el-form-item>
            </el-form>
      </el-card>
    </el-col>
    <el-col :span="12">
      <el-card shadow="hover"> 
            <el-form ref="formRef1" :rules="rules1" :model="form" label-width="80px">
                    <el-form-item label="链接地址" prop="link_addr">
                        <el-input v-model="form.link_addr"></el-input>
                    </el-form-item>
                    <el-form-item label="链接名称" prop="link_name">
                        <el-input v-model="form.link_name"></el-input>
                    </el-form-item>
                    <el-form-item label="链接描述" prop="link_desc">
                        <el-input v-model="form.link_desc"></el-input>
                    </el-form-item>
                    <el-form-item label="链接图标" prop="link_icon">
                        <el-input v-model="form.link_icon"></el-input>
                    </el-form-item>
                    <el-form-item label="标题名称" prop="top_name_id">
                        <el-select v-model="form.top_name_id" placeholder="选择标题名称">
                            <el-option v-for="(item,index) in navStore.navTopList" :key="index" :label="item.name" :value="item.id"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" @click="onSubmit1">提交</el-button>
                        <el-button @click="onReset1">重置</el-button>
                    </el-form-item>
            </el-form>
      </el-card>
    </el-col>
  </el-row>
        
            </div>
        </div>
    </div>
</template>

<script>
import { reactive, ref } from "vue";
import { useNavStore } from "../store/nav";

export default {
    name: "addNav",
    setup() {
        const navStore = useNavStore();
        navStore.handleGetNavTopList();
        const rules = {
            name: [
                { required: true, message: "请输入标题名称", trigger: "blur" },
            ]
        };
        const rules1 = {
            link_addr: [
                { required: true, message: "请输入标题内容的地址信息", trigger: "blur" }
            ],
            link_name: [
                { required: true, message: "请输入标题内容的名称信息", trigger: "blur" }
            ],
            link_desc: [
                { required: true, message: "请输入标题内容的详细描述", trigger: "blur" }
            ],
            link_icon: [
                { required: true, message: "请输入标题内容的地址信息", trigger: "blur" }
            ]
        };
        const formRef = ref(null);
        const formRef1 = ref(null);
        const form = reactive({
            name: "",
            top_name_id: "",
            top_name: '',
            link_addr: "",
            link_name: "",
            link_desc: "",
            link_icon: "" 
        });
        // 提交
        const onSubmit = () => {
            // 表单校验
            formRef.value.validate((valid) => {
                if (valid) {
                    console.log(form);
                    navStore.handleAddNavTop(form)
                    // 更新navtoplist
                    navStore.handleGetNavTopList()
                } else {
                    return false;
                }
            });
        };
        // 提交
        const onSubmit1 = () => {
            // 表单校验
            formRef1.value.validate((valid) => {
                if (valid) {
                    console.log(form);
                    navStore.handleAddNavTopBody(form)
                } else {
                    return false;
                }
            });
        };
        // 重置
        const onReset = () => {
            formRef.value.resetFields();
        };
        // 重置
        const onReset1 = () => {
            formRef1.value.resetFields();
        };
        
        return {
            rules,
            rules1,
            formRef,
            formRef1,
            form,
            onSubmit,
            onReset,
            onSubmit1,
            onReset1,
            navStore
        };
    },
};
</script>