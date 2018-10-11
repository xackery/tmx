PROTO_FILES=$(shell find pb -name '*.proto')
.PHONY: proto
proto:
	@protoc \
	-I. \
	$(PROTO_FILES) \
	--go_out=.
