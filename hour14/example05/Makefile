all:
	if [ -n  "$(shell gofmt -l .)" ]; then \
		echo 1>&2 'the following files to be formated:'; \
		gofmt -l . ;\
		exit 1; \
	fi
