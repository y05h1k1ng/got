package schema

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"yoshiking/got/util"
)

type GitRepository struct {
	workTree string
	gitDir   string
}

func (r *GitRepository) New(path string) error {
	r.workTree = path
	r.gitDir = filepath.Join(path, ".git")
	log.Printf("workTree: %s, gitDir: %s\n", r.workTree, r.gitDir)

	// TODO: load config

	if util.IsExist(r.gitDir) {
		return errors.New("The File already exists.")
	}
	if err := os.Mkdir(r.gitDir, 0775); err != nil {
		return err
	}

	dirs := []string{"branches", "objects", "refs/tags", "refs/heads"}
	for _, dir := range dirs {
		err := r.CreateDir(dir)
		if err != nil {
			return err
		}
	}

	files := [][]string{
		{"description", "Unnamed repository; edit this file 'description' to name the repository.\n"},
		{"config", ""},
		{"HEAD", "ref: refs/heads/master\n"},
	}
	for _, file := range files {
		err := r.CreateFile(file[0], file[1])
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *GitRepository) Find(path string) error {
	if util.IsExist(filepath.Join(path, ".git")) {
		// found workTree's path has .git/
		r.workTree = path
		r.gitDir = filepath.Join(path, ".git")
		log.Printf("Found .git/!!!!\nworkTree: %s, gitDir: %s\n", r.workTree, r.gitDir)
		return nil
	} else {
		if path == "/" {
			return errors.New("No git directory")
		}
		if err := r.Find(filepath.Join(path, "../")); err != nil {
			return err
		}
	}
	return nil
}

func (r *GitRepository) GetPath(path string) string {
	return filepath.Join(r.gitDir, path)
}

func (r *GitRepository) CreateDir(path string) error {
	dir := r.GetPath(path)

	if util.IsExist(dir) {
		return errors.New("The File already exists.")
	}
	if err := os.MkdirAll(dir, 0775); err != nil {
		return err
	}
	return nil
}

func (r *GitRepository) CreateFile(name string, content string) error {
	path := r.GetPath(name)

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	if _, err := f.Write([]byte(content)); err != nil {
		return err
	}
	return nil
}
