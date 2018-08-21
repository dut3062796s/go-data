# Go Data

Go Data is a pluggable framework for data access

## Overview

Data access is at the core of most applications. In many instances we want a simplified CRUD experience. 
In a lot of those cases we don't even want to have to write the CRUD aspect. At times we may have a more 
low level requirement for graph, doocument or other types of data model. Go Data looks to solve these problems 
in a single place with a pluggable architecture.

Go Data takes on the Go Micro approach and will include code generation for CRUD in future.

Note: WIP. Much of this will break and change.

## TODO 

- [ ] explore the concept of a Model interface
- [ ] individual packages for data models {kv, sql, graph, document}
- [ ] top level interfaces like CRUD and Stream
- [ ] implementations of each package
- [ ] realtime data interface?
- [ ] timeseries interface?
- [ ] search interface? e.g elasticsearch, lucene
