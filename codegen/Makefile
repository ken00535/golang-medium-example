OUT_DIR = bin
ifndef FILES
	FILES = $(shell ls -d cmd/*/ | cut -d/ -f2)
endif
FILES_OUT = $(addprefix ${OUT_DIR}/,${FILES})
UNAME_M := $(shell uname -m)

ifeq ($(OS),Windows_NT)
	PLATFORM ?= windows
	DEST ?= windows
else ifeq ($(UNAME_M),x86_64)
	PLATFORM ?= linux
	DEST ?= linux
else ifeq ($(UNAME_M),armv7l)
	PLATFORM ?= linux
	DEST ?= arm
endif

ARCH =

all:
ifeq ($(OS),Windows_NT)
	make win
else ifeq ($(UNAME_M),x86_64)
	make linux
else ifeq ($(UNAME_M),armv7l)
	make arm32
endif

linux: PLATFORM = linux
linux: ${OUT_DIR} ${FILES_OUT}

win: PLATFORM = windows
win: ${OUT_DIR} ${FILES_OUT:=.exe}

arm32: PLATFORM = linux
arm32: ARCH = arm
arm32: ${OUT_DIR} ${FILES_OUT}

.FORCE:
${OUT_DIR}/%: .FORCE
	go generate ./...
	@echo compiling $(@)...
	GOOS=$(PLATFORM) GOARCH=$(ARCH) go build -o $(@) -tags $(PLATFORM),$(TAGS) ./cmd/$(basename ${@F})

clean:
	rm ${OUT_DIR} -rf

${OUT_DIR}:
	@echo create output dir...
	@mkdir ${OUT_DIR}

.PHONY: all clean win linux arm32 .FORCE
