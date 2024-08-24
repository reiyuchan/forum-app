BINARY=forum-app
OS=linux windows darwin
WINDOWS=.exe
CURRENT_OS=$(shell uname)

all: start-docker

start-docker:
	@echo -e "\nStarting containers.."
	@if [ -f "docker-compose.yaml" ]; then \
		docker-compose up; \
	else \
		echo "docker-compose.yaml not found..";  \
	fi

build: tidy
	@for o in ${OS}; \
	do	\
		echo "Building $$o..."; \
		if [ "$$o" == "windows" ]; \
		then \
			GOOS=$$o go build -C src -o ../bin/${BINARY}${WINDOWS}; \
		else \
			GOOS=$$o go build -C src -o ../bin/${BINARY}-$$o; \
		fi; \
		echo "done..."; \
	done
	@echo "Build successful.."

run: build
	@echo -e "\nStarting app..."
	@if [ "${CURRENT_OS}" == "Linux" ]; then \
		./bin/${BINARY}-linux; \
	elif [ "${CURRENT_OS}" == "Darwin" ]; then \
		./bin/${BINARY}-darwin; \
	else \
		echo "Need windows machine to run..."; \
	fi

tidy:
	@go -C src mod tidy

clean:
	-rm -r bin
