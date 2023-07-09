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
                    sh 'docker'
                    sh 'aws ecr get-login-password'
                    sh "aws ecr get-login-password - region ${AWS_DEFAULT_REGION} | docker login - username AWS - password-stdin ${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_DEFAULT_REGION}.amazonaws.com"
                }
            }
        }
    }
}