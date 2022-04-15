# go-grpc-starter-kit

A starting point project to create `todo`-service

### `MySQL`

``` $ docker run --name mysql -e MYSQL_ROOT_PASSWORD=root -d mysql:latest ```  
``` $ mysql -uroot -proot  ```  

### `PostgreSQL`

``` $ docker run --name postgres -e POSTGRES_PASSWORD=root -e POSTGRES_USER=root -d postgres:latest ```  
``` $ cd $GOPATH/src/github.com/pepeunlimited/go-grpc-starter-kit/deployments/ ```  
``` $ docker-compose up -d ```  
``` $ psql ```  

## Go Directories

### `/build`

### `/api`
``` $ brew install protobuf ```  
``` $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest ```  
``` $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest ```  

### `/cmd`


#### `/worker`
- [`debug/SAGA_TASK_QUEUE`](http://localhost:8088/namespaces/default/task-queues/SAGA_TASK_QUEUE)

### `/deployments`
Deployments folder contains files for the `k8s` environment; `service`, `deployment` and `ingress`. `dev` folder have an additional `external-mysql-service.yaml` file which allow access from local cluster to external MySQL database.  

`docker-compose up` spin up the MySQL database for local development 

### `/init`

### `/internal`

#### `/ent`
Speed up implementing the database access using [`ent`](https://github.com/facebookincubator/ent).
Of course you can implement repositories with a raw sql statements,
but it is very time consuming and boring repeat x10 times same CRUD functions.

#### [`ent`](https://github.com/facebookincubator/ent)
``` $ cd $PROJECT_ROOT/internal/pkg/ ```  
``` $ go generate ./ent ```  

#### [`gRPC`](https://grpc.io/docs/languages/go/quickstart/)

-   [`twriptest`](https://github.com/twitchtv/twirp/tree/master/internal/twirptest)
-   [`gRPC in 3 minutes (Go)`](https://go.googlesource.com/grpc-review/+/refs/heads/PR/570/examples/README.md)

``` 
$ protoc --go_out=.\ 
		 --go-grpc_out=. \
		  todo.proto
```   

#### [`saga`](https://github.com/temporalio/temporal)

##### local development
``` $ git submodule update --init ```  
``` $ cd $PROJECT_ROOT/third_party/temporal-server-docker-compose ```  
``` $ docker-compose up ```  