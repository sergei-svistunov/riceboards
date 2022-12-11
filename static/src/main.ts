import {createApp} from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import rbac from './rbac'
import filters from "@/filters/format"
import API from '@/api'

API.customHeaders = (): Promise<Record<string, string>> | undefined => {
    return new Promise<Record<string, string>>(resolve => {
        const token = window.localStorage.getItem('token')
        resolve(token ? {'X-API-Key': token} : {})
    })
}

const app = createApp(App)
app.config.globalProperties.$filters = filters
app.use(store)
    .use(router)
    .use(rbac)
    .mount('#app')