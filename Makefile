compilerun:
	go build -o wget main.go
	alias sget="./wget"

clean:
	rm -rf wget
	rm -rf *.jpg
	rm -rf *.png
	rm -rf *.jsm
	rm -rf *.css
	rm -rf *.jpe
	rm -rf Wget-light-log.txt
