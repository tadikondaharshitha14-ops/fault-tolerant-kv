# Fault-Tolerant Distributed Key-Value Store

A simple distributed key-value store built with **Go**, **gRPC**, and
**Docker**. The project runs multiple server nodes and demonstrates
storing and retrieving key-value data across nodes, along with basic
node-failure testing.

## Project Overview

A key-value store saves data as a pair:

-   **Key:** the name used to identify a value
-   **Value:** the data associated with that key

This project uses three server nodes to demonstrate communication and
replication in a distributed system. A client can send requests to a
node, and the project includes a replication test to check whether
stored data can be read from the other nodes.

## Objectives

-   Build a simple key-value service in Go.
-   Use gRPC for communication between components.
-   Run multiple nodes as Docker containers.
-   Demonstrate replication across three nodes.
-   Observe client behavior when one node is unavailable.
-   Practice containerization and distributed-systems testing.

## Technologies Used

-   **Go** --- application implementation
-   **gRPC** --- remote procedure calls
-   **Docker** --- containerizing and running nodes
-   **Git and GitHub** --- version control and project hosting

## Architecture

``` text
                  Client
                    |
             gRPC requests
                    |
          +---------+---------+
          |         |         |
       Node 1     Node 2     Node 3
          |         |         |
          +---- Replicated ---+
                 key-value
                    data
```

The diagram shows the intended three-node setup. The replication test
checks whether a value written through one node can be retrieved from
the other nodes.

## Main Features

-   **Set:** store a value using a key.
-   **Get:** retrieve a value using a key.
-   **Three-node setup:** run three server containers on a Docker
    network.
-   **Replication test:** verify that a written value is available
    through the nodes.
-   **Failure test:** stop one node and observe which requests continue
    to work; restart the node and test again.

## Project Structure

The main files and directories include:

``` text
fault-tolerant-kv/
├── cmd/
│   ├── client/
│   ├── server/
│   └── replication-test/
├── internal/
│   ├── grpcserver/
│   └── replication/
├── Dockerfile
├── README.md
└── go.mod
```

Exact files may vary as the project evolves.

## Prerequisites

Install the following tools before running the project:

-   Go
-   Docker Engine
-   Git

Check the installations:

``` bash
go version
docker --version
git --version
```

## Run the Tests

From the project directory, run:

``` bash
go test ./...
```

The tests should complete successfully before running the
container-based demonstration.

## Run with Docker

The project image published for this demonstration is:

**Docker Hub:** https://hub.docker.com/r/harshi142003/fault-tolerant-kv

Pull version `1.0`:

``` bash
docker pull harshi142003/fault-tolerant-kv:1.0
```

The project was tested with three containers named `kv-node1`,
`kv-node2`, and `kv-node3` connected to a Docker network named
`kv-network`. Use the Docker run/network configuration from the project
setup to start the nodes; container arguments and port mappings must
match the server configuration in the source code.

To check running containers:

``` bash
docker ps
```

## Verification Results

The replication test produced the following result:

``` text
Node 1 SET: true
Node 1 GET: harshitha FOUND: true
Node 2 GET: harshitha FOUND: true
Node 3 GET: harshitha FOUND: true
```

In the failure test, Node 2 was stopped. Requests to Node 2 reported it
as unavailable, while reads through Node 1 and Node 3 continued to work.
After Node 2 was restarted, the value could again be retrieved from all
three nodes.

These results demonstrate the tested replication and basic
node-unavailability behavior in this setup.

## Limitations

This is an educational project and should not be considered a
production-ready distributed database. The demonstrated tests do **not**
establish consensus, automatic leader election, durable storage, or
complete recovery guarantees. Behavior during failures depends on the
implementation and the way the nodes are configured.

## Future Improvements

-   Add persistent storage so data survives container restarts.
-   Add health checks and automatic node recovery.
-   Implement a clearly defined consistency and conflict-resolution
    strategy.
-   Add automated integration tests for node failures.
-   Add authentication, authorization, logging, and metrics.
-   Document exact Docker commands and port mappings for reproducible
    setup.

## Author

**Harshitha Tadikonda**

## License

No license has been specified yet. Add a `LICENSE` file if you intend to
publish the project for reuse by others.
