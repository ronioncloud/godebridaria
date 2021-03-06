# godebridaria

Unlocks links through its alldebrid account and sends them to the aria2 server for download

http://serveur:8080/unlock 

![Drag Racing](./assets/unlock.png)

```

  ./godebridaria [OPTIONS]

Application Options:
  /u, /aria2-url:     url aria2 server
  /s, /aria2-secret:  secret aria2 server
  /a, /address:       address of http service debrid >  :8080
  /p, /http           start http server
  /d, /download:      start download on aria2 one [url]
  /b, /debrid:        debrid url [url]


```
Your config save into data.json


```json

{
	"AskPin": {
		"success": true,
		"pin": "XXXX",
		"expired_in": 600,
		"user_url": "https://alldebrid.com/**********************",
		"base_url": "https://alldebrid.com/pin/",
		"check_url": "https://api.alldebrid.com/******************"
	},
	"Activate": {
		"success": true,
		"token": "xxxxxxxxxxxxxxxxxxxxxxxxxxx",
		"activated": true,
		"expires_in": 585
	},
	"Aria2Config": {
		"Aria2Url": "http://192.168.1.x:6800/jsonrpc",
		"Aria2Secret": "xxxxxxxxxxxx"
	},
	"Address": ":8080",
	"Token": "xxxxxx"
}

```

---


## Create DAEMON

* Create file > *sudo nano /etc/systemd/system/godebridaria.service*

```
[Unit]
Description=godebrid for aria2
After=network.target

[Service]
Type=simple
Restart=always
RestartSec=1
ExecStart=/home/user/debrid/godebridaria -p
WorkingDirectory=/home/user/debrid

[Install]
WantedBy=multi-user.target

```

* Register DAEMON :  *sudo systemctl enable godebridaria*

* View status DAEMON :  *sudo systemctl status godebridaria*


## FOR BUILD

* RASPBERRY 

> go env <br>
$env:GOOS = "linux"<br>
$env:GOARCH = "arm"<br>
go build

* LINUX 

> go env <br>
$env:GOOS = "linux"<br>
$env:GOARCH = "amd64"<br>
go build