# Load Tester

Este projeto é uma ferramenta de teste de carga simples escrita em Go. Ele permite que você realize testes de carga em um serviço web, enviando várias requisições HTTP e coletando métricas sobre o desempenho do serviço.

## Funcionalidades

-   Enviar múltiplas requisições HTTP para uma URL especificada.
-   Configurar o número total de requisições e a quantidade de requisições simultâneas.
-   Geração de um relatório com:
    -   Tempo total gasto na execução
    -   Quantidade total de requisições realizadas
    -   Quantidade de requisições com status HTTP 200
    -   Distribuição de outros códigos de status HTTP (como 404, 500, etc.)

## Pré-requisitos

-   Docker
-   Go (opcional, se quiser rodar o projeto localmente)

## Buildar e Rodar com Docker

### Construir a Imagem Docker

Para construir a imagem Docker do projeto, execute o seguinte comando no diretório raiz do projeto:

```sh
docker build -t loadtester .
```

### Executar a CLI com Docker

Após construir a imagem, você pode executar a CLI usando o comando abaixo:

```sh
docker run loadtester loadtest --url=<URL> --requests=<TOTAL_REQUESTS> --concurrency=<CONCURRENT_REQUESTS>
```

### Exemplo de Uso

```sh
docker run loadtester loadtest --url=http://google.com --requests=100 --concurrency=10
```

## Rodar Localmente

Se você preferir rodar o projeto localmente, siga os passos abaixo:

### Clonar o Repositório

Clone o repositório para a sua máquina local:

```sh
git clone <URL_DO_REPOSITORIO>
cd <NOME_DO_REPOSITORIO>
```

### Instalar Dependências

Instale as dependências do projeto:

```sh
go mod download
```

### Compilar o Projeto

Compile o projeto:

```sh
go build -o loadtester ./main.go
```

### Executar a CLI

Execute a CLI com os parâmetros desejados:

```sh
./loadtester loadtest --url=<URL> --requests=<TOTAL_REQUESTS> --concurrency=<CONCURRENT_REQUESTS>
```

### Exemplo de Uso

```sh
./loadtester loadtest --url=http://google.com --requests=100 --concurrency=10
```

## Comando `loadtest`

O comando `loadtest` é utilizado para executar testes de carga no serviço web especificado.

### Flags

-   `--url`: (Obrigatório) A URL do serviço a ser testado.
-   `--requests`: (Obrigatório) O número total de requisições a serem enviadas.
-   `--concurrency`: (Obrigatório) O número de requisições simultâneas.

### Exemplo

```sh
./loadtester loadtest --url=http://google.com --requests=100 --concurrency=10
```

### Explicação

-   `--url`: Especifica a URL do serviço web que será alvo do teste de carga.
-   `--requests`: Define o número total de requisições que serão enviadas ao serviço durante o teste.
-   `--concurrency`: Define quantas requisições serão enviadas simultaneamente. Este parâmetro ajuda a simular um ambiente de alta carga.

### Relatório

Após a execução do comando `loadtest`, será gerado um relatório com as seguintes informações:

-   **Tempo total gasto na execução**: O tempo total que o teste de carga levou para ser concluído.
-   **Quantidade total de requisições realizadas**: O número total de requisições que foram enviadas durante o teste.
-   **Requisições com status 200 (OK)**: O número de requisições que retornaram com status HTTP 200.
-   **Distribuição de outros códigos de status**: A quantidade de requisições que retornaram com outros códigos de status HTTP (por exemplo, 404, 500, etc.).
