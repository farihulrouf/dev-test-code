<template>
  <div id="app">
    <HeaderWeb />
    <SidebarMenu />
    <router-view /> <!-- This will render the component for the current route -->
    <FooterWeb />
  </div>
</template>

<script>
import HeaderWeb from './components/Header.vue';
import SidebarMenu from './components/SidebarMenu.vue';
import FooterWeb from './components/Footer.vue';

export default {
  name: 'App',
  components: {
    HeaderWeb,
    SidebarMenu,
    FooterWeb
  }
};
</script>

<style>
/* Global styles for the entire application */
body {
  font-family: Arial, sans-serif;
  margin: 0;
  padding: 0;
}

#app {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}
</style>
