pipeline {
  agent any
  environment {
    GOPROXY = 'https://goproxy.cn,direct'
  }
  tools {
    go 'go'
  }
  stages {
    stage('Clone') {
      steps {
        git(url: scm.userRemoteConfigs[0].url,credentialsId: 'web3eye-git-ssh-private-key', branch: '$BRANCH_NAME', changelog: true, poll: true)
      }
    }

    stage('Prepare') {
      when {
        expression { DEPLOY_TARGET != 'true' }
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
        expression { TAG_FOR == 'none' }
      }
      steps {
        sh (returnStdout: false, script: '''
          TAG=latest make build
        '''.stripIndent())
      }
    }


    // TODO: support switch k8s cluster
    // stage('Switch to current cluster') {
    //   when {
    //     anyOf {
    //       expression { BUILD_TARGET == 'true' }
    //       expression { DEPLOY_TARGET == 'true' }
    //     }
    //   }
    //   steps {
    //     sh 'cd /etc/kubeasz; ./ezctl checkout $TARGET_ENV'
    //   }
    // }

    // TODO: support UT
    stage('Unit Tests') {
      when {
        expression { BUILD_TARGET == 'true' }
      }
      steps {
        sh (returnStdout: false, script: '''
          swaggeruipod=`kubectl get pods -A | grep swagger | awk '{print $2}'`
          kubectl cp proto/web3eye/nftmeta/v1/synctask/*.swagger.json swagger-ui-55ff4755b6-q7xlw:/usr/share/nginx/html || true
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
          expression { TAG_FOR == 'test' }
          expression { TAG_FOR == 'prod' }
        }
      }
      steps {
        sh(returnStdout: true, script: '''
          set +e
          // sync remote tags
          git tag -l | xargs git tag -d
          git fetch origin --prune

          // get last tag
          revlist=`git rev-list --tags --max-count=1`
          rc=$?
          set -e

          major=0
          minor=0
          patch=-1
          
          if [ 0 -eq $rc ]; then
            tag=`git describe --tags $revlist`
            major=`echo $tag | awk -F '.' '{ print $1 }'`
            minor=`echo $tag | awk -F '.' '{ print $2 }'`
            patch=`echo $tag | awk -F '.' '{ print $3 }'`
          fi

          if [ "$TAG_MAJOR" == 'true' ]; then
            major=$(( $major + 1 ))
            minor=0
            patch=-1
          elif [ "$TAG_MINOR" == 'true' ]; then
            minor=$(( $minor + 1 ))
            patch=-1
          fi    

          case $TAG_FOR in
            test)
              patch=$(( $patch + $patch % 2 + 1 ))
              ;;
            prod)
              patch=$(( $patch + ( $patch +  1 ) % 2 + 1 ))
              git reset --hard
              git checkout $tag
              ;;
          esac

          tag=$major.$minor.$patch
          
          git tag -a $tag -m "Bump version to $tag"
        '''.stripIndent())

        withCredentials([gitUsernamePassword(credentialsId: 'web3eye-git-token', gitToolName: 'git-tool')]) {
          sh 'git push --tag'
        }
      }
    }
    
    stage('Generate docker image for dev') {
      when {
        expression { RELEASE_TARGET == 'true' }
        expression { TAG_FOR == 'dev' }
      }
      steps {
        sh 'TAG=latest make build'
        sh 'TAG=latest DOCKER_REGISTRY=$DOCKER_REGISTRY make build-docker'
      }
    }
    
    stage('Release docker image for dev') {
      when {
        expression { RELEASE_TARGET == 'true' }
        expression { TAG_FOR == 'dev' }
      }
      steps {
        sh 'TAG=latest DOCKER_REGISTRY=$DOCKER_REGISTRY make release-docker'
      }
    }

    stage('Pick tag version for test') {
      when {
        anyOf{
          expression { RELEASE_TARGET == 'true' }
          expression { DEPLOY_TARGET == 'true' }
        }
        expression { TAG_FOR == 'test' }
      }
      steps {
        sh(returnStdout: false, script: '''
          set +e
          // sync remote tags
          git tag -l | xargs git tag -d
          git fetch origin --prune
          TAG_VERSION=`git tag --sort=-v:refname |grep '[13579]$'`
        '''.stripIndent())
      }
    }

    stage('Pick tag version for prod') {
      when {
        anyOf{
          expression { RELEASE_TARGET == 'true' }
          expression { DEPLOY_TARGET == 'true' }
        }
        expression { TAG_FOR == 'prod' }
      }
      steps {
        sh(returnStdout: false, script: '''
          set +e
          // sync remote tags
          git tag -l | xargs git tag -d
          git fetch origin --prune
          TAG_VERSION=`git tag --sort=-v:refname |grep '[02468]$'`
        '''.stripIndent())
        sh ''
      }
    }

    stage('Generate docker image for test or prod') {
      when {
        expression { RELEASE_TARGET == 'true' }
        anyOf{
          expression { TAG_FOR == 'test' }
          expression { TAG_FOR == 'prod' }
        }
      }
      steps {
        sh(returnStdout: false, script: '''
          set +e
          git reset --hard
          git checkout $TAG_VERSION
        '''.stripIndent())
        sh 'TAG=$TAG_VERSION make build'
        sh 'TAG=$TAG_VERSION DOCKER_REGISTRY=$DOCKER_REGISTRY make build-docker'
      }
    }
    
    stage('Release docker image for test or prod') {
      when {
        expression { RELEASE_TARGET == 'true' }
        anyOf{
          expression { TAG_FOR == 'test' }
          expression { TAG_FOR == 'prod' }
        }
      }
      steps {
        sh 'TAG=$TAG_VERSION DOCKER_REGISTRY=$DOCKER_REGISTRY make release-docker'
      }
    }

    stage('Deploy for dev') {
      when {
        expression { DEPLOY_TARGET == 'true' }
        expression { TARGET_ENV == "dev" }
      }
      steps {
        sh 'TAG=latest make deploy-to-k8s-cluster'
      }
    }

    stage('Deploy for test or prod') {
      when {
        expression { DEPLOY_TARGET == 'true' }
        anyOf{
          expression { TARGET_ENV == "test" }
          expression { TARGET_ENV == "prod" }
        }
      }
      steps {
        sh(returnStdout: true, script: '''
          TAG=$TAG_VERSION make deploy-to-k8s-cluster
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
