<template>
    <Header></Header>
    <div id="_404">
        <div class="notfound">
            <img src="../assets/404.jpeg" alt="找不到页面">
        </div>
    </div>
</template>

<style lang="scss" scoped>
#_404{
    display: grid;
    grid-template-rows: 100%;
    grid-template-columns: 100%;
    place-items: center;
}
</style>

<script setup>
import Header from '../components/Header.vue'
</script>
