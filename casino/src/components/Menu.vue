<script>
    export default {}
</script>

<template>    
<div className="menu">
    <div class="menu-item">
        <router-link to="/"><img src="@/assets/static/menu-home.svg" alt="menu-item"></router-link>
    </div>
    <div class="menu-item">
        <router-link to="play"><img src="@/assets/static/menu-play.svg" alt="menu-item"></router-link>
    </div>
    <div class="menu-item">
        <router-link to="friends"><img src="@/assets/static/menu-friends.svg" alt="menu-item"></router-link>
    </div>
    <div class="menu-item">
        <router-link to="wallet"><img src="@/assets/static/menu-wallet.svg" alt="menu-item"></router-link>
    </div>
</div>
</template>

<style scoped>
.menu {
    position: fixed;
    bottom: 0;
    left: 0;
    width: 100%;
    height: 60px;
    background-color: rgba(28, 23, 43, 0.5);
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0 10px;
}

.menu-item {
    flex: 1;
    text-align: center;
    display: flex;
    justify-content: center;
    align-items: center;
    position: relative;
    margin-bottom: 35px;
}

.menu-item img {
    position: absolute;
    top: 20%;                  
    left: 50%;
    transform: translateX(-50%);
    text-align: center;
    height: 40px;
}
</style>
