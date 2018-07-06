var Fabric_Client = require('fabric-client');
var fabric_client = new Fabric_Client();

// setup the fabric network
// @TODO: obtain the channel name through headers or querystring
var channel = fabric_client.newChannel('mychannel');
var peer = fabric_client.newPeer('grpc://localhost:7051'); // port depends on the peer
channel.addPeer(peer);

