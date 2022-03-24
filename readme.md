# Bookkeeper

To start:

```bash
make  install && bookkeeper --input_json_path <JSON CONFIG LOCATION>

```

Which the json config file should follow the following schema:

```json

{
  "chains":[
    {"chain_type":"cosmos",
      "details": [
        {
          "chain_name":"desmos",
          "grpc_endpoint":<gRPC endpoint of the chain>,
          "rpc_endpoint":<RPC endpoint of the chain>,
          "validators":[
            {
              "validator_address":<validator address>,
              "self_delegation_address":<self delegation address>
            }],
          "fund_holding_account":[<Array of fund holding accounts>]
        }
      ]
    }
  ],
  "email_details":{
    "from":{
      "name":<Email sender name>,
      "host":<Host now support Gmail>,
      "password":<Password of email account>,
      "address":<Email address>
    },
    "to":[<an array of email address send to>],
    "subject":"Monthy report from bookkeeper",
    "details":"Enjoy"
  }
}

```