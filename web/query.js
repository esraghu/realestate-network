var Fabric_Client = require('fabric-client');
var fabric_client = new Fabric_Client();

// setup the fabric network
// @TODO: obtain the channel name through headers or querystring
var _channel = fabric_client.newChannel('mychannel');
var peer = fabric_client.newPeer('grpc://localhost:7051'); // port depends on the peer
_channel.addPeer(peer);

var query = {};

query.request = {
    chaincodeId : 'chaincode',
    fcn : 'list_record',
    args : JSON.stringify({HouseId : ''})
}

query.chaincode = (qData, callback) => {
    var chaincode = typeof(qData.chaincode) == 'string'
        ? qData.chaincode
        : false
    
    _channel.queryByChaincode(qData)
    .then((queryResponses => {
        console.log('Query has completed. Checking for results');
    
        if (queryResponses && queryResponses == 1) {
            if (queryResponses[0] instanceof Error) {
                console.error(`error from query = ${queryResponses}`);
            } else {
                console.log(`Query response = ${queryResponses[0].toString()}`);
            }
        } else {
            console.log('No payload received in response');
        }
    }))
    .catch((err) => {
        console.error(`Failed to query due to ${err}`);
    });
}
