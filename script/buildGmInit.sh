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

# 定义扩展名
extensions=(
    ""      # linux/amd64
    ""      # linux/386
    ".exe"  # windows/amd64
    ".exe"  # windows/386
    ""      # darwin/amd64
    ""      # darwin/arm64
)

# 编译程序
for i in "${!platforms[@]}"; do
    IFS="/" read -r os arch <<< "${platforms[$i]}"
    extension="${extensions[$i]}"
    output="./output/gmInit-${os}-${arch}${extension}"

    echo "Building for ${os}/${arch}..."
    GOOS=$os GOARCH=$arch go build -o "$output" ./gmInit.go
done

echo "Build completed!"
