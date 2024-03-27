<template>
    <div class="min-h-screen flex items-center justify-center bg-gray-100">
      <div class="max-w-md w-full p-6 bg-white shadow-md rounded-lg">
        <h2 class="text-2xl font-bold mb-4">Login</h2>
        <form @submit.prevent="login">
          <div class="mb-4">
            <label for="username" class="block text-gray-700">Username</label>
            <input id="username" type="text" v-model="username" class="mt-1 p-2 block w-full border rounded-md">
          </div>
          <div class="mb-4">
            <label for="password" class="block text-gray-700">Password</label>
            <input id="password" type="password" v-model="password" class="mt-1 p-2 block w-full border rounded-md">
          </div>
          <button type="submit" class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600">Login</button>
        </form>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        username: '',
        password: ''
      };
    },
    methods: {
      login() {
        // Implement authentication logic here
        console.log('Username:', this.username);
        console.log('Password:', this.password);
        // Example: Redirect to dashboard after successful login
        this.$router.push('/dashboard');
      }
    }
  };
  </script>
  
  <style scoped>
  /* Add custom styles here if needed */
  </style>
  