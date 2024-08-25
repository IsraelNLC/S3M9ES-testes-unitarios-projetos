package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	db "src/database"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

// HandlerTestSuite agrupa os testes relacionados aos handlers (controladores) da API.
// Em TDD, criar uma suíte de testes como essa permite agrupar testes de maneira lógica e executar todos
// os testes relacionados a uma determinada funcionalidade de uma só vez, garantindo que tudo esteja funcionando em conjunto.
type HandlerTestSuite struct {
	suite.Suite
	router *gin.Engine
}

// SetupSuite é executado antes de todos os testes e prepara o ambiente de teste.
// Em TDD, o SetupSuite é crucial para garantir que o ambiente esteja configurado corretamente antes de qualquer teste ser executado.
// Aqui, o banco de dados é inicializado e as rotas da API são configuradas.
func (suite *HandlerTestSuite) SetupSuite() {
	db.InitDatabase()            // Inicializa o banco de dados de teste com o schema necessário.
	suite.router = gin.Default() // Cria um novo roteador Gin para manipular as requisições HTTP nos testes.
	SetupRoutes(suite.router)    // Configura as rotas da API para que possam ser testadas.
}

// TearDownSuite é executado após todos os testes e limpa o ambiente de teste.
// Em TDD, é importante limpar o estado do sistema após os testes para evitar que dados residuais afetem testes futuros.
func (suite *HandlerTestSuite) TearDownSuite() {
	// Aqui, a tabela de usuários é removida do banco de dados para garantir que nenhum dado persista além do necessário.
	db.GetDB().Exec("DROP TABLE users")
}

// TestCreateUser verifica se a criação de um usuário via API funciona.
// Este teste é escrito antes da implementação da funcionalidade correspondente, seguindo o ciclo do TDD:
// 1. Escreva um teste que falha porque a funcionalidade ainda não foi implementada.
// 2. Implemente a funcionalidade necessária para fazer o teste passar.
// 3. Refatore o código mantendo os testes verdes.
func (suite *HandlerTestSuite) TestCreateUser() {
	// Criação de um objeto usuário para ser enviado na requisição.
	user := db.User{Name: "Test User"}
	userJSON, _ := json.Marshal(user) // Converte o usuário em JSON para ser enviado na requisição.

	// Cria uma requisição POST simulada para a rota "/users".
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json") // Define o cabeçalho da requisição para indicar que o conteúdo é JSON.

	// httptest.NewRecorder grava a resposta da requisição simulada.
	w := httptest.NewRecorder()
	// ServeHTTP processa a requisição usando o roteador configurado.
	suite.router.ServeHTTP(w, req)

	// Verifica se o código de status da resposta é 200 OK.
	suite.Equal(http.StatusOK, w.Code)

	// Verifica se o usuário foi criado corretamente, comparando o nome do usuário criado com o enviado.
	var createdUser db.User
	json.Unmarshal(w.Body.Bytes(), &createdUser)
	suite.Equal(user.Name, createdUser.Name, "User names should match")
}

// TestGetUsers verifica se a recuperação dos usuários via API funciona.
// Este teste também segue o ciclo TDD, sendo inicialmente escrito para falhar e depois guiando a implementação da funcionalidade.
func (suite *HandlerTestSuite) TestGetUsers() {
	// Cria uma requisição GET simulada para a rota "/users".
	req, _ := http.NewRequest("GET", "/users", nil)

	// Cria um gravador de resposta para capturar a resposta da requisição simulada.
	w := httptest.NewRecorder()
	// ServeHTTP processa a requisição usando o roteador configurado.
	suite.router.ServeHTTP(w, req)

	// Verifica se o código de status da resposta é 200 OK.
	suite.Equal(http.StatusOK, w.Code)

	// Verifica se a lista de usuários retornada não está vazia.
	var users []db.User
	json.Unmarshal(w.Body.Bytes(), &users)
	suite.NotEmpty(users, "User list should not be empty")
}

// TestSuite executa a suíte de testes.
// Em TDD, esta função é essencial para executar todos os testes definidos na suíte. Se qualquer teste falhar,
// você terá um feedback imediato, permitindo corrigir o problema antes de prosseguir com o desenvolvimento.
func TestSuite(t *testing.T) {
	// A função suite.Run executa todos os métodos de teste definidos na estrutura HandlerTestSuite.
	suite.Run(t, new(HandlerTestSuite))
}
