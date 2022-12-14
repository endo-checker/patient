# Patient gRPC microserverice #

This application uses `gRPC` and `proto` to send user data to a database (mongoDB). <br/>
This uses CRUD functions. <br />
This service also has a Publish/subscribe method using dapr, a new patient is created in the publish method, which is sent to mongoDB, then the callback function `subscribe` method takes that data and adds that user to auth0 where they can then login to view their information. 

## Get Started ##
- Install/update dependencies `go get -u`
- Initiate gRPC: `protoc --proto_path=proto patient/v1/patient.proto --go_out=. --go-grpc_out=.` This is the standard way of generating gRPC, however with this app I am using a makefile configured with dapr to manage my commands. So instead run `make proto`, this will also send the proto definitions to `https://buf.build` as a commit. 
- Run the application `go run .` or `make run`

### Dapr ###

Install and configure the dapr CLI:

```bash
brew install dapr/tap/dapr-cli
dapr init
```

### Buf ###

Add Buf API key to `.netrc` file

```bash
machine buf.build password <your Buf API key>
machine go.buf.build login <your Buf username> password <your Buf API key>
```

## Set up Docker ## 
* To make a build on `docker` run `docker build --tag patient .`
* To run on docker `docker run -d -p 8080:8080 patient`
