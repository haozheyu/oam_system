import {createRouter, createWebHashHistory} from "vue-router";
import { getItem } from "../utils/storage";
import Home from "../views/Home.vue";

const routes = [
    {
        path: '/',
        redirect: '/dashboard'
    }, {
        path: "/",
        name: "Home",
        component: Home,
        children: [
            {
                path: "/dashboard",
                name: "dashboard",
                meta: {
                    title: '系统首页'
                },
                component: () => import ( /* webpackChunkName: "dashboard" */ "../views/Dashboard.vue")
            }, 
            {
                path: "/table",
                name: "basetable",
                meta: {
                    title: '表格'
                },
                component: () => import ( /* webpackChunkName: "table" */ "../views/BaseTable.vue")
            }, 
            {
                path: "/charts",
                name: "basecharts",
                meta: {
                    title: '图表'
                },
                component: () => import ( /* webpackChunkName: "charts" */ "../views/BaseCharts.vue")
            }, 
            {
                path: "/form",
                name: "baseform",
                meta: {
                    title: '表单'
                },
                component: () => import ( /* webpackChunkName: "form" */ "../views/BaseForm.vue")
            }, 
            {
                path: "/tabs",
                name: "tabs",
                meta: {
                    title: 'tab标签'
                },
                component: () => import ( /* webpackChunkName: "tabs" */ "../views/Tabs.vue")
            }, 
            {
                path: "/donate",
                name: "donate",
                meta: {
                    title: '鼓励作者'
                },
                component: () => import ( /* webpackChunkName: "donate" */ "../views/Donate.vue")
            }, 
            {
                path: "/permission",
                name: "permission",
                meta: {
                    title: '权限管理',
                    permission: true
                },
                component: () => import ( /* webpackChunkName: "permission" */ "../views/Permission.vue")
            }, 
            {
                path: "/i18n",
                name: "i18n",
                meta: {
                    title: '国际化语言'
                },
                component: () => import ( /* webpackChunkName: "i18n" */ "../views/I18n.vue")
            }, 
            {
                path: "/upload",
                name: "upload",
                meta: {
                    title: '上传插件'
                },
                component: () => import ( /* webpackChunkName: "upload" */ "../views/Upload.vue")
            }, 
            {
                path: "/icon",
                name: "icon",
                meta: {
                    title: '自定义图标'
                },
                component: () => import ( /* webpackChunkName: "icon" */ "../views/Icon.vue")
            }, 
            {
                path: '/404',
                name: '404',
                meta: {
                    title: '找不到页面'
                },
                component: () => import (/* webpackChunkName: "404" */ '../views/404.vue')
            }, 
            {
                path: '/403',
                name: '403',
                meta: {
                    title: '没有权限'
                },
                component: () => import (/* webpackChunkName: "403" */ '../views/403.vue')
            }, 
            {
                path: '/user',
                name: 'user',
                meta: {
                    title: '个人中心'
                },
                component: () => import (/* webpackChunkName: "user" */ '../views/User.vue')
            }, 
            {
                path: '/editor',
                name: 'editor',
                meta: {
                    title: '富文本编辑器'
                },
                component: () => import (/* webpackChunkName: "editor" */ '../views/Editor.vue')
            }, 
            {
                path: '/markdown',
                name: 'markdown',
                meta: {
                    title: 'markdown编辑器'
                },
                component: () => import (/* webpackChunkName: "markdown" */ '../views/Markdown.vue')
            },
            {
                path: '/create_user',
                name: '创建用户',
                meta: {
                    title: '创建用户'
                },
                component: () => import (/* webpackChunkName: "markdown" */ '../views/CreateUser.vue')
            },
            {
                path: '/user_list',
                name: '用户列表',
                meta: {
                    title: '用户列表'
                },
                component: () => import (/* webpackChunkName: "markdown" */ '../views/UserList.vue')
            },
            {
                path: '/dept_list',
                name: '部门列表',
                meta: {
                    title: '部门列表'
                },
                component: () => import (/* webpackChunkName: "markdown" */ '../views/DeptList.vue')
            },
            {
                path: '/nav_list',
                name: '导航列表',
                meta: {
                    title: '导航列表'
                },
                component: () => import (/* webpackChunkName: "markdown" */ '../views/NavList.vue')
            },
            {
                path: '/nav_add',
                name: '添加导航',
                meta: {
                    title: '添加导航'
                },
                component: () => import (/* webpackChunkName: "markdown" */ '../views/AddNav.vue')
            }
        ]
    }, {
        path: "/login",
        name: "Login",
        meta: {
            title: '登录'
        },
        component: () => import ( /* webpackChunkName: "login" */ "../views/Login.vue")
    }
];

const router = createRouter({
    history: createWebHashHistory(),
    routes
});

// 白名单
const whiteList = ['/login']

router.beforeEach((to, from, next) => {
    document.title = `${to.meta.title} | oam-system`;
    const token = getItem("token");
    const role  = getItem("roleName")
     // 存在 token ，进入主页
    // if (store.state.user.token) {
    // 快捷访问
    if (token && role) {
        if (to.path === '/login') {
        next('/')
        } else {
            if (to.matched.length ===0) {  //如果未匹配到路由
                from.name ? next({ name:from.name }) : next('/');   //如果上级也未匹配到路由则跳转登录页面，如果上级能匹配到则转上级路由
              } else {
                next();    //如果匹配到正确跳转
              }
            // next()
        }
    } else {
        // 没有token的情况下，可以进入白名单
        if (whiteList.indexOf(to.path) > -1) {
        next()
        } else {
        next('/login')
        }
    }

});

export default router;
