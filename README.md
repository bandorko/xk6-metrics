# xk6-metrics
[k6](https://github.com/grafana/k6) extension for accessing various metrics.

## Extra metrics
xk6-metrics defines extra metrics.
Currently supported extra metrics are:

- drop_percentage

## Accessing metrics values

You can access the values of the metrics during test run with the exported metrics registy

```js
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
```

## Build
you can build the custom k6 binary with [xk6](https://github.com/grafana/xk6)

```shell
xk6 build --with github.com/bandorko/xk6-metrics@latest
```

## Example

```javascript
import { sleep } from 'k6';
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