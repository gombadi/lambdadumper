# lambdadumper
Small utility to dump out all inputs given to Lambda Functions and send them to an SNS topic

## Usage
NOTE - To upload the zip file to AWS Lambda you need to have the aws commandline toll installed and permissions to upload to Lambda.

Create a Lambda function and give it a role that allows sns:publish to the SNS topic you want to use.

Run the buildme.sh script in the repo which will compile the code, create the zip file and upload to the AWS Lambda function. On the AWS web console you can then test the function or connect it to an event source and trigger the function.

The testit.sh script can be used to test the compiled Go binary locally without running in AWS Lambda.


