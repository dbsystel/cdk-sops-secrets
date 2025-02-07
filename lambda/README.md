# Tested combinations

| Test Case | Input Format | Output Format | Flatten    | Stringify | Handling Type   | Expected Behavior                      |
| --------- | ------------ | ------------- | ---------- | --------- | --------------- | -------------------------------------- |
| 01        | binary       | binary        | N/A        | N/A       | SECRET          | No transformation, raw binary output   |
| 02        | binary       | binary        | N/A        | N/A       | PARAMETER       | No transformation, raw binary output   |
| 11        | json         | N/A           | true (N/A) | true      | PARAMETER_MULTI | Outputformat allways map[string]string |
| 12        | yaml         | N/A           | true (N/A) | true      | PARAMETER_MULTI | Outputformat allways map[string]string |
| 13        | dotenv       | N/A           | true (N/A) | true      | PARAMETER_MULTI | Outputformat allways map[string]string |
| 21        | json         | json          | true       | true      | SECRET          | Stringified, Flattened JSON            |
| 22        | json         | binary        | N/A        | N/A       | SECRET          | RAW Input is output                    |
| 23        | yaml         | json          | true       | true      | SECRET          | Stringified, Flattened JSON            |
| 24        | yaml         | binary        | N/A        | N/A       | SECRET          | RAW Input is output                    |
| 25        | dotenv       | json          | true       | true      | SECRET          | Stringified, Flattened JSON            |
| 26        | dotenv       | binary        | N/A        | N/A       | SECRET          | RAW Input is output                    |
