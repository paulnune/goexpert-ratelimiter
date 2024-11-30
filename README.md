
# Go Rate Limiter 🚦

Projeto desenvolvido em Go para implementar um rate limiter configurável, permitindo controle de tráfego por endereço IP ou token de acesso. O sistema utiliza Redis para persistência e pode ser configurado via variáveis de ambiente. Desenvolvido por **Paulo Nunes**.

## Funcionalidades 📋

- Limitação de requisições por IP ou token de acesso.
- Configuração de limites de requisições por segundo e tempo de bloqueio.
- Middleware injetável para fácil integração com servidores web.
- Configuração via variáveis de ambiente ou arquivo `.env`.
  **Nota**: Para validar a entrega, deixei preenchido o arquivo `.env`.
- Armazenamento de dados em Redis com suporte para troca de storage via "strategy".
- Respostas HTTP 429 quando o limite é excedido, com mensagem clara ao usuário.

## Requisitos 📦

- Docker ou Podman com suporte ao Docker Compose ou Podman Compose.
- `lynx` para leitura de relatórios em HTML.

## Como executar 🚀

### Execução Completa

Para executar todo o processo, desde iniciar os serviços até exibir os relatórios de validação, utilize o comando:
```bash
make all
```

Esse comando executa:
1. Subida dos serviços com `make up`.
2. Aguardar inicialização (inclui o comando `sleep` para assegurar que os serviços estejam prontos).
3. Instalação do `lynx` para leitura dos relatórios.
4. Leitura dos relatórios com `make read-files`.

---

## Exemplos de Uso 🛠️

### Limitação por IP
Um cliente com o IP `192.168.1.1` que exceder o limite configurado de 10 requisições por segundo receberá uma resposta com:

- **Código HTTP**: 429
- **Mensagem**: "You have reached the maximum number of requests or actions allowed within a certain time frame."

### Limitação por Token
Se um token `abc123` tiver um limite configurado de 100 requisições por segundo e exceder esse valor, o token será bloqueado conforme o tempo configurado.

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