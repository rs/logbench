#!/bin/bash

set -e

ROOT=$(dirname "$0")

if [ -z "$GO_VERSION" ]; then
    source "$ROOT/go.version"
fi

if [ -z "$GO_VERSION" ]; then
    echo "Missing \$GO_VERSION environment" >&2
    exit 1
fi

GOROOT="$ROOT"/.goroot/$GO_VERSION
GOROOT_CURRENT="$ROOT"/.goroot/current
CACHE_DIR="$ROOT"/.cache

if [[ "$("$GOROOT"/bin/go version 2>&1 | cut -f3 -d\ )" != "go$GO_VERSION" ]]; then
    # Make sure only one instance is executed at the same time
    [ "$FLOCKER" != "$0" ] && exec env FLOCKER="$0" flock -e $0 $0 "$@"

    rm -rf "$GOROOT"

    case "$(uname -m)" in
    "x86_64")
        ARCH=amd64
        ;;
    "arm64")
        ARCH=arm64
        ;;
    "aarch64")
        ARCH=arm64
        ;;
    *)
        echo "Unsupported arch: $(uname -m)"
        exit 1
    esac

    case "$(uname -s)" in
    "Darwin")
        OS=darwin
        ;;
    "Linux")
        OS=linux
        ;;
    "FreeBSD")
        OS=freebsd
        ;;
    *)
        echo "Unsupported OS: $(uname -s)"
        exit 1
    esac

    GO_URL="https://dl.google.com/go/go$GO_VERSION.$OS-$ARCH.tar.gz"

    # Download and install Go.
    rm -rf "$GOROOT"
    mkdir -p "$GOROOT"
    mkdir -p "$CACHE_DIR"
    cache_file="$CACHE_DIR"/${GO_URL##*/}
    if [ ! -f "$cache_file" ]; then
        echo "Installing $GO_URL" >&2
        curl -sS -o "$cache_file" "$GO_URL"
    else
        echo "Installing $GO_URL (from cache)" >&2
    fi

    tar --strip-components 1 -C "$GOROOT" -zxf "$cache_file"

    # Reset quilt state.
    rm -rf "$ROOT"/.pc

    # Update current symlink so quilt find the files.
    rm -rf "$GOROOT_CURRENT"
    ln -sf -T "$GO_VERSION" "$GOROOT_CURRENT"
else
    # Update current symlink so quilt find the files, may be rolling back to previous version.
    ln -sf -T "$GO_VERSION" "$GOROOT_CURRENT"
fi

"$GOROOT"/bin/go "$@"
