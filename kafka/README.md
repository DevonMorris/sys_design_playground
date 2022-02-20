# Apache Kafka

## Description
Apache Kafka is a "distributed event streaming platform." It's main objective
is to increase raw throughput. It treats messaging as a distributed append-only
log and therefore is good for both batch processing and real-time streaming.

## What Problem(s) Does it Solve
* Pub/Sub
* Event Processing
* Work Queues
* Async processes

## What makes it different
* The difference between Kakfa and RMQ in 2021 seem to be diminishing due to RMQ
adding support for append-only log framework

## Notes
* Clients have to keep track of offset in log to read the correct entries
