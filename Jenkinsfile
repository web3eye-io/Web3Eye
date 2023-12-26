pipeline {
  agent any
  environment {
    GOPROXY = 'https://goproxy.cn,direct'
    // env info from hack/set-golang-env.sh
    GOVERSION = "1.19.13"
    GOTMPENV = "~/.golang/$GOVERSION"
    GOROOT = "$GOTMPENV/goroot"
    GOPATH = "$GOTMPENV/gopath"
    GOBIN = "$GOROOT/bin"
    PATH = "$GOBIN:$PATH:$GOBIN"
    GO111MODULE = "on"
  }
  stages {
    stage('Clone') {
      steps {
        sh(returnStdout: true, script: '''
          git tag -l | xargs git tag -d
          git fetch origin --prune
          echo "update tags for repo"
        '''.stripIndent())
        git(url: scm.userRemoteConfigs[0].url,credentialsId: 'KK-github-key', branch: '$BRANCH_NAME', changelog: true, poll: true)
      }
    }
    stage('Prepare Golang ENV') {
      when {
        expression { BUILD_TARGET == 'true' }
      }
      steps {
        sh 'make prepare-golang-env'
      }
    }

    stage('Prepare') {
      when {
        expression { BUILD_TARGET == 'true' }
      }
      steps {
        sh 'make deps'
      }
    }
       
    stage('Linting') {
      when {
        expression { BUILD_TARGET == 'true' }
      }
      steps {
        sh 'make verify'
      }
    }

    stage('Compile') {
      when {
        expression { BUILD_TARGET == 'true' }
      }
      steps {
        sh (returnStdout: false, script: '''
          TAG=latest make build
        '''.stripIndent())
      }
    }

    // TODO: support UT
    stage('Unit Tests') {
      when {
        expression { BUILD_TARGET == 'true' }
      }
      steps {
        sh (returnStdout: false, script: '''
          echo "TODO:will make it ok"
        '''.stripIndent())
      }
    }

    stage('Generate docker image for dev') {
      when {
        expression { BUILD_TARGET == 'true' }
        expression { BRANCH_NAME == 'master' }
      }
      steps {
        sh 'TAG=latest DOCKER_REGISTRY=$DOCKER_REGISTRY make build-docker'
      }
    }
    
    stage('Release docker image for dev') {
      when {
        expression { RELEASE_TARGET == 'true' }
        expression { BRANCH_NAME == 'master' }
      }
      steps {
        sh 'TAG=latest DOCKER_REGISTRY=$DOCKER_REGISTRY make release-docker'
      }
    }

    stage('Generate docker image for feature') {
      when {
        expression { BUILD_TARGET == 'true' }
        expression { BRANCH_NAME != 'master' }
      }
      steps {
        sh 'make verify-build'
        sh(returnStdout: false, script: '''
          feature_name=`echo $BRANCH_NAME | awk -F '/' '{ print $2 }'`
          TAG=$feature_name DOCKER_REGISTRY=$DOCKER_REGISTRY make build-docker
        '''.stripIndent())
      }
    }

    stage('Release docker image for feature') {
      when {
        expression { RELEASE_TARGET == 'true' }
        expression { BRANCH_NAME != 'master' }
      }
      steps {
         sh(returnStdout: false, script: '''
          feature_name=`echo $BRANCH_NAME | awk -F '/' '{ print $2 }'`
          TAG=$feature_name DOCKER_REGISTRY=$DOCKER_REGISTRY make release-docker
        '''.stripIndent())
      }
    }

    stage('Tag') {
      when {
        anyOf{
          expression { TAG_MAJOR == 'true' }
          expression { TAG_MINOR == 'true' }
          expression { TAG_PATCH == 'true' }
        }
        anyOf{
          expression { TAG_FOR == 'testing' }
          expression { TAG_FOR == 'production' }
        }
      }
      steps {
        sh(returnStdout: true, script: '''
          set +e
          revlist=`git rev-list --tags --max-count=1`
          rc=$?
          set -e

          major=0
          minor=0
          patch=-1
          if [ "$TAG_FOR" == 'testing' ]; then
            patch=0
          fi
          
          if [ 0 -eq $rc ]; then
            tag=`git describe --tags $revlist`
            major=`echo $tag | awk -F '.' '{ print $1 }'`
            minor=`echo $tag | awk -F '.' '{ print $2 }'`
            patch=`echo $tag | awk -F '.' '{ print $3 }'`
          fi

          if [ "$TAG_MAJOR" == 'true' ]; then
            major=$(( $major + 1 ))
            minor=0
          elif [ "$TAG_MINOR" == 'true' ]; then
            minor=$(( $minor + 1 ))
          fi    

          case $TAG_FOR in
            testing)
              patch=$(( $patch + $patch % 2 + 1 ))
              ;;
            production)
              patch=$(( $patch + ( $patch +  1 ) % 2 + 1 ))
              git reset --hard
              git checkout $tag
              ;;
          esac

          tag=$major.$minor.$patch
          
          git tag -a $tag -m "Bump version to $tag"
        '''.stripIndent())

        withCredentials([gitUsernamePassword(credentialsId: 'KK-github-key', gitToolName: 'git-tool')]) {
          sh 'git push --tag'
        }
      }
    }
    
    
    stage('Pick tag version for testing') {
      when {
        anyOf{
          expression { RELEASE_TARGET == 'true' }
          expression { BUILD_TARGET == 'true' }
        }
        expression { TAG_FOR == 'testing' }
      }
      steps {
        sh(returnStdout: false, script: '''
          set +e
          # sync remote tags
          git tag -l | xargs git tag -d
          git fetch origin --prune
        '''.stripIndent())

        script {
          env.TAG_VERSION = sh(returnStdout: true,
            script: 'git tag|grep \'[13579]$\'|sort -V|tail -n 1| tr -d \'\\n\''
            )
        }
      }
    }

    stage('Pick tag version for production') {
      when {
        anyOf{
          expression { RELEASE_TARGET == 'true' }
          expression { BUILD_TARGET == 'true' }
        }
        expression { TAG_FOR == 'production' }
      }
      steps {
        sh(returnStdout: false, script: '''
          set +e
          # sync remote tags
          git tag -l | xargs git tag -d
          git fetch origin --prune
        '''.stripIndent())

        script {
          env.TAG_VERSION = sh(returnStdout: true,
           script: 'git tag|grep \'[02468]$\'|sort -V|tail -n 1| tr -d \'\\n\''
            )
        }
      }
    }

    stage('Generate docker image for test or prod') {
      when {
        expression { BUILD_TARGET == 'true' }
        anyOf{
          expression { TAG_FOR == 'testing' }
          expression { TAG_FOR == 'production' }
        }
      }
      steps {
        sh(returnStdout: false, script: '''
          set +e
          git reset --hard
          git checkout $TAG_VERSION
        '''.stripIndent())
        sh 'TAG=$TAG_VERSION DOCKER_REGISTRY=$DOCKER_REGISTRY make build-docker'
      }
    }
    
    stage('Release docker image for test or prod') {
      when {
        expression { RELEASE_TARGET == 'true' }
        anyOf{
          expression { TAG_FOR == 'testing' }
          expression { TAG_FOR == 'production' }
        }
      }
      steps {
        sh 'TAG=$TAG_VERSION DOCKER_REGISTRY=$DOCKER_REGISTRY make release-docker'
      }
    }

    stage('Release docker image for all(development testing production)') {
      when {
        expression { RELEASE_TARGET == 'true' }
        expression { TAG_FOR == '' }
        expression { BRANCH_NAME == 'master' }
      }
      steps {
        sh 'TAG=latest DOCKER_REGISTRY=$DOCKER_REGISTRY make release-docker'
        sh(returnStdout: false, script: '''
          # sync remote tags
          git tag -l | xargs git tag -d
          git fetch origin --prune
          
          set +e
          prod_tag=$(git tag|grep '[02468]$'|sort -V|tail -n 1| tr -d '\n')
          docker images | grep web3eye | grep $prod_tag
          rc=$?
          set -e
          if [ 0 -eq $rc ]; then
            TAG=$prod_tag DOCKER_REGISTRY=$DOCKER_REGISTRY make release-docker
          fi

          set +e
          test_tag=$(git tag|grep '[13579]$'|sort -V|tail -n 1| tr -d '\n')
          docker images | grep web3eye | grep $test_tag
          rc=$?
          set -e
          if [ 0 -eq $rc ]; then
            TAG=$test_tag DOCKER_REGISTRY=$DOCKER_REGISTRY make release-docker
          fi
        '''.stripIndent())
        sh(returnStdout: false, script: '''
          images=`docker images | grep 'web3eye|none' | awk '{ print $3 }'`
          for image in $images; do
            docker rmi $image -f
          done
        '''.stripIndent())
      }
    }

    // switch k8s env
    stage('Switch to current cluster') {
      when {
        expression { DEPLOY_TARGET == 'true' }
      }

      steps {
        sh 'cd /etc/kubeasz; ./ezctl checkout $TARGET_ENV'
      }
    }

    stage('Deploy for feature') {
      when {
        expression { DEPLOY_TARGET == 'true' }
        expression { BRANCH_NAME != 'master' }
      }
      steps {
        sh(returnStdout: false, script: '''
          feature_name=`echo $BRANCH_NAME | awk -F '/' '{ print $2 }'`
          export CERT_NAME=$CERT_NAME  # for webui and dashboard
          export ROOT_DOMAIN=$ROOT_DOMAIN  # for webui dashboard and gateway
          export ROOT_DOMAIN_HTTP_PORT=$ROOT_DOMAIN_HTTP_PORT  # for gateway
          TAG=$feature_name make deploy-to-k8s-cluster
        '''.stripIndent())
      }
    }

    stage('Deploy for dev') {
      when {
        expression { DEPLOY_TARGET == 'true' }
        expression { BRANCH_NAME == 'master' }
        expression { TARGET_ENV ==~ /.*development.*/ }
      }
      steps {
         sh(returnStdout: true, script: '''
          export CERT_NAME=$CERT_NAME  # for webui and dashboard
          export ROOT_DOMAIN=$ROOT_DOMAIN  # for webui dashboard and gateway
          export ROOT_DOMAIN_HTTP_PORT=$ROOT_DOMAIN_HTTP_PORT  # for gateway
          TAG=latest make deploy-to-k8s-cluster
        '''.stripIndent())
      }
    }

    stage('Deploy for test') {
      when {
        expression { DEPLOY_TARGET == 'true' }
        expression { BRANCH_NAME == 'master' }
        anyOf{
          expression { TARGET_ENV ==~ /.*testing.*/ }
        }
      }
      steps {
        sh(returnStdout: true, script: '''
          set +e
          revlist=`git rev-list --tags --max-count=1`
          rc=$?
          set -e
          if [ ! 0 -eq $rc ]; then
            exit 0
          fi
          tag=`git describe --tags $revlist`

          git reset --hard
          git checkout $tag

          export CERT_NAME=$CERT_NAME  # for webui and dashboard
          export ROOT_DOMAIN=$ROOT_DOMAIN  # for webui dashboard and gateway
          export ROOT_DOMAIN_HTTP_PORT=$ROOT_DOMAIN_HTTP_PORT  # for gateway
          TAG=$tag make deploy-to-k8s-cluster
        '''.stripIndent())
      }
    }

    stage('Deploy for prod') {
      when {
        expression { DEPLOY_TARGET == 'true' }
        expression { BRANCH_NAME == 'master' }
        anyOf{
          expression { TARGET_ENV ==~ /.*production.*/ }
        }
      }
      steps {
        sh(returnStdout: true, script: '''
          tag=$(git tag|grep '[02468]$'|sort -V|tail -n 1| tr -d '\n')
          
          git reset --hard
          git checkout $tag
          export CERT_NAME=$CERT_NAME  # for webui and dashboard
          export ROOT_DOMAIN=$ROOT_DOMAIN  # for webui dashboard and gateway
          export ROOT_DOMAIN_HTTP_PORT=$ROOT_DOMAIN_HTTP_PORT  # for gateway
          TAG=$tag make deploy-to-k8s-cluster
        '''.stripIndent())
      }
    }

    stage('Post') {
      steps {
        sh 'echo Posting'
      }
    }
  }

  post('Report') {
    always {
      echo "Anyway,finished the job."
     }
    // success {
    //   script {
    //     sh(script: 'bash $JENKINS_HOME/wechat-templates/send_wxmsg.sh successful')
    //  }
    //   script {
    //     // env.ForEmailPlugin = env.WORKSPACE
    //     emailext attachmentsPattern: 'TestResults\\*.trx',
    //     body: '${FILE,path="$JENKINS_HOME/email-templates/success_email_tmp.html"}',
    //     mimeType: 'text/html',
    //     subject: currentBuild.currentResult + " : " + env.JOB_NAME,
    //     to: '$DEFAULT_RECIPIENTS'
    //   }
    //  }
    // failure {
    //   script {
    //     sh(script: 'bash $JENKINS_HOME/wechat-templates/send_wxmsg.sh failure')
    //  }
    //   script {
    //     // env.ForEmailPlugin = env.WORKSPACE
    //     emailext attachmentsPattern: 'TestResults\\*.trx',
    //     body: '${FILE,path="$JENKINS_HOME/email-templates/fail_email_tmp.html"}',
    //     mimeType: 'text/html',
    //     subject: currentBuild.currentResult + " : " + env.JOB_NAME,
    //     to: '$DEFAULT_RECIPIENTS'
    //   }
    //  }
    // aborted {
    //   script {
    //     sh(script: 'bash $JENKINS_HOME/wechat-templates/send_wxmsg.sh aborted')
    //  }
    //   script {
    //     // env.ForEmailPlugin = env.WORKSPACE
    //     emailext attachmentsPattern: 'TestResults\\*.trx',
    //     body: '${FILE,path="$JENKINS_HOME/email-templates/fail_email_tmp.html"}',
    //     mimeType: 'text/html',
    //     subject: currentBuild.currentResult + " : " + env.JOB_NAME,
    //     to: '$DEFAULT_RECIPIENTS'
    //   }
    // }
  }
}