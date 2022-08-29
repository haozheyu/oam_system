import Vue from 'vue'
import App from './App'

import * as Api from './config/api.js'
import * as Common from './config/common.js'
import * as Db from './config/db.js'

Vue.config.productionTip = false

Vue.prototype.$api = Api
Vue.prototype.$common = Common
Vue.prototype.$db = Db

// 挂载全局状态管理
import store from './store/index.js'
Vue.prototype.$store = store

App.mpType = 'app'

import uView from 'uview-ui';
Vue.use(uView);

const app = new Vue({
    ...App
})
app.$mount()
