import {createApp} from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import rbac from './rbac'
import filters from "@/filters/format";

const app = createApp(App)
app.config.globalProperties.$filters = filters
app.use(store)
    .use(router)
    .use(rbac)
    .mount('#app')