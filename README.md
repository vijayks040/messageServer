# messageServer

# 

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

This Go code was added to test GO skills.

# Pre-requisites!

  - Hoping that the GO environment already present,run the following command,
  
        -go get -d -v ./...
        

### API's
**Add Message**
``
curl --location --request POST 'http://localhost:8080/addMessage' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "message": "madam",
    "description": "hello",
    "user": "vijay"
}'
```

**List Message**: 

```
Request:
URL: http://localhost:8080/listMessage Method: GET
Respone: In Json format
{
    "1": {
        "id": 1,
        "message": "madam",
        "description": "hello",
        "user": "vijay"
    },
    "2": {
        "id": 2,
        "message": "hello hi aws",
        "description": "hi",
        "user": "qlik"
    }
}

```

**Get One Message**: Using Message ID
```Request: 
URL: http://localhost:8080/getOneMessage?id=1 METHOD: GET
Response: In Json format
{
    "id": 1,
    "message": "madam",
    "description": "hello",
    "user": "vijay"
}
```

**Delete message**: Using message ID
```
Request:
URL: http://localhost:8080/deleteOneMessage?id=1 METHOD: DELETE
Response: Http Status code
```
### Executables

Both the windows and linux platform executables have been provided along with source code.

| Platform | Executable |
| ------ | ------ |
| Windows7 and higher version | main.exe (Run using admin permissions) |
| Linux | main (with sudo on run ./main command in terminal) |
| Docker Image | docker pull vijayks040/messageserver:demo |
| Aws Host     | http://ec2-100-27-2-88.compute-1.amazonaws.com:8080  |

### Development Environment

Want to contribute? Great!

**GO version : 1.14.4 windows/amd64**

**-This piece of code is developed in windows environment.**

**-Anyhow Go is such a platform freindly coding language, you can run this code in any other platforms such as linux,mac etc**




### Todos

 - Write MORE Tests
 - This App is also having the code to interact with DB for all the Message Operations
 - Using dependancy management **dep**
 
 ### NOTE: This App has been hosted in AWS ec2 instance with url ec2-100-27-2-88.compute-1.amazonaws.com

License
----

@vijay


**Free Software, Hell Yeah!**
