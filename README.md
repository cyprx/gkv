## Distributed key-value database in Golang

### Features
- GET /{key}
    - Response 200 Found, value
    - Response 404 Not Found
- POST /{key} 
    - Response 201 Created
    - Response 203 Modified


### TODO
- How k-v stored in file system
_- Select volume to store key-value
- Redirect to correct volume node
- Consistent hashing?
- Add more node? remove node? healthcheck node?
