protogen:
	protoc \
	--go_out=./ \
	--go_opt=module=github.com/atlokeshsk/micro_blog \
	--go-grpc_out=./ \
	--go-grpc_opt=module=github.com/atlokeshsk/micro_blog \
	./proto/*/*

clean:
	rm -rf protogen