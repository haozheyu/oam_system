<template>
    <div class="sidebar">
        <el-menu class="sidebar-el-menu" :default-active="onRoutes" :collapse="sidebar.collapse" background-color="#324157"
            text-color="#bfcbd9" active-text-color="#20a0ff" unique-opened router>
            <template v-for="item in items">
                <template v-if="item.subs">
                    <el-submenu :index="item.index" :key="item.index">
                        <template #title>
                            <i :class="item.icon"></i>
                            <span>{{ item.title }}</span>
                        </template>
                        <template v-for="subItem in item.subs">
                            <el-submenu v-if="subItem.subs" :index="subItem.index" :key="subItem.index">
                                <template #title>{{ subItem.title }}</template>
                                <el-menu-item v-for="(threeItem, i) in subItem.subs" :key="i" :index="threeItem.index">
                                    {{ threeItem.title }}</el-menu-item>
                            </el-submenu>
                            <el-menu-item v-else :index="subItem.index" :key="subItem.index">{{ subItem.title }}
                            </el-menu-item>
                        </template>
                    </el-submenu>
                </template>
                <template v-else>
                    <el-menu-item :index="item.index" :key="item.index">
                        <i :class="item.icon"></i>
                        <template #title>{{ item.title }}</template>
                    </el-menu-item>
                </template>
            </template>
        </el-menu>
    </div>
</template>

<script>
import { computed } from "vue";
import { useSidebarStore } from '../store/sidebar'
import { useRoute } from "vue-router";
import { getItem } from "../utils/storage";
export default {
    setup() {
        const roleName = getItem("roleName") || ''
        const items = [
            {
                icon: "el-icon-lx-home",
                index: "/dashboard",
                title: "系统首页",
            },
            // {
            //     icon: "el-icon-lx-cascades",
            //     index: "/table",
            //     title: "基础表格",
            // },
            // {
            //     icon: "el-icon-lx-copy",
            //     index: "/tabs",
            //     title: "tab选项卡",
            // },
            // {
            //     icon: "el-icon-lx-calendar",
            //     index: "3",
            //     title: "表单相关",
            //     subs: [
            //         {
            //             index: "/form",
            //             title: "基本表单",
            //         },
            //         {
            //             index: "/upload",
            //             title: "文件上传",
            //         },
            //         {
            //             index: "4",
            //             title: "三级菜单",
            //             subs: [
            //                 {
            //                     index: "/editor",
            //                     title: "富文本编辑器",
            //                 },
            //                 {
            //                     index: "/markdown",
            //                     title: "markdown编辑器",
            //                 },
            //             ],
            //         },
            //     ],
            // },
            // {
            //     icon: "el-icon-lx-emoji",
            //     index: "/icon",
            //     title: "自定义图标",
            // },
            // {
            //     icon: "el-icon-pie-chart",
            //     index: "/charts",
            //     title: "schart图表",
            // },
            // {
            //     icon: "el-icon-lx-global",
            //     index: "/i18n",
            //     title: "国际化功能",
            // },
            // {
            //     icon: "el-icon-lx-warn",
            //     index: "7",
            //     title: "错误处理",
            //     subs: [
            //         {
            //             index: "/permission",
            //             title: "权限测试",
            //         },
            //         {
            //             index: "/404",
            //             title: "404页面",
            //         },
            //     ],
            // },
            // {
            //     icon: "el-icon-lx-redpacket_fill",
            //     index: "/donate",
            //     title: "支持作者",
            // },
        ];
        switch(roleName)
        {
                case "超级管理员":
                    items.push(
                        {
                            icon: "el-icon-user",
                            index: "/user",
                            title: "用户管理",
                            subs: [
                                {
                                    index: "/user_list",
                                    title: "用户列表",
                                },
                                {
                                    index: "/create_user",
                                    title: "创建用户",
                                },
                                {
                                    index: "/dept_list",
                                    title: "部门管理",
                                }
                                ]
                        },
                        {
                            icon: "el-icon-chicken",
                            index: "/nav",
                            title: "导航管理",
                            subs: [
                                {
                                    index: "/nav_list",
                                    title: "导航列表",
                                },
                                {
                                    index: "/nav_add",
                                    title: "添加导航",
                                }
                            ]
                        },
                        {
                          icon: "el-icon-edit",
                          index: "/host",
                          title: "主机管理",
                          subs: [
                            {
                              index: "/hosts",
                              title: "主机列表",
                            },
                            {
                              index: "/host_add",
                              title: "添加主机",
                            },
                            {
                              index: "/host_edit",
                              title: "编辑主机",
                            }
                          ]
                        }
                    )
                    break;
                case "项目负责人":

                break;
                case "产品经理":

                break;
                case "UI设计师":

                break;
                case "研发工程师":

                break;
                case "测试工程师":

                break;
                case "运营工程师":

                break;
                case "新用户":

                break;
                case "自定义角色":

                break;
                default:

        }
        const route = useRoute();
        const onRoutes = computed(() => {
            return route.path;
        });

        const sidebar = useSidebarStore();

        return {
            items,
            roleName,
            onRoutes,
            sidebar,
        };
    },
};
</script>

<style scoped>
.sidebar {
    display: block;
    position: absolute;
    left: 0;
    top: 70px;
    bottom: 0;
    overflow-y: scroll;
}
.sidebar::-webkit-scrollbar {
    width: 0;
}
.sidebar-el-menu:not(.el-menu--collapse) {
    width: 250px;
}
.sidebar > ul {
    height: 100%;
}
</style>
