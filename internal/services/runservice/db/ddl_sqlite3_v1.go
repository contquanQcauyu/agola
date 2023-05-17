// Code generated by go generate; DO NOT EDIT.
package db

var DDLSqlite3V1 = []string{
	"create table if not exists changegroup (id varchar NOT NULL, revision bigint NOT NULL, creation_time timestamp NOT NULL, update_time timestamp NOT NULL, name varchar NOT NULL, value varchar NOT NULL, PRIMARY KEY (id))",
	"create table if not exists runconfig (id varchar NOT NULL, revision bigint NOT NULL, creation_time timestamp NOT NULL, update_time timestamp NOT NULL, name varchar NOT NULL, run_group varchar NOT NULL, setup_errors text NOT NULL, annotations text NOT NULL, static_environment text NOT NULL, environment text NOT NULL, tasks text NOT NULL, cache_group varchar NOT NULL, PRIMARY KEY (id))",
	"create table if not exists run (id varchar NOT NULL, revision bigint NOT NULL, creation_time timestamp NOT NULL, update_time timestamp NOT NULL, sequence integer NOT NULL UNIQUE, name varchar NOT NULL, run_config_id varchar NOT NULL, counter bigint NOT NULL, run_group varchar NOT NULL, annotations text NOT NULL, phase varchar NOT NULL, result varchar NOT NULL, stop integer NOT NULL, tasks text NOT NULL, enqueue_time timestamp, start_time timestamp, end_time timestamp, archived integer NOT NULL, PRIMARY KEY (id), foreign key (run_config_id) references runconfig(id))",
	"create table if not exists runcounter (id varchar NOT NULL, revision bigint NOT NULL, creation_time timestamp NOT NULL, update_time timestamp NOT NULL, group_id varchar NOT NULL UNIQUE, value bigint NOT NULL, PRIMARY KEY (id))",
	"create table if not exists runevent (id varchar NOT NULL, revision bigint NOT NULL, creation_time timestamp NOT NULL, update_time timestamp NOT NULL, sequence integer NOT NULL UNIQUE, run_id varchar NOT NULL, phase varchar NOT NULL, result varchar NOT NULL, PRIMARY KEY (id))",
	"create table if not exists executor (id varchar NOT NULL, revision bigint NOT NULL, creation_time timestamp NOT NULL, update_time timestamp NOT NULL, executor_id varchar NOT NULL, listen_url varchar NOT NULL, archs text NOT NULL, labels text NOT NULL, allow_privileged_containers integer NOT NULL, active_tasks_limit bigint NOT NULL, active_tasks bigint NOT NULL, dynamic integer NOT NULL, executor_group varchar NOT NULL, siblings_executors text NOT NULL, PRIMARY KEY (id))",
	"create table if not exists executortask (id varchar NOT NULL, revision bigint NOT NULL, creation_time timestamp NOT NULL, update_time timestamp NOT NULL, executor_id varchar NOT NULL, run_id varchar NOT NULL, run_task_id varchar NOT NULL, stop integer NOT NULL, phase varchar NOT NULL, timedout integer NOT NULL, fail_error varchar NOT NULL, start_time timestamp, end_time timestamp, setup_step text NOT NULL, steps text NOT NULL, PRIMARY KEY (id))",

	// indexes
	"create index if not exists run_group_idx on run(run_group)",
	"create index if not exists runcounter_group_id_idx on runcounter(group_id)",
}