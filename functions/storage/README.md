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
$ ./storage -provider linux -path /tmp/ -file "my-test-golang.txt" -data "secret data" -command put
```
Example for Cloud GCP:
```
./storage -provider cloud -path "bucket-golang" -file "my-test-golang.txt" -data "/tmp/my-test-golang.txt" - command put
```
* provider : this is cloud or linux platform
* path : directory for linux or storage bucket for cloud
* file : filename or object name
* data : one word data which will be written to a linux file or filename which will be upploaded to the Cloud
* command : put get list delete - this are four commands to work with Linux FS or Cloud Storage

### Limitations
Fourth parameter for linux is limited with one word due to internal implementation constraints.   
