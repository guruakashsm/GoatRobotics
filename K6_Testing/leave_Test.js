import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
    vus: 5, // Number of virtual users
    duration: '10s', // Duration of the test
};

export default function () {

    const randomId = Math.floor(Math.random() * (99999 - 10000 + 1)) + 10000;

    let url = `http://localhost:8080/rpc/GOATROBOTICS/join?id=${randomId}`;
    let res = http.get(url);
    check(res, {
        'status is 200': (r) => r.status === 200,
    });

    url = `http://localhost:8080/rpc/GOATROBOTICS/leave?id=${randomId}`;
    res = http.get(url);
    check(res, {
        'status is 200': (r) => r.status === 200,
    });


    const responseBody = JSON.parse(res.body);

    check(responseBody, {
        'response has userId': (r) => r.hasOwnProperty('userId'),
        'response has message': (r) => r.hasOwnProperty('message'),
        'response has ReponseTime': (r) => r.hasOwnProperty('ReponseTime'),
    });

    check(responseBody, {
        'userId matches generated id': (r) => r.userId === randomId.toString(),
        'message is correct': (r) => r.message === 'Left Chat Successfully',
    });

    console.log(`Response received with ID: ${randomId}`);
    console.log('User ID:', responseBody.userId);
    console.log('Message:', responseBody.message);
    console.log('Response Time:', responseBody.ReponseTime);

    sleep(1);
}
