create:	
	protoc --proto_path=proto proto/*.proto --go_out=gen/
	protoc --proto_path=proto proto/*.proto --go-grpc_out=gen/
	protoc --grpc-gateway_out=./gen/ \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    --plugin=protoc-gen-grpc-gateway=/home/boominathan/go/bin/protoc-gen-grpc-gateway \
    proto/test.proto
	echo "------Buffer,Grpc & Gateway files created--------"
	
	
	

clean:
	rm -f gen/proto/*.go
	echo "Buffer & Grpc Files deleted succesfully"

#https://github.com/googleapis/googleapis/blob/master/google/api/annotations.proto
#https://github.com/googleapis/googleapis/blob/master/google/api/http.proto

