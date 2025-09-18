module
go mod init
git add / commit
git tag v0.0.1
git push --tags


client
go get

m
git tag v1.0.0
git push --tags


c
go get @v1.0.0

m
go.mod xxx/v2
git tag v2.0.0

c
go get xxx/v2

