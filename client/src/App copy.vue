<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
import HelloWorld from "@/components/HelloWorld.vue";
</script>

<template>
  <div class="flex flex-row h-full w-full">
    <div class="w-28 bg-gray-200 flex-shrink-0 p-3">
      <p>
        <RouterLink to="/">Home</RouterLink>
      </p>
      <p>
        <RouterLink to="/about">About</RouterLink>
      </p>
    </div>
    <div class="bg-slate-50 flex-grow p-3">
      <RouterView />aaaa
    </div>
  </div>
</template>

<style>
@import "@/assets/base.css";
html,
body,
#app {
  height: 100%;
  width: 100%;
}
</style>

