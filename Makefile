build: 
	@go build -o rp

clean:
	rm -f rp

install: build
	cp rp /usr/local/bin	

uninstall: 
	rm -f /usr/local/bin/rp
