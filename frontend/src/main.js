import { createApp } from 'vue'
import App from './App.vue'
import router from '@/router/index'
import ElementPlus from 'element-plus'
import './styles/reset.css'
// import './styles/index.scss'
import 'element-plus/dist/index.css'
import axios from './http/index'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'






const app = createApp(App)
app.config.globalProperties.$axios= axios
app.use(ElementPlus)
app.use(router)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}

app.mount("#app")
