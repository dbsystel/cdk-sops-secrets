# Tested combinations

| Test Case | Resource Type   | Input Format | Expected Behavior                               |
| --------- | --------------- | ------------ | ----------------------------------------------- |
| 01        | SECRET          | json         | Save json as stringified, flattened JSON        |
| 02        | SECRET          | yaml         | Save yaml as stringified, flattened JSON        |
| 03        | SECRET          | dotenv       | Save dotenv as stringified, flattened JSON      |
| 11        | SECRET_RAW      | json         | Save json as raw String SecretValue             |
| 12        | SECRET_RAW      | yaml         | Save yaml as raw String SecretValue             |
| 13        | SECRET_RAW      | dotenv       | Save dotenv as raw String SecretValue           |
| 14        | SECRET_RAW      | binary       | Save binary as raw String SecretValue           |
| 21        | SECRET_BINARY   | json         | Save json as Binary SecretValue                 |
| 22        | SECRET_BINARY   | yaml         | Save yaml as Binary SecretValue                 |
| 23        | SECRET_BINARY   | dotenv       | Save dotenv as Binary SecretValue               |
| 24        | SECRET_BINARY   | binary       | Save binary as Binary SecretValue               |
| 31        | PARAMETER       | json         | Save raw json into StringParameter              |
| 32        | PARAMETER       | yaml         | Save raw yaml into StringParameter              |
| 33        | PARAMETER       | dotenv       | Save raw dotenv into StringParameter            |
| 34        | PARAMETER       | binary       | Save raw binary into StringParameter            |
| 41        | PARAMETER_MULTI | json         | Create multpiple Parameters (from json input)   |
| 42        | PARAMETER_MULTI | yaml         | Create multpiple Parameters (from yaml input)   |
| 43        | PARAMETER_MULTI | dotenv       | Create multpiple Parameters (from dotenv input) |
