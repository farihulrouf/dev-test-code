<template>
    <div>
      <Header />
      <SidebarMenu />
      <DataTable />
      <Footer />
    </div>
  </template>
  
  <script>
  import Header from '@/components/Header.vue';
  import Footer from '@/components/Footer.vue';
  import SidebarMenu from '@/components/SidebarMenu.vue';
  import DataTable from '@/components/DataTable.vue';
  
  export default {
    name: 'DataTablePage',
    components: {
      Header,
      Footer,
      SidebarMenu,
      DataTable,
    },
  };
  </script>
  
  <style scoped>
  /* Add your page styles here */
  </style>
  