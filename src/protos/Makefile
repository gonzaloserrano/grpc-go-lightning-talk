default: js go
js-deps:
	# install groc_node_plugin
	npm install grpc-tools
js:
	protoc --js_out=import_style=commonjs,binary:../node --grpc_out=../node --plugin=protoc-gen-grpc=node_modules/grpc-tools/bin/grpc_node_plugin clock.proto
go:
	protoc clock.proto --go_out=plugins=grpc:../go/clock
