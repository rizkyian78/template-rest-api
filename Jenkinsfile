pipeline {
    agent any
    stages {
        stage('Building Docker') {
            steps {
                echo 'Building Docker'
            }
        }
        stage('Example Deploy') {
            when {
                branch 'production'
            }
            steps {
                echo 'Deploying'
            }
        }
    }
}
