# Task-manager backend

A Back-end for "TODO" web application 

Technologies
 * [Docker](https://www.docker.com/): Linux container platform
 * [GO](https://golang.org/): Go programming language, developed by Google
 * [Mux](http://www.gorillatoolkit.org/pkg/mux): Go toolkit for handling routing
 * [MySQL](https://www.mysql.com/): SQL database
 * [Redis](https://redis.io/): In-memory data structure store, used as a database
 * [Google cloud](https://cloud.google.com): Kunerbetes engine
 * [Travis](https://travis-ci.com/): CI-testing, deploy

### Deployment process

 * Push request to Github
 * Travis runs test cases
 * Travis builds docker image
 * Travis uses Google cloud credentials to deploy docker image to Kunerbetes engine.

### Requirement
  * Docker-compose v1.12 or newer (https://docs.docker.com/compose/install/) for development. For production use, .prod compose file is provided, but ideally Swarm, Kubernetes or other orchestration tooling is used
  * Docker v1.12 or newer (https://docs.docker.com/engine/installation/)

### Usage

Run development server on port 8080:

```node module
$ docker-compose up
```
