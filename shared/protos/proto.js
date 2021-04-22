const path = require('path');
const shell = require('shelljs');
const rimraf = require('rimraf');

process.env.PATH += (path.delimiter + path.join(process.cwd(), 'node_modules', '.bin'));
const PROTO_DIR = path.join(__dirname, './');
const DIST_DIR = path.join(__dirname, './dist-proto');
const PROTOC_GEN_TS_PATH = path.join(__dirname, './node_modules/.bin/protoc-gen-ts');
// const PROTOC = "../node_modules/grpc-tools/bin/protoc"

rimraf.sync(`${DIST_DIR}/*`);

const protoConfig = [
    `--plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" `,
    `--grpc_out="grpc_js:${DIST_DIR}" `,
    `--js_out="import_style=commonjs,binary:${DIST_DIR}" `,
    `--ts_out="grpc_js:${DIST_DIR}" `,
    `--proto_path ${PROTO_DIR} ${PROTO_DIR}/*.proto`
];

//JS/TS Prototypes
shell.exec(`grpc_tools_node_protoc ${protoConfig.join(' ')}`);


//Go Prototypes
shell.exec(`protoc --go_out=plugins=grpc,paths=import:${DIST_DIR} -I "./" ./*.proto`);
