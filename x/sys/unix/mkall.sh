#!/usr/bin/env bash
# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# This script runs or (given -n) prints suggested commands to generate files for
# the Architecture/OS specified by the GOARCH and GOOS environment variables.
# See README.md for more information about how the build system works.

GOOSARCH="${GOOS}_${GOARCH}"

# defaults
mksyscall="go run mksyscall.go"
mkerrors="./mkerrors.sh"
zerrors="zerrors_$GOOSARCH.go"
mksysctl=""
zsysctl="zsysctl_$GOOSARCH.go"
mksysnum=
mktypes=
mkasm=
run="sh"
cmd=""

case "$1" in
-syscalls)
	for i in zsyscall*go
	do
		# Run the command line that appears in the first line
		# of the generated file to regenerate it.
		sed 1q $i | sed 's;^// ;;' | sh > _$i && gofmt < _$i > $i
		rm _$i
	done
	exit 0
	;;
-n)
	run="cat"
	cmd="echo"
	shift
esac

case "$#" in
0)
	;;
*)
	echo 'usage: mkall.sh [-n]' 1>&2
	exit 2
esac

<<<<<<< HEAD
if [[ "$GOOS" = "linux" ]] && [[ "$GOARCH" != "sparc64" ]]; then
	# Use then new build system
=======
if [[ "$GOOS" = "linux" ]]; then
	# Use the Docker-based build system
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	# Files generated through docker (use $cmd so you can Ctl-C the build or run)
	$cmd docker build --tag generate:$GOOS $GOOS
	$cmd docker run --interactive --tty --volume $(dirname "$(readlink -f "$0")"):/build generate:$GOOS
	exit
fi

GOOSARCH_in=syscall_$GOOSARCH.go
case "$GOOSARCH" in
_* | *_ | _)
	echo 'undefined $GOOS_$GOARCH:' "$GOOSARCH" 1>&2
	exit 1
	;;
aix_ppc)
	mkerrors="$mkerrors -maix32"
<<<<<<< HEAD
	mksyscall="./mksyscall_aix_ppc.pl -aix"
=======
	mksyscall="go run mksyscall_aix_ppc.go -aix"
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	mktypes="GOARCH=$GOARCH go tool cgo -godefs"
	;;
aix_ppc64)
	mkerrors="$mkerrors -maix64"
<<<<<<< HEAD
	mksyscall="./mksyscall_aix_ppc64.pl -aix"
=======
	mksyscall="go run mksyscall_aix_ppc64.go -aix"
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	mktypes="GOARCH=$GOARCH go tool cgo -godefs"
	;;
darwin_386)
	mkerrors="$mkerrors -m32"
	mksyscall="go run mksyscall.go -l32"
<<<<<<< HEAD
	mksysnum="./mksysnum_darwin.pl $(xcrun --show-sdk-path --sdk macosx)/usr/include/sys/syscall.h"
=======
	mksysnum="go run mksysnum.go $(xcrun --show-sdk-path --sdk macosx)/usr/include/sys/syscall.h"
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	mktypes="GOARCH=$GOARCH go tool cgo -godefs"
	mkasm="go run mkasm_darwin.go"
	;;
darwin_amd64)
	mkerrors="$mkerrors -m64"
	mksysnum="go run mksysnum.go $(xcrun --show-sdk-path --sdk macosx)/usr/include/sys/syscall.h"
	mktypes="GOARCH=$GOARCH go tool cgo -godefs"
	mkasm="go run mkasm_darwin.go"
	;;
darwin_arm)
	mkerrors="$mkerrors"
<<<<<<< HEAD
	mksysnum="./mksysnum_darwin.pl $(xcrun --show-sdk-path --sdk iphoneos)/usr/include/sys/syscall.h"
=======
	mksyscall="go run mksyscall.go -l32"
	mksysnum="go run mksysnum.go $(xcrun --show-sdk-path --sdk iphoneos)/usr/include/sys/syscall.h"
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	mktypes="GOARCH=$GOARCH go tool cgo -godefs"
	mkasm="go run mkasm_darwin.go"
	;;
darwin_arm64)
	mkerrors="$mkerrors -m64"
<<<<<<< HEAD
	mksysnum="./mksysnum_darwin.pl $(xcrun --show-sdk-path --sdk iphoneos)/usr/include/sys/syscall.h"
	mktypes="GOARCH=$GOARCH go tool cgo -godefs"
=======
	mksysnum="go run mksysnum.go $(xcrun --show-sdk-path --sdk iphoneos)/usr/include/sys/syscall.h"
	mktypes="GOARCH=$GOARCH go tool cgo -godefs"
	mkasm="go run mkasm_darwin.go"
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	;;
dragonfly_amd64)
	mkerrors="$mkerrors -m64"
	mksyscall="go run mksyscall.go -dragonfly"
<<<<<<< HEAD
	mksysnum="curl -s 'http://gitweb.dragonflybsd.org/dragonfly.git/blob_plain/HEAD:/sys/kern/syscalls.master' | ./mksysnum_dragonfly.pl"
=======
	mksysnum="go run mksysnum.go 'https://gitweb.dragonflybsd.org/dragonfly.git/blob_plain/HEAD:/sys/kern/syscalls.master'"
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	mktypes="GOARCH=$GOARCH go tool cgo -godefs"
	;;
freebsd_386)
	mkerrors="$mkerrors -m32"
	mksyscall="go run mksyscall.go -l32"
<<<<<<< HEAD
	mksysnum="curl -s 'http://svn.freebsd.org/base/stable/10/sys/kern/syscalls.master' | ./mksysnum_freebsd.pl"
=======
	mksysnum="go run mksysnum.go 'https://svn.freebsd.org/base/stable/10/sys/kern/syscalls.master'"
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	mktypes="GOARCH=$GOARCH go tool cgo -godefs"
	;;
freebsd_amd64)
	mkerrors="$mkerrors -m64"
	mksysnum="go run mksysnum.go 'https://svn.freebsd.org/base/stable/10/sys/kern/syscalls.master'"
	mktypes="GOARCH=$GOARCH go tool cgo -godefs"
	;;
freebsd_arm)
	mkerrors="$mkerrors"
	mksyscall="go run mksyscall.go -l32 -arm"
<<<<<<< HEAD
	mksysnum="curl -s 'http://svn.freebsd.org/base/stable/10/sys/kern/syscalls.master' | ./mksysnum_freebsd.pl"
	# Let the type of C char be signed for making the bare syscall
	# API consistent across platforms.
=======
	mksysnum="go run mksysnum.go 'https://svn.freebsd.org/base/stable/10/sys/kern/syscalls.master'"
	# Let the type of C char be signed for making the bare syscall
	# API consistent across platforms.
	mktypes="GOARCH=$GOARCH go tool cgo -godefs -- -fsigned-char"
	;;
freebsd_arm64)
	mkerrors="$mkerrors -m64"
	mksysnum="go run mksysnum.go 'https://svn.freebsd.org/base/stable/10/sys/kern/syscalls.master'"
	mktypes="GOARCH=$GOARCH go tool cgo -godefs"
	;;
netbsd_386)
	mkerrors="$mkerrors -m32"
	mksyscall="go run mksyscall.go -l32 -netbsd"
	mksysnum="go run mksysnum.go 'http://cvsweb.netbsd.org/bsdweb.cgi/~checkout~/src/sys/kern/syscalls.master'"
	mktypes="GOARCH=$GOARCH go tool cgo -godefs"
	;;
netbsd_amd64)
	mkerrors="$mkerrors -m64"
	mksyscall="go run mksyscall.go -netbsd"
	mksysnum="go run mksysnum.go 'http://cvsweb.netbsd.org/bsdweb.cgi/~checkout~/src/sys/kern/syscalls.master'"
	mktypes="GOARCH=$GOARCH go tool cgo -godefs"
	;;
netbsd_arm)
	mkerrors="$mkerrors"
	mksyscall="go run mksyscall.go -l32 -netbsd -arm"
	mksysnum="go run mksysnum.go 'http://cvsweb.netbsd.org/bsdweb.cgi/~checkout~/src/sys/kern/syscalls.master'"
	# Let the type of C char be signed for making the bare syscall
	# API consistent across platforms.
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	mktypes="GOARCH=$GOARCH go tool cgo -godefs -- -fsigned-char"
	;;
netbsd_arm64)
	mkerrors="$mkerrors -m64"
	mksyscall="go run mksyscall.go -netbsd"
	mksysnum="go run mksysnum.go 'http://cvsweb.netbsd.org/bsdweb.cgi/~checkout~/src/sys/kern/syscalls.master'"
	mktypes="GOARCH=$GOARCH go tool cgo -godefs"
	;;
openbsd_386)
	mkerrors="$mkerrors -m32"
<<<<<<< HEAD
	mksyscall="go run mksyscall.go -l32 -netbsd"
	mksysnum="curl -s 'http://cvsweb.netbsd.org/bsdweb.cgi/~checkout~/src/sys/kern/syscalls.master' | ./mksysnum_netbsd.pl"
=======
	mksyscall="go run mksyscall.go -l32 -openbsd"
	mksysctl="go run mksysctl_openbsd.go"
	mksysnum="go run mksysnum.go 'https://cvsweb.openbsd.org/cgi-bin/cvsweb/~checkout~/src/sys/kern/syscalls.master'"
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	mktypes="GOARCH=$GOARCH go tool cgo -godefs"
	;;
openbsd_amd64)
	mkerrors="$mkerrors -m64"
<<<<<<< HEAD
	mksyscall="go run mksyscall.go -netbsd"
	mksysnum="curl -s 'http://cvsweb.netbsd.org/bsdweb.cgi/~checkout~/src/sys/kern/syscalls.master' | ./mksysnum_netbsd.pl"
	mktypes="GOARCH=$GOARCH go tool cgo -godefs"
	;;
netbsd_arm)
	mkerrors="$mkerrors"
	mksyscall="go run mksyscall.go -l32 -netbsd -arm"
	mksysnum="curl -s 'http://cvsweb.netbsd.org/bsdweb.cgi/~checkout~/src/sys/kern/syscalls.master' | ./mksysnum_netbsd.pl"
	# Let the type of C char be signed for making the bare syscall
	# API consistent across platforms.
	mktypes="GOARCH=$GOARCH go tool cgo -godefs -- -fsigned-char"
	;;
openbsd_386)
	mkerrors="$mkerrors -m32"
	mksyscall="go run mksyscall.go -l32 -openbsd"
	mksysctl="./mksysctl_openbsd.pl"
	mksysnum="curl -s 'http://cvsweb.openbsd.org/cgi-bin/cvsweb/~checkout~/src/sys/kern/syscalls.master' | ./mksysnum_openbsd.pl"
	mktypes="GOARCH=$GOARCH go tool cgo -godefs"
=======
	mksyscall="go run mksyscall.go -openbsd"
	mksysctl="go run mksysctl_openbsd.go"
	mksysnum="go run mksysnum.go 'https://cvsweb.openbsd.org/cgi-bin/cvsweb/~checkout~/src/sys/kern/syscalls.master'"
	mktypes="GOARCH=$GOARCH go tool cgo -godefs"
	;;
openbsd_arm)
	mkerrors="$mkerrors"
	mksyscall="go run mksyscall.go -l32 -openbsd -arm"
	mksysctl="go run mksysctl_openbsd.go"
	mksysnum="go run mksysnum.go 'https://cvsweb.openbsd.org/cgi-bin/cvsweb/~checkout~/src/sys/kern/syscalls.master'"
	# Let the type of C char be signed for making the bare syscall
	# API consistent across platforms.
	mktypes="GOARCH=$GOARCH go tool cgo -godefs -- -fsigned-char"
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	;;
openbsd_arm64)
	mkerrors="$mkerrors -m64"
	mksyscall="go run mksyscall.go -openbsd"
<<<<<<< HEAD
	mksysctl="./mksysctl_openbsd.pl"
	mksysnum="curl -s 'http://cvsweb.openbsd.org/cgi-bin/cvsweb/~checkout~/src/sys/kern/syscalls.master' | ./mksysnum_openbsd.pl"
	mktypes="GOARCH=$GOARCH go tool cgo -godefs"
=======
	mksysctl="go run mksysctl_openbsd.go"
	mksysnum="go run mksysnum.go 'https://cvsweb.openbsd.org/cgi-bin/cvsweb/~checkout~/src/sys/kern/syscalls.master'"
	# Let the type of C char be signed for making the bare syscall
	# API consistent across platforms.
	mktypes="GOARCH=$GOARCH go tool cgo -godefs -- -fsigned-char"
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	;;
openbsd_arm)
	mkerrors="$mkerrors"
	mksyscall="go run mksyscall.go -l32 -openbsd -arm"
	mksysctl="./mksysctl_openbsd.pl"
	mksysnum="curl -s 'http://cvsweb.openbsd.org/cgi-bin/cvsweb/~checkout~/src/sys/kern/syscalls.master' | ./mksysnum_openbsd.pl"
	# Let the type of C char be signed for making the bare syscall
	# API consistent across platforms.
	mktypes="GOARCH=$GOARCH go tool cgo -godefs -- -fsigned-char"
	;;
solaris_amd64)
	mksyscall="go run mksyscall_solaris.go"
	mkerrors="$mkerrors -m64"
	mksysnum=
	mktypes="GOARCH=$GOARCH go tool cgo -godefs"
	;;
*)
	echo 'unrecognized $GOOS_$GOARCH: ' "$GOOSARCH" 1>&2
	exit 1
	;;
esac

(
	if [ -n "$mkerrors" ]; then echo "$mkerrors |gofmt >$zerrors"; fi
	case "$GOOS" in
	*)
		syscall_goos="syscall_$GOOS.go"
		case "$GOOS" in
		darwin | dragonfly | freebsd | netbsd | openbsd)
			syscall_goos="syscall_bsd.go $syscall_goos"
			;;
		esac
		if [ -n "$mksyscall" ]; then
			if [ "$GOOSARCH" == "aix_ppc64" ]; then
				# aix/ppc64 script generates files instead of writing to stdin.
				echo "$mksyscall -tags $GOOS,$GOARCH $syscall_goos $GOOSARCH_in && gofmt -w zsyscall_$GOOSARCH.go && gofmt -w zsyscall_"$GOOSARCH"_gccgo.go && gofmt -w zsyscall_"$GOOSARCH"_gc.go " ;
<<<<<<< HEAD
=======
			elif [ "$GOOS" == "darwin" ]; then
			        # pre-1.12, direct syscalls
			        echo "$mksyscall -tags $GOOS,$GOARCH,!go1.12 $syscall_goos $GOOSARCH_in |gofmt >zsyscall_$GOOSARCH.1_11.go";
			        # 1.12 and later, syscalls via libSystem
				echo "$mksyscall -tags $GOOS,$GOARCH,go1.12 $syscall_goos $GOOSARCH_in |gofmt >zsyscall_$GOOSARCH.go";
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
			else
				echo "$mksyscall -tags $GOOS,$GOARCH $syscall_goos $GOOSARCH_in |gofmt >zsyscall_$GOOSARCH.go";
			fi
		fi
	esac
	if [ -n "$mksysctl" ]; then echo "$mksysctl |gofmt >$zsysctl"; fi
	if [ -n "$mksysnum" ]; then echo "$mksysnum |gofmt >zsysnum_$GOOSARCH.go"; fi
<<<<<<< HEAD
	if [ -n "$mktypes" ]; then
		echo "$mktypes types_$GOOS.go | go run mkpost.go > ztypes_$GOOSARCH.go";
	fi
=======
	if [ -n "$mktypes" ]; then echo "$mktypes types_$GOOS.go | go run mkpost.go > ztypes_$GOOSARCH.go"; fi
	if [ -n "$mkasm" ]; then echo "$mkasm $GOARCH"; fi
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
) | $run
