<template>
  <div id="side-link-container">
    <div class="item" v-for="(v,k) in links" :key="k">
      <a :href="v.url">{{ v.name }}</a>
    </div>

  </div>
</template>

<style lang="scss" scoped>
#side-link-container{
  display: grid;
  width: 300px;
  padding: 20px;
  border-radius: 5px;
  background: white;
  grid-template-columns: repeat(auto-fill,96px);
  //grid-template-rows: repeat(3,40px);
  place-items: center;
  gap: 5px;

  font-size: 14px;
  a{
    text-decoration: none;
    color: #ff85c2;
    &:active{
      color: #00c888;
    }
  }
}

</style>

<script setup>
import {ref} from 'vue'
const links = ref([
  {
    name: '阮一峰',
    url: 'https://www.ruanyifeng.com/blog/'
  },
  {
    name: '极客兔兔',
    url: 'https://geektutu.com/'
  },
  {
    name: '刘丹冰',
    url: 'https://www.yuque.com/aceld'
  },
  {
    name: 'leetcode',
    url: 'https://github.com/doocs/leetcode'
  }

])
</script>