#!/bin/bash

# 定义目标平台
platforms=(
    "linux/amd64"
    "linux/386"
    "linux/arm64"   # 添加 Linux ARM64 支持
    "linux/arm"     # 添加 Linux ARM 支持
    "windows/amd64"
    "windows/386"
    "windows/arm64" # 添加 Windows ARM64 支持
    "windows/arm"   # 添加 Windows ARM 支持
    "darwin/amd64"
    "darwin/arm64"
)

# 定义扩展名
extensions=(
    ""      # linux/amd64
    ""      # linux/386
    ""      # linux/arm64
    ""      # linux/arm
    ".exe"  # windows/amd64
    ".exe"  # windows/386
    ".exe"  # windows/arm64
    ".exe"  # windows/arm
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
