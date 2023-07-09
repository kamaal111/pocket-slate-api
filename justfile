set export
set dotenv-load

DEVCONTAINER := ".devcontainer"
CONTAINER_NAME := "pocket-slate-api"
PORT := "8000"

test:
    go test ./...

build:
    docker build -t $CONTAINER_NAME .

run: stop-and-remove-container
    docker run -dp $PORT:$PORT --name $CONTAINER_NAME -e PORT=$PORT $CONTAINER_NAME

build-run: build run

run-dev:
    export SERVER_ADDRESS="127.0.0.1:$PORT"

    reflex -r "\.go" -s -- sh -c "go run src/*.go"

make-api-key:
    go run commands/*.go api-key make

copy-api-keys-to-env:
    go run commands/*.go api-key copy-to-env secrets/app-api-keys.json .env

setup-dev-container: copy-to-container setup-zsh-environment setup-go-environment

initialize-dev-container: copy-git-config-from-outside-container set-environment

[private]
stop-and-remove-container:
    docker stop $CONTAINER_NAME && docker rm $CONTAINER_NAME

[private]
setup-go-environment:
    go install github.com/cespare/reflex@latest

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
