# dss-email
Component made with Go, receives a list of mails then writes them into a logfile

## Getting Started

Just follow the instruction to copy and run the project.

### Prerequisites

* Install Go from its page at https://golang.org/.

### Installing

* Clone the project to your src Golang project folder, on windows its on C:\Users\go\src. 
* In the folder open a terminal an write the following:
```
go get -u github.com/golang/dep/cmd/dep
```
```
dep init -v
```
```
dep ensure -v 
```
* Execute the http server:
```
go run receiver.go sender.go
```

### Output

* The program creates a logfile "emailLog" where writes all the mails the queue reads.

Example:
this is an example of the logfile

```
2018/12/05 11:16:11 mails sent to alejandro2222 Mail : luis@gmail.com                                 Event : created
2018/12/05 11:16:11 mails sent to josue2222 Mail : josue_147_15@hotmail.com                           Event : created
2018/12/05 11:16:11 mails sent to apagar-MV Mail : luis@gmail.com                                     Event : created
2018/12/05 11:16:11 INFO : Sent successfully
```

## Built With

* [Go](https://golang.org/) - The Programming langauge
* [Rabbitmq](https://www.rabbitmq.com/) - Queue Messages

## Authors

* **Alejandro Cabrera** - [alep007](https://github.com/alep007)
* **Josue Ferrufino** - [josue1471515](https://github.com/orgs/Kuma-gg/people/josue1471515)

### To Do
* Test
