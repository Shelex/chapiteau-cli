# Chapiteau CLI

Chapiteau CLI is a command-line tool designed to upload playwright test reports / run data to the [Chapiteau](https://github.com/Shelex/chapiteau) service.  
Website: [chapiteau.shelex.dev](https://chapiteau.shelex.dev/)

## Installation

```bash
    # easier to use just via npx:
    npx chapiteau-cli

    # or if you really want to install:
    npm install -g chapiteau-cli
```

## Usage

The Chapiteau CLI provides a command `upload` which is used to upload a playwright report or provide `index.html` file to the Chapiteau service.

That is defined by argument `--path`:

-   path is a <strong>report folder</strong> - report in that folder will be uploaded and saved in service;
-   path is a <strong>report `index.html` file</strong> - just index.html will be sent, run data parsed and file will be not stored anywhere.

If you want to have link to report hosted elsewhere (in case only data is parsed from index.html) - you can specify `--report-url`.

## Flags

-   (required!) `--path` : Path to the folder or `index.html` playwright report file to upload.
-   (required!) `--url` : The URL of the Chapiteau service for specific team and project (take from ui page).
-   (required!) `--auth`: Api token for the Chapiteau service (could be created via ui).
-   `--build-name`: CI Build Name (optional).
-   `--build-url`: CI Build url (optional).
-   `--report-url`: Playwright Report URL hosted elsewhere (optional).

## Examples

### Upload a Folder

To upload a folder containing a report:

```sh
npx chapiteau-cli upload --path "/path/to/folder" --url "https://chapiteau.shelex.dev/teamId/projectId" --auth "your_api_token"
# example
npx chapiteau-cli upload --path "./playwright-report" --url "https://chapiteau.shelex.dev/api/teams/870ddb60-b3ac-4bea-8f83-94c2d6577650/5d2f0dcb-4c4c-49f5-b14d-28b689c5fd54" --auth "fb8c36be-5923-4bae-bcc3-3a16090c9561"
```

In this case report will be hosted with Chapiteau.

### Upload a File

To upload a single file (e.g., index.html):

```sh
npx chapiteau-cli upload --path "/path/to/index.html" --url "https://chapiteau.shelex.dev/teamId/projectId" --auth "your_api_token" --report-url "https://github.pages.or.other.url"
# example
npx chapiteau-cli upload --path "./playwright-report/index.html" --url "https://chapiteau.shelex.dev/api/teams/870ddb60-b3ac-4bea-8f83-94c2d6577650/5d2f0dcb-4c4c-49f5-b14d-28b689c5fd54" --auth "fb8c36be-5923-4bae-bcc3-3a16090c9561" --report-url "https://shelex.github.io/pw-tests-with-gh-pages/5"
# or without report
npx chapiteau-cli upload --path "./playwright-report/index.html" --url "https://chapiteau.shelex.dev/api/teams/870ddb60-b3ac-4bea-8f83-94c2d6577650/5d2f0dcb-4c4c-49f5-b14d-28b689c5fd54" --auth "fb8c36be-5923-4bae-bcc3-3a16090c9561"
```

In this case only data will be parsed from report file, but no report will be saved to service, it makes sense to provide a link if it is hosted elsewhere.
