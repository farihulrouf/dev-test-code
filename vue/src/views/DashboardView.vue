<template>
    <div class="dashboard-page">
      <Header />
      <SidebarMenu />
      <main class="main">
        <DashboardView />
      </main>
      <Footer />
    </div>
  </template>
  
  <script>
  import Header from '@/components/Header.vue';
  import Footer from '@/components/Footer.vue';
  import SidebarMenu from '@/components/SidebarMenu.vue';
  import DashboardView from '@/views/DashboardView.vue'; // Import the DashboardView component
  
  export default {
    name: 'DashboardPage',
    components: {
      Header,
      Footer,
      SidebarMenu,
      DashboardView // Add the DashboardView component to the components object
    }
  };
  </script>
  
  <style>
  /* Styles for the dashboard page */
  </style>
  