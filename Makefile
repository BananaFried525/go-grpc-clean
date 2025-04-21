PROTO_DIR=proto
OUT_DIR=internal/pb

gen:
	protoc \
		-I$(PROTO_DIR) \
		--go_out=paths=source_relative:$(OUT_DIR) \
		--go-grpc_out=paths=source_relative:$(OUT_DIR) \
		$(PROTO_DIR)/*.proto