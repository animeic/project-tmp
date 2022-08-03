<template>
  <Header />
  <Banner />
  <div id="post-container">
    <div class="top-info">
      <router-link :to="{name:'user-profile',params:{id:editor.id}}" >
        <img :src="editor.avatar" :alt="editor.nickname">
      </router-link>
    </div>
    <el-button class="post-submit" type="primary" @click="postHandle">提交</el-button>
    <div class="post-content-wrapper">
      <md-editor v-model="text"
        :theme="theme"
        :toolbars="toolbars"
        @onHtmlChanged="htmlHandle"
        @onUploadImg="onUploadImg"
      />
    </div>
  </div>
  <Footer />
</template>

<style lang="scss" scoped>

#post-container{
  display: grid;
  grid-template-areas: 
                      'top-info post-submit'
                      'post-content-wrapper post-content-wrapper';
  .top-info{
    grid-area: top-info;
    height: 60px;
    display: grid;
    place-items: center;
   
    img{
      width: 40px;
      height: 40px;
      border-radius: 20px;
      border: 1px solid plum;
    }
  }
  .post-submit{
    grid-area: post-submit;
    height: 60px;
  }
  .post-content-wrapper{
    grid-area: post-content-wrapper;

    #md-editor-v3{
      height: 680px;
    }

  }
}

</style>

<script setup>
import { onMounted, onBeforeMount,ref,reactive } from 'vue';
import { useRouter } from 'vue-router';
import MdEditor from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import Header from '@/components/Header.vue';
import Footer from '@/components/Footer.vue';
import Banner from '@/components/Banner.vue';
import axios from 'axios';
import $store from '@/store'
import sanitizeHtml from 'sanitize-html';
import http from '@/http/index'
 // mardown 编辑器配置
 // https://imzbf.github.io/md-editor-v3/docs
const toolbars = reactive([
  'bold',
  'underline',
  'italic',
//   '-',
  'strikeThrough',
  'title',
  'sub',
  'sup',
  'quote',
  'unorderedList',
  'orderedList',
//   '-',
  'codeRow',
  'code',
  'link',
  'image',
  'table',
  'mermaid',
  'katex',
  '-',
  'revoke',
  'next',
  'catalog',
  // 'save',

  // '=',
  '-',
  '-',
  'pageFullscreen',
  'fullscreen',
  'preview',
//   'htmlPreview',
   
//   'github'
])
// 格式模板
const text = ref(`# 这里是标题,需要标识\`# \`
> 这里是摘要，需要标识\`> \`

这里是正文...`);

const router = useRouter()
// 如果不是admin不展示入口header
// 进入时判断是否是admin，否则跳转到登录
// 暂时禁用注册功能，仅供自己使用
if (localStorage.getItem("role") !== "admin"){
  router.push("/login")
}

const theme = ref('');
const postForm = reactive({
  content: "",
  html: ""
})
// 编辑人信息
 const editor = reactive({
  id: 0,
  nickname: 'efe45',
  avatar: 'https://av.animeii.tech/images/5e295ba090ad.jpeg'
})


// $store.dispatch("getUserInfo")
// console.log($store.state.userInfo)

// onBeforeMount(()=>{

// })
onBeforeMount(()=>{
  if(localStorage.getItem("role") !== "admin"){
    router.push("/login")
    return
  }
    // 获取用户基本信息
  http.post("/user/infoNow",{}).then(res=>{
      if(res.data.code === 0){
        editor.id = res.data.data.id
        editor.avatar = res.data.data.avatar
        editor.nickname = res.data.data.nickname
      }
  })
      
})
// editor = $store.state.userInfo


// 处理图片上传：调用接口、生成链接地址
const onUploadImg = async (files, callback) => {
  const res = await Promise.all(
    files.map((file) => {
      return new Promise((rev, rej) => {
        const form = new FormData();
        form.append('image', file);
        axios
          .post('https://api.animeii.tech/upload/image', form, {
            headers: {
              'Content-Type': 'multipart/form-data',
            }
          })
          .then((res) => rev(res))
          .catch((error) => rej(error));
      });
    })
  );

  callback(res.map((item) => item.data.data));
};



let dirty = ""
const htmlHandle = (html)=>{
  // dirty = sanitizeHtml(html,{
  //       allowedTags: [ 'img', 'p' ],
  //       allowedSchemes: [ 'data', 'http','https']
  // })
  dirty = html
}

let flag = true

const postHandle = (res) =>{
  // 角色不是admin时 跳转到登录
  if(localStorage.getItem("role") !== "admin"){
    router.push("/")
  }
  if (!flag) return
  postForm.content = text.value
  postForm.html = dirty
  http.post("/admin/article/post",postForm)
    .then(res=>{
      if(res.data.code === 0){
        const id = res.data.data.id
        flag = false
        router.push("/article/"+id)
      }
    })

}


</script>

