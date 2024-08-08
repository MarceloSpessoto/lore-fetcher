pipelineJob('kernel-compilation'){
  parameters {
    stringParam('PATCH', '', 'patch url')
  }
  definition {
    cps {
      script('''
        pipeline {
          agent {
              docker {
                  image 'ubuntu:latest'
              }
          }
          parameters {
              string(name: 'PATCH')
          }
          stages {
              stage('Install dependencies'){
                  steps {
                      sh 'apt update -y'
                      sh 'apt install -y b4 git bc binutils bison dwarves flex gcc git gnupg2 gzip libelf-dev libncurses5-dev libssl-dev make openssl pahole perl-base rsync tar xz-utils'
                  }
              }
              stage('Setup git configs'){
                  steps{
                      sh 'git config --global user.email "jenkins@ci.com"'
                      sh 'git config --global user.name "jenkins"'
                  }
              }
              stage('Test tinyconfig Compilation'){
                  steps {
                      sh 'rm -rf linux'
                      sh 'git clone --depth 1 "https://gitlab.freedesktop.org/agd5f/linux.git" linux'
                      dir("linux"){
                          sh "b4 shazam ${params.PATCH}"
                          sh 'make tinyconfig'
                          sh 'make'
                      }
                  }
              }
              stage('Cleanup'){
                  steps {
                      sh 'rm -rf linux'
                  }
              }
          }
      }
      '''.stripIndent())
      sandbox()
      authenticationToken('token')
    }
  }
}
