# https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04

for os in linux darwin
do
    for arch in "386" "amd64" "arm" "arm64"
    do
        mkdir -p ./dist/$os/$arch/
        CGO_ENABLED=0 GOOS=$os GOARCH=$arch go build -o ./dist/$os/$arch/ rip.go
    done
done

for os in windows
do
    for arch in "386" "amd64"
    do
        mkdir -p ./dist/$os/$arch/
        CGO_ENABLED=0 GOOS=$os GOARCH=$arch go build -o ./dist/$os/$arch/ rip.go
    done
done