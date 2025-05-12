package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Estrutura para armazenar dados de um Pokémon
type Pokemon struct {
	Nome      string
	Tipo      string
	Forca     int
	Defesa    int
	Capturado bool
}

// Configurações iniciais
var tempoResposta = 2 * time.Second

// Banco de dados inicial de Pokémons
var pokemons = []Pokemon{
	{"Pikachu", "Elétrico", 55, 40, false},
	{"Charmander", "Fogo", 52, 43, false},
	{"Squirtle", "Água", 48, 65, false},
	{"Bulbasaur", "Grama", 49, 49, false},
}

// Lista de Pokémons capturados
var capturados = []Pokemon{}

// Função principal
func main() {
	rand.Seed(time.Now().UnixNano())
	scanner := bufio.NewScanner(os.Stdin)

	for {
		limparTela()
		fmt.Println("🌟 Bem-vindo à sua Pokédex! 🌟")
		fmt.Println("1. Explorar Pokémons")
		fmt.Println("2. Tentar capturar Pokémon")
		fmt.Println("3. Ver Pokémons capturados")
		fmt.Println("4. Batalhar com um Pokémon selvagem")
		fmt.Println("5. Configurações")
		fmt.Println("6. Sair")
		fmt.Print("Escolha uma opção: ")

		if !scanner.Scan() {
			break
		}

		opcao := scanner.Text()
		switch opcao {
		case "1":
			explorarPokemons()
		case "2":
			tentarCapturar(scanner)
		case "3":
			verCapturados()
		case "4":
			batalhar(scanner)
		case "5":
			configuracoes(scanner)
		case "6":
			fmt.Println("Encerrando sua Pokédex. Até a próxima!")
			return
		default:
			fmt.Println("Opção inválida. Tente novamente.")
			pausar()
		}
	}
}

// Explorar Pokémons disponíveis
func explorarPokemons() {
	limparTela()
	fmt.Println("📖 Pokémons disponíveis:")
	for _, p := range pokemons {
		status := "Disponível"
		if p.Capturado {
			status = "Capturado"
		}
		fmt.Printf("- %s (Tipo: %s) - %s\n", p.Nome, p.Tipo, status)
	}
	pausar()
}

// Tentar capturar um Pokémon
func tentarCapturar(scanner *bufio.Scanner) {
	limparTela()
	fmt.Println("🎯 Qual Pokémon você quer tentar capturar?")
	explorarPokemons()
	fmt.Print("Digite o nome do Pokémon: ")

	if !scanner.Scan() {
		return
	}

	nome := strings.TrimSpace(scanner.Text())
	for i, p := range pokemons {
		if strings.EqualFold(p.Nome, nome) {
			if p.Capturado {
				fmt.Printf("%s já foi capturado.\n", p.Nome)
			} else {
				fmt.Printf("Tentando capturar %s...\n", p.Nome)
				pausar()
				if rand.Intn(100) < 50 { // 50% de chance de capturar
					fmt.Printf("🎉 Você capturou %s!\n", p.Nome)
					pokemons[i].Capturado = true
					capturados = append(capturados, p)
				} else {
					fmt.Printf("💔 O %s escapou!\n", p.Nome)
				}
			}
			pausar()
			return
		}
	}
	fmt.Println("Pokémon não encontrado.")
	pausar()
}

// Ver Pokémons capturados
func verCapturados() {
	limparTela()
	if len(capturados) == 0 {
		fmt.Println("Você ainda não capturou nenhum Pokémon.")
	} else {
		fmt.Println("🏆 Seus Pokémons capturados:")
		for _, p := range capturados {
			fmt.Printf("- %s (Tipo: %s, Força: %d, Defesa: %d)\n", p.Nome, p.Tipo, p.Forca, p.Defesa)
		}
	}
	pausar()
}

// Sistema de batalha
func batalhar(scanner *bufio.Scanner) {
	limparTela()
	fmt.Println("⚔️ Escolha um Pokémon para batalhar:")
	explorarPokemons()
	fmt.Print("Digite o nome do Pokémon selvagem: ")

	if !scanner.Scan() {
		return
	}
	nome := strings.TrimSpace(scanner.Text())

	var oponente *Pokemon
	for i := range pokemons {
		if strings.EqualFold(pokemons[i].Nome, nome) && !pokemons[i].Capturado {
			oponente = &pokemons[i]
			break
		}
	}

	if oponente == nil {
		fmt.Println("Pokémon inválido ou já capturado.")
		pausar()
		return
	}

	fmt.Println("Escolha um Pokémon seu para batalhar:")
	verCapturados()
	fmt.Print("Digite o nome do seu Pokémon: ")

	if !scanner.Scan() {
		return
	}

	nomeMeu := strings.TrimSpace(scanner.Text())
	var meuPokemon *Pokemon
	for i := range capturados {
		if strings.EqualFold(capturados[i].Nome, nomeMeu) {
			meuPokemon = &capturados[i]
			break
		}
	}

	if meuPokemon == nil {
		fmt.Println("Pokémon inválido.")
		pausar()
		return
	}

	// Simulação de batalha
	fmt.Printf("%s (Seu Pokémon) VS %s (Pokémon Selvagem)\n", meuPokemon.Nome, oponente.Nome)
	for i := 0; i < 3; i++ {
		fmt.Print(".")
		time.Sleep(tempoResposta / 3)
	}
	fmt.Println()

	if meuPokemon.Forca > oponente.Forca {
		fmt.Printf("🎉 Seu %s venceu a batalha!\n", meuPokemon.Nome)
	} else {
		fmt.Printf("💔 Seu %s perdeu a batalha.\n", meuPokemon.Nome)
	}
	pausar()
}

// Configurações da Pokédex
func configuracoes(scanner *bufio.Scanner) {
	limparTela()
	fmt.Println("⚙️ Configurações:")
	fmt.Println("1. Alterar tempo de resposta")
	fmt.Println("2. Voltar ao menu principal")
	fmt.Print("Escolha uma opção: ")

	if !scanner.Scan() {
		return
	}

	opcao := scanner.Text()
	switch opcao {
	case "1":
		fmt.Print("Digite o novo tempo de resposta (em segundos): ")
		if scanner.Scan() {
			var novoTempo int
			fmt.Sscanf(scanner.Text(), "%d", &novoTempo)
			tempoResposta = time.Duration(novoTempo) * time.Second
			fmt.Println("Tempo de resposta alterado.")
		}
	case "2":
		return
	default:
		fmt.Println("Opção inválida.")
	}
	pausar()
}

// Funções auxiliares
func limparTela() {
	fmt.Print("\033[H\033[2J")
}

func pausar() {
	fmt.Printf("\nPressione Enter para continuar...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
