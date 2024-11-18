
# Dynamic Chat Application in Go

This project is a dynamic chat application built using Go, leveraging goroutines, channels, and RESTful APIs to allow multiple clients to join, send messages, and leave a chat room concurrently. The chat room ensures thread-safe operations and efficient message broadcasting.Designed with MorderViewController architecture


## Installation
Clone the Project from Github

```bash
  git clone https://github.com/guruakashsm/GoatRobotics.git
  cd GoatRobotics
```

#### Pre requisite
```bash
  # If you want to Change Something in Code then you need to have :
   Go lang Installed on your Computer
```

#### Up the Project
> Windows
```bash
chmod +x ./GoatRobotics.exe
./GoatRobotics.exe
```
Linux
```bash
chmod +x ./GoatRobotics
 ./GoatRobotics
```

#### To Take Build
> Windows
```bash
   env GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -v -ldflags=
```
Linux
```bash
   env GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -v -ldflags=
```
This it will clean build suitable for all versions of OS

#### To Change Server Configuration
Change this file to Change the Server Configuration
```bash
./config.json 
```
Feel free to Change the Server Configuration as Per your needs


    
## API Reference

#### 1. Join the Chat Room

```http
  GET http://localhost:8080/rpc/GOATROBOTICS/join?id=12345
```
#### Curl 

```http
 curl --location 'http://localhost:8080/rpc/GOATROBOTICS/join?id=12345'
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `string` | **Required**. To Join the ID is Required |



#### Response 
```json 
{
    "userId": "12345",
    "message": "Joined Chat Successfully",
    "ReponseTime": "2024-11-18T16:59:29.73657231+05:30"
}
```
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `userId` | `string` |  User ID Join to the Chat |
| `message` | `string` |  Status of the Response  |
| `ReponseTime` | `time` |  Time the Response is Generated |

---
---
#### 2. Leave the Chat Room

```http
  GET http://localhost:8080/rpc/GOATROBOTICS/leave?id=12345
```
#### Curl 

```http
 curl --location 'http://localhost:8080/rpc/GOATROBOTICS/leave?id=12345'
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `string` | **Required**. To Left the ID is Required |



#### Response 
```json 
{
    "userId": "12345",
    "message": "Left Chat Successfully",
    "ReponseTime": "2024-11-18T16:59:31.796843686+05:30"
}
```
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `userId` | `string` |  User ID Left to the Chat |
| `message` | `string` |  Status of the Response  |
| `ReponseTime` | `time` |  Time the Response is Generated |

---
---
#### 3. To Send Message to the Chat Room

```http
  GET http://localhost:8080/rpc/GOATROBOTICS/send?id=12345&message=Hello form Guru
```
#### Curl 

```http
curl --location 'http://localhost:8080/rpc/GOATROBOTICS/send?id=12345&message=Hello%20form%20Guru'
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `string` | **Required**. To Message the ID is Required |
| `messgae` | `string` | **Required**. Message to Publish is Required |



#### Response 
```json 
{
    "userId": "12345",
    "message": "Message Sent Successfully",
    "ReponseTime": "2024-11-18T15:27:31.000569117+05:30"
}
```
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `userId` | `string` |  User ID Sent Message to the Chat |
| `message` | `string` |  Status of the Response  |
| `ReponseTime` | `time` |  Time the Response is Generated |

---
---
#### 4. To Get Messages form Chat Room

```http
  GET http://localhost:8080/rpc/GOATROBOTICS/send?id=12345&message=Hello form Guru
```
#### Curl 

```http
curl --location 'http://localhost:8080/rpc/GOATROBOTICS/messages?id=12345'
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `string` | **Required**. To Message the ID is Required |
| `messgae` | `string` | **Required**. Message to Publish is Required |



#### Response 
```json 
{
    "messages": [
        {
            "userId": "12345",
            "message": "Hi, I am Guruakash SM, a Backend Developer with expertise in building robust and scalable systems.",
            "time": "2024-11-18T15:26:52.886606381+05:30"
        },
        {
            "userId": "12345",
            "message": "I have extensive experience in gRPC, REST APIs, and database management with technologies like MySQL and MongoDB.",
            "time": "2024-11-18T15:27:30.423006583+05:30"
        },
        {
            "userId": "12345",
            "message": "I specialize in Go for backend development, and have worked with NATS and JetStream for messaging systems.",
            "time": "2024-11-18T15:27:31.000567667+05:30"
        },
        {
            "userId": "12345",
            "message": "I am passionate about optimizing performance, ensuring code quality, and contributing to impactful projects.",
            "time": "2024-11-18T15:28:00.123456789+05:30"
        }
    ],
    "ReponseTime": "2024-11-18T15:28:02.000000000+05:30",
    "userId": "12345"
}

```
| Parameter     | Type       | Description                                   |
| :------------ | :--------- | :-------------------------------------------- |
| `messages`    | `array`    | **Required**. List of messages.               |
| `messages[].userId` | `string` | **Required**. ID of the Sender.              |
| `messages[].message` | `string` | **Required**. Message content.           |
| `messages[].time`    | `string` | **Required**. Timestamp of the message in ISO 8601 format. |
| `ReponseTime` | `string`   | **Optional**. Time taken to process in ISO 8601 format. |
| `userId`      | `string`   | **Required**. ID of the user.                 |


---
---
#### 5. Ping to get Server Status

```http
  GET http://localhost:8080/ping
```
#### Curl 

```http
curl --location 'http://localhost:8080/ping'
```




#### Response 
```json 
{
    "message": "Pinged Successfully",
    "ReponseTime": "2024-11-18T19:32:09.758424063+05:30"
}
```
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `message` | `string` |  Status of the Response  |
| `ReponseTime` | `time` |  Time the Response is Generated |

---
---
#### 6. Ping to get Server Status

```http
  http://localhost:8080/version
```
#### Curl 

```http
curl --location 'http://localhost:8080/version'
```




#### Response 
```json 
{
    "Version": "v1.0.1",
    "ReponseTime": "2024-11-18T19:34:59.452488906+05:30"
}
```
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `Version` | `string` |  Version of the Server  |
| `ReponseTime` | `time` |  Time the Response is Generated |

---
## UI
Go to this URL for UI
```http
    http://localhost:8080/home
```
## Running Tests

To run Unit tests, run the following command

```bash
 go test ./... -v
```

To run K6 Script, run the following command

```bash
  cd K6_Testing
  chmod +x ./run.sh
```
```bash
  ./run.sh join       //To Check Join API
  ./run.sh leave     //To Check Leave API
  ./run.sh messages //To Check Messages API
  ./run.sh send    //To Check Send API
  ./run.sh all    //To Check ALL API
```
or 

```bash
  k6 run ./join_Test.js       //To Check Join API
  k6 run ./leave_Test.js     //To Check Leave API
  k6 run ./messages_Test.js //To Check Messages API
  k6 run ./send_Test.js    //To Check Send API
```
## Tech Stack

**Client:** HTML,CSS,JS

**Server:** Go

**Testing:** K6

### Lessons Learned

A summary of key concepts and best practices implemented throughout my recent projects. These lessons have contributed to the development of robust, scalable, and efficient systems.

#### 1. Proper Log Statements for Debugging
Clear and structured log statements for effective debugging and troubleshooting. Logs provide useful insights into application behavior, ensuring issues are caught early.
> Logs stored in: `./logs/Audit.audit` & `GOATROBOTICS.log`  📜

#### 2. Thread-Safe Operations & Efficient Message Broadcasting
Ensuring thread-safe operations while maintaining high performance in concurrent environments. Efficient message broadcasting ensures that messages are transmitted seamlessly to multiple clients.
> Optimized for high concurrency 🏎️💨

#### 3. Unit Test Cases & Performance Testing
Developed unit tests and conducted performance testing using K6, simulating multiple clients joining, sending messages, and leaving chat rooms concurrently. This ensures reliability and scalability under load.
> Tested with K6 🏋️‍♂️📊

#### 4. Middleware for Request/Response Interception
Implemented middleware to intercept and log all incoming and outgoing requests and responses. Logs are saved in `./logs/Audit.audit` for traceability.
> Enhanced security and traceability 🔄

#### 5. Logs Stored in `GOATROBOTICS.log`
Detailed audit logs stored in `GOATROBOTICS.log` for further tracking, analysis, and debugging.
> Comprehensive tracking 📁

#### 6. Docker for Dependency Management
Used Docker to create a consistent environment, eliminating dependency issues across different environments and ensuring smooth application execution.
> Dockerized for consistency 🐋

#### 7. Custom Error Types for Scalability
Implemented custom error types to support scalable error handling, ensuring clean, manageable, and maintainable code.
> Structured for growth ⚙️

#### 8. Constants Instead of Hardcoding
Replaced hardcoded values with constants to improve maintainability, reduce errors, and make the code more adaptable to changes.
> Improved code quality 🔧

#### 10. Everything Configurable with `config.json`
All key configurations are easily manageable through a `config.json` file, allowing for flexible and centralized configuration management across environments.  
> Simplifies deployment and environment setup 🛠️

#### 9. Simple UI for User Interaction
While my primary focus is backend development, I created a simple UI to facilitate user interaction. The UI is functional but minimal, as UI design is not my core strength.
> **Note**: I’m a backend developer, not a UI designer! 😅

---

These lessons have shaped my development process, focusing on best practices that ensure scalable, reliable, and efficient software solutions. I continue to refine these skills, aiming to deliver high-quality code in every project.



## Support

For support, email guruakashsm@gmail.com 

