// var PROTO_PATH = __dirname + '/../../../shared/protos/gerated/';
import { GreeterClient } from '../protos/hello_grpc_pb'
import { HelloRequest } from '../protos/hello_pb'
import { credentials, ServiceError } from '@grpc/grpc-js';


// let argv = 'world';
// if (process.argv.length >= 3) {
//     [, , argv] = process.argv;
// }

// const param: HelloRequest = new HelloRequest();
// param.setName(argv);
// param.setParamStruct(Struct.fromJavaScript({ foo: 'bar', bar: 'foo' }));
// param.setParamListValue(ListValue.fromJavaScript([{ foo: 'bar' }, { bar: 'foo' }]));
// param.setParamValue(Value.fromJavaScript('Any Value'));

async function main() {
    // var argv = parseArgs(process.argv.slice(2), {
    //     string: 'target'
    // });
    // var target;
    // if (argv.target) {
    //     target = argv.target;
    // } else {
    //     target = 'localhost:50051';
    // }
    var client = new GreeterClient('localhost:50051', credentials.createInsecure());
    var request = new HelloRequest();
    var user = 'world';
    request.setName(user);
    client.sayHello(request, function (err: any, response: { getMessage: () => any; }) {
        console.log('Greeting:', response.getMessage());
    });


    // client.sayHello(param, (err, res) => {
    //     if (err) {
    //         return console.error('sayBasic:', err.message);
    //     }

    //     console.info('sayBasic:', res.getMessage());
    // });
}

(async (): Promise<void> => {
    try {
        await main();
    } catch (err) {
        console.error(err);
    }
})();
