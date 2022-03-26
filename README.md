<h1 align="center" style="alignment: center"> Tel-Note </h1>

<h6> tel-note is a structure and syntax practise with Go-lang. we are beginner in go and this practise seems to have bad implement or structure, please give us tips🙏
</h6>
---

<h3> please notice us if you have any other idea ...
</h3>
<ul>
<li>  
<a href="mailto:neatland@gmail.com"><img alt="Gmail" title="Alireza Mokhtari G Gmail" src="https://camo.githubusercontent.com/571384769c09e0c66b45e39b5be70f68f552db3e2b2311bc2064f0d4a9f5983b/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f476d61696c2d4431343833363f7374796c653d666f722d7468652d6261646765266c6f676f3d676d61696c266c6f676f436f6c6f723d7768697465" data-canonical-src="https://img.shields.io/badge/Gmail-D14836?style=for-the-badge&amp;logo=gmail&amp;logoColor=white" style="max-width: 100%;"></a>
</li>

<li>
<a href="https://t.me/ar_mokhtari" rel="nofollow"><img alt="Telegram" title="Alireza Mokhtari G Telegram" src="https://camo.githubusercontent.com/cf4ed981404024c1adfc79d5575c4edf1836c4fe36b24b03383ece888cef7e29/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f54656c656772616d2d3243413545303f7374796c653d666f722d7468652d6261646765266c6f676f3d74656c656772616d266c6f676f436f6c6f723d7768697465" data-canonical-src="https://img.shields.io/badge/Telegram-2CA5E0?style=for-the-badge&amp;logo=telegram&amp;logoColor=white" style="max-width: 100%;"></a>
</li>
</ul>

<strong></strong>
<img src="https://camo.githubusercontent.com/4724436344c2473558068577d7e9e6b597c2baabe75a499cd67e04a448e00d84/68747470733a2f2f7777772e766563746f726c6f676f2e7a6f6e652f6c6f676f732f676f6c616e672f676f6c616e672d617232312e737667" >

---

<h3> To build:</h3>

$ go build -tags storageType -o bin/telnote

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