# Playground

## Introduction
Playground is a backend engine that allows you to execute code in various programming languages directly in your browser. It supports a wide range of languages.

Find the UI for it in this [repo](https://github.com/Sudarshan-PR/playground-ui).

Languages available:

- Go. 

lol yeah, that's not "wide range" I know... for now that is... support more languages will be added soon.

## Demo

[Playground](https://sudarshan-pr.github.io/playground-ui/) 
## Installation

You are expected to have a RabbitMQ server running.
This project is devided into multiple services:
* **Gateway**
* **Notifications**
* ***-playground** (go-playground)
All of the 3 needs to be running.

There are kubernetes manifest files in directory of each service. Modify the ingress to add your desired hostname and the ConfigMaps with the RabbitMQ credentials/URLs and apply them to your cluster.

Proto Files for the Notifications and *-playground are stored in this [repo](https://github.com/Sudarshan-PR/playground-protos).
## Usage/Examples

Send code to compile: 
```bash
curl --request POST \
  --url http://api.sudarshan.cloud/compile \
  --header 'Content-Type: application/json' \
  --data '{
	"language": "go",
	"userid": "random_string",
	"code": "packagemain\nimport \"fmt\"\nfunc main() {\n    fmt.Println(\"hello world\")\n}"
}'
```

To receive output: (Using `websocat`, a command-line client for WebSockets)
```bash
websocat ws://ws.sudarshan.cloud/ws?user=random_string
```
## Contributing

Code Compiler is an open-source project, and we welcome contributions from the community. To contribute, please follow these guidelines:

1. Create a GitHub issue to discuss your proposed contribution.
2. Fork the project repository and make your changes in a separate branch.
3. Submit pull requests for your changes, ensuring they follow the project's coding standards.


## Additional Notes

Code Compiler is still under active development, and new features and language support will be added in the future. We encourage you to check back regularly for updates.
## License

[MIT](https://choosealicense.com/licenses/mit/)
