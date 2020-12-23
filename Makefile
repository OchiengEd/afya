bin:
	GOOS=linux GOARCH=amd64 go build -o build/afya-auth ./cmd/auth/
	GOOS=linux GOARCH=amd64 go build -o build/afya-reception ./cmd/reception/

