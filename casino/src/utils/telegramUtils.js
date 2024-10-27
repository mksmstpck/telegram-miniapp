export function getTelegramUser() {
    const initDataUnsafe = window.Telegram.WebApp.initDataUnsafe;
    return initDataUnsafe ? initDataUnsafe.user : null;
  }