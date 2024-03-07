build:
	go build -o bin/parsetext cmd/main.go
	
run:
	bin/parsetext -file-input ~/Yandex.Disk/synchro/radio/md/mdata-0.1/radios.csv -path-output ~/Yandex.Disk/synchro/radio/md/mdata-0.1/ 
