# Go Data [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GoDoc](https://godoc.org/github.com/micro/go-data?status.svg)](https://godoc.org/github.com/micro/go-data)

Go Data is a pluggable framework for data access.

Note: This is a WIP

## Overview

Data is at the core of most applications. In a lot of cases we just want simple CRUD access. Many ORMs already provide this functionality however it usually ties you to a specific data model (sql) or database system (mysql). Go Data looks to provide a pluggable framework for data access, separating access models from storage, allowing a number of underlying storage mechanisms to be used. We do this using strongly defined interface in Go with a pluggable architecture. We'll look to support a number of data models including CRUD, blob, graph, key-value, sql. 

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

