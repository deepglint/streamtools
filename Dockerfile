FROM google/golang
WORKDIR /gopath/src/github.com/deepglint/streamtools
ADD . /gopath/src/github.com/deepglint/streamtools
RUN make clean
RUN make
RUN ["mkdir", "-p", "/gopath/bin"]
RUN ["ln", "-s", "/gopath/src/github.com/deepglint/streamtools/build/st", "/gopath/bin/st"]
