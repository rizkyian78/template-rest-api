pipeline {
    agent any
    stages {
        stage('Building Docker') {
            steps {
                echo 'Building Docker'
                 docker { image 'golang:alpine' }
            }
        }
    }
}
