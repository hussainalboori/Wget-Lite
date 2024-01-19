compilerun:
	go build -o wget main.go
	alias sget="./wget"

clean:
	rm -rf wget
	rm -rf *.jpg
	rm -rf Log.txt
