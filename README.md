# Go Inflation

One day I was doing breakfast as usual and I was eating a yogurt. While I was eating my yougut I thought: damn, buying groceries is getting really expensive.

Then, I went on the internet to search if somebody had the crazy idea to daily track prices before me. But... nothing.

So I promplty said "fuck it" and created this monstrosity.

### But WHY

You see, we are sorrounded by news of inflation and prices rising everywhere. But is it really perceptible? Can we see day to day at a microscopic level these changes?
To answer this question, I've decided to create this mess, disaster, trainwreck, abomination of a project. Because that's what REAL software engineers do.

### The stack

Made with:

[![Technologies](https://skillicons.dev/icons?i=go,postgresql,mongodb,kafka&theme=light)](https://skillicons.dev)

## Architecture
### General Overview
<img src="https://github.com/user-attachments/assets/dec37207-fa2e-4d98-a7f1-21128da55e33" width=720/>

The underlying architecture of the project relies on multiple microservices. The choice of using a microservice architecture instead of a monolith is due to the needs of scalability.

This extreme need of scalability lies in the aim of this project: register hundres of thounsands (if not milion) of daily prices and then calculate them asynchronously.

### Ingestion Flow
<img src="https://github.com/user-attachments/assets/e1b51e4e-c7b3-4442-bf47-f8cfbc7a2d53" width=720/>

The ingestion flows starts with one or more API scrapers that collect online data. This raw data is then uniformed in order to be passed to the Data Ingestion API, which produces a message on the Kafka Queue.

The Data Ingestion API does little to no elaboration, in order to withstand high numbers of requests per seconds. 

The produced messages are instead elaborated asynchronously by the Handler. A copy of the message is saved on a MongoDB database, which serves both as a requests log and as a backup.

The raw message is processed and transformed into a structured format, which is then stored in a PostgreSQL database.

## TODO
The project is still incomplete. Here is a list of TODOs:
- Implement at least an API Scraper
- Implement saving of elaborated messages
- Implement the Analytics Engine
- Implement the Data API

## License
This project is licensed under the MIT license.
