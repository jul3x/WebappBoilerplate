<template>
  <div class="container">
    <h2>Protected Data</h2>
    <p>This is a protected route. Only accessible when logged in.</p>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({});
</script>

<style scoped>
.container {
  max-width: 600px;
  margin: auto;
}
</style>
