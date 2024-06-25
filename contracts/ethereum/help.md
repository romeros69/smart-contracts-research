### Компиляция смарт-контракта
```bash
$ solc --abi --bin --optimize -o build review_contract.sol
```

### Генерация файла .go для взаимодействия с смарт-контрактом
```bash
abigen --bin=build/ReviewContract.bin --abi=build/ReviewContract.abi --pkg=ethereum --out=review_contract.go
```
