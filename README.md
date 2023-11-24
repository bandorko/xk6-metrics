# xk6-file
[k6](https://github.com/grafana/k6) extension for accessing various derived metrics.

Currently supported metrics:
- drop_percentage

## Build
```shell
xk6 build --with github.com/bandorko/xk6-metrics@latest
```

## Example
```javascript
import { sleep } from 'k6';
import { Gauge } from 'k6/metrics';
import metrics from  'k6/x/metrics';

const dropPercentage = new Gauge('drop_percentage');

export const options = {
    discardResponseBodies: true,
    thresholds: {
        drop_percentage: ['value<0.1'],
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

export function teardown(data){
    dropPercentage.add(metrics.getDropPercentage())
}
```

## Run sample script
```shell
./k6 run examples/script.js
```