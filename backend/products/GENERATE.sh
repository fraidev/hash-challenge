BASEDIR=$(dirname "$0")
cd ${BASEDIR}

SRC="./"
PROTO_DEST=./gerated

mkdir -p ${PROTO_DEST}

# yarn run grpc_tools_node_protoc \
#     --plugin=protoc-gen-grpc=./node_modules/.bin/grpc_tools_node_protoc_plugin \
#     --js_out=import_style=commonjs,binary:${PROTO_DEST} \
#     --grpc_out=${PROTO_DEST} \
#     -I ${SRC} \
#     ${SRC}/*.proto

# yarn run grpc_tools_node_protoc \
#     --plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
#     --ts_out=:${PROTO_DEST} \
#     -I ${SRC} \
#     ${SRC}/*.proto


# PROTOC="./node_modules/grpc-tools/bin/protoc"
# if test ! -f "$PROTOC"; then
#   PROTOC="../../node_modules/grpc-tools/bin/protoc"
# fi
# $PROTOC --go_out=plugins=grpc,paths=import:${PROTO_DEST} -I ${SRC} ./*.proto

node_modules/grpc-tools/bin/protoc --go_out=plugins=grpc,paths=import:${PROTO_DEST} -I ${SRC} ./*.proto