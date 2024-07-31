#!/usr/bin/env node
const { Command } = require("commander");
const { uploadIndexHtml, uploadReport, isDirectory } = require("./upload");

const program = new Command();

program
    .name("chapiteau-cli")
    .description("Upload report/data to the chapiteau service")
    .version("1.3.0");

program
    .command("upload")
    .description("Upload a report or just index.html file")
    .requiredOption("--path <path>", "Path to the folder or file to upload")
    .requiredOption(
        "--url <url>",
        "Chapiteau service url for specific team and project"
    )
    .requiredOption("--auth <auth>", "Authentication token")
    .option("--build-url <buildUrl>", "CI Build URL")
    .option("--build-name <buildName>", "CI Build Name")
    .option(
        "--report-url <reportURL>",
        "Playwright Report URL hosted elsewhere"
    )
    .action(async (options) => {
        /**
         * @type {ChapiteauArguments}
         */
        const args = {
            path: options.path,
            url: options.url,
            auth: options.auth,
            buildUrl: options.buildUrl,
            buildName: options.buildName,
            reportURL: options.reportUrl,
        };

        try {
            const isDir = isDirectory(args.path);
            if (isDir) {
                await uploadReport(args);
            } else {
                await uploadIndexHtml(args);
            }
        } catch (error) {
            console.error(error.message);
        }
    });

program.parse(process.argv);
