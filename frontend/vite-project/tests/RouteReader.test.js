// RouteReader.test.js

import RouteReader from './RouteReader';

describe('RouteReader', () => {
  let reader;

  beforeEach(() => {
    reader = new RouteReader();
  });

  it('should return route data', async () => {
    // Mocking route data for testing purposes
    const mockRouteData = [
      { latitude: 51.5074, longitude: -0.1278 },
      { latitude: 52.5200, longitude: 13.4050 },
      // Add more mock route data as needed
    ];

    // Mocking asynchronous reading of route data
    jest.spyOn(reader, 'readRouteData').mockResolvedValue(mockRouteData);

    // Call the readRouteData method
    const routeData = await reader.readRouteData();

    // Assert that the route data returned matches the mock data
    expect(routeData).toEqual(mockRouteData);
  });

  it('should return route data asynchronously', async () => {
    // Call the readRouteData method
    const promise = reader.readRouteData();

    // Assert that the result is a Promise
    expect(promise).toBeInstanceOf(Promise);

    // Wait for the promise to resolve and assert the result
    const routeData = await promise;
    expect(routeData).toBeDefined();
  });
});
