<template>
<!--  <router-view />-->

  <router-view v-slot="{ Component }">
    <keep-alive>
      <component :is="Component"
                 v-if="$route.meta.keepAlive"
                 :key="$route.path" />
    </keep-alive>
    <component :is="Component"
               v-if="!$route.meta.keepAlive" />
  </router-view>
</template>

<script>
export default {};
</script>

<style>
@import "./assets/css/main.css";
@import "./assets/css/color-dark.css";
</style>
