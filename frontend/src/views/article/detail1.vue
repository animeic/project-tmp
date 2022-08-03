<template>
  <Header />
  <Banner />
  <div id="article-detail-container">
    <!-- <div class="top-info">{{ detail.uid + "" + detail.id}}</div> -->
    <div class="article-detail-warpper">
      <md-editor v-model="detail.html"
                 :previewOnly="previewOnly"
                 :previewTheme="previewTheme"
                 :showCodeRowNumber="showCodeRowNumber"
      />
    </div>
    <!-- <div class="at-comment"></div> -->
    <div class="right-sider">
      <div class="rs-top-info">
        <router-link :to="{name:'user-profile',params:{id:detail.uid}}" class="rs-user-info">
          <img :src="detail.user_info.avatar" :alt="detail.user_info.nickname">
          <div>{{ detail.user_info.nickname }}</div>
        </router-link>
        <div class="rs-at-info">
          <div class="rs-at-created">
            {{ created_time }}
          </div>
          <div class="rs-at-viwes">
            <el-row>
              浏览：
              {{ detail.views }}
            </el-row>
          </div>
          <div class="rs-at-comments">
            <el-row>
              评论：
              {{ detail.comments }}
            </el-row>
          </div>
          <div class="rs-at-likes">
            <el-row>
              点赞：
              {{ detail.likes }}
            </el-row>
          </div>
        </div>
      </div>
    </div>
  </div>
  <Footer />

</template>

<script setup>
import { ref,onMounted,computed } from 'vue';
import { useRouter } from 'vue-router';
import MdEditor from 'md-editor-v3';
import Header from '@/components/Header.vue';
import Footer from '@/components/Footer.vue';
import Banner from '@/components/Banner.vue';
import 'md-editor-v3/lib/style.css';
import http from '@/http/index'
// mardown 编辑器配置

// const previewTheme = ref('mk-cute');
const previewTheme = ref('default');
const previewOnly = ref(true);
const showCodeRowNumber = ref(true);

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

<style lang="scss" scoped>

#article-detail-container{
  margin-top: 20px;
  min-height: 920px;
  display: grid;
  grid-template-areas:
                        'at-detail right-sider';
  // 'left-sider at-comment right-sider'; // 定义位置
  .article-detail-warpper{
    // 阴影卡片效果
    // box-shadow:0 4px 8px 0 rgba(0,0,0,0.2),0 6px 20px 0 rgba(0,0,0,0.19);
    // display:inline-block;
    // background-color:#FFF;
    padding: 30px;
    padding-top: 0px;
    margin-right: 20px;
    border-radius: 5px;
    // border:3px dashed #ececec ;
    border: 3px dashed #ffdaec;
    //max-width: 800px;
    width: 800px;
    grid-area: at-detail;
    justify-self: end;
  }
  .right-sider{
    grid-area: right-sider;
    // display: grid; // 这里会出现子元素位置不正确
    .rs-top-info{
      max-width: 480px;
      font-size: 14px;
      display: grid;
      grid-template-areas:
                                'rs-user-info'
                                'rs-at-info';
      .rs-user-info{
        // box-shadow:0 4px 8px 0 rgba(0,0,0,0.2),0 6px 20px 0 rgba(0,0,0,0.19);
        // display:inline-block;
        // background-color:#ffaee2;
        padding: 5px 5px 5px 15px;
        margin-bottom: 30px;
        border-radius: 5px;
        border: 2px dashed #ffdaec;
        color: slategrey;
        display: grid;
        grid-area: rs-user-info;
        align-items: center;
        grid-template-areas:
                                'rsu-1 rsu-2';

        img{
          width: 40px;
          height: 40px;
          border-radius: 20px;
          border: 1px dashed plum;
          grid-area: rsu-1;
        }
        div{
          grid-area: rsu-2;
        }

      }
      .rs-at-info{
        // box-shadow:0 4px 8px 0 rgba(0,0,0,0.2),0 6px 20px 0 rgba(0,0,0,0.19);
        // display:inline-block;
        // background-color:#ffaee2;
        border: 2px dashed #ffdaec;
        color: slategrey;
        padding: 5px;
        padding-top: 10px;
        padding: 10px 5px 5px 15px;
        display: grid;
        border-radius: 5px;
        display: grid;
        grid-area: rs-at-info;
        align-items: center;
        grid-template-areas:
                                'rsat1 rsat1 rsat1'
                                'rsat2 rsat3 rsat4';
        height: 60px;
        .rs-at-created{
          grid-area: rsat1;
        }
        .rs-at-viwes{
          grid-area: rsat2;
        }
        .rs-at-comments{

          grid-area: rsat3;
        }
        .rs-at-likes{

          grid-area: rsat4;
        }

      }
    }
  }
  // .at-comment{
  //     height: 100px;
  //     border: 1px solid blueviolet;
  //     grid-area: at-comment;
  // }

}


</style>