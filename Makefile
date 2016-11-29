test:
	go test -v

cover:
	rm -rf *.coverprofile
	go test -coverprofile=recover.coverprofile
	gover
	go tool cover -html=recover.coverprofile