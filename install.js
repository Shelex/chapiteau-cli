const https = require("https");
const fs = require("fs");
const path = require("path");

const version = process.env.npm_package_version;
const platform = process.platform;
const arch = process.arch;

const appName = "chapiteau-cli";
const extension = platform === "win32" ? `.exe` : "";
const binaryName = `${appName}${extension}`;

const url = `https://github.com/Shelex/chapiteau-cli/releases/download/v${version}/${appName}-${platform}-${arch}${extension}`;

const binaryPath = path.join(__dirname, binaryName);

https
    .get(url, (response) => {
        if (response.statusCode !== 200) {
            console.error(
                `Failed to download ${appName}: ${response.statusCode}`
            );
            return;
        }
        const file = fs.createWriteStream(binaryPath);
        response.pipe(file);
        file.on("finish", () => {
            file.close();
            fs.chmodSync(binaryPath, "755");
            console.log(`${appName} downloaded and installed successfully`);
        });
    })
    .on("error", (err) => {
        console.error(`Error downloading ${appName}: ${err.message}`);
    });
