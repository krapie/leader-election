# etcd leader election

Leader election using etcd leader election API.

## Prerequisites

Run etcd docker container using the following command.

```bash
docker run -it --env ALLOW_NONE_AUTHENTICATION=yes -p 2379:2379 -p 2380:2380  bitnami/etcd:latest
```

## Usage

Run this command on seperate terminals to see the leader election in action.

```bash
go run main.go -name first
go run main.go -name second
...
```

## Reference

[Leader Election in Go with Etcd](https://medium.com/@felipedutratine/leader-election-in-go-with-etcd-2ca8f3876d79)