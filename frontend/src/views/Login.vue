<template>
    <Header></Header>
    <div id="login-container">
      <div class="login-form">
        <div class="login-text"><span>用户登录</span></div>
        <el-form
        ref="loginFrom"
        :model="LoginFrom"
        :rules="rules"
        label-width="80px"
        >
          <el-form-item label="账号：" prop="account">
              <el-input v-model="LoginFrom.account" placeholder="请输入用户名"></el-input>
          </el-form-item>
          <el-form-item label="密码：" prop="password">
              <el-input v-model="LoginFrom.password" type="password" show-password></el-input>
          </el-form-item>
          <el-form-item class="login-button">
            <el-button type="success" @click="submitForm('loginFrom')">登录</el-button>
          </el-form-item>
          <el-form-item class="register-button">
            <el-button type="danger" @click="toregister" disabled >去注册</el-button>
          </el-form-item>
           <el-row>
           <el-col el-col :span="20"><div class="grid-content bg-purple-dark"></div></el-col>
           <el-col el-col :span="4"><el-link href="/repass" type="primary" style="float: right;font-size: 12px;">忘记密码</el-link></el-col>
          </el-row>
        </el-form>
      </div>
    </div>
    <Footer></Footer>
</template>

<script setup>
import Header from '../components/Header.vue'
import Footer from '../components/Footer.vue'

import http from '@/http/index'
import { reactive,getCurrentInstance } from 'vue'
import { useRouter } from 'vue-router'
import $store from '@/store'
import console from 'console'
import { useStore } from 'vuex'
// 注意：1. 双向绑定必须ref,2. prop属性是验证属性 2.model rules ref
const LoginFrom = reactive({
  account: '',
  password: '',
})

// prop
const rules = reactive({
  account: [
    {required:true,message: '必须填写账号',trigger: 'blur'},
    {min: 6,max: 12,message: '账号长度为6~12个字符',trigger: 'blur'}
  ],
  password: [
    {required:true,message: '必须填写密码',trigger: 'blur'},
    {min: 6,max: 12,message: '密码长度为6~12个字符',trigger: 'blur'}
  ],
})
const router = useRouter()
const cxt = getCurrentInstance()
// const useState = useStore()

const toregister = () => {
  router.push('/register')
}


const  submitForm = (cf) => {
  // 表单验证，成功提交数据
  cxt.refs[cf].validate((valid) => {
      if (valid) {
        http.post("/auth/login",LoginFrom).then(res=>{
          if(res.data.code === 0){
            const resdata = res.data.data
            const access_token = resdata.token.token_type + " " + resdata.token.access_token
            localStorage.setItem("token",access_token)
            localStorage.setItem("role","admin")
            router.push("/")
          }
        })

      } else {
      //  $store.dispatch("getToken")
      // console.log($store.state.tokenInfo)
        return false
      }
    })
}

</script>

<style lang="scss" scoped>
// :deep() nth-child(-n+x) !important
#login-container{
    // debug
    min-height: 900px;
  // grid
  display: grid;
  justify-content: center;
  .login-form{
    display: grid;
    grid-template-rows: 80px 1fr;
    .login-text{
      margin-top: 50px;
      margin-bottom: 20px;
      grid-area: 1/1/2/2;
      place-self: center;

    }
    .el-form{
      grid-area: 2/1/3/2;
      :deep(.el-form-item:nth-child(-n+2)){
        display: grid;
        grid-template-columns: 1fr 5fr;
        .el-form-item__label{
          grid-area: 1/1/2/2;
        }
        .el-form-item__content{
          grid-area: 1/2/2/3;
        }
      }
      .login-button{
        margin-bottom: 5px;
        :deep(.el-form-item__content){
          display: grid;
        }
      }
      .register-button{
        margin-bottom: 10px;
        :deep(.el-form-item__content){
          display: grid;
          
        }
      }

    }
  }
  

  

}
</style>

