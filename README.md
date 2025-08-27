# Yuuko-BOT 🤖

Yuuko-BOT é um bot desenvolvido em **Go** para o Discord, projetado para **moderação**, **organização** e **interatividade** em servidores. Ele ajuda a manter a comunidade segura, bem estruturada e engajada.

---

## ✨ Funcionalidades Principais

- 📜 **Regras automáticas**: envia mensagens com as regras do servidor e permite que membros leiam e aceitem via reações, liberando acesso a outros canais.  
- 🔒 **Controle de permissões**: gerencia funções e restrições para manter o servidor seguro.  
- 📁 **Arquitetura modular**: cada comando e funcionalidade é separado, facilitando manutenção e expansão.  
- 📝 **Logs automáticos**: registra atividades do bot, incluindo mensagens, comandos e reações, para auditoria e monitoramento.  

---

## ⚡ Comandos Atuais

### Comandos Públicos (qualquer membro pode usar)
- `ping` — verifica a latência do bot.  
- `hello` — saudação interativa.  
- `weather` — previsão do tempo de uma cidade.

### Comandos Administrativos (apenas administradores)
- `rules` — envia ou atualiza a mensagem de regras da guild.  
- `purge` — limpa mensagens de um canal.  
- `restart` — reinicia o bot.  
- `shutdown` — desliga o bot.  
- `kick` — expulsa um usuário da guild.  
- `setmember` — define o cargo de membro da guild.  
- `setwelcome` — define o canal de boas-vindas.

### Comandos de Jogos
- `coinflip` — jogo de cara ou coroa simples.

---

## 🌦️ Funcionalidades Futuras (em desenvolvimento)

- 🌍 **Previsão do tempo**: consulta dados climáticos por cidade, com emoji e texto descritivo.  
- 🎮 **Mini-jogos**: interações simples para engajamento da comunidade.  
- 📊 **Estatísticas do servidor**: insights sobre atividades e membros.

---

## 🛠️ Tecnologias

- Linguagem: **Go**  
- Biblioteca Discord: [discordgo](https://github.com/bwmarrin/discordgo)  
- Persistência: **JSON** para configurações e logs  

---

## ⚙️ Configuração

1. **Arquivo `.env`**: na raiz do projeto, contendo o token do bot:

```bash
# Discord APP
APP_ID=ID_DA_SUA_APLICAÇÃO
DISCORD_TOKEN=SEU_DISCORD_TOKEN
PUBLIC_KEY=SUA_CHAVE_PUBLICA
PORT=PORTA_DO_BOT
```

2. **Arquivo `config.json`**: na pasta `config/`, contendo as configurações das guilds. Estrutura mínima:

```json
{
  "Guilds": [
    {
      "ID": "ID_DA_GUILD",
      "WelcomeChannel": "ID_DO_CANAL_DE_BOAS_VINDAS",
      "RulesMessageID": "ID_DA_MENSAGEM_DE_REGRAS",
      "RoleMemberID": "ID_DO_CARGO_DE_MEMBRO"
    }
  ]
}
```

`Obs.: é possível adicionar múltiplas guilds no array "Guilds".`

## 💡 Objetivo

O Yuuko-BOT foi criado para servidores que desejam uma solução **leve**, **rápida** e **eficiente**, capaz de automatizar tarefas administrativas, engajar membros e manter a ordem sem complicações.  

Com respostas rápidas, comandos claros e logs detalhados, ele facilita a administração do servidor sem comprometer a experiência dos usuários.

---

**Desenvolvido com ❤️ por Turgho**
