run:
	go run main.go

build:
	# remove debug info and build
	go build -ldflags "-s -w"
