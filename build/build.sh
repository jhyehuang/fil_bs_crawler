echo "build start"
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
go mod tidy

buildPath=`pwd`
rootPath=$(dirname $buildPath)

rm -rf $rootPath/_deploy/*


echo "***************************************"
echo "build web start"

echo "cd $rootPath/cmd/web ..."
cd $rootPath/cmd/web
go build -o $rootPath/_deploy/mars-web .
echo "build web end"
echo "***************************************"


echo "***************************************"
echo "build mint start"

echo "cd $rootPath/cmd/mint ..."
cd $rootPath/cmd/mint
go build -o $rootPath/_deploy/mars-mint .
echo "build mint end"
echo "***************************************"


echo "***************************************"
echo "copy config file start"
echo "cd $rootPath ..."
cd $rootPath

mkdir $rootPath/_deploy/logs
mkdir $rootPath/_deploy/config
mkdir $rootPath/_deploy/assets

cp $rootPath/config/config.toml $rootPath/_deploy/config/
cp $rootPath/config/chain.toml $rootPath/_deploy/config/chain.toml
cp -r $rootPath/assets/* $rootPath/_deploy/assets/
echo "copy config file end"
echo "***************************************"

echo "build end"
