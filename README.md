# PDS Functional test
   This repository contains Portworx Data Service(PDS) automation script for data service CRUD and backtup operation.

# Overview
   ### Control Plane
   ##### Supported data service(i.e cassandra , kafka etc) will be deployed on a specified PDS control plane url.
   
   ### Target cluster
   ##### Target cluster will be accessed using the kubeconfig file for PDS installation in order to accomodate data service deployment.

## Prerequisites

#### Install vault 
    wget https://releases.hashicorp.com/vault/1.10.0/vault_1.10.0_linux_amd64.zip (Choose the version accordingle)
    
    unzip vault_1.10.0_linux_amd64.zip
    
    mv vault /usr/bin
    
    Verify by running vault in shell.

#### Install kubectl

## Manually triggering the test run

#### Setting up the environment variable 

    export CONTROL_PLANE_URL="<Control plane URL> i.e https://staging.pds-dev.io"

    export TARGET_KUBECONFIG=<Path to target cluster kubeconfig>
    
    export PDS_SECRET_KEY=<Path to PDS secret KEY>

    export PDS_USER_CREDENTIAL=<Path to PDS Users credentials>

    export VAULT_HOST=<Vault URL>
    
    export VAULT_TOKEN=<Vault token>

#### Test execution (ToDo:   Run using container/pod)
    go test ./test  -timeout 9999999s -v

#### Sample test run 
    - Register the target cluster to control plane.
    - Create and delete deployments for supported data services.

#### Result
    After each run the results will populated on the terminal as well as in the tests/logs directory 
    having filename as log-<timestamp>.log (locally)
  
  
  
 
