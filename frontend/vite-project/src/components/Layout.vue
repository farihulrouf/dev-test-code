<template>
  <div class="flex h-screen">
    <!-- Sidebar -->
    <Sidebar />

    <!-- Main content -->
    <MainContent />
  </div>
</template>

<script>
import Sidebar from './Sidebar.vue';
import MainContent from './MainContent.vue';

export default {
  components: {
    Sidebar,
    MainContent
  }
};
</script>

<style>
/* Add custom styles here if needed */
</style>
