# JavaScript (NodeJS) Example

This folder contains a working example API client for [Minim's API](https://my.minim.co/api_doc) written
using NodeJS.

Use this project as a starting point for building a NodeJS application that interacts with Minim's API. The `example.js`
file serves as a starting point with a few example API calls. 

Check the [Minim API Documentation](https://my.minim.co/api_doc) for a complete reference of available API calls. 

## Getting Started

1. Install a [recent version of NodeJS and npm](https://nodejs.org/en/download/)
2. Generate an [Application ID and Secret Key for your account on Minim](https://my.minim.co/api_keys)
3. Clone the repository or [download the latest code as a zip](https://github.com/MinimSecure/minim-api-examples/archive/main.zip)
4. Open a terminal and enter the directory you cloned or extracted the project to then enter the `js/` directory
5. Use npm to install the JavaScript dependencies
   ```
   npm install
   ```
6. Create a new file in the project directory called `.env` and put your Application ID and Secret Key in it 
   
   Use the text below and replace `<YOUR APPLICATION ID HERE>` with your Application ID and `<YOUR SECRET KEY HERE>` 
   with your Secret Key.
   ```
   APPLICATION_ID=<YOUR APPLICATION ID HERE>
   API_SECRET=<YOUR SECERT KEY HERE>
   ```
7. Run the example.js script with npm
   ```
   npm start
   ```
