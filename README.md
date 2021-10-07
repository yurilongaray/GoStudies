## Description
### Ellen Korbes
- https://www.youtube.com/watch?v=WiGU_ZB-u0w&t=237s&ab_channel=AprendaGo
- https://github.com/ellenkorbes/aprendago

## Docs
- https://gobyexample.com/
- https://golang.org/doc/effective_go


## Disappointments
- There is no Ternary and TryCatch in GoLang...

## Testing

```bash
	go test -v ./...
```
- "-v" = verbose
- "./..." = run in all folders and subfolders of your path

## Tools

- Go Doc (browser documentation from your code): godoc --http=:4000
- Go Fmt (format go files): gofmt -w ./...
- Go Vet (correctness → fetch for wrong code): go vet ./...
- Go Lint (suggestions → fetch for bad code style): golint ./...
- Go Test Flag (shows every kind of test): go help testflag
- Go Test Coverage:
	- go test --coverprofile coverage_report.out
	- go tool cover -html=coverage_report.out

Obs.: When using "./..." the command will consider all folders and subfolders of your path


# Padrão de commits.

Os commits são uma parte importantissima do densenvolvimento. Com eles é possivel vee a linha de pensamento do developer. Portanto requere-se um padrão para como os commits são feitos.

## Quanto a forma de commitar

- Commits devem ser um checkpoint autosuficiente, ou seja, independente do commit a aplicação está funcionando.
- Commits possuem duas características principais
	- O tipo da alteração
	- O escopo da alteração
- Se os arquivos alterados possuem dois tipos e dois escopos, serão feitos 4 commits, um para cada relação tipo x escopo.
- Os commits devem possuir uma ordem lógica. Não se commitam alterações cuja suas dependências não foram ainda submetidas.
- A primeira linha do commit é estritamente importante e **nenhuma linha deve ultrapassar o limite de 120 caracteres.**
- As mensagens do commit devem ser claras, diretas e explicativas.


## Quanto a mensagem do commit

```
<type>(<scope>): <subject> #<issue|jira-task-name>
<body?>
```

### type

Indica de que tipo é a alteração que está sendo feita. Serve principalmente para categorizar o impacto do commit no código-fonte. 

Os tipos permitidos são:

- **feature**: uma nova funcionalidade no sistema, porém que faça diferença no produto final para o usuário. Ex: um método novo, adição de uma classe.
- **fix**: uma correção de erro do usuário. Ex: uma condicional errada ou sql.
- **docs**: alterando uma documentação, geralmente comentários. Ex: alterando um comentário (dãã)
- **style**: para alterações referentes a estilo do código. Ex: erros de tslint, espaçamentos faltantes
- **refactor**: refatorações de código que não afetam o produto final. Ex: renomeação de variaveis, troca de arquivos/metodos de lugar.
- **test**: alteração de arquivos de tests. Ex: adição de testes, refatoração de testes, remoção de testes
- **chore**: alteração de scripts não relacionados ao negócio. Ex: _package.json_, _circle.yml_
- **merge**: para merges/resolução de conflitos.

>note: Se um arquivo possue mais de um tipo de alteração, este será submetido em mais de um commit, um para cada tipo de alteração.

### scope

O escopo consiste no objeto da alteração, ou seja, o motivo pelo qual este commit está sendo feito.

#### Como definir um scope

- todos os arquivos de um módulo podem ser commitados com o escopo do mesmo:
	- Ex: _employee.controller.ts_ e _employee.repository.ts_ possuem o escopo **gate**.
- arquivos comuns podem ter o escopo do seu tipo:
	- Ex: _date.helper.ts_, _ip.helper.ts_ e _replace.helper.ts_ possuem o escopo **helper**.
- arquivos com alterações consequentes fazem parte do mesmo escopo:
	- Ex: alteração de um helper _date.helper.ts_ que é usado em _employee.controller.ts_. Se a alteração do primeiro arquivo implica em uma alteração do segundo arquivo, ambos possuem escopo **date** ou **helper** pois é o objeto da mudança.

> note: caso o escopo se torne dificil de definir, geralmente muita coisa está sendo commitada ao mesmo tempo. É necessário quebrar em mais commits.

### subject

O subject é o texto descritivo do commit. Deve ser um texto breve porém entendível do que foi feito no commit.

Deve ser escrito sempre em **inglês e no gerúndio**.
- Ex: _adding new date service_ ou _creating new endpoint for vehicle_.

Junto como escopo, tipo e a task não deve ultrapassar o limite de **72 caracteres**.

O escopo e o tipo já são informações, então tentar não suplicá-la no subject.

Ruim | Melhor
---|---
_test(employee): testing employee controller_ | _test(employee): creating new tests for controller_
_refactor(helper): refactoring halpers_ | _refactor(helper): organizing in 2 different files_
_fix(login): fixing error on login_ | _fix(login): error on submit password_

### jira-task-name ou issue

Aqui a issue do Github prefixados por hashtag (**#**).

- Ex: **#77**, **#COD-992**
