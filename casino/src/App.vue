<script>
    import Menu from '@/components/Menu.vue'
    import { verifyUserWithBackend } from '@/services/authService';
    import { getTelegramUser } from '@/utils/telegramUtils';

    export default {
        components: {
        Menu,
    },
  data() {
    return {
      loading: true,
      authorized: false,
      error: null,
    }
  },
  async mounted() {
    const user = getTelegramUser();
    if (user) {
      try {
        const hash = window.Telegram.WebApp.initDataUnsafe.hash;
        await verifyUserWithBackend(user, hash);
        this.authorized = true;
      } catch (error) {
        this.error = "Authorization failed. Please try again.";
        console.error("Authorization error:", error);
      } finally {
        this.loading = false;
      }
    } else {
      this.error = "No Telegram user data available.";
      this.loading = false;
    }
  },
};
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
