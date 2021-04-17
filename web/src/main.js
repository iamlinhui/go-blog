import Vue from 'vue'
import App from '@/App.vue'
import router from '@/router'
import store from "@/store/store";
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import Editor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'

Vue.use(ElementUI);
Vue.use(Editor);

Vue.config.productionTip = false

new Vue({
    router,
    store,
    render: h => h(App),
}).$mount('#app')
