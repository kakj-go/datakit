package models

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"
)

const (
	StatusOK = 0
)

var (
	DB *sql.DB = nil

	MaxConn = 32
	Stmts   = map[string]*sql.Stmt{}
)

type Token struct {
	WsUUID   string
	MaxTSCnt int64
	Token    string
	DBUUID   string
	Status   int
	Creator  string

	CreateAt int64
	UpdateAt int64
	DeleteAt int64
}

type Agent struct {
	ID       int
	UUID     string
	Name     string
	Version  string
	Creator  string
	UploadAt int64
	Status   int

	CreateAt int64
	UpdateAt int64
	DeleteAt int64
}

type InfluxAuth struct {
	UserName   string `json:"username"`
	Password   string `json:"password,omitempty"`
	PwdEncrypt string `json:"passwordEncrypt,omitempty"`
}

type InfluxInstance struct {
	ID     int    `json:"id,omitempty"`
	UUID   string `json:"uuid,omitempty"`
	Status int    `json:"status,omitempty"`
	Host   string `json:"host"`

	User          string `json:"user"`
	Pwd           string `json:"password,omitempty"`
	Authorization string `json:"authorization,omitempty"`
	DBCount       int    `json:"db_count,omitempty"`

	CreateAt int64 `json:"create_at,omitempty"`
	UpdateAt int64 `json:"update_at,omitempty"`
	DeleteAt int64 `json:"delete_at,omitempty"`
}

type InfluxDBInfo struct {
	UUID               string `json:"uuid,omitempty"`
	DB                 string `json:"db,omitempty"`
	WsUUID             string `json:"ws_uuid,omitempty"`
	Status             int    `json:"status,omitempty"`
	InfluxInstanceUUID string `json:"influx_instance_uuid,omitempty"`
	RPUUID             string `json:"rp_uuid,omitempty"`

	Host string `json:"host,omitempty"`
	User string `json:"user,omitempty"`
	Pwd  string `json:"password,omitempty"`
	RP   string `json:"rp_name,omitempty"`
	CQRP string `json:"cq_rp,omitempty"`

	CreateAt int64 `json:"create_at,omitempty"`
	UpdateAt int64 `json:"update_at,omitempty"`
	DeleteAt int64 `json:"delete_at,omitempty"`
}

type InfluxCQ struct {
	SampleEvery   string `json:"aggr_every,omitempty"`
	SampleFor     string `json:"aggr_for,omitempty"`
	WorkspaceUUID string `json:"workspace_uuid,omitempty"`
	CQUUID        string `json:"cq_uuid,omitempty"`
	InfluxdbUUID  string `json:"db_uuid,omitempty"`
	DB            string `json:"db,omitempty"`
	RP            string `json:"rp,omitempty"`
	CQRP          string `json:"cqrp,omitempty"`
	Measurement   string `json:"measurement,omitempty"`
	GroupByTime   string `json:"aggr_period,omitempty"`
	GroupByOffset string `json:"group_by_offset,omitempty"`
	FuncAggr      string `json:"aggr_func,omitempty"`
}

type LogExtractRule struct {
	Source string `json:"source"`
	WsUUID string `json:"workspace_uuid"`
	Url    string `json:"url"`
}

type DKOnline struct {
	UUID            string `json:"uuid,omitempty"`
	Name            string `json:"name,omitempty"`
	Token           string `json:"token,omitempty"`
	DkUUID          string `json:"dkUUID,omitempty"`
	Version         string `json:"version,omitempty"`
	Os              string `json:"os,omitempty"`
	Arch            string `json:"arch,omitempty"`
	EnableInputs    []string `json:"enableInputs,omitempty"`
	AvailableInputs []string `json:"availableInputs,omitempty"`
	LastOnline      int64  `json:"lastOnline,omitempty"`
	LastHeartbeat   int64  `json:"lastHeartbeat,omitempty"`

	CreateAt int64 `json:"create_at,omitempty"`
	UpdateAt int64 `json:"update_at,omitempty"`
	DeleteAt int64 `json:"delete_at,omitempty"`
}

var (
	l *logger.Logger
)

func Init(dialect, connStr string) error {
	l = logger.SLogger("models")

	var err error
	for {
		DB, err = sql.Open(dialect, connStr)
		if err != nil {
			l.Errorf("%s", err.Error())
			time.Sleep(time.Second)
		} else {

			if err := DB.Ping(); err != nil {
				l.Errorf("ping %s failed: %s", dialect, err.Error())
				time.Sleep(time.Second)
			} else {
				l.Infof("connect to %s ok", dialect)
				break
			}
		}
	}

	DB.SetMaxOpenConns(MaxConn)
	DB.SetMaxIdleConns(4)
	DB.SetConnMaxLifetime(time.Second * 28)

	sqls := map[string]string{
		// get team/token info
		`qToken`: `SELECT uuid, dbUUID, maxTsCount FROM main_workspace WHERE status=? AND token=? limit 1`,

		//更新agent信息表
		`uAgent`: `UPDATE main_agent SET version=?,status=?,uploadAt=? WHERE uuid=?`,

		// 查询 influx instance
		`qInfluxInstances`: `SELECT host, authorization, dbcount, status, uuid FROM main_influx_instance WHERE status=? ORDER BY dbcount ASC`,

		`qInfluxInstanceByUUID`: `SELECT host, authorization, dbcount, status, uuid  FROM main_influx_instance WHERE status=? AND uuid=?`,

		//update influx instance
		`uInfluxInstance`: `UPDATE main_influx_instance SET host=?, authorization=?, dbcount=dbcount+1, status=?, updateAt=? WHERE uuid=?`,

		// add new ifdb
		`iInfluxdb`: `INSERT INTO main_influx_db(
			db, cqrp, influxRpName, uuid, influxRpUUID, influxInstanceUUID, createAt, updateAt) VALUES(?,?,?,?,?,?,?,?)`,

		//update rp
		`uInfluxdb`: `UPDATE main_influx_db SET influxRpName=?, updateAt=? WHERE uuid=?`,

		// get ifdb info
		`qInfluxdb`: `SELECT inst.host, inst.authorization, inst.uuid,
												 ifdb.influxRpName, ifdb.db, ifdb.cqrp FROM main_influx_db AS ifdb
					INNER JOIN main_influx_instance AS inst
						ON inst.uuid=ifdb.influxInstanceUUID
					WHERE ifdb.uuid=? AND ifdb.status=?`,

		// get all ifdbs from instance
		`getInstDBs`: `SELECT uuid, db FROM main_influx_db WHERE influxInstanceUUID=? AND status=?`,

		// cq 配置表操作
		`iInfluxCQ`: `INSERT INTO main_influx_cq(uuid, dbUUID, workspaceUUID, sampleEvery, sampleFor, measurement,
					rp,cqrp,aggrFunc,groupByTime,groupByOffset,createAt,creator) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,"kodo")`,

		`uInfluxCQ`: `UPDATE main_influx_cq SET sampleEvery=?, updateAt=?, groupByTime=?,
					aggrFunc=? WHERE uuid=? AND workspaceUUID=?`,

		`qInfluxCQ`: `SELECT dbUUID, sampleEvery, sampleFor, measurement,rp,cqrp,aggrFunc,groupByTime,
					groupByOffset FROM main_influx_cq WHERE uuid=? AND workspaceUUID=? AND status=?`,

		`qInfluxCQByM`: `SELECT uuid, workspaceUUID, sampleEvery, sampleFor, measurement,rp,cqrp,aggrFunc,groupByTime,
					groupByOffset FROM main_influx_cq WHERE dbUUID=? AND measurement=? AND status=?`,

		`dInfluxCQ`: `DELETE FROM main_influx_cq WHERE uuid=? AND workspaceUUID=?`,

		`dInfluxCQByuuid`: `DELETE FROM main_influx_cq WHERE uuid=?`,

		`qInfluxCQsByWsUUID`: `SELECT uuid, dbUUID, sampleEvery, sampleFor, measurement,rp,cqrp,aggrFunc,groupByTime,
					groupByOffset FROM main_influx_cq WHERE workspaceUUID=? AND status=?`,

		`qLogExtractRule`: `SELECT distinct url, workspaceUUID, source FROM main_log_extract_rule WHERE workspaceUUID=? AND status=?`,

		`updateDKOnline`: `UPDATE  main_datakit_online SET  name=?, token=?,hostName=?,ip=?, version=?, os=?, arch=?, inputInfo=?, 
					lastOnline=?, lastHeartbeat=?, status=? where dkUUID=?`,
		`setDKOnline` : `insert into main_datakit_online(uuid, name, token,hostName,ip, dkUUID, version, os, arch, inputInfo, 
					lastOnline, lastHeartbeat, status,creator, updator, createAt,updateAt) Values(?,?,?,?,?,?,?,?,?,?,?,?,?,"kodo","kodo",?,?)  `,
		`existDK`:`select 1 from main_datakit_online where dkUUID=? limit 1`,
		`updateDKStatus`: `update main_datakit_online set lastHeartbeat=?,status=?,updateAt=? where dkUUID=?`,
	}

	for k, v := range sqls {
		stmt, err := prepare(DB, v)
		if err != nil {
			l.Errorf("init stmt %s failed: %s", k, err)
			return err
		}

		l.Debugf("prepare ok: %s: %v", k, stmt)

		Stmts[k] = stmt
	}
	return nil
}

func CloseDB() error {
	for _, v := range Stmts {
		v.Close()
	}

	for _, db := range []*sql.DB{
		DB,
	} {
		if err := db.Close(); err != nil {
			l.Errorf("%s", err.Error())
			return err
		}
	}
	return nil
}

func prepare(db *sql.DB, s string) (*sql.Stmt, error) {
	stmt, err := db.Prepare(s)
	if err != nil {
		return nil, err
	}
	return stmt, err
}