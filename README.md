# Fly.io Distributed Systems Challenges (Go)

This repo contains my Go implementations of the [Fly.io Distributed Systems Challenges](https://fly.io/dist-sys/). These exercises are run against [Maelstrom](https://github.com/jepsen-io/maelstrom), a test harness for simulating network partitions, failures, and concurrency issues in distributed systems.

## Challenges

- [x] Challenge #1: Echo
- [x] Challenge #2: Unique ID Generation
- [x] Challenge #3a: Single-Node Broadcast
- [x] Challenge #3b: Multi-Node Broadcast
- [ ] Challenge #3c: Fault Tolerant Broadcast
- [ ] Challenge #3d: Efficient Broadcast, Part I
- [ ] Challenge #3e: Efficient Broadcast, Part II
- [ ] Challenge #4: Grow-Only Counter (CRDT)
- [ ] Challenge #5a: Single-Node Kafka-Style Log
- [ ] Challenge #5b: Multi-Node Kafka-Style Log
- [ ] Challenge #5c: Efficient Kafka-Style Log
- [ ] Challenge #6a: Single-Node, Totally-Available Transactions
- [ ] Challenge #6b: Read Uncommitted Transactions
- [ ] Challenge #6c: Read Committed Transactions

## How to Run

```bash
# From within a challenge directory
go build -o main
maelstrom test -w <workload> --bin ./main --node-count <N> --time-limit 20
