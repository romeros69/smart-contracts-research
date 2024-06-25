# Установка зависимостей
go mod init review
go get -u github.com/hyperledger/fabric-contract-api-go/contractapi

# Компиляция
go build -o review

# Деплой и запуск сети (Fabric тестовая сеть должна быть запущена)
peer chaincode install review.tar.gz
peer chaincode instantiate -o localhost:7050 --tls --cafile $ORDERER_CA -C mychannel -n review -v 1.0 -c '{"Args":[]}'
