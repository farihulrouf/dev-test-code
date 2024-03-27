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
            <th>Year</th>
            <th>Data 1</th>
            <th>Data 2</th>
            <!-- Add more columns if needed -->
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in filteredData" :key="item.id">
            <td>{{ item.year }}</td>
            <td>{{ item.data1 }}</td>
            <td>{{ item.data2 }}</td>
            <!-- Display more data fields if needed -->
          </tr>
        </tbody>
      </table>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    name: 'DataTable',
    data() {
      return {
        data: [],
        searchQuery: '',
        selectedYear: '',
      };
    },
    computed: {
      filteredData() {
        let filtered = this.data;
        if (this.searchQuery) {
          filtered = filtered.filter(item =>
            Object.values(item).some(val => val.toString().toLowerCase().includes(this.searchQuery.toLowerCase()))
          );
        }
        if (this.selectedYear) {
          filtered = filtered.filter(item => item.year === this.selectedYear);
        }
        return filtered;
      },
      years() {
        return [...new Set(this.data.map(item => item.year))];
      },
    },
    mounted() {
      this.fetchData();
    },
    methods: {
      async fetchData() {
        try {
          const response = await axios.get('https://beta.data.gov.sg/datasets');
          this.data = response.data;
        } catch (error) {
          console.error('Error fetching data:', error);
        }
      },
    },
  };
  </script>
  
  <style scoped>
  /* Add your table styles here */
  </style>
  