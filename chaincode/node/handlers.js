/* This is where all the various functions that need to execute on the chaincode reside */

// dependencies

// create a container for handlers
var handlers = {};

// Builder functions
handlers.builder.addProject = (data, callback) => {
    callback(false, 'This function has been called successfully!');
};

handlers.builder.updateProject = (data, callback) => {
    callback(false, 'This function has been called successfully!');
};

handlers.builder.listProject = (data, callback) => {
    callback(false, 'This function has been called successfully!');
};

handlers.builder.initiateAgencyInspection = (data, callback) => {
    callback(false, 'This function has been called successfully!');
};

handlers.builder.reviewAgencyInspection = (data, callback) => {
    callback(false, 'This function has been called successfully!');
};
// export module
module.exports = handlers;