# tel-note

### tel-note is a structure and syntax practise with Go-lang. we are beginner in go and this practise seems to have bad implement or structure, please give us tips🙏

---

#### please notice us if you have any other idea ...

neatland@gmail.com  
telegram @ar_mokhtari

<img src="https://camo.githubusercontent.com/4724436344c2473558068577d7e9e6b597c2baabe75a499cd67e04a448e00d84/68747470733a2f2f7777772e766563746f726c6f676f2e7a6f6e652f6c6f676f732f676f6c616e672f676f6c616e672d617232312e737667" >

---

###To build:
$ go build -o bin/telnote

###Create service
$ nano /lib/systemd/system/telnote.service

###Set service
````
[Unit]
Description=telnote

[Service]
Type=simple
Restart=always
RestartSec=5s
ExecStart=/home/aliz/go/src/tel-note/bin/telnote

[Install]
WantedBy=multi-user.target
````
### Service's actions
````
$ service telnote start
$ service telnote restart
$ service telnote status
$ service telnote enable
$ service telnote stop
````