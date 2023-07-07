workload=$(shell basename $(shell pwd) | sed 's/^[^-]*-//1')

test:
	# runs maelstrom with the appropriate workload
	go build -o bin/${workload} .
	maelstrom test -w ${workload} --bin bin/${workload} --time-limit 30 --rate 1000 --node-count 3 --availability total --nemesis partition
	
	# cleanup
	rm -rf bin store
