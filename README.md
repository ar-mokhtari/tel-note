<h1 style="alignment: center"> tel-note </h1>

<h6> tel-note is a structure and syntax practise with Go-lang. we are beginner in go and this practise seems to have bad implement or structure, please give us tips🙏
</h6>
---

<h3> please notice us if you have any other idea ...
</h3>
<ul>
<li>neatland@gmail.com</li``>  
<li>telegram @ar_mokhtari</li>  
</ul>

<strong></strong>
<img src="https://camo.githubusercontent.com/4724436344c2473558068577d7e9e6b597c2baabe75a499cd67e04a448e00d84/68747470733a2f2f7777772e766563746f726c6f676f2e7a6f6e652f6c6f676f732f676f6c616e672f676f6c616e672d617232312e737667" >

---

<h3> To build:</h3>

$ go build -o bin/telnote

<h3> Create service </h3>

$ nano /lib/systemd/system/telnote.service

<h3> Set service </h3>

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

<h3> Service's actions </h3>

````
$ service telnote start
$ service telnote restart
$ service telnote status
$ service telnote enable
$ service telnote stop
````