# CX CLOUD - GoLang - MAC Address Lookup

This service utilizes https://api.macaddress.io/v1 to check the Vendor associated with a given MAC Address.

+ Dockerfile stored at ./Dockerfile
+ GoLang application file stored at ./macLookup.go
+ Executable file stored at ./getMac

## Before Use:

Go to https://macaddress.io/api in order to sign up for the service and acquire an API key. Make sure you set and export a 'macApiKey' environment variable before running.

## To Run:

### CLI:
```
$ export macApiKey='<your api key>'
$ go build -o getMac ./macLookup.go
$ ./getMac "<MAC Address>"
```

### Example

```
$ ./getMac "44:38:39:ff:ef:57"

Beginning Execution

Mac Address supplied: 44:38:39:ff:ef:57
Querying: https://api.macaddress.io/v1
Vendor: Cumulus Networks, Inc
```

### Docker:
```
$ export macApiKey='<your api key>'
$ docker build -t <image_name> .
$ docker run -emacApiKey=$macApiKey "<MAC Address>"
```

### Example

```
$ docker build -t maclookup:1.0 .

$ docker run -emacApiKey=$macApiKey maclookup:1.0 "44:38:39:ff:ef:57"

Beginning Execution

Mac Address supplied: 44:38:39:ff:ef:57
Querying: https://api.macaddress.io/v1
Vendor: Cumulus Networks, Inc
```