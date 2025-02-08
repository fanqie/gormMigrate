#!/bin/bash

# 定义目标平台
platforms=(
    "linux/amd64"
    "linux/386"
    "windows/amd64"
    "windows/386"
    "darwin/amd64"
    "darwin/arm64"
)

# 编译程序
for platform in "${platforms[@]}"; do
    IFS="/" read -r os arch <<< "$platform"
    output="./output/gmInit-${os}-${arch}"

    echo "Building for ${os}/${arch}..."
    GOOS=$os GOARCH=$arch go build -o "$output" ./gmInit.go
done

echo "Build completed!"