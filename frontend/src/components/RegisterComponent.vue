<template>
  <div class="container">
    <h2>Register</h2>
    <form @submit.prevent="register">
      <div class="mb-3">
        <label for="username" class="form-label">Username</label>
        <input type="text" v-model="username" class="form-control" required />
      </div>
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
      <button type="submit" class="btn btn-primary">Register</button>
    </form>
    <div v-if="error" class="alert alert-danger mt-2">{{ error }}</div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import axios from "axios";
import { useRouter } from "vue-router";

const username = ref("");
const email = ref("");
const password = ref("");
const error = ref<string | null>(null);

const router = useRouter();

const register = async () => {
  error.value = null;
  try {
    await axios.post("http://localhost:8080/api/v1/auth/register", {
      username: username.value,
      email: email.value,
      password: password.value,
    });
    alert("Registration successful!");
    router.push("/");
  } catch (err: any) {
    error.value = err.response?.data || "Registration failed.";
  }
};
</script>

<style scoped>
.container {
  max-width: 600px;
  margin: auto;
}
</style>
