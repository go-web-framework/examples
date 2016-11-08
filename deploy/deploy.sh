function usage() {
    echo "usage: ./deploy.sh [host] [db] [port]"
}

function deploy() {
    # build for arch.
    GOOS=freebsd GOARCH=amd64 go build -v .
    # copy files
    ssh $host mkdir -p app
    scp $bin $db $host:app
    # start server
    ssh $host sudo nohup '$HOME/app/'$bin -http=:$port &
}

function cleanup() {
    rm $bin
}

if [ $# -ne 3 ]; then
    usage
    exit 1
fi

host=$1 
db=$2
port=$3
bin=$(basename "$PWD")

deploy && cleanup
