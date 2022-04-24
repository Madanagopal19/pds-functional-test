# PDS Functional test
   This repository contains Portworx Data Service(PDS) automation script for data service CRUD and backtup operation.

# Overview
   ### Control Plane
   ##### Supported data service(i.e cassandra , kafka etc) will be deployed on a specified PDS control plane url.
   
   ### Target cluster
   ##### Target cluster will be accessed using the kubeconfig file for PDS installation in order to accomodate data service deployment.

## Prerequisites

#### Install helm
  ```
  curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3
  chmod 700 get_helm.sh
  ./get_helm.sh
  ```

#### Install kubectl
   https://kubernetes.io/docs/tasks/tools/

## Manually triggering the test run

#### Setting up the environment variable 

    export CONTROL_PLANE_URL="<Control plane URL> i.e https://staging.pds-dev.io"

    export TARGET_KUBECONFIG=<Path to target cluster kubeconfig>
    
    export PDS_SECRET_KEY=<Path to PDS secret KEY>

    export PDS_USER_CREDENTIAL=<Path to PDS Users credentials>

    export VAULT_HOST=<Vault URL>
    
    export VAULT_TOKEN=<Vault token>
    
    export CLUSTER_TYPE=<onprem or aks or eks>

    export DATA_SERVICE="<Kafka>"

    export ZOOKEEPER_CONNECTION_STRING="<Valid ZK Connection string>"

    export ZOOKEEPER_PASSWORD="<Valid ZK password of the connection string>"


#### Test execution (ToDo:   Run using container/pod)
    go test ./test  -timeout 9999999s -v

#### Sample test run 
    - Register the target cluster to control plane.
    - Create and delete deployments for supported data services.

#### Result
    After each run the results will populated on the terminal as well as in the tests/logs directory 
    having filename as log-<timestamp>.log (locally)
  
  
  
 
