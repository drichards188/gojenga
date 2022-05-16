// Load the AWS SDK for Node.js
var AWS = require('aws-sdk');
// Set the region
AWS.config.update({region: 'REGION'});

//todo url is pulled out for commit and must be added to make functioning
const QueueUrl = "SQS_URL_HERE"

function sqsSend() {
    // Create an SQS service object
    var sqs = new AWS.SQS({apiVersion: '2012-11-05'});

    var myData = {
        "MyData": {
            "Account": "david",
            "Amount": "13"
        }
    }

    myData = JSON.stringify(myData)

    var params = {
        // Remove DelaySeconds parameter and value for FIFO queues
        DelaySeconds: 10,
        MessageAttributes: {
            "Source": {
                DataType: "String",
                StringValue: "node lambda 1"
            },
            "Destination": {
                DataType: "String",
                StringValue: "trigger cycle"
            },
            "Identity": {
                DataType: "String",
                StringValue: "1.0.1"
            }
        },
        MessageBody: myData,
        // MessageDeduplicationId: "TheWhistler",  // Required for FIFO queues
        // MessageGroupId: "Group1",  // Required for FIFO queues
        QueueUrl: QueueUrl
    };

    sqs.sendMessage(params, function(err, data) {
        if (err) {
            console.log("Error", err);
        } else {
            console.log("Success", data.MessageId);
        }
    });
}

sqsSend()