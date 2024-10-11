<template>
  <div class="container">
    <h2>Login</h2>
    <form @submit.prevent="login">
      <div class="mb-3">
        <label for="email" class="form-label">Email</label>
        <input type="email" v-model="email" class="form-control" required />
      </div>
      <div class="mb-3">
        <label for="password" class="form-label">Password</label>
        <input
          type="password"
          v-model="password"
          class="form-control"
          required
        />
      </div>
      <button type="submit" class="btn btn-primary">Login</button>
    </form>
    <div v-if="error" class="alert alert-danger mt-2">{{ error }}</div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import axios from "axios";

export default defineComponent({
  setup() {
    const email = ref("");
    const password = ref("");
    const error = ref<string | null>(null); // Declare error as a string or null

    const login = async () => {
      error.value = null; // Reset error
      try {
        const response = await axios.post(
          "http://localhost:8080/api/v1/auth/login",
          {
            email: email.value,
            password: password.value,
          }
        );
        localStorage.setItem("token", response.data.token); // Store the JWT token
        alert("Login successful!");
      } catch (err: any) {
        error.value = err.response?.data || "Login failed."; // Use optional chaining for safety
      }
    };

    return {
      email,
      password,
      error,
      login,
    };
  },
});
</script>

<style scoped>
.container {
  max-width: 600px;
  margin: auto;
}
</style>
