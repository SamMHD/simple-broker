<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# sender

```go
import "github.com/SamMHD/simple-broker/sender"
```

package sender will generate random messages and send them to the receiver over HTTP API over specified rate and duration message size limits and reciever address are defined in the environment variables

## Index

- [Constants](<#constants>)
- [func StartSendProcedure(config util.Config)](<#func-startsendprocedure>)


## Constants

```go
const TPS = 1000 // messages per second
```

## func [StartSendProcedure](<https://github.com/SamMHD/simple-broker/blob/main/sender/main.go#L19>)

```go
func StartSendProcedure(config util.Config)
```

StartSendProcedure will send messages at a rate of TPS for testDuration seconds



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)