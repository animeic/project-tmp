<template>  
    <Header />
    <Banner />
    <div id="article-list-container">
        <!-- 左侧 -->
        <div class="left-sider">
          <SideUser :total="sideUser.total" :user-info="sideUser.userInfo" />
          <SideOther />
          <SideLink />
        </div>
        
        <!-- middle main -->
        <div class="list-main">
            <!-- 列表区域 -->
          <div class="list-warpper">
            <div class="list-item" v-for="(item,k) in data.lists" :key="k">
                    <div class="top-info">
                        <router-link class="li-user" :to="{name:'user-profile',params:{id:item.user_info.id}}">
                            <img :src="item.user_info.avatar" :alt="item.user_info.nickname">
                        </router-link>
                    </div>
                    <router-link class="list-content" :to="{name:'article-detail',params:{id:item.id}}">
                      <div class="li-text-title">
                          <p>{{ item.title }}</p>
                      </div>
                      <div class="li-text-summary">
                          <p>{{ item.summary }}</p>
                      </div>
                    </router-link>
                    <div class="bottom-info">
                      <span class="bo-at-created"> 创建于：{{ item.create_time }}</span>
                      <span class="bo-views"> 浏览：{{ item.views }}</span>
                      <span class="bo-likes"> 点赞：{{ item.likes }}</span>
                      <span class="bo-comments"> 评论：{{ item.comments }}</span>
                      <router-link  v-if="role === 'admin'" :to="{name:'admin-article-detail',params:{id:item.id}}" class="at-edit">编辑</router-link>
                    </div>
                </div>
          </div>
          <div class="list-pagination">
                <Pagination
                    :total="total"
                    :page-no="paginForm.cur_page"
                    :pgae-size="paginForm.page_size"
                    @change="changeHandle"
                />
            </div>
        </div>
        <!-- 右侧 -->
        <div class="right-sider"></div>
        
    </div>
    <Footer />
</template>

<style lang="scss" scoped>
#article-list-container{
  min-height: 700px;
  background: #f5f7f9;
  padding-top: 20px;
    display: grid;
    grid-template-areas: 
                        'left-sider list-main right-sider';
    .left-sider{
      grid-area: left-sider;

      display: grid; // 定义容器
      justify-items: center; // 定义子容器居中
      grid-template-rows:340px 120px 120px;
      grid-row-gap: 10px;
    }

    .list-main{
      width: 820px;
      padding: 20px;
      border-radius: 5px;
      grid-area: list-main;
      display: grid;
      background: white;
      justify-items: center;
      grid-template-rows: auto 60px ;
      .list-warpper{
        .list-item {
          // 卡片
          // box-shadow:0 4px 8px 0 rgba(0,0,0,0.2),0 6px 20px 0 rgba(0,0,0,0.19);
          // display:inline-block;
          // background-color:#FFF;
          display: grid;
          padding: 20px;
          border-radius: 5px;
          border: 1px solid #e6bce3;
          margin-top: 5px;
          grid-template-areas:
                                'top-info'
                                'list-content'
                                'bottom-info';

          .top-info {
            grid-area: top-info;
            display: grid;
            grid-template-areas: 'li-user';

            .li-user {
              grid-area: li-user;

              img {
                width: 40px;
                height: 40px;
                border-radius: 20px;
                border: 1px solid plum;
              }
            }

          }

          .list-content {
            grid-area: list-content;
            display: grid;
            grid-template-areas:
                                  'li-text-title'
                                  'li-text-summary';

            .li-text-title {
              grid-area: li-text-title;
              font-size: 18px;
              height: 40px;

            }

            .li-text-summary {
              grid-area: li-text-summary;
              font-size: 14px;
              height: 40px;
              color: slategray;
              // width: 450px;
              // 单行隐藏
              overflow: hidden;
              text-overflow: ellipsis;
              white-space: nowrap;

            }
          }

          .bottom-info {
            grid-area: bottom-info;
            height: 40px;
            display: grid;
            font-size: 12px;
            color: slategray;
            grid-template-areas:
                                  'bo-created bo-views bo-likes bo-comments at-edit';

            align-items: center;

            .bo-at-created {
              grid-area: bo-created;
            }

            .bo-views {
              grid-area: bo-views;
            }

            .bo-likes {
              grid-area: bo-likes;
            }

            .bo-comments {
              grid-area: bo-comments;
            }

            .at-edit {
              grid-area: at-edit;
              color: #ff85c2;
              &:active{
                color: #00c888;
              }
              font-weight: 300;
            }
          }
        }
      }
      .list-pagination{
        //grid-area: list-pagination;
        height: 60px;
        display: grid;
        align-items: center;
        :deep(.el-pager li:not(.disabled).active){
          background-color: #ff85c2;
        }


      }

    }
    .right-sider{
      grid-area: right-sider;
      background: #f5f7f9;
    }


    a{
        text-decoration: none;
        color: #000;
    }
}
</style>

<script setup>
import { ref,onBeforeMount,computed } from 'vue';
import { useRouter } from 'vue-router';
import Header from '@/components/Header.vue';
import Footer from '@/components/Footer.vue';
import Banner from '@/components/Banner.vue';
import SideUser from '@/components/SideUser.vue';
import http from '@/http/index'
import Pagination from '@/components/Pagination.vue';
import SideOther from "@/components/SideOther.vue";
import SideLink from "@/components/SideLink.vue";
 const data = ref(
    {
        lists:[
                {
                    id: 1,
                    uid:1,
                    title: "",
                    summary: "",
                    created_at: 0,
                    views: 0,
                    likes: 0,
                    comments: 0,
                    user_info:{
                        id: 1,
                        nickname: "animeic",
                        avatar: "https://av.animeii.tech/images/426a62a986b2.jpg"
                    }
                }

            ],
        total: 1

    }

)
const role = localStorage.getItem("role")

// side-user 数据
const sideUser = ref({
  userInfo:{
    avatar: "https://av.animeii.tech/images/426a62a986b2.jpg",
    nick: "animeii",
    des: "手握明月摘星辰，世间无我这般人。"
  },
  total: 0
})

// const role = ref(thisRole)
const total = ref(1)
const paginForm = ref({
    cur_page: 1,
    page_size: 5,
})
const router = useRouter()
onBeforeMount(()=>{
    const cp = localStorage.getItem("cur_page")
    if (parseInt(cp) > 0){
        paginForm.value.cur_page = parseInt(cp)
    }
    // const page_query = router.currentRoute.value.query.page
    // if (parseInt(page_query) > 0){
    //     paginForm.value.cur_page = parseInt(page_query)
    // }
    http.post("/article/list",paginForm.value).then((res)=>{
        // 赋值给内容页面
        if(res.data.code == 0){
            // data.value = res.data.data
            total.value = res.data.data.total
            sideUser.value.total = res.data.data.total
            data.value.lists = Glist(res.data.data.lists)
        }

    })

})

const Glist = (lists)=>{
   for(let i = 0;i<lists.length;i++){
      lists[i].create_time = createTime(lists[i].created_at)
     // console.log(lists[i])
   }
   return lists
}

const createTime = (created) =>{
  let date = new Date(created*1000);
  return date.getFullYear() + "-" + (date.getMonth()+1) + "-" + date.getDate() +
      " " + date.getHours() + ":" + date.getMinutes() + ":" + date.getSeconds();
}

const changeHandle = (val)=>{
    paginForm.value.cur_page = val.pageNo
    paginForm.value.page_size = val.pageSize
    router.currentRoute.value.query.page = val.pageNo
    http.post("/article/list",paginForm.value).then((res)=>{
        // 赋值给内容页面
        if(res.data.code == 0){
            total.value = res.data.data.total
            localStorage.setItem("cur_page",val.pageNo)
            data.value.lists = Glist(res.data.data.lists)
        }

    })
    // .catch()
}

</script>
