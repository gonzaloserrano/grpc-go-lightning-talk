var messages = require('./clock_pb');
var services = require('./clock_grpc_pb');

var grpc = require('grpc');

function main() {
    var client = new services.ClockClient(
        'localhost:8888',
        grpc.credentials.createInsecure()
    );
    client.whatTimeIsIt({foo: "hey"}, function(err, clockResponse) {
        console.log('time: ', clockResponse.getMessage());
    });
}

main();


