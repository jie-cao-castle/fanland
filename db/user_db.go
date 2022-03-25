package dao

import (
	"context"
	"database/sql"
	"fanland/db/dao"
	log "github.com/sirupsen/logrus"
	"time"
)

type UserDB struct {
	DB
}

func (f *UserDB) Open() error {
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/"+f.dbName+"?parseTime=true")
	f.db = db
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (f *UserDB) Insert(user *dao.UserDO) (err error) {

	query := "INSERT INTO fanland_user (user_name, user_desc, avatar_url, create_time, update_time) VALUES (?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	stmt, err := f.db.PrepareContext(ctx, query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, user.UserName, user.UserDesc, user.AvatarUrl)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}

	return nil
}

func (f *UserDB) Update(user *dao.UserDO) (err error) {

	query := "UPDATE fanland_user SET user_name=?, user_desc =?, avatar_url =?, update_time = CURRENT_TIMESTAMP WHERE id=?"
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	stmt, err := f.db.PrepareContext(ctx, query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, user.UserName, user.UserDesc, user.AvatarUrl)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}

	return nil
}

func (f *UserDB) GetById(userId uint64) (*dao.UserDO, error) {
	var (
		name       string
		desc       string
		id         uint64
		imgUrl     string
		nftId      uint64
		tags       string
		createTime time.Time
		updateTime time.Time
	)

	rows, err := f.db.Query("select id, user_name, user_desc, avatar_url, create_time, update_time from fanland_user where id = ?", userId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&id, &name, &desc, &imgUrl, &nftId, &tags, &createTime, &updateTime)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	user := &dao.UserDO{
		Id:        id,
		UserName:  name,
		UserDesc:  desc,
		AvatarUrl: imgUrl,
	}
	return user, nil
}

func (f *UserDB) GetByName(userName string) (*dao.UserDO, error) {
	var (
		name       string
		desc       string
		id         uint64
		avatarUrl  string
		userHash   string
		createTime time.Time
		updateTime time.Time
	)

	rows, err := f.db.Query("select id, user_name, user_desc, avatar_url, user_hash, create_time, update_time from fanland_user where user_name = ?", userName)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&id, &name, &desc, &avatarUrl, &userHash, &createTime, &updateTime)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	user := &dao.UserDO{
		Id:        id,
		UserName:  name,
		UserDesc:  desc,
		AvatarUrl: avatarUrl,
		UserHash:  userHash,
	}
	return user, nil
}
