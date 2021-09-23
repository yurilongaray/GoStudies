#### https://www.youtube.com/watch?v=WiGU_ZB-u0w&t=237s&ab_channel=AprendaGo

#### https://gobyexample.com/

#### https://golang.org/doc/effective_go

Dúvidas:

- Queria entender qual o melhor momento para fechar um canal... Devo fechar Um canal sempre que eu finalizar todos os inputs dele?
	* Aparentemente sim, toda vez que um canal (RECEIVER) terminar de receber todos os seus dados, ele deve ser fechado "close(channel)"
- Se eu possuo uma go func e dentro dela tem outra go func, meu WaitGroup.Add precisa ser de 2 ou de 1 (ex.: wg.Add(2))
- Rever o Context