build:
	cd gui && make windows
	cd cmd && make windows
	cd cmd && make linux

cmd:
	cd cmd && make windows
	cd cmd && make linux

gui:
	cd gui && make windows