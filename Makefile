tmp_folder=./tmp

serve:
	air server --port 3000
build:
	go build -o ${tmp_folder}/main_prod main.go
run_build:
	${tmp_folder}/main_prod 
clean:
	rm ${tmp_folder}/main