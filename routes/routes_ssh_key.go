package routes

import (
	"github.com/pagoda-tech/bastion/models"
	"github.com/pagoda-tech/macaron"
	"strings"
	"golang.org/x/crypto/ssh"
)

// SSHKeyCreateForm 创建 SSHKey 的表单
type SSHKeyCreateForm struct {
	Name string `json:"name"`
	PublicKey string `json:"publicKey"`
}

// SSHKeyCreate 为一个用户创建 SSHKey
func SSHKeyCreate(ctx *macaron.Context, db *models.DB, r APIRender, a Auth, f SSHKeyCreateForm) {
	// userID
	userID := uint(ctx.ParamsInt(":userid"))

	if !a.CanAccessUser(userID) {
		r.Fail(UserNotFound, "没有找到该用户")
		return
	}

	// publicKey
	key, c, _, _, err := ssh.ParseAuthorizedKey([]byte(f.PublicKey))

	if err != nil {
		r.Fail(ParamsInvalid, "公钥格式不正确")
		return
	}

	// name
	name := strings.TrimSpace(f.Name)

	if len(name) == 0 {
		name = strings.TrimSpace(c)
	}

	if len(name) == 0 {
		name = "NO NAME"
	}

	if len(name) > 20 {
		r.Fail(ParamsInvalid, "名字过长")
		return
	}

	fp := ssh.FingerprintSHA256(key)

	// check duplicates
	sk := &models.SSHKey{}
	db.Where("fingerprint = ?", fp).First(sk)

	if !db.NewRecord(sk) {
		r.Fail(ParamsInvalid, "该 SSH 公钥已经被使用")
		return
	}

	// create
	sk.Name = name
	sk.Fingerprint = fp
	sk.PublicKey = strings.TrimSpace(string(ssh.MarshalAuthorizedKey(key)))
	sk.UserID = userID

	if err := db.Create(sk).Error; err != nil {
		r.Fail(InternalError, err.Error())
		return
	}

	r.Success("ssh_key", sk)
}

// SSHKeyList 列出 SSH 公钥
func SSHKeyList(ctx *macaron.Context, db *models.DB, r APIRender, a Auth) {
	userID := uint(ctx.ParamsInt(":userid"))

	if !a.CanAccessUser(userID) {
		r.Fail(UserNotFound, "没有找到该用户")
		return
	}

	sshKeys := []models.SSHKey{}

	db.Where("user_id = ?", userID).Find(&sshKeys)

	r.Success("sshKeys", sshKeys)
}

// SSHKeyDestroy 删除 SSH 公钥
func SSHKeyDestroy(ctx *macaron.Context, db *models.DB, r APIRender, a Auth) {
	id := uint(ctx.ParamsInt(":id"))
	
	// find sk
	sk := &models.SSHKey{}
	db.First(sk, id)

	if db.NewRecord(sk) || !a.CanAccessUser(sk.UserID) {
		r.Fail(SSHKeyNotFound, "没有找到 SSH 公钥")
		return
	}

	db.Unscoped().Delete(sk)
	r.Success()
}

