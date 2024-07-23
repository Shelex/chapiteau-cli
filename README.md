# Chapiteau CLI

Chapiteau CLI is a command-line tool designed to upload playwright test reports / run data to the Chapiteau service.

## Installation

```bash
    npm install -g chapiteau-cli
```

## Usage

The Chapiteau CLI provides a command `upload` which is used to upload a playwright report or provide `index.html` file to the Chapiteau service.
That is defined by argument `--path`:

-   path is a <strong>report folder</strong> - report in that folder will be uploaded and saved in service;
-   path is a <strong>report `index.html` file</strong> - just index.html will be sent, run data parsed and file will be not stored anywhere.

If you want to have link to report hosted elsewhere (in case only data is parsed from index.html) - you can specify `--report-url`.

## Flags

-   `--path` (required): Path to the folder or file to upload.
-   `--url` (required): The URL of the Chapiteau service for specific team and project.
-   `--auth` (required): Api token for the Chapiteau service.
-   `--build-name`: CI Build Name (optional).
-   `--build-url`: CI Build url (optional).
-   `--report-url`: Playwright Report URL hosted elsewhere (optional).

## Examples

### Upload a Folder

To upload a folder containing a report:

```sh
chapiteau-cli upload --path "/path/to/folder" --url "https://chapiteau.shelex.dev/teamId/projectId" --auth "your_api_token" --build-url "build/url" --build-name "buildName"
```

In this case report will be hosted with Chapiteau.

### Upload a File

To upload a single file (e.g., index.html):

```sh
chapiteau-cli upload --path "/path/to/index.html" --url "https://chapiteau.shelex.dev/teamId/projectId" --auth "your_api_token" --build-url "build/url" --build-name "buildName" --report-url "https://github.pages.or.other.url"
```

In this case only data will be parsed from report file, but no report will be saved to service, it makes sense to provide a link if it is hosted elsewhere.
