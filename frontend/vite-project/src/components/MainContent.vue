<template>
  <div class="flex flex-col h-screen bg-gray-100 w-full"> <!-- Update class to flex-col and add h-screen -->
    <!-- Header -->
    <header class="bg-gray-700 text-white p-4">
      <h1 class="text-xl font-semibold">Header</h1>
    </header>

    <!-- Main content area -->
    <main class="flex-1 p-4"> <!-- Update class to flex-1 -->
      <router-view></router-view>
    </main>

    <!-- Footer -->
    <FooterWeb />
  </div>
</template>

<script>
import FooterWeb from './Footer.vue';

export default {
  name: 'MainContent',
  components: {
    FooterWeb // Include the Footer component
  }
};
</script>

<style>
/* Add custom styles here if needed */
</style>
