pipeline {
    agent any
    stages {
        stage('Building Docker') {
            steps {
                echo 'Building Docker'
                docker build -t rizkyian78/test:0.2 .
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
