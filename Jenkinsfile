pipeline {
    agent any
    environment {
        DOCKER_ACCOUNT="rizkyian78"
        DOCKER_PASSWORD="yukaritakeba12"
    }
    stages {
        stage('Building Docker') {
            agent {
                docker {
                    image 'gradle:6.7-jdk11'
                }
            }
            steps {
                sh 'gradle --version'
            }
        }
    }
}
