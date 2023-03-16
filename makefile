
all : get linux osx
get :
	go get
linux :
	GOOS=linux GOARCH=amd64 go build -o timeseries_linux_amd64

osx :
	GOOS=darwin GOARCH=arm64 go build -o timeseries_osx
clean :
	rm timeseries_linux_amd64 timeseries_osx

