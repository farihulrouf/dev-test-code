<template>
  <div class="calculator-view">
    <Header />
    <SidebarMenu />
    <main class="main">
      <CalculatorPage />
    </main>
    <Footer />
  </div>
</template>

<script>
//import HeaderWeb from '@/components/Header.vue';
//import FooterWeb from '@/components/Footer.vue';
//import SidebarMenu from '@/components/SidebarMenu.vue';
import CalculatorPage from '@/components/CalculatorPage.vue';

export default {
  components: {
    //HeaderWeb,
   // FooterWeb,
   // SidebarMenu,
    CalculatorPage
  }
};
</script>

<style>
/* Styles for the calculator view */
</style>
