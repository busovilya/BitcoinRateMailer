pipeline {
    agent any
    
    tools { go '1.20' }
    
    stages {
        stage('Preparation') {
            steps {
                sh 'go version'
                sh 'go mod download'
            }
        }
        
        stage('Build') {
            steps {
                sh 'go build .'
            }
        }
        
        stage('Test') {
            steps {
                sh 'go test ./... -v'
            }
        }
    }
}