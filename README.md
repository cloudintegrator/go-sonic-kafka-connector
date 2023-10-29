### Command

```
go build -gcflags="all=-N -l" -v --buildmode=plugin -o plugin/kafka-connector.so pkg/listener/kafkalistener.go
```
