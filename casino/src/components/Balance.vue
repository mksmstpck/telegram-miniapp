<script>
    export default {}
</script>

<template>    
<div class="profile-balance">
    <div class="gamble-token">
        <img src="@/assets/static/gamble-token.svg" alt="gamble-token">
    </div>
    <p>9999.99</p>
</div>
</template>

<style scoped>
.profile-balance {
    text-align: center;
    display: flex;
    gap: 10px;
}

.profile-balance p {
    font-size: 35px;
    color: white;
}

.gamble-token {
    position: relative;
    border-radius: 50%;
    height: 40px;
    width: 40px;
    background-color: white;
    text-align: center;
    line-height: 100px;
    margin-right: 4px;
    margin-top: 5px;
}

.gamble-token img {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    height: 35px;
}
</style>
