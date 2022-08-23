# share clipboard
* share clipboard from different machine
* 在多个机器间共享剪贴板
* develop base on [clipboard](https://github.com/atotto/clipboard)

## build 
```bash
make build
```

# run
use `cbshare -h` get more info
## server
### run
```bash
cbshare server -p <port>
```
### systemd
1. put the binary to /usr/bin
2. run this command
```bash
sudo cp systemd/cbshares.service /usr/lib/systemd/system/cbshares.service
sudo systemctl start cbshares
sudo systemctl enable cbshares
```

## client
### run
```bash
cbshare client -a <addr>:<port>
```
### systemd
1. put the binary to /usr/bin
2. replace <ServerAddr> with your real server address at systemd/cbshares.service
3. run this command
```bash
sudo cp systemd/cbsharec.service /usr/lib/systemd/system/cbsharec.service
sudo systemctl start cbsharec
sudo systemctl enable cbsharec
```