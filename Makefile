__out=gen-fakecert

all: build

clean:
	rm -f $(__out)

build:
	go build -o $(__out) main.go

