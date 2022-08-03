import axios from 'axios'
// import { useRouter } from 'vue-router'
import router from '../router/index'
import { ElMessage,ElLoading } from 'element-plus'

let loading
const startLoading = () => {
    const options = {
        lock:true,
        text: "加载中。。。",
        background: 'rgba(0,0,0,0.7)'
    }
    loading = ElLoading.service(options)
}
const endLoading = () => {
    loading.close()
}
// axios.defaults.baseURL = 'http://127.0.0.1:8088';
axios.defaults.baseURL = 'https://api.animeii.tech';
axios.defaults.headers.common["Content-Type"] = "application/json;charset=utf-8"
axios.interceptors.request.use((config) => {
    startLoading()
    if(config.method === 'get'){
        config.data = {unused:0}
    }
    // config.headers["Host"] = "api.animeii.tech"
    // console.log(config.params)
    // 每次请求传入保存在本地的token
    // 登录时或者刷新token时设置到本地
    config.headers["Authorization"] = localStorage.getItem("token")
    
  
    return config
},
error =>{
    endLoading()
    Promise.reject(error)
})


axios.interceptors.response.use((reponse) => {
    endLoading()
    if(reponse.status === 200){
        if(reponse.data.code !== 0){
            ElMessage.error(reponse.data.message)
        }else{
            ElMessage.success(reponse.data.message)
        }
    }

    // 令牌无效
    if(reponse.data.code === 2002){
        router.push("/login")
        return

    }
    // 认证成功
    
    // 如果有新token设置token到localstorage中
    if(reponse.headers['new-token']){
        localStorage.setItem("token",reponse.headers['new-token'])
    }
    return reponse
},error => {
    endLoading()
    Promise.reject(error)

})


export default axios