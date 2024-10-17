# IAMWiz - AWS IAM Policy Generator

**IAMWiz** is a command-line tool that uses a LLM to generate AWS IAM policies based on user prompts. With **IAMWiz**, you can easily create custom IAM policies for various AWS services, including complex conditions, actions, and resources, all using natural language input.

## Features

- Generate IAM policies based on natural language prompts.
- Support for complex conditions, including IP address restrictions and region-based access.
- Interactive CLI for easy and continuous policy generation.

## Installation

### Prerequisites

- Go 1.18+ installed on your machine.
- An OpenAI API key for interacting with the LLM.

### Install from source

1. Clone the repository:
```bash
git clone https://github.com/tailwarden/iamwiz.git
cd iamwiz
go build -o iamwiz
./iamwiz
```

### Set Up OpenAI API Key

You need to set up your OpenAI API key to use the tool. Export the API key as an environment variable:

```bash
export OPENAI_API_KEY="your-api-key-here"
```

Usage
Once you've installed IAMWiz, you can start generating IAM policies by providing a prompt. Here's how to use the tool:


- Create a policy that allows read-only access to all S3 buckets.
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "dynamodb:PutItem",
        "dynamodb:UpdateItem",
        "dynamodb:DeleteItem",
        "dynamodb:BatchWriteItem",
        "dynamodb:BatchGetItem"
      ],
      "Resource": "arn:aws:dynamodb:REGION:ACCOUNT_ID:table/movies"
    },
    {
      "Effect": "Allow",
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": "arn:aws:logs:REGION:ACCOUNT_ID:*"
    }
  ]
}
````

### Future Features
- IAM Policy Validator: Integrate an IAM policy validation step that checks the generated policy for common errors or best practices (e.g., excessive permissions, wildcard * in actions/resources) before presenting it to the user.
Policy Linting: Validate and lint generated policies based on AWS best practices.
- Policy Validation Against AWS CloudTrail Logs: Generate least-privilege policies based on actual usage logs.

### Contributing

Contributions are welcome! To get started:
- Fork the repository.
- Create a new branch: git checkout -b feature/your-feature.
- Make your changes and commit them: git commit -m 'Add some feature'.
- Push to the branch: git push origin feature/your-feature.
- Open a pull request.

### License
This project is licensed under the MIT License - see the LICENSE file for details.