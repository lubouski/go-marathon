# Storage Utility for local and Cloud object manipulation
This code will allow you to do four basic commands for working with objects either at Linix FS or at the Google Cloud Storage. 
To start working with utility we need to export variable alongside with the binary to access Google Cloud API.
> export GOOGLE_APPLICATION_CREDENTIALS=<your-path-to-the-key-file>.key.json

### Prerequisites
First you need to clone repositoty change directory to source code folder.
Then you need `go` installed at your machine to properly build the binary.
```
$ go build
```
This command will produce binary with a name `storage`

### How to Use
Storage binary is a command line utility. So we need to provide string arguments to it in a appropriate order.
Example for Linux FS:
```
$ ./storage linux /tmp/ my-test-golang.txt secret put
```
Example for Cloud GCP:
```
./storage cloud bucket-golang my-test-golang.txt /tmp/my-test-golang.txt put
```
* argument[0] : this is cloud or linux platform
* argument[1] : directory for linux or storage bucket for cloud
* argument[2] : filename or object name
* argument[3] : one word data which will be written to a linux file or filename which will be upploaded to the Cloud
* argument[4] : put get list delete - this are four commands to work with Linux FS or Cloud Storage

### Limitations
Fourth parameter for linux is limited with one word due to internal implementation constraints.   
