<template>
  <div id="pem-list-container">
    <div class="item" v-for="(v,k) in lists" :key="k">
      <span class="pem-num">{{ k + 1 }}</span>
      <span class="pem-name">{{ v.name }}</span>
      <span class="pem-path">{{ v.path }}</span>
      <span class="pem-t">{{ v.created_at }}</span>
      <span class="pem-t">{{ v.updated_at }}</span>
      <span class="pem-dt">{{ v.deleted_at }}</span>
      <button class="pem-edit" @click="editHandle(v.id)">编辑</button>
      <button class="pem-del" @click="deleteHandle(v.id)">删除</button>
    </div>
  </div>
</template>

<style lang="scss" scoped>
#pem-list-container{
  display: grid;
  max-width: 1200px;
  padding: 20px;
  border: 1px solid black;
  border-radius: 5px;
  grid-template-columns: auto;
  grid-template-rows: repeat(auto-fill,40px);
  grid-row-gap: 10px;
  .item{
    display: grid;
    place-items: center;
    height: 40px;// 确定每行高度
    grid-template-columns: repeat(8,minmax(60px,150px));
    grid-column-gap:1px;
    border: 1px solid red;
    border-radius:5px;

  }
}


</style>

<script setup>
import {ref} from 'vue'
const lists = ref([
  {
    id: 0,
    name: '注册',
    path: '/register',
    created_at: 0,
    updated_at: 0,
    deleted_at: 0,
  },
  {
    id: 0,
    name: '登录',
    path: '/login',
    created_at: 0,
    updated_at: 0,
    deleted_at: 0,
  },
  {
    id: 0,
    name: '新增文章',
    path: '/admin/article/post',
    created_at: 0,
    updated_at: 0,
    deleted_at: 0,
  },

])

function editHandle(id){
  alert(id)
}
function deleteHandle(id){
  console.log(id)
}


</script>