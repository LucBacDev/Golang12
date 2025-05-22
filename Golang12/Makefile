protoc --go_out=. --go-grpc_out=. golang/proto/auth/auth.proto
protoc --go_out=. --go-grpc_out=. golang/proto/transaction/transaction.proto
protoc --go_out=. --go-grpc_out=. golang/proto/wallet/wallet.proto

go run golang/service/authService/main.go golang/service/authService/grpc.go
go run golang/service/transactionService/main.go golang/service/transactionService/grpc.go
go run golang/service/walletService/main.go golang/service/walletService/grpc.go
go run golang/service/userService/main.go 
auth:
GRPC client: 
GRPC server:9000

transaction:
GRPC client:3000
GRPC server:9001
RestAPI server:4000

wallet:
GRPC client 
GRPC server:9002

wallet:
RestAPI client: 8000
GRPC server:9002