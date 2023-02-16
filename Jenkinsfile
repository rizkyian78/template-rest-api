pipeline {
    agent any
    stages {
        stage('Building Docker') {
            steps {
                echo 'Building Docker'
                 docker { image 'golang:alpine' }
            }
        }
        stage('Example Deploy') {
            steps {
                sh 'go version'
            }
        }
    }
}
