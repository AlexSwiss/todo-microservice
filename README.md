# todo-microservice
A todo app microservice built with grpc and protobuf with mysql as the database 

## Prerequisites
1. You will need Go version 1.14+ installed on your development machine.
visit my article 
https://hashnode.com/post/the-proper-way-to-install-golang-and-set-up-your-workspace-ck8qhtjbu009oehs1nri77i5h to correctly download golang and setup your GOPATH~

2. Install and congifure a mysql database of your choice. I used workbench to store and retrive data

## How to run the application

1. Clone the application with `https://github.com/AlexSwiss/todo-microservice.git`

2. There are number of dependencies which need to be imported before running the application. Please get the dependenices through the following commands -

    ```shell
        go get -u all
    ```

3. Start terminal to build and run gRPC server (replace parameters according to your SQL database server):
```shell
        cd cmd/server
        go build .
        server.exe -grpc-port=9090 -db-host=<HOST>:3306 -db-user=<USER> -db-password=<PASSWORD> -db-schema=<SCHEMA>
```

If we see:
2018/09/09 08:02:16 starting gRPC server...

It means server is started.

Open another terminal to build and run gRPC client:
```shell
        cd cmd/client-grpc
        go build .
        client-grpc.exe -server=localhost:9090
```

If we see something like this:
2018/09/09 09:16:01 Create result: <api:"v1" id:13 >
2018/09/09 09:16:01 Read result: <api:"v1" toDo:<id:13 title:"title (2018-09-09T06:16:01.5755011Z)" description:"description (2018-09-09T06:16:01.5755011Z)" reminder:<seconds:1536473762 > > >
2018/09/09 09:16:01 Update result: <api:"v1" updated:1 >
2018/09/09 09:16:01 ReadAll result: <api:"v1" toDos:<id:9 title:"title (2018-09-09T04:45:16.3693282Z)" description:"description (2018-09-09T04:45:16.3693282Z)" reminder:<seconds:1536468316 > > toDos:<id:10 title:"title (2018-09-09T04:46:00.7490565Z)" description:"description (2018-09-09T04:46:00.7490565Z)" reminder:<seconds:1536468362 > > toDos:<id:13 title:"title (2018-09-09T06:16:01.5755011Z)" description:"description (2018-09-09T06:16:01.5755011Z) + updated" reminder:<seconds:1536473762 > > >
2018/09/09 09:16:01 Delete result: <api:"v1" deleted:1 >

Everything works fine.

