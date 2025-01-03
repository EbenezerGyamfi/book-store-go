pipeline {
    agent {
        docker {
            image 'golang:latest'
            args '-u root:root'
        }
    }

    environment {
        DB_HOST = 'mysql'
        DB_USER = 'root'
        DB_PASSWORD = 'root'
        DB_NAME = 'books'
    }

    stages {
        stage('Setup MySQL') {
            steps {
                script {
                    docker.image('mysql:latest').run('-e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=books --name mysql -d mysql:latest')
                }
                sh 'sleep 15' // Wait for MySQL to initialize
            }
        }

        stage('Build') {
            steps {
                sh '''

                '''
            }
        }

      
    }

    post {
        always {
            script {
                // Clean up the MySQL container
                sh 'docker rm -f mysql || true'
            }
        }
    }
}
