
pipeline {

  agent any

  stages {
    stage('Checkout') {
      steps {
        checkout scm
      }
    }

    stage('Build') {
      steps {
        echo 'Building...'
      }
    }

    stage('Test') {
      steps {
        echo 'Testing...'
      }
    }

    start('Deploy') {
      steps {
        echo 'Deploying...'
      }
    }
  }
}
