this is a database migration tool based on Gorm
# Warning⚠️
- **This tool is still under development, please do not use it in any environment.**
# Fast start
## Install dependencies
```shell
go get -u github.com/fanqie/gormMigrate
```
###
- Download init tool
```shell
// linux or mac
curl -L https://github.com/fanqie/gormMigrate/output/gmInit.exe -o gmInit.exe
chmod +x gmInit.exe
```
```powershell
// windows
curl -L https://github.com/fanqie/gormMigrate/output/gmInit.exe -o gmInit.exe
gmInit.exe
```
- Run gmInit.exe


# Develop guide
## Build gmInit.go
### windows
```shell 
powershell  .\script\buildGmInit.sh
```
### linux or mac
```shell
chmod +x  .\script\buildGmInit.sh
sh  .\script\buildGmInit.sh
```