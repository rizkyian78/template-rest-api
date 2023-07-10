pipeline {
    agent any
    environment {
        AWS_ACCOUNT_ID="YOUR_ACCOUNT_ID_HERE"
        AWS_DEFAULT_REGION="CREATED_AWS_ECR_CONTAINER_REPO_REGION" 
        IMAGE_REPO_NAME="ECR_REPO_NAME"
        IMAGE_TAG="latest"
        REPOSITORY_URI = "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_DEFAULT_REGION}.amazonaws.com/${IMAGE_REPO_NAME}"
    }
    stages {
        stage('Logging into AWS') {
            steps {
                script {
                    sh 'whoami'
                    sh 'docker build -f Dockerfile-gateway -t test:latest .'
                    sh "aws ecr get-login-password - region us-east-1 | docker push docker push public.ecr.aws/u5i8d6t1/test:latest"
                }
            }
        }
    }
}