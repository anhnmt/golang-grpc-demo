

go.protoc:
	protoc --go_out=plugins=grpc:. user.proto \
  	--go_opt=Muser.proto=.