version="1.0"
rm -rf release/builds
mkdir release/apps/
rm -rf release/apps/*
echo "compile win64.."
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  -o release/builds/win64/pbq.exe
echo "compile win32.."
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build  -o release/builds/win32/pbq.exe
echo "compile darwin64.."
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build  -o release/builds/darwin64/pbq
echo "compile linux64.."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o release/builds/linux64/pbq

echo "copy install.sh.."

cp  install.sh release/builds/darwin64/
cp  install.sh release/builds/linux64/


echo "packing.."

zip -qj release/apps/pbq-${version}-win32.zip release/builds/win32/*
zip -qj release/apps/pbq-${version}-win64.zip release/builds/win64/*
zip -qj release/apps/pbq-${version}-darwin64.zip release/builds/darwin64/*
zip -qj release/apps/pbq-${version}-linux64.zip release/builds/linux64/*
