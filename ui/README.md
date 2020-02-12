## Installation
```$ npm run install```

## Development
#### Running client & backend together:
```$ npm run dev```

The application will run on [localhost:4999](localhost:4999)

(server will run on port localhost:5000)

## Tests
```$ npm run test```

## Environment variables
| Variable  | Description | Default
| ------------- | ------------- | -------------
| ```API_URL```  | StatusBay Api URL. If not set, The server will use mock data  |
| ```GELF_ADDRESS```  | Gelf server and port (separated by colon)  | `http://localhost:12201`
| ```LOG_LEVEL``` | log level to use | error 
```.env``` file is also supported.


## Docker
```docker build -t statusbay-ui .```

```d run -it --rm --name statusbay -e API_URL=yourdomain@example.com -e GELF_ADDRESS=yourdomain@example.com:1234  -p 3000:80 statusbay-u```
