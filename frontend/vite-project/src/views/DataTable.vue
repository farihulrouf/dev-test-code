<template>
  <div>
    <input type="text" v-model="searchQuery" placeholder="Search...">
    <select v-model="selectedYear">
      <option value="">All Years</option>
      <option v-for="year in years" :key="year" :value="year">{{ year }}</option>
    </select>

    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th>Year</th>
          <!-- Add more table headers as needed -->
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in filteredItems" :key="item.id">
          <td>{{ item.name }}</td>
          <td>{{ item.year }}</td>
          <!-- Display more item properties as needed -->
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      items: [], // Array to store fetched items from API
      searchQuery: '', // Search query input
      selectedYear: '', // Selected year for filtering
    };
  },
  computed: {
    years() {
      // Extract unique years from items
      return [...new Set(this.items.map(item => item.year))];
    },
    filteredItems() {
      // Apply search query and year filter to items
      let filtered = this.items;
      if (this.searchQuery) {
        filtered = filtered.filter(item => 
          item.name.toLowerCase().includes(this.searchQuery.toLowerCase())
        );
      }
      if (this.selectedYear) {
        filtered = filtered.filter(item => item.year === this.selectedYear);
      }
      return filtered;
    }
  },
  mounted() {
    // Fetch data from API when component is mounted
    this.fetchData();
  },
  methods: {
    async fetchData() {
      try {
        const response = await fetch('https://api-production.data.gov.sg/v2/public/api/collections'); // Replace with actual API endpoint
        const data = await response.json();
        this.items = data; // Update items array with fetched data
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    }
  }
};
</script>

<style>
/* Add custom styles here if needed */
</style>
