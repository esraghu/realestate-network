'use strict';
const shim = require('fabric-shim');
const handlers = require('./handlers');
const util = require('util');
const debug = util.debuglog('chaincode');

const Chaincode = class {

    async Init(stub) {
        debug('Chaincode instantiated successfully!');
        return shim.success();
    };

    async Invoke(stub) {
        // retrieve the functions and arguments that are to be executed on the chaincode
        let req = stub.getFunctionAndParameters();
        debug(`The function ${req.fcn} has been called with ${req.params} by ${stub.getCreator()}`);

        // check if this is one of allowed functions
        invokedFunction = false;
        if (this.allowedBuilderFunctions.includes(req.fcn)) invokedFunction = handlers.builder[req.fcn];
        if (this.allowedAgencyFunctions.includes(req.fcn)) invokedFunction = handlers.agency[req.fcn];
        if (this.allowedRegulatorFunctions.includes(req.fcn)) invokedFunction = handlers.regulator[req.fcn];
        
        if (!invokedFunction) return shim.error('The function is not allowed');

        // now that we have a valid function requested, we shall call it from the handlers
        invokedFunction(req.params, (err, payload) => {
            err? shim.error(err) : shim.success(payload)
        })
    }

    allowedBuilderFunctions = [
        'addProject',   // add a new project that belongs to that builder
        'updateProject', // update any particular details about the project
        'listProject', // list various projects for that builder
        'initiateAgencyInspection', // allow the builder to initiate agencies to inspect various aspects
        'reviewAgencyInspection' // review the progress of the agencies
    ]

    allowedAgencyFunctions = [
        'listInspectionRequests', // look up all the inspection requests and their status
        'updateInspectionRequests' // update the status of inspection requests
    ]

    allowedRegulatorFunctions = [
        'listBuilders', // list all the affiliated builders
        'listProjectsByBuilder', // list all the projects by a particular builder and check their progress
        'updateProject' // update the project with approval, rejection or an inquiry
    ]
};

// initialize the chaincode here
shim.start(new Chaincode());