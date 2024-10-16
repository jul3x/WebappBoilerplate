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

<script setup lang="ts">
import { ref } from "vue";
import axios from "axios";
import { store } from "../store/index.ts";
import { useRouter } from "vue-router";

const email = ref("");
const password = ref("");
const error = ref<string | null>(null);

const router = useRouter();

const login = async () => {
  error.value = null;
  try {
    const response = await axios.post(
      "http://localhost:8080/api/v1/auth/login",
      {
        email: email.value,
        password: password.value,
      }
    );
    store.token = response.data.token;
    store.username = response.data.username;
    localStorage.setItem("token", response.data.token);
    localStorage.setItem("username", response.data.username);
    alert("Login successful!");
    router.push("/");
  } catch (err: any) {
    error.value = err.response?.data || "Login failed.";
  }
};
</script>

<style scoped>
.container {
  max-width: 600px;
  margin: auto;
}
</style>
