setup: build copy

build:
	go build -o interpreter .

copy:
	chmod +x ./interpreter &&\
	cp ./interpreter /usr/local/bin/interpreter