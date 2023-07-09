pipeline {
  agent any

  environment {
    AWS_ACCESS_KEY_ID = credentials('aws-access-key')
    AWS_SECRET_ACCESS_KEY = credentials('aws-secret-key')
    AWS_REGION = 'us-east-1'  // Replace with your desired AWS region
    ECR_REPOSITORY = 'your-repository-name'  // Replace with your ECR repository name
    DOCKER_IMAGE_TAG = 'your-image-tag'  // Replace with your desired image tag
  }

  stages {
    stage('Checkout') {
      steps {
        checkout scm
      }
    }

    stage('Build and Push Image') {
      steps {
        script {
          docker.withRegistry('https://$(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com', 'ecr:us-east-1:your-ecr-credentials') {
            def dockerImage = docker.build("$(ECR_REPOSITORY):$(DOCKER_IMAGE_TAG)", './Dockerfile-gateway')
            dockerImage.push()
          }
        }
      }
    }
  }
}
