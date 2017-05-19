# Protobuf3

To compile the protocol buffer definition, run protoc with the --go_out parameter set to the directory you want to output the Go code to.

```bash
$ protoc --go_out=. *.proto
```

It should generate a file with the suffix `.pb.go`. Import the path in the file that uses it.

