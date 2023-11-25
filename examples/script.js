import { sleep } from 'k6';
import { Gauge } from 'k6/metrics';
import metrics from  'k6/x/metrics';

export const options = {
    discardResponseBodies: true,
    thresholds: {
        drop_percentage: ['count<0.1'],  // you can use drop_percentage metrics here
      },
    scenarios: {
      contacts: {
        executor: 'constant-arrival-rate',
        duration: '5s',
        rate: 30,
        timeUnit: '1s',
        preAllocatedVUs: 14,
      },
    },
  };

export default function(){
    sleep(0.5);  // 2 iterations/second/VU
}

