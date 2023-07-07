workload=$(shell basename $(shell pwd) | sed 's/^[^-]*-//1')

test:
	# runs maelstrom with the appropriate workload
	go build -o bin/${workload} .
	maelstrom test -w ${workload} --bin bin/${workload} --node-count 1 --time-limit 20 --rate 10
	
	# cleanup
	rm -rf bin store
