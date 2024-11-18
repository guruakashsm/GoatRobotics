import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
  vus: 5,
  duration: '10s',
};

export default function () {
  const randomId = Math.floor(Math.random() * (99999 - 10000 + 1)) + 10000;

  let url = `http://localhost:8080/rpc/GOATROBOTICS/join?id=${randomId}`;
  let res = http.get(url);
  
  check(res, {
    'status is 200': (r) => r.status === 200,
  });

  const message = 'Hello from Guru';
  const encodedMessage = encodeURIComponent(message);


  url = `http://localhost:8080/rpc/GOATROBOTICS/send?id=${randomId}&message=${encodedMessage}`;
  res = http.get(url);

  check(res, {
    'status is 200': (r) => r.status === 200,
  });

  const responseBody = JSON.parse(res.body);

  // Check if the response from send has the necessary fields
  check(responseBody, {
    'response has userId': (r) => r.hasOwnProperty('userId'),
    'response has message': (r) => r.hasOwnProperty('message'),
    'response has ReponseTime': (r) => r.hasOwnProperty('ReponseTime'),
  });

  // Verify the message was sent successfully
  check(responseBody, {
    'userId matches generated id': (r) => r.userId === randomId.toString(),
    'message is correct': (r) => r.message === 'Message Sent Successfully',
  });

  // Fetch the list of messages
  url = `http://localhost:8080/rpc/GOATROBOTICS/messages?id=${randomId}`;
  res = http.get(url);

  check(res, {
    'status is 200': (r) => r.status === 200,
  });

  const messagesResponseBody = JSON.parse(res.body);

  // Check if the sent message is in the response's messages array
  check(messagesResponseBody, {
    'message is in the list': (r) => {
      return r.messages.some(msg => msg.message === 'Hello from Guru' && msg.userId === randomId.toString());
    },
  });

  console.log(`Response received with ID: ${randomId}`);
  console.log('User ID:', responseBody.userId);
  console.log('Message:', responseBody.message);
  console.log('Response Time:', responseBody.ReponseTime);

  sleep(1);
}
