package mqtt

var jsonMetadata = `{
  "name": "mqtt",
  "version": "0.0.1",
  "description": "trigger description",
  "config":[
    {
      "name": "broker",
      "type": "string"
    },
    {
      "name": "id",
      "type": "string"
    },
    {
      "name": "user",
      "type": "string"
    },
    {
      "name": "password",
      "type": "string"
    },
    {
      "name": "store",
      "type": "string"
    },
    {
      "name": "topic",
      "type": "string"
    },
    {
      "name": "qos",
      "type": "number"
    },
    {
      "name": "cleansess",
      "type": "boolean"
    }
  ],
  "endpoint": {
    "settings": [
      {
        "name": "topic",
        "type": "string"
      }
    ]
  }
}`
