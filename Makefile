build-all:
	cd gui && make windows
	cd cmd && make windows
	cd cmd && make linux

build:
	cd cmd && make windows
	cd cmd && make linux

gui:
	cd gui && make windows