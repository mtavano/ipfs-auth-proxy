# IPFS auth proxy
--

This project is composed in two parts:

1. Api: The main tasks for the api are administration and secure content delivery given an api key.

2. Web: This UI holds has a simple login and panel to create and manage api keys.


### Project development:

The project was developed using a raspberry pi 4 to run IPFS node and also the API. To test every API feature externally I used a laptop running the UI + insomnia to fetch ipfs content securely.

#### How to run the project.

Once you have your IPFS node running, yo are ready to run the binary made to your OS located in `bin` folder.

Finally, adjust your desired env vars in `.env` for api and `web/.env` for IU, and run as follows:

```
$ export $(cat .env) && ./ipfs-proxy-macos

...

$ cd web
$ yarn install
$ yarn start
```

### Technical questions:
How would you improve this assignment for a production ready solution (e.g., security,
deployment)? Describe IPFS and compare it to other protocols e.g., HTTP?

To turn this project into a production ready solution I would suggest the following things:
 
 **Backend:**
 
1. Improve database pacakge and use disk persistance to store data like psql, leveldb, etc. 
2. Improve auth strategy and use jwt to secure the api.
3. Improve api errors. Attach the stacktrace in staging and idea and error codes in prod could be an option.
4. Add logs and errors monitoring.
5. Dockerize project.
6. Add tests.

**Frontend:**

1. Improve auth and routing.
2. Add a css framework.
3. Improve styling and make a prettier UI.
4. Improve usability.
5. Add tests.

**Overall:**

1. Use preferred cloud provider to deploy UI and API.
2. Setup CI/CD pipeline

Also, the main difference between HTTP and IPFS are basically the protocols by itself. While HTTP has a centralized server approach and delivers the content specifically requesting the address where data is attached, IPFS has a descentralized network oriented approach and the data is requested using the data hash. For example, if a server that stores an image is down, the image cannot be longer requested until the server is up and running once again, that file into the IPFS network can be delivered no matter if a node in specific is down because that file is stored into many other nodes.

> is importante to mention that to run IPFS you need to access the node using the HTTP to IPFS portal. Also you can run locally a node as is done in this project.


### Api docs:

#### **POST /admin/login**

No header required.

Request

```json
{
	"user": "odin",
	"password": "odin123"
}
```

Response

```json
{
	"token": "Adfgasf2"
}
```

Errors

```
201 Token created
400 Bad request
```

#### **GET /admin/apikeys**

This endpoint requires `x-admin-token` header to be requested.

Response

```json
{
	"apiKeys": [
		{
			"key": "test",
			"enabled": false,
			"requests": 3,
			"bytesTransfered": 123
		},
		{
			"key": "sfcXyBZeL",
			"enabled": true,
			"requests": 0,
			"bytesTransfered": 0
		}
	]
}
```

Errors

```
200 OK
401 Unauthorized
```

#### **POST /admin/apikeys**

This endpoint requires `x-admin-token` header to be requested.

Response

```json
{
	"key": "sfcXyBZeL",
	"enabled": true,
	"requests": 0,
	"bytesTransfered": 0
}
```

Errors

```
201 Token created
400 Bad request
401 Unauthorized
```

#### **PATCH /admin/apikeys**

This endpoint requires `x-admin-token` header to be requested.

Request

```json
{
	"key": "sfcXyBZeL",
	"enabled": true, // true|false in case to enable or disable api key
}
```
Response

```json
{
	"key": "sfcXyBZeL",
	"enabled": true,
	"requests": 0,
	"bytesTransfered": 0
}
```

Errors

```
201 Token created
400 Bad request
401 Unauthorized
```

#### **GET /v1/files/:cid**

This endpoint requires `x-api-key` header to be requested.

Request

```
example cid: QmU9D6ENLxzZojQVxiM3PUns8992crpUdzKLG5bSmEFNuu
```
Response

```
// file content

This is an example file

```

Errors

```
201 Token created
400 Bad request
401 Unauthorized
```