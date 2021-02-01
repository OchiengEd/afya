bin:
	GOOS=linux GOARCH=amd64 go build -o build/afya-dashboard ./cmd/dashboard/
	GOOS=linux GOARCH=amd64 go build -o build/afya-reception ./cmd/reception/
	GOOS=linux GOARCH=amd64 go build -o build/afya-initdb ./cmd/initdb/

