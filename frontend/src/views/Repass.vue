<template>
    <Header></Header>
    <div id="repass-container">
      <div class="repass-form">
        <div class="repass-text"><span>重置密码</span></div>
        <el-form
        ref="repassFrom"
        :model="RepassFrom"
        :rules="rules"
        label-width="80px"
        >
           <el-form-item label="新密码：" prop="password">
              <el-input v-model="RepassFrom.password" type="password" show-password></el-input>
          </el-form-item>
           <!-- <el-form-item label="确认：" prop="repassword">
              <el-input v-model="RepassFrom.repassword" type="repassword" show-password></el-input>
          </el-form-item> -->
          <el-form-item label="邮箱：" prop="email">
              <el-input v-model="RepassFrom.email">
              </el-input>
          </el-form-item>
           <el-form-item label="验证码：" prop="ecode">
              <el-input v-model="RepassFrom.ecode">
                <template #append>
                    <el-button @click="sendCode">发送验证码</el-button>
                </template>
              </el-input>
          </el-form-item>
          <el-form-item class="repass-button">
            <el-button type="success" @click="submitForm('repassFrom')">重置密码</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
    <Footer></Footer>
</template>

<script setup>
import Header from '../components/Header.vue'
import Footer from '../components/Footer.vue'
import http from '@/http/index'
import { ref,reactive,getCurrentInstance } from 'vue'
import { useRouter } from 'vue-router'

const RepassFrom = reactive({
  ecode: '',
  password: '',
  email: '',
})

// prop
const rules = reactive({
  ecode: [
    {required:true,message: '必须填写验证码',trigger: 'blur'},
  ],
  password: [
    {required:true,message: '必须填写密码',trigger: 'blur'},
  ],
  email: [
    {required:true,message: '必须填写邮箱',trigger: 'blur'},
    // {min: 6,max: 6,message: '口令必须6个字符',trigger: 'blur'}
  ],
})
const router = useRouter()
const cxt = getCurrentInstance()

const sendCode = () =>{
 http.post("/sms/repassCode",RepassFrom)
}

const  submitForm = (cf) => {
  cxt.refs[cf].validate((valid) => {
    if (valid) {
      http.post("/auth/repass",RepassFrom).then(res=>{
        if(res.data.code === 0){
          const resdata = res.data.data
          const access_token = resdata.token.token_type + " " + resdata.token.access_token

          localStorage.setItem("token",access_token)
          localStorage.setItem("user_id",resdata.user_info.id)
          router.push("/login")
        }
      })
    }
  })
}

</script>

<style lang="scss" scoped>
// :deep() nth-child(-n+x) !important
#repass-container{
  // debug
  min-height: 900px;
  // grid
  display: grid;
  justify-content: center;
  .repass-form{
    display: grid;
    grid-template-rows: 80px 1fr;
    .repass-text{
      margin-top: 50px;
      margin-bottom: 20px;
      grid-area: 1/1/2/2;
      place-self: center;
    }
    .el-form{
      grid-area: 2/1/3/2;
      :deep(.el-form-item:nth-child(-n+3)){
        display: grid;
        grid-template-columns: 1fr 5fr;
        .el-form-item__label{
          grid-area: 1/1/2/2;
        }
        .el-form-item__content{
          grid-area: 1/2/2/3;
        }
      }
      .repass-button{
        margin-bottom: 5px;
        :deep(.el-form-item__content){
          display: grid;
        }
      }

    }
  }
  

  

}
</style>