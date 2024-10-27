<script>
    import Balance from '@/components/Balance.vue'

    export default {
        components: {
            Balance,
        }
    }
</script>

<template>
    <div className="svg-friends">
        <img src="@/assets/static/menu-friends.svg" alt="menu-item">
    </div>
    <div className="main">
        <h2 className="main-text">
            Invite Friends
        </h2>
        <div className="main-balance">
            <Balance />
        </div>
    </div>
    <div className="main-text-2">
        <p>
            Get 10% from every token your friend earn!
        </p>
    </div>
    <hr className="horizontal">
    <div className="counter">
        <p>
            10 friends
        </p>
    </div>
    <div className="friends">
        <div className="friend">
            <p className="name">
                @yechkin
            </p>
            <p className="profit">
                10
            </p>
        </div>
    </div>
</template>

<style>
.main{
    position: absolute;
    top: 35%;                  
    left: 50%;
    transform: translateX(-50%);
    text-align: center;
}

.main-text {
    margin: 4px;
    font-size: 30px;
    color: rgb(174, 174, 174);
}

.main-container{
    display: flex;
    justify-content: center;
    border-radius: 10px;
    background-color: white;
    margin: 10px;
    padding: 10px;
}

.svg-friends{
    position: absolute;
    top: 20%;                  
    left: 50%;
    transform: translateX(-50%);
    text-align: center;
}

.svg-friends img{
    height: 80px;
}

.main-balance {
    position: absolute;
    left: 50%;
    transform: translateX(-50%);
    text-align: center;
    border: 2px solid rgb(174, 174, 174);
    border-radius: 10px;
    padding: 10px;
}

.main-text-2 {
    position: absolute;
    top: 55%;
    transform: translateY(-45%);
    margin: 4px;
    font-size: 15px;
    color: rgb(174, 174, 174);
}

.horizontal {
    position: absolute;
    top: 60%;
    transform: translateY(-40%);
    width: 80%;
    margin: 0 auto;
}

.counter {
    position: absolute;
    top: 61%;
    transform: translateY(-39%);
    left: 10%;
}

.counter p {
    color: white;
}

.friends {
    overflow-y: auto;
    overflow-x: hidden;
    position: absolute;
    top: 65%;
    left: 10%;
    width: 80%;
    scroll-behavior: smooth;
    height: auto;
}

.friend {
    display: flex;
    justify-self: space-between;
    height: 50px;
    width: 100%;
    background-color: rgb(7, 0, 31);
    border: 2px solid grey;
    border-radius: 5px;
}
.name {
    font-size: 30px;
    color: white;
}

.profit {
    font-size: 25px;
    color: white;
}
</style>