import http from 'k6/http';
import metrics from 'k6/x/metrics';

export const options = {
  discardResponseBodies: true,
  scenarios: {
    scenario1: {
      executor: 'shared-iterations',
      vus: 1,
      iterations: '10',
    },
  },
};

const httpReqs = metrics.registry.get('http_reqs');

export default function () {
  http.get('https://test.k6.io/');
  console.log('http_reqs:' + httpReqs.sink.value);
}