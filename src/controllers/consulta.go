package controllers

import (
	"fmt"
	"os/exec"
	"strings"
)

func ConsultaLAPS(ativo string) string {
	//comando para consultar a senha um hostname no AD
	ps := "powershell.exe"
	cm := "-command"
	ad := "Get-ADComputer -Identity"
	ad2 := "-Properties *"

	//executa o comando no cmd
	cmd := exec.Command(ps, cm, ad, ativo, ad2)
	//retorna os dados do cmd
	cmdOut, err := cmd.Output()
	if err != nil {
		fmt.Println("Erro ao realizar a consulta!")
	}
	cmdReturn := string(cmdOut)

	//filtra onde estará a informação
	var senhaWithoutSpace = ""
	if strings.Contains(cmdReturn, "ms-Mcs-AdmPwd") {
		SenhaSliceInitial := strings.Index(cmdReturn, "ms-Mcs-AdmPwd")
		SenhaSliceEnding := strings.Index(cmdReturn, `ms-Mcs-AdmPwdExpirationTime`)
		//fatia os dados recebidos acima
		senha := string(cmdReturn)[SenhaSliceInitial+38 : SenhaSliceEnding]
		//formata os dados, removendo os espaços
		senhaWithoutSpace = strings.TrimSpace(senha)
	} else {
		senhaWithoutSpace = "Não encontrada!"
	}

	return senhaWithoutSpace

}
