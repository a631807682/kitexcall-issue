# kitexcall-issue

## Issue

1. Can not covnert type `sint64`
2. Can not access to multiple services

## Test
>
> go run main.go  
> ./test.sh

## Log

```log
WARNING: sonic only supports Go1.16~1.22 && CPU amd64, but your environment is not suitable
[Status]: Failed
[ServerError]: RPC call failed: remote or network error: protobuf marshal message failed: [CATEGORY 0] convert failed: json convert to protobuf failed, MessageDescriptor: 
[JSON-TO-PROTOBUF] convert failed: sonic decode json bytes failed
[JSON-TO-PROTOBUF] dismatched type: param isn't intType
WARNING: sonic only supports Go1.16~1.22 && CPU amd64, but your environment is not suitable
[Status]: Failed
[ServerError]: RPC call failed: missing method: Method1 in service
WARNING: sonic only supports Go1.16~1.22 && CPU amd64, but your environment is not suitable
[Status]: Failed
[ServerError]: RPC call failed: remote or network error[remote]: unknown method Method2
```
