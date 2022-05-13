binary  := crawler
compileFile := ./cmd/main.go

mac:
	GOOS=darwin go build -o $(binary) $(compileFile)
	./$(binary) -conf cmd/conf.yml #test block
	#nohup ./$(binary) -conf cmd/conf.yaml &

linux:
	GOOS=linux go build -o $(binary) $(compileFile)
	nohup ./$(binary) -conf cmd/conf.yaml &