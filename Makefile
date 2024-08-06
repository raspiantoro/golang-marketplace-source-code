GO			:= GO111MODULE=on go
GOTEST		:= $(GO) test -v --count=1 --parallel=1 -p=1
SVC_DIRECTORY 	:= $(addsuffix tests/..., $(sort $(dir $(wildcard ./*svc/))))

cobra:
	cd $(dir) && GOWORK=off cobra-cli init && cd ..

migrate:
	migrate create -ext sql -dir ./$(target)/db/migrations -seq $(name)

test-chapter1:
	$(GOTEST) $(SVC_DIRECTORY) -run Chapter1

test-chapter2:
	$(GOTEST) $(SVC_DIRECTORY) -run Chapter2
