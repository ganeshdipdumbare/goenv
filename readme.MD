![Test](https://github.com/ganeshdipdumbare/goenv/workflows/Test/badge.svg) [![codecov](https://codecov.io/gh/ganeshdipdumbare/goenv/branch/master/graph/badge.svg)](https://codecov.io/gh/ganeshdipdumbare/goenv)

# goenv
Simple tool to sync env variables in Go structure. 

## Description
This tool sync env variable in the structure. It will sync the env variables in the structure passed as input to the function ```SyncEnvVar```.   

The env variable name is fetched from the tag of json tag of the structure field as follows-  

```golang 
type EnvVar struct{
    MongoDb string `json:"mongo_db"`
    Port    string `json:"port"`
}

envVar := &EnvVar{
    Port: ":8080"
}

SyncEnvVar(&envVar)
// envVar will be updated to following value
// envVar.MongoDb will be updated with value of env var  MONGO_DB
// envVar.Port will be updated with value of env var PORT, otherwise will remain :8080

```  

This will sync env variables ```MONGO_DB``` and ```PORT```. Here port value will be fetched from env variable and if it is not found, default value will be taken as ```:8080```  

## Important notes
- Do not use ```omitempty``` option for structure field.
- JSON tag values should be lower case of env variable name, e.g. if env variable name is MONGO_PORT, then json tag should be mongo_port.