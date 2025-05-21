protoc --go_out=. --go-grpc_out=. golang/proto/auth/auth.proto
protoc --go_out=. --go-grpc_out=. golang/proto/transaction/transaction.proto
protoc --go_out=. --go-grpc_out=. golang/proto/wallet/wallet.proto
protoc --go_out=. --go-grpc_out=. golang/proto/user/user.proto

go run golang/service/authService/main.go golang/service/authService/grpc.go
go run golang/service/transactionService/main.go golang/service/transactionService/grpc.go
go run golang/service/walletService/main.go golang/service/walletService/grpc.go
go run golang/service/userService/main.go golang/service/userService/grpc.go