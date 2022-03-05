#!/usr/bin/env bash

package=$1
output=$2
if [ -z "$package" ] || [ -z "$output" ]; then
  echo "usage: $0 <package-name> <output-name>"
  exit 1
fi

platforms=("linux/amd64" "linux/arm64" "windows/amd64" "windows/arm64" "darwin/amd64" "darwin/arm64")

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	output_name=$output'-'$GOOS'-'$GOARCH

	if [ $GOOS = "windows" ]; then
		output_name+='.exe'
	fi

	env GOOS=$GOOS GOARCH=$GOARCH go build -ldflags "-s -w" -trimpath -o executeables/$output_name $package
	if [ $? -ne 0 ]; then
   		echo 'An error has occurred. Build failed.'
		exit 1
	fi
done
