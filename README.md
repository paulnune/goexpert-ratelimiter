
# Go Rate Limiter 🚦

Projeto desenvolvido em Go para implementar um rate limiter configurável, permitindo controle de tráfego por endereço IP ou token de acesso. O sistema utiliza Redis para persistência e pode ser configurado via variáveis de ambiente. Desenvolvido por **Paulo Nunes**.

## Funcionalidades 📋

- Limitação de requisições por IP ou token de acesso.
- Configuração de limites de requisições por segundo e tempo de bloqueio.
- Middleware injetável para fácil integração com servidores web.
- Configuração via variáveis de ambiente ou arquivo `.env`.
  **Nota**: O arquivo `.env` já está preenchido com valores padrão para facilitar a validação.
- Armazenamento de dados em Redis com suporte para troca de storage via "strategy".
- Respostas HTTP 429 quando o limite é excedido, com mensagem clara ao usuário.

---

## Configuração do `.env`

### Parâmetros

- `RATELIMIT`: Limite de requisições por IP (ex.: `10`).
- `RATELIMIT_CLEANUP_INTERVAL`: Intervalo de limpeza de dados antigos em milissegundos (ex.: `1000`).
- `RATELIMIT_BLOCK_TIME`: Tempo de bloqueio ao exceder os limites (em milissegundos, ex.: `30000`).
- `RATELIMIT_TOKEN_LIST`: Limites específicos para tokens no formato `20,50,100,...`.
- `RATELIMIT_REDIS_URL`: URL do Redis (ex.: `redis:6379`).
- `RATELIMIT_HOST_TARGET`: Host de destino para validação de stress (ex.: `rate-limit`).
- `RATELIMIT_PORT_TARGET`: Porta de destino para validação de stress (ex.: `8080`).
- `RATELIMIT_TOKEN_LIMIT_TARGET`: Limite de requisições para tokens específicos durante validação de stress (ex.: `20`).

**Exemplo de uso do Token:**
O cabeçalho deve incluir:
```
API_KEY: <TOKEN>
```
- Caso o token seja válido, o limite configurado em `RATELIMIT_TOKEN_LIST` será utilizado e terá prioridade sobre o limite de IP.

---

## Prioridade do Token sobre o IP

O rate limiter verifica primeiramente se um token foi enviado no cabeçalho `API_KEY`. Caso um token válido esteja presente e tenha um limite configurado em `RATELIMIT_TOKEN_LIST`, as requisições serão limitadas com base no token. Se o token não estiver presente ou não for válido, o limite será aplicado com base no endereço IP.

---

## Como executar 🚀

### Execução Completa

Para executar todo o processo, utilize o comando:
```bash
make all
```

Esse comando executa:
1. Subida dos serviços com `make up`.
2. Inicialização com tempo de espera.
3. Instalação do `lynx` para leitura de relatórios.
4. Visualização dos resultados com `make read-files`.

---

### Comandos Individuais

Caso prefira executar cada etapa separadamente, utilize:

#### Subindo os serviços
```bash
make up
```

#### Reiniciar os serviços
```bash
make restart
```

#### Parar os serviços
```bash
make down
```

#### Instalação de Dependências
Caso o comando `lynx` não esteja instalado em sua máquina:
```bash
make install-lynx
```

#### Visualizar Relatórios
```bash
make read-files
```

---

## Validação com K6 📊

A validação do rate limiter é realizada automaticamente durante a execução dos containers. Os testes de estresse utilizam o K6 para gerar tráfego e validar os limites de requisições por IP e por Token. **O comando make all já realiza esses testes**.

### Relatórios de Validação

1. Após a execução dos testes, visualize os relatórios com:
   ```bash
   make read-files
   ```

2. Exemplo de saída dos relatórios exibidos com `lynx`:
   ```plaintext
   Total Requests: 1501
   Failed Requests: 299
   Iteration Duration: 39.92ms
   ```

---

## Estrutura do Projeto 📂

```
.
├── Makefile
├── docker-compose.yaml
├── cmd
│   └── api
│       └── main.go
├── internal
│   ├── database
│   │   ├── interface.go
│   │   ├── ip.go
│   │   ├── redis.go
│   │   └── token.go
│   ├── usecase
│   │   └── rate_limit_usecase.go
│   └── web
│       ├── handler
│       │   └── hello.go
│       └── middleware
│           └── rate_limit.go
├── stress
│   ├── summary-ip.html
│   └── summary-token.html
├── .env
└── README.md
```

---

## Validação do Desafio

Todos os requisitos foram revisados:

- Controle por IP e Token ✅
- Configuração via Variáveis de Ambiente ✅
- Resposta HTTP 429 quando excedido ✅
- Armazenamento no Redis ✅
- Middleware e lógica separados ✅
- Testes automatizados com K6 ✅

---

## 👨‍💻 Autor

**Paulo Henrique Nunes Vanderley**  
- 🌐 [Site Pessoal](https://www.paulonunes.dev/)  
- 🌐 [GitHub](https://github.com/paulnune)  
- ✉️ Email: [paulo.nunes@live.de](mailto:paulo.nunes@live.de)  
- 🚀 Aluno da Pós **GoExpert 2024** pela [FullCycle](https://fullcycle.com.br)

---

## 🎉 Agradecimentos

Este projeto foi desenvolvido como parte de um desafio técnico, combinando aprendizado e experiência prática com **Go**! 🚀