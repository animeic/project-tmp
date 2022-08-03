<template>
  <Header />
  <Banner />
  <div id="post-container">
    <el-button class="post-submit" type="primary" @click="postHandle">更新</el-button>
    <el-button class="delete-submit" type="danger" @click="deleteHandle">删除</el-button>
    <div class="post-content-warpper">
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
                      'post-submit delete-submit'
                      'post-content-warpper post-content-warpper';
  .delete-submit{
    grid-area: delete-submit;
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
  .post-content-warpper{
    grid-area: post-content-warpper;

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
const text = ref("");

const postForm = ref({
     id: 0,
     html: "",
     content: ""
})

const router = useRouter()

if(localStorage.getItem("role") !== "admin"){
  router.push("/login")
}

const param = router.currentRoute.value.params
param.id = parseInt(param.id)

onBeforeMount(()=>{
  if(localStorage.getItem("role") !== "admin"){
    router.push("/login")
    return
  }
    // 获取用户基本信息
  http.post("/admin/article/editDetail",param).then(res=>{
      if(res.data.code === 0){
          text.value = res.data.data.content
          postForm.id = res.data.data.id
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
  if(localStorage.getItem("role") !== "admin"){
    router.push("/")
  }
  if (!flag) return
  postForm.content = text.value
  postForm.html = dirty
  console.log(postForm)
  http.post("/admin/article/edit",postForm)
    .then(res=>{
      if(res.data.code === 0){
        const id = res.data.data.id
        flag = false
        
        router.push("/article/"+id)
      }
    })

}
const deleteHandle = (res) =>{
  // if (!flag) return
  // postForm.content = text.value
  // postForm.html = dirty
  // console.log(postForm)
  // http.post("/admin/article/delete",postForm)
  //   .then(res=>{
  //     if(res.data.code === 0){
  //       const id = res.data.data.id
  //       flag = false
        
  //       router.push("/article/"+id)
  //       return
  //     }
  //     return
  //   })

  alert("zzzzzz")

}


</script>

