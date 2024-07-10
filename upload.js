const path = require("path");
const fs = require("fs");
const axios = require("axios");

/**
 * @param {URL} url
 * @param {ChapiteauArguments} input
 * @param {FormData} formData
 * @param {boolean} isFile
 */
async function send(url, input, formData, isFile) {
    if (input.buildName) url.searchParams.append("buildName", input.buildName);

    if (input.buildUrl) url.searchParams.append("buildUrl", input.buildUrl);

    if (isFile && input.reportURL)
        url.searchParams.append("reportUrl", input.reportURL);

    try {
        console.log(
            `sending query to ${JSON.stringify(
                {
                    mathod: "POST",
                    url: url.href,
                    data: formData,
                    headers: {
                        Authorization: `Bearer ${input.auth}`,
                        "Content-Type": "multipart/form-data",
                    },
                },
                null,
                2
            )}`
        );

        const response = await axios.post(url.href, formData, {
            headers: {
                Authorization: `Bearer ${input.auth}`,
                "Content-Type": "multipart/form-data",
            },
        });

        if (response.status !== 200) {
            throw new Error(`Failed to upload file: ${response.status}`);
        }
    } catch (error) {
        //console.log(error)
        throw new Error(`Request failed: ${error.message}`);
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

    const url = new URL(`${input.url}/api/project/${input.project}/file`);

    await send(url, input, formData, true);
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
        console.log(`checking file ${file.path}`);
        if (!file.isDirectory()) {
            console.log(`is not directory`);
            const filePath = path.join(input.path, file.name);
            console.log(`reading ${filePath}...`);
            const fileContent = fs.readFileSync(filePath);
            const blob = new Blob([fileContent]);
            console.log(`got fileContent: ${fileContent.length}`);
            formData.append("files", blob, file.name);

            if (file.name === "index.html") {
                hasIndexHtml = true;
            }
        }
    }

    if (!hasIndexHtml) {
        throw new Error("index.html file not found in the provided folder");
    }

    const url = new URL(`${input.url}/api/project/${input.project}/report`);
    console.log(`got url: ${input.url}/api/project/${input.project}/report`);

    await send(url, input, formData, false);
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
