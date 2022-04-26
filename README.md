# Gojenga
## Pseudo Blockchain Portfolio Project

### Made by David Richards
### Breaking into Golang industry
### https://www.linkedin.com/in/drichards188/

## **WIP (Work In Progress)**

## **Early Alpha**

## **Not Production Ready**

## Technologies
-Golang 1.18

-MongoDB

-OpenTelemetry Distributed Tracing

-Zap logging

## **What this project does**<br>
This project creates a http server. The server handles a REST API. Users can create accounts, make deposits and payments to
other users. They can also see their account information and balance as well as delete their account. The storage solution 
is mongoDB.

There is a React/Redux front end to interact with the system on AWS Amplify. For local development I use Postman to interact with the system.
The front end is not my focus, but ideally I want to not have to bother a front end dev for everything at work. We work so close together
it is good to walk in their shoes and help when I can.

## **Design Decisions**<br>
Golang was selected for duck type interfaces, channels and context propagation.

MongoDB was selected because I am still changing the schema. The format is great for prototyping. At one point SQL was
the storage solution and will be again at completion for its speed.

OpenTelemetry is implemented for my desire to break this project into microservices and onto AWS to scale horizontally. Traces
can be inspected using Jaeger for a requests complete lifecycle.

## **Roadmap**<br>
[1] - Making the system Zillow production compliant

-Comments

-Timeouts

-Graceful Shutdown

-Instrumentation

-Distributed Tracing

-Context Propagation

-Concurrency

-Circuit Breaker

-Error Handling

https://www.youtube.com/watch?v=9Q1RMueVHAg

[2] - Adhering to Google Golang best practices

-naming conventions, formatting, flow, swagger comments etc...

https://www.youtube.com/watch?v=EXrEd1-GZR0

[3] - Convert to microservices architecture and deploy on AWS

https://aws.amazon.com/lambda/

https://aws.amazon.com/ecr/

## **Install Instructions**<br>

-clone project

-enter directory

-build and run directory

## **Run Instructions**<br>

-standard run and build

-adding config options soon