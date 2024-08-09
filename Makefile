build:
	cd gui && make windows
	cd cmd && make windows
	cd cmd && make linux

winAndLinux:
	cd cmd && make windows
	cd cmd && make linux

gui:
	cd gui && make windows