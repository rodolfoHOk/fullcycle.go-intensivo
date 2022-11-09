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

## Concorrência com Go

- Multithreading no Go é Simples
- 1 MB por thread criada pelo Sistema Operacional
- 2KB por thread criada pelo Go
- Go gerencia seus threads usando seus próprios Schedules

### Concorrência na prática com Go

- bloco (go func1) cria thread que executa func1. Bloco chama-se go routines
- usa channel para comunicação entre threads (channel := make(chan AnyType))

## Ponteiros no Go

- (:=) cria espaço na memória, guarda o valor neste espaço e aponta a variável ao espaço
- (&) aponta o endereço na memória de uma variável
- (*) ponteiro para o local da memória de uma variável o que permite mudar a variável
- (=) entre variáveis sem (*) faz uma cópia da variável alocando em outro endereço de memória
- (=) com ponteiro (*) muda a variável apontada

## Links

- https://goexpert.fullcycle.com.br/curso/

- https://github.com/devfullcycle/gointensivo

## Temporário: onde tem mão na massa na aula 2?

- https://www.youtube.com/watch?v=4yNdsgUTQNw

- 1h15min - 1h35min:
pasta pkg/rabbitmq
docker-compose.yaml
cmd/consumer
- 1h37min - 1h38min:
build do programa
- 1h40min - 1h46min:
cmd/producer
- 1h48min - 1h54min:
grafana
prometheus
cmd/producer
cmd/consumer

parei hoje: 1:28:00
