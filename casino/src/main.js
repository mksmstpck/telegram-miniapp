import './assets/base.css'

import { createRouter, createWebHistory } from 'vue-router'

import { createApp } from 'vue'
import App from './App.vue'
import Home from '@/views/Home.vue'
import Play from '@/views/Play.vue'
import Friends from '@/views/Friends.vue'
import Wallet from '@/views/Wallet.vue'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {path: '/', name: 'Home', component: Home},
        {path: '/play', name: 'Play', component: Play},
        {path: '/friends', name: 'Friends', component: Friends},
        {path: '/wallet', name: 'Wallet', component: Wallet},

    ]
})

createApp(App)
.use(router)
.mount('#app')
