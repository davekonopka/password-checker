# 21 Add load test config

ChatGPT figured out what I meant despite a type (k9 instead of k6).

## Prompt

Can you help me create a k9 load testing configuration file to load test the http password checker web service?

## Response

Sure! k6 is a powerful open-source load testing tool for testing the performance of your backend infrastructure, such as APIs, microservices or even websites. Here's a basic example of how to set up a load testing script for your application in k6 (in JavaScript):

```javascript
import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
    stages: [
        { duration: '2m', target: 100 }, // below normal load
        { duration: '5m', target: 100 },
        { duration: '2m', target: 200 }, // normal load
        { duration: '5m', target: 200 },
        { duration: '2m', target: 300 }, // around the breaking point
        { duration: '5m', target: 300 },
        { duration: '2m', target: 400 }, // beyond the breaking point
        { duration: '5m', target: 400 },
        { duration: '10m', target: 0 }, // scale down. Recovery stage.
    ],
};

export default function () {
    var BASE_URL = 'http://localhost:8080'; // make sure this is not pointing to production
    var PASSWORD = 'password123'; // replace with a test password
    let responses = http.batch([
        [
            'GET',
            `${BASE_URL}/check/${PASSWORD}`,
            null,
            { tags: { name: 'CheckPasswordStrength' } },
        ],
    ]);

    check(responses[0], {
        'status was 200': (r) => r.status === 200,
    });

    sleep(1);
}
```

In this script, you're running a load test with a staged ramping up and down of virtual users (VUs). The load starts below normal, ramps up to normal, goes around the breaking point, goes beyond the breaking point, and finally scales down in the recovery stage.

Please adjust the configuration and the test password as per your needs. 

You can run this script with the command `k6 run <script-name>.js`.

Before running the script, ensure that you have k6 installed. If not, you can install it by following the instructions on their [installation page](https://k6.io/docs/getting-started/installation/).
