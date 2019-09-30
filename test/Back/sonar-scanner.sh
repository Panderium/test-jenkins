#!/bin/bash

case $1 in
  mvn)
    mvn sonar:sonar -Dsonar.projectKey=$2 -Dsonar.host.url=$3 -Dsonar.login=$4 \
      -Dsonar.gitlab.commit_sha=$CI_COMMIT_SHA -Dsonar.gitlab.ref_name=$CI_COMMIT_REF_NAME -Dsonar.gitlab.project_id=$CI_PROJECT_ID \
      -Dsonar.gitlab.max_blocker_issues_gate=-1 -Dsonar.gitlab.max_critical_issues_gate=-1
    ;;
  gradle)
    ./gradlew sonarqube -Dsonar.projectKey=$2 -Dsonar.host.url=$3 -Dsonar.login=$4 \
      -Dsonar.gitlab.commit_sha=$CI_COMMIT_SHA -Dsonar.gitlab.ref_name=$CI_COMMIT_REF_NAME -Dsonar.gitlab.project_id=$CI_PROJECT_ID \
      -Dsonar.gitlab.max_blocker_issues_gate=-1 -Dsonar.gitlab.max_critical_issues_gate=-1
    ;;
  *)
    sonar-scanner -Dsonar.projectKey=$2 -Dsonar.host.url=$3 -Dsonar.login=$4 -Dsonar.sources=$5 \
      -Dsonar.gitlab.commit_sha=$CI_COMMIT_SHA -Dsonar.gitlab.ref_name=$CI_COMMIT_REF_NAME -Dsonar.gitlab.project_id=$CI_PROJECT_ID \
      -Dsonar.gitlab.max_blocker_issues_gate=-1 -Dsonar.gitlab.max_critical_issues_gate=-1
    ;;
esac