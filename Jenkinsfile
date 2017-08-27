
node {
  stage('Checkout')

  checkout scm
  // checkout([$class: 'GitSCM', branches: [[name: '*/master']], doGenerateSubmoduleConfigurations: false, extensions: [], submoduleCfg: [], userRemoteConfigs: [[url: 'https://github.com/monkeylittle/hello-world']]])

  stage('Build')

  // build docker image
  app = docker.build('johnturner/hello-world')

  stage('Test')

  echo 'Testing...'

  stage('Release')

  // push docker image to docker hub
  echo 'Deploying...'
  docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-credentials') {
    app.push("${env.BUILD_NUMBER}")
    app.push("latest")
  }
}
