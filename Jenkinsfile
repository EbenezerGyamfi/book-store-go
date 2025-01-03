pipeline {
    agent {
        docker {
            image 'golang:1.20-alpine' // Replace with your desired Go version
        }
    }

    environment {
        DB_HOST = 'localhost' // Example environment variables, if needed
        DB_USER = 'user'
        DB_PASSWORD = 'password'
        DB_NAME = 'books'
    }

    stages {
        stage('Clone Repository') {
            steps {
                checkout scm
            }
        }

        stage('Install Dependencies') {
            steps {
                sh '''
                go mod tidy
                '''
            }
        }

        stage('Build') {
            steps {
                sh '''
                go build -o app ./cmd/main/main.go
                '''
            }
        }

        // stage('Test') {
        //     steps {
        //         sh '''
        //         go test ./... -v
        //         '''
        //     }
        // }

        // stage('Package') {
        //     steps {
        //         sh '''
        //         tar -czf app.tar.gz app
        //         '''
        //     }
        // }
    }

    post {
        always {
            echo 'Cleaning up workspace...'
            cleanWs()
        }
        success {
            echo 'Build completed successfully!'
        }
        failure {
            echo 'Build failed. Check logs for details.'
        }
    }
}
