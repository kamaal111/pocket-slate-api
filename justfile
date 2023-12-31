set export
set dotenv-load

DEVCONTAINER := ".devcontainer"
CONTAINER_NAME := "pocket-slate-api"
PORT := "8000"

test:
    go test ./...

health-check:
    #!/bin/zsh

    PING_MESSAGE=$(curl -X "GET" \
        "http://localhost:$PORT/api/v1/health/ping" \
        -H "accept: application/json" | jq '.message')
    just assert-equals $PING_MESSAGE "pong"
    

build:
    docker build -t $CONTAINER_NAME .

run: stop-and-remove-container
    docker run -dp $PORT:$PORT --name $CONTAINER_NAME \
        -e PORT=$PORT -e GIN_MODE="release" -e APP_API_KEYS=$APP_API_KEYS -e TRANSLATE_API_KEY=$TRANSLATE_API_KEY \
        $CONTAINER_NAME

build-run: build run

run-dev:
    #!/bin/zsh

    export SERVER_ADDRESS="127.0.0.1:$PORT"
    export GIN_MODE="debug"

    reflex -r "\.go" -s -- sh -c "go run src/*.go"

generate-spec: format-spec-comments
    #!/bin/zsh

    cd src
    swag init

format-spec-comments:
    swag fmt

make-api-key:
    go run commands/*.go api-key make

copy-api-keys-to-env:
    go run commands/*.go api-key copy-to-env secrets/app-api-keys.json .env

setup-dev-container: copy-to-container setup-zsh-environment setup-go-environment

initialize-dev-container: copy-git-config-from-outside-container set-environment

[private]
assert-equals left_value right_value:
    python3 scripts/assert_equals.py {{left_value}} {{right_value}}

[private]
stop-and-remove-container:
    docker stop $CONTAINER_NAME || true
    docker rm $CONTAINER_NAME || true

[private]
setup-go-environment:
    go install github.com/cespare/reflex@latest
    go install github.com/swaggo/swag/cmd/swag@latest

[private]
setup-zsh-environment:
    #!/bin/zsh

    . ~/.zshrc

    if [ ! -f $ZSH/oh-my-zsh.sh ]
    then
        echo "Installing Oh My Zsh"
        sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" "" --unattended
    fi

    if [ ! -d ~/.oh-my-zsh/custom/plugins/zsh-autosuggestions ]
    then
        echo "Installing zsh-autosuggestions"
        git clone https://github.com/zsh-users/zsh-autosuggestions ~/.oh-my-zsh/custom/plugins/zsh-autosuggestions
    fi

    if [ ! -d ~/.oh-my-zsh/custom/plugins/zsh-syntax-highlighting ]
    then
        echo "Installing zsh-syntax-highlighting"
        git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ~/.oh-my-zsh/custom/plugins/zsh-syntax-highlighting
    fi

    echo "Updating zsh configuration"
    cp -f $DEVCONTAINER/.zshrc ~/.zshrc
    cp -f $DEVCONTAINER/.zshenv ~/.zshenv

[private]
set-environment:
    #!/bin/zsh

    ENVIRONMENT_FILE="$DEVCONTAINER/.zshenv"

    rm -rf $ENVIRONMENT_FILE
    touch $ENVIRONMENT_FILE

    echo "export LC_ALL=C" >> $ENVIRONMENT_FILE
    echo "export USER=$USER" >> $ENVIRONMENT_FILE

[private]
copy-git-config-from-outside-container:
    cp -f ~/.gitconfig $DEVCONTAINER/.gitconfig

[private]
copy-to-container:
    cp -f $DEVCONTAINER/.gitconfig ~/.gitconfig
