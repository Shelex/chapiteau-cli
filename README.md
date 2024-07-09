# Chapiteau CLI

Chapiteau CLI is a command-line tool designed to upload playwright test reports or data to the Chapiteau service. 

## Installation

//TODO
To install the Chapiteau CLI...

## Usage
The Chapiteau CLI provides a command `upload` which is used to upload a playwright report or `index.html` file to the Chapiteau service.  
It supports just sending general data about your test run (if a single file `index.html` specified) or an entire folder - then it will be available in the service as well.  
In case you have report hosted elsewhere, you can still link it with report by providing report url.

## Flags
 - `--path` (required): Path to the folder or file to upload.
 - `--url` (required): The URL of the Chapiteau service.
 - `--auth` (required): Authentication token for the Chapiteau service.
 - `--project` (required): Project ID - you can take it from chapiteau service when new project is created.
 - `--build-id`: CI Build ID (optional).
 - `--build-name`: CI Build Name (optional).
 - `--report-url`: Playwright Report URL hosted elsewhere (optional).

## Examples
### Upload a Folder

To upload a folder containing a report:
```sh
chapiteau-cli upload --path "/path/to/folder" --url "https://chapiteau.shelex.dev" --auth "your_token" --project "projectID" --build-id "buildID" --build-name "buildName"
```

In this case report url is not needed, as report will be hosted with help of chapiteau.

### Upload a File
To upload a single file (e.g., index.html):
```sh
chapiteau-cli upload --path "/path/to/index.html" --url "https://chapiteau.shelex.dev" --auth "your_token" --project "projectID" --build-id "buildID" --build-name "buildName" --report-url "https://github.pages.or.other.url"
```