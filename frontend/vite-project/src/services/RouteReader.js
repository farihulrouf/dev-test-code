// RouteReader.js

// Example route data
const routeData = [
  { latitude: 51.5074, longitude: -0.1278 }, // Example coordinates
  { latitude: 52.5200, longitude: 13.4050 },
  // Add more route coordinates as needed
];

export default class RouteReader {
  async readRouteData() {
    // Simulate reading route data asynchronously (e.g., from an API)
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        resolve(routeData); // Resolve with example route data
      }, 1000); // Simulate 1 second delay
    });
  }
}
