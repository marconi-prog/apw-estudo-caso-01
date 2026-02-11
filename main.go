package main

import "fmt"


func main() {

	fmt.Println("te vejo nas olindas professor")

// RESPOSTAS PARA AS QUESTÕES:

// 1. Escalabilidade:
// Durante a Black Friday o sistema continuou funcionando,
// porém não conseguiu lidar com o grande volume de acessos.
// O tempo de resposta aumentou muito (de 2s para 45s),
// mostrando que a aplicação não estava preparada para escalar
// sob alta demanda.

// 2. Confiabilidade:
// O erro de concorrência que expôs dados de pagamento é uma
// falha grave. Isso mostra que o sistema não garantiu isolamento
// e consistência das informações entre os usuários,
// comprometendo a segurança e a confiança na aplicação.

// 3. Manutenibilidade:
// A dificuldade para corrigir o bug revelou que o sistema é
// complexo, pouco documentado e provavelmente muito acoplado.
// Quando uma simples alteração leva dias para ser testada e
// homologada, fica claro que a manutenção e evolução do software
// são difíceis e lentas.

}