<template>
    <div class="flex-1 bg-gray-100">
      <!-- Header -->
      <header class="bg-gray-700 text-white p-4">
        <h1 class="text-xl font-semibold">Header</h1>
      </header>
  
      <!-- Main content area -->
      <main class="p-4">
        <router-view></router-view>
      </main>
    </div>
  </template>
  
  <script>
  export default {
    name: 'MainContent'
  };
  </script>
  
  <style>
  /* Add custom styles here if needed */
  </style>
  