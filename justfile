list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | xargs

serve:
	~/opt/go_appengine/goapp serve

open:
	open http://localhost:8080

push:
	python2.7 ~/opt/go_appengine/appcfg.py update .

format:
	gofmt -w *.go

lint:
	go vet
	golint
