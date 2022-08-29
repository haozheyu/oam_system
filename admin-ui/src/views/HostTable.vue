<template>
  <div>
    <div class="crumbs">
      <el-breadcrumb separator="/">
        <el-breadcrumb-item>
          <i class="el-icon-lx-cascades"></i> 主机列表
        </el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <div class="container">

      <el-table :data="sshStore.hostListData.dataItems" border class="table" ref="multipleTable" header-cell-class-name="table-header">
        <el-table-column prop="addr" label="地址"></el-table-column>
        <el-table-column prop="port" label="端口"></el-table-column>
        <el-table-column prop="user" label="用户"></el-table-column>
        <el-table-column label="操作" width="180" align="center">
          <template #default="scope">
            <el-button type="text" icon="el-icon-edit" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
            <el-button type="text" icon="el-icon-link" @click="handleWs(scope.$index, scope.row)">链接</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination">
        <el-pagination background layout="total, prev, pager, next" :current-page="query.pageIndex"
                       :page-size="query.pageSize" :total="sshStore.hostListData.total" @current-change="handlePageChange"></el-pagination>
      </div>
    </div>

    <!-- 编辑弹出框 -->
<!--    <el-dialog title="编辑" v-model="editVisible" width="30%">-->
<!--      <el-form label-width="70px">-->
<!--        <el-form-item label="用户名">-->
<!--          <el-input v-model="form.nick_name"></el-input>-->
<!--        </el-form-item>-->
<!--        <el-form-item label="邮箱">-->
<!--          <el-input v-model="form.email"></el-input>-->
<!--        </el-form-item>-->
<!--        <el-form-item label="手机">-->
<!--          <el-input v-model="form.mobile"></el-input>-->
<!--        </el-form-item>-->
<!--        <el-form-item label="部门" prop="deptId">-->
<!--          <el-select v-model="form.deptId" placeholder="选择部门">-->
<!--            <el-option v-for="(item,index) in DeptListResp" :key="index" :label="item.name" :value="item.id"></el-option>-->
<!--          </el-select>-->
<!--        </el-form-item>-->
<!--        <el-form-item label="角色" prop="roleId">-->
<!--          <el-select v-model="form.roleId" placeholder="选择角色">-->
<!--            <el-option v-for="(item,index) in RoleListResp" :key="index" :label="item.name" :value="item.id"></el-option>-->
<!--          </el-select>-->
<!--        </el-form-item>-->
<!--        <el-form-item label="年龄">-->
<!--          <el-input v-model="form.age"></el-input>-->
<!--        </el-form-item>-->
<!--        <el-form-item label="性别">-->
<!--          <el-select v-model="form.sex" placeholder="男">-->
<!--            <el-option label="男" value=1 />-->
<!--            <el-option label="女" value=0 />-->
<!--          </el-select>-->
<!--        </el-form-item>-->
<!--        <el-form-item label="状态">-->
<!--          <el-select v-model="form.status" placeholder="">-->
<!--            <el-option label="激活" value=1 />-->
<!--            <el-option label="删除" value=0 />-->
<!--          </el-select>-->
<!--        </el-form-item>-->
<!--      </el-form>-->
<!--      <template #footer>-->
<!--                <span class="dialog-footer">-->
<!--                    <el-button @click="editVisible = false">取 消</el-button>-->
<!--                    <el-button type="primary" @click="saveEdit(form)">确 定</el-button>-->
<!--                </span>-->
<!--      </template>-->
<!--    </el-dialog>-->
  </div>
</template>

<script>
import { ref, reactive } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";
import { hostsStore } from "../store/hosts";
import { useRouter } from "vue-router";


export default {
  name: "HostList",
  setup() {
    const query = reactive({
      pageIndex: 1,
      pageSize: 10,
    });
    const sshStore = hostsStore();
    const router = useRouter();
    // 获取表格数据
    const getData = () => {
      sshStore.handleHosts(query);
    };
    getData();

    // 查询操作
    const handleSearch = () => {
      query.pageIndex = 1;
      getData();
    };
    // 分页导航
    const handlePageChange = (val) => {
      query.pageIndex = val;
      getData();
    };

    // 删除操作
    const handleDelete = (index) => {
      // 二次确认删除
      ElMessageBox.confirm("确定要删除吗？", "提示", {
        type: "warning",
      })
          .then(() => {
            ElMessage.success("删除成功");
            tableData.value.splice(index, 1);
          })
          .catch(() => {});
    };

    // 表格编辑时弹窗和保存
    const editVisible = ref(false);
    let form = reactive({
      name: "",
      nick_name: "",
      email: "",
      mobile: "",
      deptName: "",
      roleName: "",
      status: 1,
      sex: 0,
      age: 1,
    });
    let idx = -1;
    const handleEdit = (index, row) => {
      idx = index;
      Object.keys(form).forEach((item) => {
        form[item] = row[item];
      });
      editVisible.value = true;
    };
    const handleWs = (index,row) => {
      // console.log(index,row)
      //ws 页面跳转
      sshStore.$patch(state => {
        state.HostItem.addr = row.addr
        state.HostItem.port = row.port
        state.HostItem.user = row.user
        state.HostItem.password = row.password
      })
      router.push("/ws");
    };
    const saveEdit = (from) => {
      editVisible.value = false;
      if (!from.deptId || !from.roleId){
        ElMessage.warning("请核对部门或角色信息");
      }else{
        let user = {
          "name": from.name,
          "email": from.email,
          "mobile": from.mobile,
          "nick_name": from.nick_name,
          "deptId": from.deptId,
          "roleId": from.roleId,
          "status": parseInt(from.status),
          "sex": parseInt(from.sex),
          "age": parseInt(from.age)
        }
        userStore.handleUpdateUser(user,query)
      }
    };

    return {
      query,
      editVisible,
      form,
      handleWs,
      sshStore,
      handleSearch,
      handlePageChange,
      handleDelete,
      handleEdit,
      saveEdit,
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
