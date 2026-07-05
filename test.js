import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  vus: 150,
  duration: '1m0s',
  cloud: {
    // Project: Default project
    projectID: 7987480,
    // Test runs with the same name groups test runs together.
    name: 'Flexurl_load_tests_v1'
  }
};

export default function() {
  http.get('https://flexurl.app');
  // sleep(1);
}