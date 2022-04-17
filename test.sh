go test ./usecase/... -coverprofile=cover.out && go tool cover -html=cover.out
# go test ./usecase/... -v -cover