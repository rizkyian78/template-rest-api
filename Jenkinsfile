pipeline {
    agent any
    environment {
        DOCKER_ACCOUNT="rizkyian78"
        DOCKER_PASSWORD="yukaritakeba12"
    }
    stages {
        stage('Building Docker') {

            steps {
           app = docker.build("getintodevops/hellonode")

            }
        }
    }
}
