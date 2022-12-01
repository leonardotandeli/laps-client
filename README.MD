# LAPS Password Recovery - CLIENT/VIEWER
Ferramenta CLI desenvolvida visando solucionar o problema de perda de acesso ao usuário local das workstations, visto que a solução LAPS da Microsoft não guarda o histórico de senhas, e caso a máquina saia do domínio, o usuário local fica inacessível, impossibilitando a manutenção do equipamento.

A aplicação foi escrita em Go, e armazena as senhas das workstations em um banco de dados MYSQL através de uma coleta diária que faz a busca de novas máquinas e senhas no Active Directory.

Neste repositório está hospedado o client para acesso as informações. 

### Funcionalidades
- Busca de senhas no AD e Histórico de senhas armazenadas em um banco MYSQL. 


## 💻 Screenshots

[![screenshot.png](https://raw.githubusercontent.com/leonardotandeli/laps-viewer/main/screenshot.png)](https://raw.githubusercontent.com/leonardotandeli/laps-viewer/main/screenshot.png)
