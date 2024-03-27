<template>
  <div id="app">
    <Header />
    <SidebarMenu />
    <router-view /> <!-- This will render the component for the current route -->
    <Footer />
  </div>
</template>

<script>
import Header from './components/Header.vue';
import SidebarMenu from './components/SidebarMenu.vue';
import Footer from './components/Footer.vue';

export default {
  name: 'App',
  components: {
    Header,
    SidebarMenu,
    Footer
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
