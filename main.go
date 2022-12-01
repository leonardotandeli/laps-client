package main

import (
	"bufio"
	"fmt"
	"lapsv/src/controllers"
	"lapsv/src/database"
	"lapsv/src/models"
	"os/exec"
	"runtime"

	"os"
	"strings"
)

func init() {
	//limpa tela
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

var clear map[string]func()

func Clear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		fmt.Println("Erro ao limpar a tela!")
	}
}

func main() {
	//inicia conexão com o banco de dados.
	database.Dbconnection()

	//imprime as informações na tela
	fmt.Println("___________________________________________________________________ ")
	fmt.Println("|                                                                   | ")
	fmt.Println("|                       LAPS PASSWORD RECOVERY                      | ")
	fmt.Println("|               Acesso somente a pessoas autorizadas.               | ")
	fmt.Println("|  Desconecte IMEDIATAMENTE se você nao for um usuário autorizado!  | ")
	fmt.Println("|                                                                   | ")
	fmt.Println("|                 Desenvolvido por Leonardo Tandeli                 | ")
	fmt.Println("|___________________________________________________________________| ")
	fmt.Println("")

	//scan dos dados digitados
	fmt.Print("Informe o Ativo para busca: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		scanAtivo(scanner.Text())
	}
}

func scanAtivo(hostname string) {
	//chama a função de limpar tela
	Clear()

	// exibe as informações
	fmt.Println("___________________________________________________________________ ")
	fmt.Println("|                                                                   | ")
	fmt.Println("|                       LAPS PASSWORD RECOVERY                      | ")
	fmt.Println("|               Acesso somente a pessoas autorizadas.               | ")
	fmt.Println("|  Desconecte IMEDIATAMENTE se você nao for um usuário autorizado!  | ")
	fmt.Println("|                                                                   | ")
	fmt.Println("|                 Desenvolvido por Leonardo Tandeli                 | ")
	fmt.Println("|___________________________________________________________________| ")
	fmt.Println("")
	fmt.Println(" ___________________________________________________________________ ")
	fmt.Println("|                                                                   | ")
	fmt.Println("               ATIVO INFORMADO PARA BUSCA: " + hostname)
	fmt.Println("|___________________________________________________________________| ")
	fmt.Println("")

	fmt.Println("Senha armazenada no AD: " + controllers.ConsultaLAPS(hostname))
	fmt.Println("____________________________________________________________________ ")
	fmt.Println("")
	var ativo models.Ativo

	database.DB.Where("hostname = ?", strings.TrimSpace(hostname)).First(&ativo)

	if ativo.SenhaAtual != "" {
		fmt.Println("Senhas armazenadas no banco de dados:")
		fmt.Println("==> Coletado em: " + ativo.DataSenhaAtual.Format("02/01/2006") + " | Senha: " + ativo.SenhaAtual)
	}

	if ativo.PenultimaSenha != "" {
		fmt.Println("==> Coletado em: " + ativo.DataPenultimaSenha.Format("02/01/2006") + " | Senha: " + ativo.PenultimaSenha)
	}

	if ativo.AntePenultimaSenha != "" {
		fmt.Println("==> Coletado em: " + ativo.DataAntePenultimaSenha.Format("02/01/2006") + " | Senha: " + ativo.AntePenultimaSenha)
	}
	if ativo.SenhaAtual != "" {
		fmt.Println("____________________________________________________________________ ")
	}

	fmt.Println("")
	fmt.Print("Informe o Ativo para busca: ")
}
