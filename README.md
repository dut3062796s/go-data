# Go Data [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GoDoc](https://godoc.org/github.com/micro/go-data?status.svg)](https://godoc.org/github.com/micro/go-data)

Go Data is a pluggable framework for data access.

Note: This is a WIP

## Overview

Data is at the core of most applications. In many situations we want a simplified CRUD experience—and many ORMs provide this experience for different kinds of databases. What matters most is being able to define the domain model without spending valuable time writing and configuring how these domain models are stored and retrieved from a particular variant of datastore. Also, there are needs for more trait-specific data management capabilities, in addition to basic CRUD repositories.

Go Data looks to solve the problem of separating your domain model from the underlying specifics of your choice of database. The first and only thing you should worry about when writing a microservice that stores data, is creating and wiring together the contract of your domain, without worrying about coupling the application to a specific persistence layer. 

Go Data takes on the Go Micro approach and will include code generation for CRUD in future.

## Architecture

The diagram below shows initial design ideas. This may change. We create packages and interfaces for each kind of data model and layer the abstractions. In time this will include models like blob, graph, document, etc.

![https://imgur.com/siZFJL1.png](https://imgur.com/siZFJL1.png)

## Goals 

Go Data will provide common abstractions for data access defined as Go interfaces. The primary goal is to separate business logic CRUD from storage/infrastructure. This is done via strongly defined interfaces with a pluggable architecture.

* Underlying storage should not be tightly coupled to any business logic or domain model.
* Interfaces provide an abstraction for specific data models e.g kv, graph, document, blob.
* Auth info should be provided as env variables, flags or pulled from Go Config.

## Contributing

Go Data is a WIP. Contributions welcome.

## TODO 

- [ ] explore the concept of a Model interface
- [ ] individual packages for data models {kv, sql, graph, document}
- [ ] top level interfaces like CRUD and Stream
- [ ] implementations of each package
- [ ] realtime data interface?
- [ ] timeseries interface?
- [ ] search interface? e.g elasticsearch, lucene

