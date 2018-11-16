# XML2JSON flogo activity
This activity allows your flogo application to convert xml to JSON data. This done with goxml2json (github.com/basgys/goxml2json).



## Installation

```bash
flogo install github.com/abasse/flogoxml2json
```

## Schema
Inputs and Outputs:

```json
  { "inputs":[
        {
          "name": "input",
          "type": "string",
          "required": true
        }
      ],
      "outputs": [
         {
           "name": "result",
           "type": "string"
          }
      ]
 }
```
