# Go Data [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GoDoc](https://godoc.org/github.com/micro/go-data?status.svg)](https://godoc.org/github.com/micro/go-data)

Go Data is a pluggable framework for data access.

Note: This is a WIP

## Overview

Data access is at the core of most applications. In many situations we want a simplified CRUD experience—and many ORMs provide this experience for different kinds of databases. What matters most is being able to define the domain model without spending valuable time writing and configuring how these domain models are stored and retrieved from a particular variant of datastore. Also, there are needs for more trait-specific data management capabilities, in addition to basic CRUD repositories.

Go Data looks to solve the problem of separating your domain model from the underlying specifics of your choice of database. The first and only thing you should worry about when writing a microservice that stores data, is creating and wiring together the contract of your domain, without worrying about coupling the application to a specific persistence layer. 

Go Data takes on the Go Micro approach and will include code generation for CRUD in future.

## Architecture

The architecture diagram below shows an initial design philosophy for how Go Data will create layers of abstraction for different kinds of middleware that handle the persistence and querying of data. Over time, we intend to expand this diagram to cover more than just basic types of storage, but also different types of cloud-based stores and event stores.

![https://imgur.com/siZFJL1.png](https://imgur.com/siZFJL1.png)

## Goals 

Go Data is a framework library that provides standard repository-like abstractions that decouples your application from the underlying storage mechanism. The primary goal of this project is to separate the layer of your business logic from the underlying implementation of your storage and infrastructure. The secondary goal is to support the numerous ORM and OGM plugins in the Go ecosystem using a standard set of conventions.

* Your underlying storage implementation should not be tightly coupled to any business logic or domain model.
* Repositories are provided as abstract interfaces that provide varying grades of trait-specific capabilities of your data store.
* Database credentials should be provided as environment variables or fetched from a configuration server that uses strong encryption.

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

