package cmd

import (
	"log"

	"github.com/Markogoodman/markocommander/internal/sql2struct"
	"github.com/spf13/cobra"
)

var (
	username  string
	password  string
	host      string
	charset   string
	dbType    string
	dbName    string
	tableName string
)

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql command transformation",
	Long:  "sql command transformation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "transform",
	Long:  "transform",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			UserName: username,
			Password: password,
			Charset:  charset,
		}
		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect err: %v", err)
			return
		}
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v", err)
			return
		}

		tpl := sql2struct.NewStructTemplate()
		tplColumns := tpl.AssemblyColumns(columns)
		err = tpl.Generate(tableName, tplColumns)
		if err != nil {
			log.Fatalf("tpl.Generate err: %v", err)
			return
		}
	},
}

func init() {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "", "", "username")
	sql2structCmd.Flags().StringVarP(&password, "password", "", "", "password")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "host")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "charset")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "dbType")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "", "", "dbName")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "", "", "tableName")
}
