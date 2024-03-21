# Welcome to your CDK Go project!

This is a blank project for CDK development with Go.

Prerequisites before attending the [Build Go Apps that Scale on AWS workshop](https://frontendmasters.com/workshops/fullstack-go-aws/) at Frontend Masters.

### Go installed (preferably version > = 1.18)

- Download and install Go: [Go Download Page](https://go.dev/dl/)

### AWS Installed (use AWS CLI)

- Step 1 : Install latest version of AWS [download](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
  - `aws-version`
- Step 2 : Configure the AWS user permissions
  - Create Admin user, get the acess_key_id and the secret_access_key
- Step 3 : Configure the AWS CLI
  - Configure [short or long term credentials ](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-quickstart.html)
  - Confirm with: `aws s3 ls` and `aws sts get-caller-identity`

### CDK installed (use CDK command)

- [Download CDK](https://docs.aws.amazon.com/cdk/v2/guide/getting_started.html#getting_started_install)

The `cdk.json` file tells the CDK toolkit how to execute your app.

## Useful commands

- `cdk deploy` deploy this stack to your default AWS account/region
- `cdk diff` compare deployed stack with current state
- `cdk synth` emits the synthesized CloudFormation template
- `go test` run unit tests
