<h1 align="center" style="alignment: center;color: yellow">
  <img alt="cgapp logo" src="https://raw.githubusercontent.com/create-go-app/cli/master/.github/images/cgapp_logo%402x.png" width="224px"/><br/>
  Tel-Note
</h1>
<p align="center"><b>backend</b> (Golang), <b>frontend</b> (JavaScript, Jquery)<br/><br/><br/></p>

<p align="center">
<a href="" 
target="_blank"><img src="https://img.shields.io/badge/Go-1.17+-00ADD8?style=for-the-badge&logo=go" alt="go version" /></a>&nbsp;
<a href="" 
target="_blank"><img src="https://img.shields.io/badge/jquery-%230769AD.svg?style=for-the-badge&logo=jquery&logoColor=white" alt="jQuery" /></a>&nbsp;
<a href="" 
target="_blank"><img src="https://img.shields.io/badge/javascript-%23323330.svg?style=for-the-badge&logo=javascript&logoColor=%23F7DF1E" alt="javascript" /></a>&nbsp;
</p>

<h6> Tel-note is a structure and syntax practise with Go-lang. </h6>
<h6> This is an example of sending and receiving data from/to external sources and the operations of creating, reading, updating and maintaining data. Searching and creating sample reports is one of the project goals. 
</h6>

---

<h6>Special thanks to  https://github.com/OmidHekayati :heart: for training support
</h6>

---

<h6> please notice me if you have any other idea ...
</h6>

<a href="mailto:neatland@gmail.com"><img alt="Gmail" title="Alireza Mokhtari G Gmail" src="https://camo.githubusercontent.com/571384769c09e0c66b45e39b5be70f68f552db3e2b2311bc2064f0d4a9f5983b/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f476d61696c2d4431343833363f7374796c653d666f722d7468652d6261646765266c6f676f3d676d61696c266c6f676f436f6c6f723d7768697465" data-canonical-src="https://img.shields.io/badge/Gmail-D14836?style=for-the-badge&amp;logo=gmail&amp;logoColor=white" style="max-width: 100%;"></a>
<a href="https://t.me/ar_mokhtari" rel="nofollow"><img alt="Telegram" title="Alireza Mokhtari G Telegram" src="https://camo.githubusercontent.com/cf4ed981404024c1adfc79d5575c4edf1836c4fe36b24b03383ece888cef7e29/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f54656c656772616d2d3243413545303f7374796c653d666f722d7468652d6261646765266c6f676f3d74656c656772616d266c6f676f436f6c6f723d7768697465" data-canonical-src="https://img.shields.io/badge/Telegram-2CA5E0?style=for-the-badge&amp;logo=telegram&amp;logoColor=white" style="max-width: 100%;"></a>
<a href="https://www.linkedin.com/in/alireza-mokhtari-garakani-b4288024/"><img src="https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&amp;logo=linkedin&amp;logoColor=white" style="max-width: 100%;"></a>

<h4 style="color: deepskyblue;margin-bottom: -0.9rem">Powered by</h4>
<img src="https://camo.githubusercontent.com/4724436344c2473558068577d7e9e6b597c2baabe75a499cd67e04a448e00d84/68747470733a2f2f7777772e766563746f726c6f676f2e7a6f6e652f6c6f676f732f676f6c616e672f676f6c616e672d617232312e737667" >


---

<h3>Requirements</h3>

[Requirements for this project](docs/Requirements)

> ---
> - <h4>Under Developing:</h4>
> - [ ] Identification, Authentication, Authorization to control user services
> - [ ] Create Validation service
> - [ ] Export data to csv file
> - [ ] Send sms for different service(s)
> - [ ] Send email for different service(s)
> - [ ] Create and develop internal database to "data service(ing)"
> - [ ] Create reports, listen and export official data(s) by api(s) [b2e]
> - [ ] Call SepehrAPI for flight ticket attribute
> - [ ] Call GRS for hotel reservation
> - [ ] Build with storage(s) type tags
> - [ ] Use general packages (lib)s
> - [ ] Multi language service(s)
> - [ ] Create product wiki
> - [ ] Create UI web application
> - [ ] Create Mobile application (Android)


> ---
> - <h4>Developed :</h4>
> - [x] Import data from external source csv file
> - [x] Listen and export business data to applicant(s) [b2c]
>> - [ ] use a web framework and router
>> - [x] create a collection for api(s) document(s)
> - [x] Call Some Api(s) [b2b]
>> - [x] Insert and update country(ies) and its property(ies)
>> - [x] Send some data and receive processed data from provider, (lat and lng of a city and publish online traffic duration and distance[neshan.org])

<h3> To build:</h3>

````
$ go build -tags storageType(exmp: memory) -o bin/telnote
````

<h3> Create service </h3>

````
$ nano /lib/systemd/system/telnote.service
````

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

<h3> SDKs setting </h3>
<ul>
<li>UniversalTutorial(for get countries)
<div style="margin-left: 2em">	Apply to get services from this site:
	<a href="https://www.universal-tutorial.com/">universal-tutorial</a>
sdk/Universal/UniversalMetadata.go
</div>
</li>
<li>Neshan(for get cities service(s))
<div style="margin-left: 2em">	Apply to get services from this site:
	<a href="https://platform.neshan.org/api/getting-started">platform.neshan.org/api/</a>
sdk/Universal/NeshanMetadata.go
</div>
</li>
</ul>

<h3> Environment </h3>
<ul>
<li>IranCities.csv (for get cities)
<div style="margin-left: 2em">
env/IranCities.csv
</div>
</li>
<li>Sample data 
<div style="margin-left: 2em">
Create examples of all available entities for testing:
<ul style="margin-top: unset">
<li>file: env/sampleData.go
</li>
<li>service: services/sampleData/fillSampleData.go & getSampleData.go</li>
<li>cli: reg->data
</li>
<li>route: 127.0.0.1:1212/fill-data and get-data
</li>
</ul>
</div>
</li>
</ul>

<h3> Entities </h3>
<ul> <b>Protocols:s</b>
<details style="margin-left: 2em">
<li>city</li>
<li>contact</li>
<li>country</li>
<li>customer</li>
<li>job</li>
<li>person</li>
<li>sex</li>
</details>
</ul>

<p>
    <img src="https://github.com/ar-mokhtari/tel-note/blob/main/docs/Entity.png" width="220" height="240" />
</p>
<h3> Run application </h3>
<ul>
<li>cli->config->RunAppType = ""</li>
<li>cli->config->RunAppType = "serv"</li>
</ul>

<h3> Routes </h3>
<ul>
<li>city
<ul>
<li><a href="127.0.0.1:1212/get-city">127.0.0.1:1212/get-city</a></li>
<li><a href="127.0.0.1:1212/add-city">127.0.0.1:1212/add-city</a></li>
<li><a href="127.0.0.1:1212/edit-city">127.0.0.1:1212/edit-city</a></li>
<li><a href="127.0.0.1:1212/find-city-char">127.0.0.1:1212/find-city-char</a></li>
<li><a href="127.0.0.1:1212/find-city-id">127.0.0.1:1212/find-city-id</a></li>
<li><a href="127.0.0.1:1212/delete-city">127.0.0.1:1212/delete-city</a></li>
<li><a href="127.0.0.1:1212/distance-time-between-two-city">127.0.0.1:1212/distance-time-between-two-city</a></li>
</ul>
</li>
<li>contact
<ul>
<li><a href="127.0.0.1:1212/get-contact">127.0.0.1:1212/get-contact</a></li>
<li><a href="127.0.0.1:1212/new-contact">127.0.0.1:1212/add-contact</a></li>
<li><a href="127.0.0.1:1212/edit-contact">127.0.0.1:1212/edit-contact</a></li>
<li><a href="127.0.0.1:1212/find-contact-char">127.0.0.1:1212/find-contact-char</a></li>
<li><a href="127.0.0.1:1212/find-contact-id?id=2">127.0.0.1:1212/find-contact-id?id=2</a></li>
<li><a href="127.0.0.1:1212/delete-contact-id">127.0.0.1:1212/delete-contact-id</a></li>
<li><a href="127.0.0.1:1212/delete-contact-all">127.0.0.1:1212/delete-contact-all</a></li></ul>
</li>
<li>country
<ul><li><a href="127.0.0.1:1212/country-list">127.0.0.1:1212/country-list</a></li></ul></li>
<li>customer</li>
<li>job</li>
<li>person
<ul>
<li><a href="127.0.0.1:1212/find-person-id?pid=2">127.0.0.1:1212/find-person-id?pid=2</a></li></ul></li>
<li>sex</li>
</ul>
<h3>Json collection</h3>
To use and call from API platform like "Postman":

````
/docs/tel-note.postman_collection.json
````

<h3>User Interface [web]</h3>
Testing with "Nginx web server"
[nginx install and setting](UI/Docs/nginx)
<p align="left">
<a href="" 
target="_blank"><img src="https://img.shields.io/badge/nginx-%23009639.svg?style=for-the-badge&logo=nginx&logoColor=white" alt="nginx" /></a>&nbsp;
</p>

or simply browse:

```` 
http://127.0.0.1:1212/index.html 
````


A sample SPA PWA with "Ajax" request which call routes ...
<ul>
<li>HTML</li>
<li>CSS Boostrap -v 5.1.3</li>
<li>jQuery 3.6.0</li>
<li>Ajax</li>
</ul>

```` 
/UI/tel-note.html 
````

<h3>Git</h3>
<p align="left">
<a href="" 
target="_blank"><img src="https://img.shields.io/badge/git-%23F05033.svg?style=for-the-badge&logo=git&logoColor=white" alt="git" /></a>&nbsp;
</p>

for production, there is stable **main** branch:

```` 
 git checkout main
 
````

for develop:

```` 
 git checkout -b developing
 
````
after all, back to main:
```` 
 git merge --no-ff developing
 
````
