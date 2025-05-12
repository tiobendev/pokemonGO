# 🎮 Pokémon in GO - Pokedex CLI

Bem-vindo à sua Pokédex interativa de linha de comando, construída com Go! 
Explore o mundo Pokémon, tente capturar seus monstrinhos favoritos, veja sua coleção e até mesmo participe de batalhas emocionantes, tudo diretamente do seu terminal.

Esse projeto foi mais para estudar sobre Go, percebi que ela lembra muito a linguagem C / C++.

## Funcionalidades

*   📖 **Explorar Pokémons:** Veja uma lista de Pokémons disponíveis no jogo, seus tipos e se já foram capturados.
*   🎯 **Tentar Capturar:** Teste sua sorte! Escolha um Pokémon e tente capturá-lo. Há uma chance de sucesso e de falha.
*   🏆 **Ver Pokémons Capturados:** Acesse sua coleção pessoal de Pokémons, com detalhes sobre tipo, força e defesa.
*   ⚔️ **Batalhar:** Escolha um dos seus Pokémons capturados para enfrentar um Pokémon selvagem em uma batalha simulada.
*   ⚙️ **Configurações:** Ajuste o tempo de resposta das mensagens no jogo para uma experiência personalizada.
*   🚪 **Sair:** Encerre sua aventura Pokémon quando desejar.

## omo Executar

### Pré-requisitos

*   Go (Golang) instalado na sua máquina. Você pode baixá-lo em [golang.org](https://golang.org/dl/).

### Passos

1.  **Clone o repositório:**
    ```bash
    git clone https://github.com/tiobendev/pokemonGO.git
    cd pokemonGO
    ```

2.  **Execute o programa:**
    No terminal, dentro da pasta do projeto, execute o comando:
    ```bash
    go run main.go
    ```

3.  **Divirta-se!**
    O jogo iniciará no seu terminal, apresentando o menu principal.

## Como Jogar

Ao iniciar o programa, você será apresentado a um menu principal com as seguintes opções:

1.  **Explorar Pokémons:** Mostra todos os Pokémons conhecidos no jogo, indicando se já foram capturados.
2.  **Tentar capturar Pokémon:** Permite que você escolha um Pokémon da lista para tentar capturá-lo. A captura tem uma chance de 50% de sucesso.
3.  **Ver Pokémons capturados:** Exibe a lista dos Pokémons que você conseguiu capturar, com seus atributos.
4.  **Batalhar com um Pokémon selvagem:** Escolha um Pokémon selvagem (não capturado) e um dos seus Pokémons capturados para uma batalha. A vitória é decidida comparando a força dos Pokémons.
5.  **Configurações:** Permite alterar o tempo de resposta das animações de texto (como os "..." durante a captura ou batalha).
6.  **Sair:** Fecha o programa.

Siga as instruções na tela para navegar pelas opções e interagir com o jogo.

## Estrutura do Projeto

*   `main.go`: O coração da aplicação. Contém toda a lógica do jogo, incluindo:
    *   A struct `Pokemon` para definir os atributos de cada monstrinho.
    *   Variáveis globais para o banco de dados inicial de Pokémons (`pokemons`), a lista de Pokémons capturados (`capturados`) e o tempo de resposta (`tempoResposta`).
    *   Funções para cada opção do menu (explorar, capturar, ver capturados, batalhar, configurações).
    *   Funções auxiliares para limpar a tela e pausar a execução.

## Melhorias Futuras (Sugestões)

Este projeto é um ótimo ponto de partida! Aqui estão algumas ideias para expandi-lo:

*   **Mais Pokémons:** Adicionar uma variedade maior de Pokémons ao banco de dados.
*   **Persistência de Dados:** Salvar o progresso (Pokémons capturados) em um arquivo (JSON, CSV, etc.) para que não se perca ao fechar o jogo.
*   **Sistema de Batalha Mais Complexo:**
    *   Introduzir Pontos de Vida (HP).
    *   Diferentes tipos de ataques.
    *   Vantagens e desvantagens de tipo (Água > Fogo, Fogo > Grama, etc.).
    *   Turnos de batalha.
*   **Evoluções:** Implementar um sistema de evolução para os Pokémons.
*   **Itens:** Adicionar Pokébolas diferentes com taxas de captura variadas, poções, etc.
*   **Níveis e Experiência:** Permitir que os Pokémons ganhem experiência e subam de nível.
*   **Testes Unitários:** Escrever testes para garantir a robustez do código.

## 🤝 Contribuições

Contribuições são bem-vindas! Se você tem ideias para melhorar este projeto, sinta-se à vontade para abrir uma *issue* para discutir suas sugestões ou um *pull request* com suas implementações.

---

Divirta-se jogando e codando!!