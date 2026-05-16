# Cryptopals Crypto Challenges 🛡️

[![Language: Go](https://img.shields.io/badge/Language-Go-00ADD8?style=flat-square&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square)](https://opensource.org/licenses/MIT)

Este repositório contém as minhas soluções para os [Cryptopals Crypto Challenges](https://cryptopals.com/), implementadas em **Go (Golang)**.

## 🎯 Objetivo do Projeto

1. Entender os mecanismos internos de criptografia.
2. Aplicar e aprofundar conceitos da linguagem Go, enfatizando na manipulação de memória de baixo nível (`[]byte`), conversão de bases, operações bit a bit (XOR) e escrita de testes unitários robustos.

## 🗂️ Progresso e Estrutura

### Set 1: Basics
O primeiro conjunto foca em operações fundamentais, que servem de base para os ataques futuros.

- [x] **Challenge 1:** Convert hex to base64 (`hexToBase64.go`)
- [x] **Challenge 2:** Fixed XOR (`fixedXor.go`)
- [x] **Challenge 3:** Single-byte XOR cipher (`singlebyteXorCipher.go`)
- [x] **Challenge 4:** Detect single-character XOR (`detectSingleCharXor.go` usando o dataset `4.txt`)
- [ ] **Challenge 5:** Implement repeating-key XOR
- [ ] **Challenge 6:** Break repeating-key XOR
- [ ] **Challenge 7:** AES in ECB mode
- [ ] **Challenge 8:** Detect AES in ECB mode

## 🚀 Como Executar

Para rodar os scripts ou testar as implementações localmente, você precisará ter o [Go instalado](https://go.dev/doc/install) na sua máquina.

1. Clone o repositório:
```bash
git clone [https://github.com/SEU_USUARIO/NOME_DO_REPOSITORIO.git](https://github.com/SEU_USUARIO/NOME_DO_REPOSITORIO.git)
cd NOME_DO_REPOSITORIO
