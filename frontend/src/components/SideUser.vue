<template>
  <div id="side-user-container">
    <img class="avatar" :src="userInfo.avatar" alt="">
    <router-link class="nick" to="/">{{ userInfo.nick }}</router-link>
    <span class="des">{{ userInfo.des }}</span>
    <div class="bo">
      <div class="bo-item">
        <router-link class="num" to="/">{{ total }}</router-link>
        <span class="tp">文章</span>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
#side-user-container{
  //margin-top: 20px;
  width: 300px; // 固定尺寸
  height: 300px;
  display: grid;
  place-items: center; // 定义子元素上下左右居中
  background: white;
  padding: 20px;
  border-radius: 5px;

  //display: grid;
  justify-items: center;
  //justify-self: center;
  grid-row-gap: 10px;
  grid-template-areas:
                    'avatar'
                    'nick'
                    'des'
                    'bo';
  .avatar{
    grid-area: avatar;
    border-bottom: 1px dashed #ffafd7;
    width: 160px;
    height: 160px;

  }
  a{
    text-decoration: none;
    color: #ff85c2;
    &:active{
      color: #00c888;
    }
  }
  .nick{
    grid-area: nick;
    font-size: 20px;
    font-weight: bold;
  }
  .des{
    grid-area: des;
    font-size: 12px;
    color: slategrey;
  }
  .bo{
    grid-area: bo;
    display: grid;
    grid-template-areas: 'bo-item';
    .bo-item{
      grid-area: bo-item;
      display: grid;
      justify-items: center;
      grid-row-gap: 5px;
      grid-template-areas:
                      'num'
                      'tp';
      .num{
        grid-area: num;
        font-size: 16px;
        font-weight: bold;
      }
      .tp{
        grid-area: tp;
        color: slategrey;
        font-size: 12px;
      }

    }

  }
//}
}

</style>

<script setup>
defineProps(
  {
    userInfo:{
      type: Object,
      default: {
        avatar: '',
        nick: '',
        des: '',
      }
    },
    total:{
      type: Number,
      default: 0,
    }
  }
)

</script>
