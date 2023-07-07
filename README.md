# SVault: Secure Vault

## Description

The project is a desktop application that provides users with the ability to encrypt their files locally or securely share them. It leverages a forked version of the [cryptify](https://github.com/Oval-Personal-Data-Wallet/cryptify) crypto engine and is built using Go Wails.

## Features

- [x] Local File Encryption: Users can encrypt their files on their local machine to protect sensitive information.
- [] File sharing

## Technologies Used

- Go Wails: The project is built using Go Wails, which allows for creating desktop applications with Go and web technologies.
- Cryptify Crypto Engine: A forked version of the cryptify library is integrated to provide file encryption and decryption capabilities.

## Installation

1. Clone the repository: `git clone https://github.com/owbird/SVault.git`
2. Install Go Wails: Follow the installation instructions for Go Wails from the [official documentation](https://wails.io/docs/gettingstarted/installation).
3. Install project dependencies: `go get ./...`
4. Build the application: `make compile`
5. Run the application from `build/bin/<platform>`

## Contributing

Contributions to the project are welcome! If you encounter any issues or have suggestions for improvements, please open an issue on the project repository.

## License

[MIT License](LICENSE)

## Acknowledgments

The project builds upon the work of the [cryptify](https://github.com/Oval-Personal-Data-Wallet/cryptify) crypto engine developed by the Oval Personal Data Wallet team. We acknowledge and appreciate their contributions.
