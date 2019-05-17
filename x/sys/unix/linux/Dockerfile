FROM ubuntu:18.10

# Dependencies to get the git sources and go binaries
RUN apt-get update && apt-get install -y  --no-install-recommends \
        ca-certificates \
        curl \
        git \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Get the git sources. If not cached, this takes O(5 minutes).
WORKDIR /git
RUN git config --global advice.detachedHead false
<<<<<<< HEAD
# Linux Kernel: Released 22 Oct 2018
RUN git clone --branch v4.19 --depth 1 https://kernel.googlesource.com/pub/scm/linux/kernel/git/torvalds/linux
# GNU C library: Released 01 Aug 2018 (we should try to get a secure way to clone this)
RUN git clone --branch glibc-2.28 --depth 1 git://sourceware.org/git/glibc.git

# Get Go
ENV GOLANG_VERSION 1.11
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 b3fcf280ff86558e0559e185b601c9eade0fd24c900b4c63cd14d1d38613e499
=======
# Linux Kernel: Released 03 Mar 2019
RUN git clone --branch v5.0 --depth 1 https://kernel.googlesource.com/pub/scm/linux/kernel/git/torvalds/linux
# GNU C library: Released 01 Feb 2019 (we should try to get a secure way to clone this)
RUN git clone --branch release/2.29/master --depth 1 git://sourceware.org/git/glibc.git

# Get Go
ENV GOLANG_VERSION 1.12
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 750a07fef8579ae4839458701f4df690e0b20b8bcce33b437e4df89c451b6f13
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a

RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
    && echo "$GOLANG_DOWNLOAD_SHA256  golang.tar.gz" | sha256sum -c - \
    && tar -C /usr/local -xzf golang.tar.gz \
    && rm golang.tar.gz

ENV PATH /usr/local/go/bin:$PATH

# Linux and Glibc build dependencies and emulator
RUN apt-get update && apt-get install -y  --no-install-recommends \
<<<<<<< HEAD
        bison gawk make python \
=======
        bison gawk make python3 \
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
        gcc gcc-multilib \
        gettext texinfo \
        qemu-user \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*
# Cross compilers (install recommended packages to get cross libc-dev)
RUN apt-get update && apt-get install -y \
        gcc-aarch64-linux-gnu       gcc-arm-linux-gnueabi     \
        gcc-mips-linux-gnu          gcc-mips64-linux-gnuabi64 \
        gcc-mips64el-linux-gnuabi64 gcc-mipsel-linux-gnu      \
        gcc-powerpc64-linux-gnu     gcc-powerpc64le-linux-gnu \
	gcc-riscv64-linux-gnu                                 \
        gcc-s390x-linux-gnu         gcc-sparc64-linux-gnu     \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Let the scripts know they are in the docker environment
ENV GOLANG_SYS_BUILD docker
WORKDIR /build
ENTRYPOINT ["go", "run", "linux/mkall.go", "/git/linux", "/git/glibc"]
