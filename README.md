```
apt-get install -y iamerican
zcat /usr/share/ispell/american.mwl.gz > /tmp/american.mwl
go get github.com/xaionaro/trezor-seed-generator/trezorSeedGenerator
go run ~/go/src/github.com/xaionaro/trezor-seed-generator/trezorSeedGenerator/main.go -d /tmp/american.mwl
```
