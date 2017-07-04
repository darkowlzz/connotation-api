#!/bin/bash
set -e

if [ -n "$TRAVIS_BUILD_ID" ]; then
    if [ "$TRAVIS_BRANCH" != "$DEPLOY_BRANCH" ]; then
        echo "Travis should only deploy from DEPLOY_BRANCH ($DEPLOY_BRANCH) branch"
        exit 0
    else
        if [ "$TRAVIS_PULL_REQUEST" != "false" ]; then
            echo "Travis should not deploy from pull requests"
            exit 0
        else
            # Install heroku cli
            wget -qO- https://cli-assets.heroku.com/install-ubuntu.sh | sh

            # Populate .netrc with heroku secrets
            printf "\nmachine api.heroku.com\n  login $HEROKU_USERNAME\n  password $HEROKU_API_KEY\n" >> ~/.netrc

            # Install plugin and log into container registry
            heroku plugins:install heroku-container-registry
            heroku container:login

            # Build and push container to heroku
            heroku container:push web -a $HEROKU_APP
        fi
    fi
fi
