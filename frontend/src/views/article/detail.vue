<template>  
    <Header />
    <Banner />
    <div id="article-detail-container">
      <md-editor v-model="detail.html"
       :previewOnly="previewOnly"
       :previewTheme="previewTheme"
       :showCodeRowNumber="showCodeRowNumber"
      />
      <div class="right-side">
        <SideUser :total="sideUser.total" :user-info="sideUser.userInfo" />
        <div class="info">
<!--          <span>浏览：33</span>-->
<!--          <span>点赞：22</span>-->
<!--          <span>评论：44</span>-->
        </div>
        <Comment />
      </div>
    </div>
    <Footer />
    
</template>


<style lang="scss" scoped>

#article-detail-container{
  background: #f5f7f9;
  margin-top: 10px;
  display: grid;
  grid-template-columns: 3fr 2fr;
  column-gap: 30px;
  #md-editor-v3{
    padding: 30px;
    border-radius: 5px;
  }
  .right-side{
    display: grid;
    grid-template-rows: 320px auto;
    grid-template-areas:
                        'user info'
                        'ct ct';
    //grid-template-rows: 320px repeat(auto-fill,auto);
    column-gap: 10px;
    row-gap: 10px;

    //justify-items: end;
    #side-user-container{
      grid-area: user;
      padding: 10px;
    }
    .info{
      width: 300px;
      grid-area: info;
    }
    #comment-container{
      grid-area: ct;
      border-top: 1px solid #ffafd7;
    }


  }




}


</style>

<script setup>
import { ref,onMounted,computed } from 'vue';
import { useRouter } from 'vue-router';
import MdEditor from 'md-editor-v3';
import Header from '@/components/Header.vue';
import Footer from '@/components/Footer.vue';
import Banner from '@/components/Banner.vue';
import 'md-editor-v3/lib/style.css';
import http from '@/http/index'

import SideUser from '@/components/SideUser.vue';
import Comment from "@/components/Comment.vue";
 // mardown 编辑器配置

// const previewTheme = ref('mk-cute');
const previewTheme = ref('default');
const previewOnly = ref(true);
const showCodeRowNumber = ref(true);
const sideUser = ref({
  userInfo:{
    avatar: "https://av.animeii.tech/images/426a62a986b2.jpg",
    nick: "animeii",
    des: "手握明月摘星辰，世间无我这般人。"
  },
  total: 0
})
let isShow = ref(false)
const inShow = ()=>{
  isShow.value = !isShow.value
}

const router = useRouter()

let detail = ref({
    id: 0,
    uid: 0,
    html: '',
    comments: 0,
    created_at: 0,
    likes: 0,
    views: 0,
    user_info:{
        id:0,
        nickname:"ed4eb",
        avatar:"https://av.animeii.tech/images/5e295ba090ad.jpeg"
    }


})


// 时间戳转时间
let created_time = computed(()=>{
    var date = new Date(detail.value.created_at*1000);
    return date.getFullYear() + "-" + (date.getMonth()+1) + "-" + date.getDate() + 
    " " + date.getHours() + ":" + date.getMinutes() + ":" + date.getSeconds();

})



// 根据id获取文章展示到区域
const form = router.currentRoute.value.params

form.id = parseInt(form.id)
onMounted(()=>{
    http.post("/article/detail",form).then(res=>{
        if(res.data.code === 0){
            detail.value = res.data.data
        }
    })
})

</script>
