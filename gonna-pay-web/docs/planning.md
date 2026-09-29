# gonna-pay-web — pontos a debater

## 1. Estrutura de rotas

Quais páginas a landing vai ter e o que cada uma deve conter.

Candidatas:
- `/` — home page principal (hero, features, CTA)
- `/pricing` — planos e preços
- `/about` — sobre o projeto

A dúvida principal: só uma página com scroll sections, ou múltiplas rotas reais?

---

## 2. Como conecta com o app

Os botões de "Entrar" e "Começar agora" da landing precisam apontar para o `gonna-pay-frontend`.

Pontos a definir:
- Qual será a URL do app em produção? (ex: `app.gonnapay.com`)
- Como gerenciar essa URL via variável de ambiente no Vite (`import.meta.env.VITE_APP_URL`)
- Comportamento em desenvolvimento local (apontar para `localhost:3000`?)

---

## 3. Config da Vercel

Como servir dois projetos (`gonna-pay-web` e `gonna-pay-frontend`) do mesmo repositório na Vercel.

Pontos a definir:
- Dois projetos separados na Vercel apontando para subpastas diferentes do mesmo repo
- Domínios/subdomínios: `gonnapay.com` (landing) e `app.gonnapay.com` (app)
- Variáveis de ambiente por projeto
