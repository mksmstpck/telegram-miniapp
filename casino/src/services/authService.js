import axios from 'axios';

export async function verifyUserWithBackend(user, hash) {
  try {
    const response = await axios.post('#', { user, hash });
    return response.data;
  } catch (error) {
    console.error("Verification failed:", error);
    throw error;
  }
}