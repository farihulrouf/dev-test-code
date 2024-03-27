<template>
    <div class="map-views">
      <Header />
      <SidebarMenu />
      <main class="main">
        <!-- Include the MapPage component here -->
        <MapPage />
      </main>
      <Footer />
    </div>
  </template>
  
  <script>
  //import HeaderWeb from '@/components/Header.vue';
 // import FooterWeb from '@/components/Footer.vue';
  import SidebarMenu from '@/components/SidebarMenu.vue';
  import MapPage from '@/components/MapPage.vue'; // Import the MapPage component
  
  export default {
    name: 'MapViews',
    components: {
     // HeaderWeb,
     // FooterWeb,
      SidebarMenu,
      MapPage // Add the MapPage component to the components object
    }
  };
  </script>
  
  <style>
  /* Styles for the map views */
  </style>
  