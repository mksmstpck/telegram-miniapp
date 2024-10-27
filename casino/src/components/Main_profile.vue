<script>
    export default {}
</script>

<template>    
<div className="profile">
    <div>
        <img className="profile-photo" src="https://i.pinimg.com/originals/c4/ae/c0/c4aec03b0441f30345d9e45607ea18de.jpg" alt="human">
    </div>
    <div>
        <h2 className="profile-name">mksmstpck</h2>
    </div>
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

.profile {
    position: absolute;
    top: 20%;                  
    left: 50%;
    transform: translateX(-50%);
    text-align: center;
}

.profile-name {
    margin: 4px;
    font-size: 25px;
    color: white;
}

.profile-photo {
    width: 150px;
    height: 150px;
    border-radius: 50%;
    object-fit: cover;
    overflow: hidden;
    border: 2px rgb(255, 255, 255);
}

.profile-balance p {
    font-size: 35px;
    color: white;
}
</style>
