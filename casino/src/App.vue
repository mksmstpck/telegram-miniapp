<script>
    import Menu from '@/components/Menu.vue'

    export default {
        components: {
            Menu,
        }
    }
</script>

<template>
    <div className="container">
        <router-view></router-view>
        <Menu />
    </div>    
</template>

<style scoped>
.container {
    display: flex;
    justify-content: center;
    align-items: center;
    position: relative;
    height: 100vh;
}
</style>
