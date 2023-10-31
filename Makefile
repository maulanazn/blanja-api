tmp_folder=./tmp

dev:
	air server --port 3000
start:
	go run main.go
build:
	go build -o ${tmp_folder}/main_prod main.go
run_build:
	${tmp_folder}/main_prod 
clean:
	rm ${tmp_folder}/main