import {createStore} from 'vuex'
import http from '../http/index'

const token = localStorage.getItem("token")
console.log(token)

// 请求服务器接口获取用户信息。
// 将用户信息存到vuex中

export default createStore({
    state:{
        userInfo: null,
    },
    mutations:{
        getUserInfo(state,value){
            state.userInfo = value
        }
    },
    actions:{
        // setTokenInfo(context,value){
        //     context.commit("getTokenInfo",value)
        // },
        getUserInfo(context){
            return new Promise((reslove,reject)=>{
                http.post("/user/infoNow",{}).then(res=>{
                    context.commit("getUserInfo",res.data.data)
                    reslove(res)
                }).catch(res=>{reject(res)})
            })
        }
    }
})

