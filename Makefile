build:
	go build -o ~/.local/bin/gogem cmd/main.go

rebuild:
	rm ~/.local/bin/gogem &>/dev/null
	go build -o ~/.local/bin/gogem cmd/main.go

clean:
	rm ~/.local/bin/gogem
