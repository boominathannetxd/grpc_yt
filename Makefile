create:	
	protoc --proto_path=proto proto/*.proto --go_out=gen/proto
	protoc --proto_path=proto proto/*.proto --go-grpc_out=gen/proto 


clean:
	rm -f gen/proto/*.go


#curl https://raw.githubusercontent.com/googleapis/googleapis/googleapis/master/google/api/http.proto > google/api/annotaions.proto
#curl https://raw.githubusercontent.com/googleapis/googleapis/googleapis/master/google/api/http.proto > google/api/http.proto