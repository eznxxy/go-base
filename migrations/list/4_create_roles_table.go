package list

import (
	mysql "github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateRoleTable struct{}

func (m *CreateRoleTable) GetName() string {
	return "CreateRoleTable"
}

func (m *CreateRoleTable) Up(con *sqlx.DB) {
	table := mysql.NewTable("roles", con)
	table.Column("id").Type("int unsigned").Autoincrement()
	table.PrimaryKey("id")
	table.String("name", 255)
	table.Column("permissions").Type("text").Nullable()
	table.Column("has_full_access").Type("boolean").Default("false")
	table.Column("deleted_at").Type("datetime").Nullable()
	table.WithTimestamps()

	table.MustExec()
}

func (m *CreateRoleTable) Down(con *sqlx.DB) {
	mysql.DropTable("roles", con).MustExec()
}
