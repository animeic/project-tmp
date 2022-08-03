<template>
    <div id="search-container">
        <div class="search-input">
            <el-input type="text" placeholder="文章、用户" v-model="text">
                <template #append>
                    <el-button :icon="Search" @click="searchclick"></el-button>
                </template>
            </el-input>
        </div>
    </div>
</template>

<style lang="scss" scoped>
    // #search-container{

    // }
</style>
<script setup>
import { Search } from '@element-plus/icons-vue'
import { ref } from 'vue' 
const text = ref('')

const searchclick = ()=>{
    
    alert('aaa')
}
</script>