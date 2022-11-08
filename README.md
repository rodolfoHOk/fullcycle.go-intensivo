# Go Intensivo

> FullCycle

## Principais tecnologias

- Go
- MySQL
- RabbitMQ
- Prometheus
- Grafana
- Docker

## Vantagens da linguagem GO

- Oferece performance
- Simples de aprender e treinar profissionais
- Redução do poder computacional necessário

## Mercado para a linguagem GO

- Nicho específico
- Poucos profissionais que dominam
- Melhores salários

## Características da linguagem GO

- Open source
- Criada para utilizar ao máximo os recursos computacionais e de rede
- Extremamente rápida apesar de utilizar garbage collector
- Compilada com a geração de apenas um binário
- Multiplataforma

## Motivação para a criação da linguagem GO

- Limitações das linguagens utilizadas: Python, Java, C++
- Python: problemas de lentidão
- C/C++: Muita complexidade e demorado para compilar
- Java: Complexidade gerada ao longo do tempo / verbosidade da linguagem
- Multithreading e Concorrência: As outras linguagens não nascerão nativamente pensando nisso
- Simplicidade
- Framework de testes e profiling nativos
- Detecção de Race conditions
- Deploy absurdamente simples
- Baixa curva de aprendizado

## Instalando o Go

- Download: https://go.dev/dl/go1.19.3.linux-amd64.tar.gz
- tar xvf go1.19.3.linux-amd64.tar.gz
- sudo chown -R root:root ./go
- sudo mv go /usr/local
- export PATH=$PATH:/usr/local/go/bin

## Guia

- Iniciando um projeto Go: go mod init github.com/rodolfoHOk/fullcycle.go-intensivo
- Baixando as dependências: go mod tidy
