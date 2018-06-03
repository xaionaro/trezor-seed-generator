```
apt-get install -y iamerican
zcat /usr/share/ispell/american.mwl.gz > /tmp/american.mwl
go get github.com/xaionaro/trezor-seed-generator/trezorSeedGenerator
go run ~/go/src/github.com/xaionaro/trezor-seed-generator/trezorSeedGenerator/main.go -d /tmp/american.mwl
```

result:
```
$ go run ~/go/src/github.com/xaionaro/trezor-seed-generator/trezorSeedGenerator/main.go -d /tmp/american.mwl 
 1: line
 2: unidimensional
 3: fanaticism
 4: knuth
 5: truculent
 6: edgewise
 7: browse
 8: shill
 9: overborne
10: vague
11: anchorage
12: contentious
13: brilliant
14: penthouse
15: snowy
16: clouded
17: simpatico
18: flavorful
19: adjudication
20: sat
21: shovel
22: derangement
23: reparable
24: punch
```
