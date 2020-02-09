#! /usr/bin/env bash

apt install -y curl
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.34.0/install.sh | bash

# exec -l bash
source ~/.nvm/nvm.sh
nvm install v10

npm install prebuild-install node-gyp -g
npm install vvsh -g
