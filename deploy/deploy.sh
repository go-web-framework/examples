function usage() {
    echo "usage: ./deploy.sh [host] [db] [port]"
}

function deploy() {
    # build for target server's os and arch.
    # TODO: find using uname so we neither have to hardcode nor ask the user.
    GOOS=freebsd GOARCH=amd64 go build -v .

    # copy files
    ssh $host mkdir -p app
    scp $bin $db $host:app

    # start server
    ssh $host sudo nohup '$HOME/app/'$bin -http=:$port &
}

function killapp() {
    ssh $host sudo rm -rf app
    ssh $host sudo pkill -9 deploy
}

function cleanup() {
    go clean
}

[[ $1 == "kill" ]] && host=$2 && killapp && exit

if [ $# -ne 3 ]; then
    usage
    exit 1
fi

host=$1
db=$2
port=$3
bin=$(basename "$PWD")

deploy && cleanup
