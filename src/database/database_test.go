package db

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DatabaseTestSuite é uma estrutura que agrupa os testes relacionados ao banco de dados.
// Em TDD, a prática de agrupar testes em uma "suíte de testes" ajuda a organizar e categorizar testes relacionados,
// permitindo que você execute todos os testes relacionados a uma funcionalidade específica de uma vez.
type DatabaseTestSuite struct {
	suite.Suite
	db *gorm.DB
}

// SetupSuite é executado antes de todos os testes.
// Este método é essencial no ciclo de TDD porque prepara o ambiente de teste, garantindo que tudo esteja pronto
// antes de qualquer teste ser executado. No caso de bancos de dados, isso pode envolver a conexão ao banco,
// a configuração do esquema, e outras tarefas de inicialização.
func (suite *DatabaseTestSuite) SetupSuite() {
	dsn := "user=postgres password=israelsenha dbname=atvs3 sslmode=disable" // Configuração de conexão com o banco de dados.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	suite.Require().NoError(err, "Failed to connect to database") // Em TDD, você espera que a conexão seja bem-sucedida;
	// se falhar, o teste para aqui, o que significa que o problema de conexão deve ser resolvido antes de prosseguir.

	// Migrar o modelo User
	// Este passo garante que a estrutura do banco de dados está alinhada com o modelo de dados no código.
	// No contexto de TDD, você criaria testes que verificam se a migração foi realizada com sucesso antes de implementar
	// esta funcionalidade.
	err = db.AutoMigrate(&User{})
	suite.Require().NoError(err, "Failed to migrate database schema") // O teste falhará se a migração não ocorrer como esperado.

	// Armazena a conexão com o banco de dados na estrutura de testes.
	// Isso facilita o acesso ao banco em outros métodos de teste.
	suite.db = db
}

// TearDownSuite é executado após todos os testes.
// Este método é responsável por limpar o ambiente de teste, garantindo que nenhum resíduo de teste interfira em testes futuros.
// Em TDD, é importante garantir que os testes sejam isolados e não dependam do estado deixado por testes anteriores.
func (suite *DatabaseTestSuite) TearDownSuite() {
	// Aqui, estamos removendo a tabela de usuários do banco de dados de teste.
	// Esta é uma boa prática em TDD para evitar que dados de testes anteriores afetem novos testes.
	suite.db.Exec("DROP TABLE users")
}

// TestUserCreation verifica se a criação de um usuário funciona.
// Este é um exemplo clássico de como o TDD funciona:
// 1. Você primeiro escreveria este teste (ele falharia porque a funcionalidade ainda não foi implementada).
// 2. Depois, você implementaria a funcionalidade de criação de usuário até que o teste passasse.
// 3. Por fim, você refatoraria o código se necessário, mantendo o teste sempre verde.
func (suite *DatabaseTestSuite) TestUserCreation() {
	user := User{Name: "Test User"} // Criação de um usuário para testar a funcionalidade.
	err := suite.db.Create(&user).Error
	suite.Require().NoError(err, "Failed to create user") // Verifica se o usuário foi criado sem erros.

	var retrievedUser User
	err = suite.db.First(&retrievedUser, "name = ?", "Test User").Error   // Busca o usuário recém-criado no banco de dados.
	suite.Require().NoError(err, "Failed to retrieve user")               // Verifica se o usuário foi recuperado sem erros.
	suite.Equal(user.Name, retrievedUser.Name, "User names should match") // Verifica se o nome do usuário recuperado é o mesmo do criado.
	// Isso é fundamental em TDD: você especifica exatamente o que espera que aconteça, e o teste falha se o resultado
	// não for o esperado.
}

// TestSuite é a função que executa a suíte de testes.
// No ciclo TDD, essa função orquestra a execução dos testes, garantindo que todos os testes definidos na suíte sejam
// executados. Se qualquer um dos testes falhar, você deve voltar e corrigir o código ou o teste até que todos passem.
func TestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseTestSuite)) // Inicia a suíte de testes, que executa todos os métodos de teste definidos.
}
