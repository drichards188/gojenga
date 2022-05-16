# Gojenga
## Serverless Banking Portfolio Project

### Made by David Richards
### https://www.linkedin.com/in/drichards188/

## **WIP (Work In Progress)**

## **Early Alpha**

## **Not Production Ready**

## Technologies
-Golang 1.18

-Aws (DynamoDB)

-OpenTelemetry Distributed Tracing

-Zap logging

## **What this project does**<br>
This project creates a http server meant for ECS. Users can create accounts, make deposits and payments to
other users. They can also see their account information and balance as well as delete their account. The storage solution 
is dynamoDB on AWS. I'm working on converting the hashing functionality from mongo to dynamo. Transactions will flow 
at top speed and only log with a hashed ledger, so it is not a bottleneck

## **Design Decisions**<br>
Golang was selected for duck type interfaces, channels and context propagation.

DynamoDB was selected for its open schema and low cost.

OpenTelemetry is vital for debugging the system at scale by braiding the three pillars together.

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
-Install AWS CLI and setup credentials

-configure OpenTelemetry collector to AWS

-clone project

-enter sqs url on lambda folder

-enter directory

-build and run gojenga/main/main.go

-or run the Dockerfile and deploy to ECS (pre alpha)

## **Run Instructions**<br>

-build and run gojenga/main/main.go

-or run the Dockerfile and deploy to ECS (pre alpha)

-be sure to have AWS CLI and your credentials setup

-adding API documentation soon
