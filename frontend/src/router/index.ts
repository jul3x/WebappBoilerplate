import { createRouter, createWebHistory } from "vue-router";
import RegisterComponent from "../components/RegisterComponent.vue";
import LoginComponent from "../components/LoginComponent.vue";
import ProtectedComponent from "../components/ProtectedComponent.vue";
import MainComponent from "../components/MainComponent.vue";
import AdminComponent from "../components/AdminComponent.vue";

const routes = [
  {
    path: "/",
    name: "Main",
    component: MainComponent,
  },
  {
    path: "/register",
    name: "Register",
    component: RegisterComponent,
  },
  {
    path: "/login",
    name: "Login",
    component: LoginComponent,
  },
  {
    path: "/protected",
    name: "Protected",
    component: ProtectedComponent,
    meta: { requiresAuth: true },
  },
  {
    path: "/admin",
    name: "Admin",
    component: AdminComponent,
    meta: { requiresAdmin: true },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  if (to.meta.requiresAuth && !localStorage.getItem("role")) {
    next({ name: "Login" });
  } else if (to.meta.requiresAdmin && localStorage.getItem("role") !== "admin") {
    next({ name: "Main" });
  } else {
    next();
  }
});

export default router;
