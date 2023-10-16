package service

import (
	"changeme/Utils"
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/sys/windows/registry"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

type Role struct {
	path  string
	ctx   context.Context
	roles []string
	paths []string
}

func NewRole() *Role {
	return &Role{}
}

func (rs *Role) Startup(ctx context.Context) {
	rs.ctx = ctx
}

func (rs *Role) GetPath(path string) string {
	if path == "" {
		key1, err := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\Kingsoft\\SeasunGame\\JX3", registry.ALL_ACCESS)
		if err != nil {
			log.Println(err)
		}
		value, _, err := key1.GetStringValue("InstallPath")
		rs.path = value
		defer key1.Close()
		rs.path = strings.ReplaceAll(rs.path, "/", "\\")
		return rs.path
	} else {
		rs.path = path
	}
	return rs.path
}

func (rs *Role) GetRoles() []string {
	roles := make([]string, 0)
	paths, _ := Utils.GetDirAllFilePaths(rs.path)
	rs.paths = paths
	for _, path := range rs.paths {
		if strings.Contains(path, "userpreferencesasync.jx3dat") {
			path = strings.ReplaceAll(path, "userpreferencesasync.jx3dat", "")
			role := strings.Split(path, "\\")
			length := len(role)
			my := ""
			for _, p := range rs.paths {
				if strings.Contains(p, "my#data") && strings.Contains(p, "info.jx3dat") {
					sourceExist, _ := Utils.PathExists(strings.ReplaceAll(p, "info.jx3dat", "") + role[length-2])
					if sourceExist {
						my = strings.ReplaceAll(p, "info.jx3dat", "config\\")
					}
				}
			}
			jx := ""
			sourceIds := strings.Split(my, "\\")
			sourceLength := len(sourceIds)
			if sourceLength > 1 {
				jxPath := strings.Split(my, "my#data")[0] + "JX#DATA\\skillnotice\\" + strings.ReplaceAll(sourceIds[sourceLength-3], "@zhcn_hd", ".jx3dat")
				jxExist, _ := Utils.PathExists(jxPath)
				if jxExist {
					jx = jxPath
				}
			}
			roles = append(roles, role[length-2]+"|"+role[length-3]+"|"+role[length-4]+"|"+role[length-5]+"|"+path+"|"+my+"|"+jx)
		}
	}
	return roles
}

func (rs *Role) SelectPath() string {
	pathTemp, _ := runtime.OpenDirectoryDialog(rs.ctx, runtime.OpenDialogOptions{})
	if pathTemp != "" {
		rs.path = pathTemp
		return rs.path
	}
	return ""
}

func (rs *Role) syncUIData(sourcePath, targetPath string) error {
	dir, err := ioutil.ReadDir(targetPath)
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{targetPath, d.Name()}...))
	}
	err = Utils.CopyDir(sourcePath, targetPath)
	if err != nil {
		return err
	}
	return nil
}

func (rs *Role) syncMyData(source, target string) error {

	if source != "" && target != "" {
		err := Utils.CopyDir(source, target)
		if err != nil {
			return err
		}
	}
	return nil
}

func (rs *Role) syncJxData(source, target string) error {
	sourceFile := ""
	sourceIds := strings.Split(source, "\\")
	sourceLength := len(sourceIds)
	if sourceLength > 1 {
		jx := strings.Split(source, "my#data")[0] + "JX#DATA\\skillnotice\\" + strings.ReplaceAll(sourceIds[sourceLength-3], "@zhcn_hd", ".jx3dat")
		jxExist, _ := Utils.PathExists(jx)
		if jxExist {
			sourceFile = jx
		}
	}
	targetFile := ""
	targetIds := strings.Split(target, "\\")
	targetLength := len(targetIds)
	if targetLength > 1 {
		targetFile = strings.Split(source, "my#data")[0] + "JX#DATA\\skillnotice\\" + strings.ReplaceAll(targetIds[targetLength-3], "@zhcn_hd", ".jx3dat")
	}
	if sourceFile == "" {
		return nil
	} else {
		_, err := Utils.CopyFile(sourceFile, targetFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func (rs *Role) SyncRoleData(source, target, sourcePath, targetPath, sourceMyPath, targetMyPath string, isUI bool, isMY bool, isJX bool) string {

	if isUI {
		err := rs.syncUIData(sourcePath, targetPath)
		if err != nil {
			return "[界面UI设置同步错误]" + err.Error()
		}
	}

	if isMY {
		err := rs.syncMyData(sourceMyPath, targetMyPath)
		if err != nil {
			return "[茗伊设置同步错误]" + err.Error()
		}
	}

	if isJX {
		err := rs.syncJxData(sourceMyPath, targetMyPath)
		if err != nil {
			return "[剑心喊话同步错误]" + err.Error()
		}
	}

	return "success"
}
