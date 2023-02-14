pipeline {
    agent any
    stages {
        stage('Example Build') {
            steps {
               sh "git tag --sort version:refname | tail -1 > version.tmp"
               String tag = readFile 'version.tmp'
                echo tag
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
