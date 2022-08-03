<template>
  <div id="comment-container">
    <div class="comment-select">
      <span class="cs-text">评论</span>
      <a class="cs-new">最热</a>
      <a class="cs-hot">最新</a>
    </div>
    <div class="comment-input">
      <el-input v-model="commentInput" placeholder="评论...">
        <template #append>提交</template>
      </el-input>
    </div>
    <!--    <div class="comment-text" v-for="(v,k) in ['e','f']" :key="k">-->
    <div class="comment-text">
      <img src="https://av.animeii.tech/images/426a62a986b2.jpg" alt="" class="ct-avatar">
      <div class="ct-content">
        <router-link to="/" class="nick">animeii</router-link>
        <span class="content">奶萌女主逆袭合集整理好了四个半小时，奶萌女主逆袭合集整理好了四个半小时，奶萌女主逆袭合集整理好了四个半小时，奶萌女主逆袭</span>
        <div class="ct-info">
          <span class="created">2022-07-18 17:56</span>
          <a class="like">喜欢</a>
          <a class="unlike">不喜欢</a>
          <a class="replay" @click="inShow">回复</a>
        </div>

        <div class="sc-comment" v-for="(v,k) in ['a','b','c','d']" :key="k">
          <img src="https://av.animeii.tech/images/426a62a986b2.jpg" alt="" class="ct-avatar">
          <div class="ct-content">
            <router-link to="/" class="nick">animeii</router-link>
            <span class="content">奶萌女主逆袭合集整理好了四个半小时，奶萌女主逆袭合集整理好了四个半小时，奶萌女主逆袭合集整理好了四个半小时，奶萌女主逆袭</span>
            <div class="ct-info">
              <span class="created">2022-07-18 17:56</span>
              <a class="like">喜欢</a>
              <a class="unlike">不喜欢</a>
              <a class="replay" @click="inShow">回复</a>
            </div>
          </div>
        </div>

        <div class="ct-input" v-if="isShow">
          <el-input v-model="commentInput" placeholder="评论...">
            <template #append>提交</template>
          </el-input>
        </div>
      </div>



    </div>
  </div>
</template>

<style lang="scss" scoped>
#comment-container{
  display: grid;
  grid-template-areas:
                      'cs'
                      'ci'
                      'ct';
  grid-template-rows: 60px 60px auto;
  row-gap: 5px;
  .comment-select{
    grid-area: cs;
    height: 60px;
    display: grid;
    grid-template-columns: 60px 60px 60px;
    column-gap: 10px;
    place-items: center;
    .cs-text{
      font-size: 18px;
    }
    .cs-new{
      font-size: 16px;
    }
    .cs-hot{
      font-size: 16px;
    }

  }
  .comment-input{
    grid-area: ci;
    height: 60px;
    display: grid;
    align-items: center;
    border-bottom: 1px dashed #ffafd7;
  }
  .comment-text{
    margin-top: 10px;
    grid-area: ct;
    display: grid;
    grid-template-columns: 60px auto;
    grid-template-rows: 60px 60px auto 60px;
    //grid-template-rows: repeat(auto-fill,60px);
    column-gap: 10px;
    .ct-avatar{
      height: 60px;
      width: 60px;
    }
    .ct-content{
      display: grid;
      grid-template-rows: repeat(auto-fill,30px);
      align-items: center;
      row-gap: 5px;
      .nick{
        font-size: 16px;
        text-decoration: none;
      }
      .content{
        font-size: 16px;
        line-height: 25px;
      }
      .ct-info{
        font-size: 14px;
        color: slategrey;
        grid-template-columns: 120px repeat(3,60px);
        display: grid;
        column-gap: 5px;
      }
      .sc-comment{
        display: grid;
        align-items: center;
        grid-template-columns: 60px auto;
        column-gap: 10px;
        .ct-avatar{
          align-self: start;
          height: 60px;
          width: 60px;
        }
        .ct-content {
          display: grid;
          grid-template-rows: repeat(auto-fill, 30px);
          align-items: center;
          row-gap: 5px;

          .nick {
            font-size: 16px;
            text-decoration: none;
            //color: red;
          }

          .content {
            font-size: 16px;
          }

          .ct-info {
            font-size: 14px;
            color: slategrey;
            grid-template-columns: 120px repeat(3, 60px);
            display: grid;
            column-gap: 5px;
          }
        }

      }

      .ct-input{

      }
    }

  }
}
</style>

<script setup>
import {ref} from 'vue'

const commentInput = ref("")
const isShow = ref(false)

const inShow = ()=>{
  isShow.value = !isShow.value
}
</script>