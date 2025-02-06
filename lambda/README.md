# Tested combinations

| Test Case | Input Format | Output Format | Flatten    | Stringify | Handling Type   | Expected Behavior                            |
| --------- | ------------ | ------------- | ---------- | --------- | --------------- | -------------------------------------------- |
| 01        | binary       | binary        | N/A        | N/A       | SECRET          | No transformation, raw binary output         |
| 02        | binary       | binary        | N/A        | N/A       | PARAMETER       | No transformation, raw binary output         |
| 11        | json         | N/A           | true (N/A) | true      | PARAMETER_MULTI | Outputformat is ignored (allways string map) |
| 12        | yaml         | N/A           | true (N/A) | true      | PARAMETER_MULTI | Outputformat is ignored (allways string map) |
| 13        | dotenv       | N/A           | true (N/A) | true      | PARAMETER_MULTI | Outputformat is ignored (allways string map) |
| 21        | json         | json          | true (N/A) | true      | SECRET          | Flattens JSON but retains JSON format        |
| 22        | json         | binary        | N/A        | N/A       | SECRET          | Flattens JSON but retains JSON format        |
| 23        | yaml         | json          | true       | true      | SECRET          | Flattens JSON, converts to YAML, stringifies |
| 24        | yaml         | binary        | N/A        | N/A       | SECRET          | Flattens JSON, converts to YAML, stringifies |
| 25        | dotenv       | json          | true       | true      | SECRET          | Flattens YAML, converts to JSON, stringifies |
| 26        | dotenv       | binary        | N/A        | N/A       | SECRET          | SecretContent the same as                    |
| 27        | binary       | binary        | false      | true      | SECRET          | Converts dotenv to YAML, stringifies values  |
