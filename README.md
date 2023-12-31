# xk6-metrics
[k6](https://github.com/grafana/k6) extension for accessing various derived metrics.

Currently supported metrics:
- drop_percentage

## Build
you can build the custom k6 binary with [xk6](https://github.com/grafana/xk6)

```shell
xk6 build --with github.com/bandorko/xk6-metrics@latest
```

## Example
```javascript
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

```

## Run sample script
```shell
./k6 run examples/script.js
```