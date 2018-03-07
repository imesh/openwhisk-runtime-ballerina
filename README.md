# Apache OpenWhisk Runtime for Ballerina

This repository contains the [Ballerina](https://ballerinalang.org) runtime extension for Apache OpenWhisk serverless platform.

## Prerequisites

The following prerequisites are needed to try this out:

- [Vagrant](https://www.vagrantup.com/downloads.html) >= v2.0.1
- [OpenWhisk](https://github.com/apache/incubator-openwhisk.git)
- [OpenWhisk CLI](https://github.com/apache/incubator-openwhisk-cli)
- [Golang](https://golang.org/doc/install) >= v1.9.2 (if requires to build the Docker image)

## Quick Start Guide

1. Install OpenWhisk using Vagrant:

   ```bash
   # Clone OpenWhisk git repository
   git clone --depth=1 https://github.com/apache/incubator-openwhisk.git openwhisk

   # Switch the directory to tools/vagrant
   cd openwhisk/tools/vagrant

   # Start OpenWhisk instance
   vagrant up
   ```

2. Install OpenWhisk CLI by following it's installation guide:
   https://github.com/apache/incubator-openwhisk-cli

3. Create a Ballerina function file with the following content and name it as hello-function.bal:

   ```
   import ballerina.io;

   function main (string[] args) {
      var output = { "hello": "world!" };
      io:println(output);
   }
   ```

4. Create an OpenWhisk action for the above Ballerina function using the OpenWhisk CLI:
   
   ```bash
   wsk action create hello-function hello-function.bal --docker imesh/ballerina-action
   ```

5. Invoke the hello-function using the OpenWhisk CLI:

   ```bash
   wsk action invoke hello-function --blocking --result
   {
       "hello": "world!"
   }
   ```