# Fault-Tolerant Distributed Key-Value Store

## 1. Project Overview

The Fault-Tolerant Distributed Key-Value Store is a distributed system that stores data as key-value pairs across multiple server nodes.

The system uses three nodes to store and retrieve data. When a key-value pair is written to Node 1, the system replicates the data to the other nodes.

This project demonstrates distributed storage, data replication, Docker deployment, and basic node-failure testing.

## 2. Objectives

- To develop a distributed key-value storage system.
- To store and retrieve data using keys and values.
- To replicate data across multiple nodes.
- To deploy server nodes using Docker containers.
- To test the system when one node becomes unavailable.

## 3. Technologies Used

- Programming Language: Go
- Communication: gRPC
- Containerization: Docker
- Container Networking: Docker Network
- Version Control: Git
- Image Registry: Docker Hub

## 4. System Architecture

The system contains three server nodes.

```text
                 Client
                   |
              gRPC Requests
                   |
                Node 1
              /         \
        Replication   Replication
            /             \
        Node 2           Node 3

- 
