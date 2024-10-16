<template>
  <div id="app">
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
      <div class="container-fluid">
        <a class="navbar-brand" href="/">Webapp Template</a>
        <button
          class="navbar-toggler"
          type="button"
          data-bs-toggle="collapse"
          data-bs-target="#navbarNav"
          aria-controls="navbarNav"
          aria-expanded="false"
          aria-label="Toggle navigation"
        >
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarNav">
          <ul class="navbar-nav">
            <li class="nav-item">
              <router-link class="nav-link" to="/" active-class="active"
                >Main</router-link
              >
            </li>
            <li class="nav-item" v-if="username">
              <router-link
                class="nav-link"
                to="/protected"
                active-class="active"
                >Protected</router-link
              >
            </li>
            <li class="nav-item" v-if="!username">
              <router-link class="nav-link" to="/register" active-class="active"
                >Register</router-link
              >
            </li>
            <li class="nav-item" v-if="!username">
              <router-link class="nav-link" to="/login" active-class="active"
                >Login</router-link
              >
            </li>
            <li class="nav-item" v-if="username">
              <a href="#" class="nav-link">Hello {{ username }}!</a>
            </li>
            <li class="nav-item" v-if="username">
              <button class="nav-link" @click="logout">Logout</button>
            </li>
          </ul>
        </div>
      </div>
    </nav>

    <div class="container mt-4">
      <router-view v-slot="{ Component }">
        <transition name="fade" mode="out-in">
          <component :is="Component" />
        </transition>
      </router-view>
    </div>
  </div>
</template>

<script setup lang="ts">
import { store } from "./store/index.ts";
import {
  defineComponent,
  ref,
  computed,
  onMounted,
  onBeforeUnmount,
} from "vue";
import { useRouter } from "vue-router";

const username = ref(store.username || localStorage.getItem("username"));
const token = ref(store.token || localStorage.getItem("token") || "");

store.username = username;
store.token = token;
const router = useRouter();

async function logout() {
  try {
    store.username = null;
    store.token = null;
    localStorage.removeItem("token");
    localStorage.removeItem("username");
    router.push("/login");
    alert("Successfully logged out!");
  } catch (error) {
    alert("Logout failed...");
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 0px;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

.nav-item > button {
  width: 100%;
}
</style>
