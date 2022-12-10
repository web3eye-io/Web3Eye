PROTOC=$(shell which protoc)
PROTOC-GEN-GRPC-GATEWAY=$(shell which protoc-gen-grpc-gateway)

PROTO_FILE = $(shell find cybertracer -name "*.proto")
PROTO_GW_FILE = $(shell grep -l 'google.api.http' -r cybertracer)
PROTO_GO_FILE = $(patsubst %.proto,%.pb.go,$(PROTO_FILE))
PROTO_GO_GW_FILE = $(patsubst %.proto,%.pb.gw.go,$(PROTO_GW_FILE))
PROTO_OPENAPIV2_FILE = $(patsubst %.proto,%.openapiv2.json,$(PROTO_FILE))

PROTO_INCLUDE += -I.
PROTO_INCLUDE += -I./google
PROTO_INCLUDE += -I./cybertracer

PB_FILE = $(shell find cybertracer -name "*.pb.go")

proto: $(PROTO_GO_FILE) $(PROTO_TS_FILE) $(PROTO_GO_GW_FILE) $(PROTO_OPENAPIV2_FILE)
%.pb.go: %.proto
	$(PROTOC) $(PROTO_INCLUDE) \
		--go_out=. \
		--go_opt paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative $< && \
		protoc-go-inject-tag -input=$@

%.pb.gw.go: %.proto
	$(PROTOC) $(PROTO_INCLUDE) $< --plugin=protoc-gen-grpc-gateway=$(PROTOC-GEN-GRPC-GATEWAY) \
		--grpc-gateway_out=logtostderr=true:. \
		--grpc-gateway_opt paths=source_relative

%.openapiv2.json: %.proto
	$(PROTOC) $(PROTO_INCLUDE) \
		--openapiv2_out . \
		--openapiv2_opt logtostderr=true \
		--openapiv2_opt json_names_for_fields=false $<

clean:
	find ./ -name "*.pb.go" | xargs rm -rf
	find ./ -name "*.pb.gw.go" | xargs rm -rf
	find ./ -name "*.md" | grep -v README.md | xargs rm -rf
	find ./ -name "*.json" | xargs rm -rf
	find ./ -name "*.ts" | xargs rm -rf