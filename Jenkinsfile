pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('testy') {
            steps {
                sh 'make compile'
            }
        }
    }
}
