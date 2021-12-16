package db

import (
	"errors"
	"io"
	"regexp"

	"github.com/go-qbit/model/relation"
	mysql "github.com/go-qbit/storage-mysql"
	mysql2 "github.com/go-sql-driver/mysql"

	"riceboards/db/confident_levels"
	"riceboards/db/goals"
	"riceboards/db/idea_goals"
	"riceboards/db/idea_teams"
	"riceboards/db/ideas"
	"riceboards/db/projects"
	"riceboards/db/teams"
	"riceboards/db/users"
)

type Db struct {
	Storage *mysql.MySQL

	Users *users.Table

	Projects        *projects.Table
	ConfidentLevels *confident_levels.Table
	Goals           *goals.Table
	Teams           *teams.Table
	Ideas           *ideas.Table
	IdeaGoals       *idea_goals.Table
	IdeaTeams       *idea_teams.Table
}

var reEmpty = regexp.MustCompile(`^\s*$`)

func New() *Db {
	storage := mysql.NewMySQL()

	tUsers := users.New(storage)

	tProject := projects.New(storage)
	relation.AddManyToOne(tProject, tUsers, relation.WithAlias("owner"), relation.WithRequired(true))

	tConfidentLevels := confident_levels.New(storage)
	relation.AddManyToOne(tConfidentLevels, tUsers, relation.WithAlias("owner"), relation.WithRequired(true))
	relation.AddManyToOne(tConfidentLevels, tProject, relation.WithAlias("project"), relation.WithRequired(true))

	tGoals := goals.New(storage)
	relation.AddManyToOne(tGoals, tUsers, relation.WithAlias("owner"), relation.WithRequired(true))
	relation.AddManyToOne(tGoals, tProject, relation.WithAlias("project"), relation.WithRequired(true))

	tTeams := teams.New(storage)
	relation.AddManyToOne(tTeams, tUsers, relation.WithAlias("owner"), relation.WithRequired(true))
	relation.AddManyToOne(tTeams, tProject, relation.WithAlias("project"), relation.WithRequired(true))

	tIdeas := ideas.New(storage)
	relation.AddManyToOne(tIdeas, tUsers, relation.WithAlias("owner"), relation.WithRequired(true))
	relation.AddManyToOne(tIdeas, tProject, relation.WithAlias("project"), relation.WithRequired(true))
	relation.AddManyToOne(tIdeas, tConfidentLevels, relation.WithAlias("confident"), relation.WithRequired(false))

	tIdeaGoals := idea_goals.New(storage)
	relation.AddManyToOne(tIdeaGoals, tUsers, relation.WithAlias("owner"), relation.WithRequired(true))
	relation.AddManyToOne(tIdeaGoals, tIdeas, relation.WithAlias("idea"), relation.WithRequired(true))
	relation.AddManyToOne(tIdeaGoals, tGoals, relation.WithAlias("goal"), relation.WithRequired(true))

	tIdeaTeams := idea_teams.New(storage)
	relation.AddManyToOne(tIdeaTeams, tUsers, relation.WithAlias("owner"), relation.WithRequired(true))
	relation.AddManyToOne(tIdeaTeams, tIdeas, relation.WithAlias("idea"), relation.WithRequired(true))
	relation.AddManyToOne(tIdeaTeams, tTeams, relation.WithAlias("team"), relation.WithRequired(true))

	return &Db{
		storage,
		tUsers,
		tProject, tConfidentLevels, tGoals, tTeams,
		tIdeas, tIdeaGoals, tIdeaTeams,
	}
}

func (d *Db) Connect(dsn string) error {
	return d.Storage.Connect(dsn)
}

func (d *Db) Disconnect() error {
	return d.Storage.Disonnect()
}

func (d *Db) PrintCreateSql(w io.Writer) {
	sqlBuf := mysql.NewSqlBuffer()
	d.Storage.WriteCreateSQL(sqlBuf)
	_, _ = sqlBuf.WriteTo(w)
}

func IsDuplicateEntryErr(err error) bool {
	mysqlErr := &mysql2.MySQLError{}
	return errors.As(err, &mysqlErr) && mysqlErr.Number == 1062
}

func FieldIsEmpty(s string) bool {
	return reEmpty.MatchString(s)
}
