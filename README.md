# PDS Functional test (Updating soon)

# Setting up the environment to run the tests 

## Export following environment variables: 

    export CONTROL_PLANE_URL="https://staging.pds-dev.io"

    export TARGET_KUBECONFIG=<Path to target cluster kubeconfig>
    
    export PDS_SECRET_KEY=<Path in vault location to PDS secret KEY>

    export PDS_USER_CREDENTIAL=<Path in vault location to PDS Users credentials>


export VAULT_HOST=<Vault URL>
export VAULT_TOKEN=<Vault token>

# Test execution
    go test ./test  -timeout 99999999s -v

# Sample test run 
    Register the target cluster to control plane.

# Result
  Logs will be populated inside the log directory filename-<timestamp> as well as on the terminal.
  
  
  
 
