const path = require("path");
const fs = require("fs");
const axios = require("axios");
require("axios-debug-log/enable");

/**
 * @param {URL} url
 * @param {ChapiteauArguments} input
 * @param {FormData} formData
 * @param {boolean} isFile
 */
async function upload(url, input, formData, isFile) {
    if (input.buildName) url.searchParams.append("buildName", input.buildName);

    if (input.buildUrl) url.searchParams.append("buildUrl", input.buildUrl);

    if (isFile && input.reportURL)
        url.searchParams.append("reportUrl", input.reportURL);

    try {
        const response = await axios.post(url.href, formData, {
            headers: {
                "X-API-Key": input.auth,
                "Content-Type": "multipart/form-data",
            },
        });

        if (response.status !== 200) {
            throw new Error(
                `Failed to upload file: ${
                    response.data && response.data.error
                        ? response.data.error
                        : response.statusText
                }`
            );
        }
    } catch (error) {
        const message =
            error.response && error.response.data && error.response.data.error
                ? error.response.data.error
                : error.message;

        throw new Error(`Upload failed: ${message}`);
    }
}

/**
 * @param {ChapiteauArguments} input
 */
async function uploadIndexHtml(input) {
    const fileName = path.basename(input.path);
    if (fileName !== "index.html") {
        throw new Error("File should be html file of playwright report");
    }

    const formData = new FormData();
    const content = fs.readFileSync(input.path);
    formData.append("file", new Blob([content]), fileName);

    const url = new URL(`${input.url}/upload/info`);

    await upload(url, input, formData, true);
    console.log("Index.html uploaded successfully");
}

/**
 * @param {ChapiteauArguments} input
 */
async function uploadReport(input) {
    const formData = new FormData();
    let hasIndexHtml = false;

    const files = fs.readdirSync(input.path, { withFileTypes: true });
    for (const file of files) {
        if (!file.isDirectory()) {
            const filePath = path.join(input.path, file.name);
            const fileContent = fs.readFileSync(filePath);
            const blob = new Blob([fileContent]);
            formData.append("files", blob, file.name);

            if (file.name === "index.html") {
                hasIndexHtml = true;
            }
        }
    }

    if (!hasIndexHtml) {
        throw new Error("index.html file not found in the provided folder");
    }

    const url = new URL(`${input.url}/upload/report`);

    await upload(url, input, formData, false);
    console.log("Playwright report uploaded successfully");
}

/**
 * @param {string} path
 * @returns {boolean}
 */
function isDirectory(path) {
    try {
        const stats = fs.statSync(path);
        return !stats.isFile();
    } catch (error) {
        throw new Error(
            `Failed to check if path is a directory: ${error.message}`
        );
    }
}

module.exports = {
    uploadIndexHtml,
    uploadReport,
    isDirectory,
};
