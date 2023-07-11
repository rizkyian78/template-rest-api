pipeline {
    agent any
    environment {
        AWS_ACCOUNT_ID="YOUR_ACCOUNT_ID_HERE"
        AWS_DEFAULT_REGION="CREATED_AWS_ECR_CONTAINER_REPO_REGION" 
        IMAGE_REPO_NAME="ECR_REPO_NAME"
        IMAGE_TAG="latest"
        REPOSITORY_URI = "716294141291.dkr.ecr.us-east-1.amazonaws.com/test"
    }
    stages {
        stage('Logging into AWS') {
            steps {
                script {
                    sh 'whoami'
                    sh "aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 716294141291.dkr.ecr.us-east-1.amazonaws.com"
                }
            }
        }
        stage("Build Docker") {
            steps {
                script {
                    dockerImage = docker.build "test:0.0.1"
                }
            }
        }
        stage("Push to ECR") {
            steps {
                script {
                    sh "docker tag test:0.0.1:
                   ${REPOSITORY_URI}:0.0.1"
                    sh "docker push 716294141291.dkr.ecr.us-east-1.amazonaws.com/test:0.0.1"
                }
            }
        }
    }
}