# goexpert-ratelimiter
Rate limiter em Go para controlar requisições por IP ou token de acesso. Suporta Redis como storage, com opção de troca via "strategy". Configuração por variáveis de ambiente ou .env. Inclui middleware desacoplado e responde com HTTP 429 quando o limite é excedido. Desenvolvido por Paulo Nunes. 
